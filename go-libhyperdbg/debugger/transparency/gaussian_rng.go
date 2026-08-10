// Package transparency implements the HyperDbg anti-detection module.
//
// This file mirrors HyperDbg/hyperdbg/libhyperdbg/code/debugger/transparency/
// gaussian-rng.cpp. It provides:
//   - GaussianRng: a Gaussian (normal) random number generator based on the
//     polar form of the Box-Muller transform, identical to C++ Randn().
//   - Statistical helpers (Median, Average, CalculateStandardDeviation,
//     MedianAbsoluteDeviationTest, GaussianGenerateRandom) used by the
//     transparency measurement code.
//
// To produce bit-identical sequences to the C++ version (which calls rand()
// with the default MSVCRT seed of 1), the GaussianRng embeds a Go port of the
// MSVCRT rand() linear congruential generator:
//
//	seed = seed * 214013 + 2531011
//	return (seed >> 16) & 0x7fff
//
// RAND_MAX = 32767.
package transparency

import (
	"math"
	"sort"
	"sync"
)

// randMax is MSVCRT's RAND_MAX (see stdlib.h).
const randMax = 32767

// msvcRand replicates the MSVCRT rand() LCG so the Go GaussianRng produces the
// same sequence as C++ Randn() (which uses rand() with the default seed of 1,
// because the transparency code never calls srand()).
//
// See: https://learn.microsoft.com/en-us/cpp/c-runtime-library/reference/rand
type msvcRand struct {
	seed uint32
}

// next returns the next pseudo-random int32 in [0, randMax], matching MSVCRT
// rand() exactly.
func (r *msvcRand) next() int32 {
	r.seed = r.seed*214013 + 2531011
	return int32((r.seed >> 16) & 0x7fff)
}

// float64 returns the next pseudo-random value in [0, 1], matching the C++
// expression `((double)rand() / RAND_MAX)`.
func (r *msvcRand) float64() float64 {
	return float64(r.next()) / float64(randMax)
}

// GaussianRng is a Gaussian (normal) random number generator using the polar
// form of the Box-Muller transform. It mirrors C++ Randn(mu, sigma):
//
//   - Each iteration produces two independent samples (X1, X2); the second is
//     cached and returned on the following call (the C++ static `Call` toggle).
//   - The PRNG is the MSVCRT rand() LCG seeded with 1 (the C++ default).
//
// Concurrent use is safe: a mutex guards the cached state.
type GaussianRng struct {
	mu    float64
	sigma float64

	mu_  sync.Mutex
	rng  msvcRand
	call bool // toggles between returning the cached X2 and generating fresh
	x1   float64
	x2   float64
}

// NewGaussianRng creates a GaussianRng with the given mean (mu) and standard
// deviation (sigma). The PRNG is seeded with 1 to match C++ rand() default
// behaviour (the transparency code never calls srand()).
func NewGaussianRng(mean, stddev float64) *GaussianRng {
	return &GaussianRng{
		mu:    mean,
		sigma: stddev,
		rng:   msvcRand{seed: 1},
	}
}

// Seed sets the underlying PRNG seed. Use this to reproduce a specific C++
// rand() sequence after srand() is called. Most callers should leave the
// default seed of 1.
func (g *GaussianRng) Seed(seed uint32) {
	g.mu_.Lock()
	defer g.mu_.Unlock()
	g.rng.seed = seed
	g.call = false
	g.x1 = 0
	g.x2 = 0
}

// Next returns the next Gaussian-distributed sample. The polar Box-Muller
// transform generates two samples (X1, X2) per iteration; the second is cached
// and returned on the following call, exactly as in C++ Randn().
//
// The algorithm:
//
//	do {
//	    U1 = -1 + rand()/RAND_MAX * 2   // uniform in [-1, 1)
//	    U2 = -1 + rand()/RAND_MAX * 2
//	    W  = U1*U1 + U2*U2
//	} while (W >= 1 || W == 0)
//	mult = sqrt(-2 * log(W) / W)
//	X1 = U1 * mult
//	X2 = U2 * mult
//	return mu + sigma * X1   // (or X2 on the next call)
func (g *GaussianRng) Next() float64 {
	g.mu_.Lock()
	defer g.mu_.Unlock()

	// C++: if (Call == 1) { Call = !Call; return (mu + sigma * X2); }
	if g.call {
		g.call = !g.call
		return g.mu + g.sigma*g.x2
	}

	var u1, u2, w, mult float64
	for {
		// C++: U1 = -1 + ((double)rand() / RAND_MAX) * 2;
		u1 = -1 + g.rng.float64()*2
		u2 = -1 + g.rng.float64()*2
		w = u1*u1 + u2*u2
		if w < 1 && w != 0 {
			break
		}
	}

	mult = math.Sqrt((-2 * math.Log(w)) / w)
	g.x1 = u1 * mult
	g.x2 = u2 * mult
	g.call = !g.call

	return g.mu + g.sigma*g.x1
}

// NextUint64 returns Next() truncated to uint64. This matches the C++ pattern
// `(UINT64)Randn(...)` used by GuassianGenerateRandom when computing the
// average/standard-deviation/median statistics.
func (g *GaussianRng) NextUint64() uint64 {
	return uint64(g.Next())
}

// ----------------------------------------------------------------------------
// Statistical helpers — pure functions mirroring gaussian-rng.cpp.
// ----------------------------------------------------------------------------

// Median returns the median of the given samples. Mirrors C++ Median().
//
// The input slice is not modified; a sorted copy is used.
func Median(cases []float64) float64 {
	n := len(cases)
	if n == 0 {
		// C++ returns 0 for an empty vector ("Undefined, really").
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, cases)
	sort.Float64s(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

// Average returns the arithmetic mean of vec. Mirrors C++ Average<T>().
//
// Note: the C++ template returns 0.0 when the size is 1 (an upstream quirk);
// this Go port replicates that behaviour for parity with C++ measurements.
func Average(vec []float64) float64 {
	n := len(vec)
	if n == 1 {
		return 0.0
	}
	if n == 0 {
		return 0.0
	}
	sum := 0.0
	for _, v := range vec {
		sum += v
	}
	return sum / float64(n)
}

// CalculateStandardDeviation returns the population standard deviation of v.
// Mirrors C++ CalculateStandardDeviation<T>().
//
// Uses the identity `stdev = sqrt(E[X^2] - E[X]^2)` exactly as the C++ code
// (std::inner_product for the sum of squares, then sqrt of the difference).
func CalculateStandardDeviation(v []float64) float64 {
	n := len(v)
	if n == 0 {
		return 0
	}
	sum := 0.0
	for _, x := range v {
		sum += x
	}
	mean := sum / float64(n)

	sqSum := 0.0
	for _, x := range v {
		sqSum += x * x
	}
	variance := sqSum/float64(n) - mean*mean
	if variance < 0 {
		// Floating-point round-off can push the variance slightly negative
		// when all values are identical; clamp to zero for a real stdev.
		return 0
	}
	return math.Sqrt(variance)
}

// MedianAbsoluteDeviationTest returns 1.4826 * MAD of the data. Mirrors C++
// MedianAbsoluteDeviationTest().
//
// The C++ function takes `vector<double> Data` by value (a copy), computes the
// median of the original, replaces each element with abs(element - median),
// then returns 1.4826 * median of the modified copy. This Go port reproduces
// that exactly (the input slice is not modified by the caller-visible state).
func MedianAbsoluteDeviationTest(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	medianData := Median(data)
	absDevs := make([]float64, n)
	for i, x := range data {
		absDevs[i] = math.Abs(x - medianData)
	}
	return 1.4826 * Median(absDevs)
}

// GaussianGenerateRandom computes the average, standard deviation (with +5
// added for variance) and median of the data after removing outliers via the
// 3*MAD test. Mirrors C++ GuassianGenerateRandom().
//
// Returns (average, standardDeviation, median) as uint64, matching the C++
// output parameters. The +5 on the standard deviation is intentional in C++
// ("this value might be 0 or 1 so we need more variance") and is preserved.
func GaussianGenerateRandom(data []float64) (average, stddev, median uint64) {
	if len(data) == 0 {
		// C++ would push nothing into FinalData and divide by zero; we return
		// safe defaults so callers don't panic.
		return 0, 5, 0
	}

	// C++: vector<double> OriginalData = Data;  (copy)
	original := make([]float64, len(data))
	copy(original, data)

	// C++: Mad = MedianAbsoluteDeviationTest(ChangableData);
	// (ChangableData == OriginalData at this point; the move is irrelevant
	// because MedianAbsoluteDeviationTest takes its argument by value.)
	mad := MedianAbsoluteDeviationTest(data)
	medians := Median(original)

	// C++: filter outliers — keep items where |item - medians| <= 3 * Mad.
	// (The C++ condition is `item > (3*Mad)+medians || item < -(3*Mad)+medians`.)
	finalData := make([]float64, 0, len(original))
	for _, item := range original {
		if item > (3*mad)+medians || item < -(3*mad)+medians {
			// outlier — skip
			continue
		}
		finalData = append(finalData, item)
	}

	if len(finalData) == 0 {
		return 0, 5, 0
	}

	stddevF := CalculateStandardDeviation(finalData)
	avgF := Average(finalData)
	medianF := Median(finalData)

	average = uint64(avgF)
	// C++: *StandardDeviationOfData = (UINT64)StandardDeviation + 5;
	stddev = uint64(stddevF) + 5
	median = uint64(medianF)
	return
}
