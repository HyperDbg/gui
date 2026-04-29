package transparency

import (
	"fmt"
	"maps"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/common"
)

type TransparencyMode int

const (
	ModeDisabled TransparencyMode = iota
	ModeEnabled
)

type GaussianRng struct {
	mu sync.Mutex
	r  *rand.Rand
}

func NewGaussianRng() *GaussianRng {
	return &GaussianRng{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *GaussianRng) NextFloat() float64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	u1 := g.r.Float64()
	if u1 < 1e-10 {
		u1 = 1e-10
	}
	u2 := g.r.Float64()
	return math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
}

func (g *GaussianRng) NextInt(min, max int) int {
	val := g.NextFloat()
	range_ := max - min
	result := min + int(float64(range_)*val)
	if result > max {
		result = max
	}
	if result < min {
		result = min
	}
	return result
}

type TransparencyManager struct {
	mode        TransparencyMode
	processes   map[uint32]TransparencyProcessState
	gaussianRng *GaussianRng
	mu          sync.RWMutex
}

type TransparencyProcessState struct {
	Pid       uint32
	Name      string
	IsHidden  bool
	HookCount int
}

func NewTransparencyManager() *TransparencyManager {
	return &TransparencyManager{
		mode:        ModeDisabled,
		processes:   make(map[uint32]TransparencyProcessState),
		gaussianRng: NewGaussianRng(),
	}
}

func (tm *TransparencyManager) Enable(pid uint32, processName string, usePid bool) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var targetPid uint32
	if usePid {
		targetPid = pid
	} else {
		targetPid = pid
	}

	state := TransparencyProcessState{
		Pid:      targetPid,
		Name:     processName,
		IsHidden: true,
	}
	tm.processes[targetPid] = state
	tm.mode = ModeEnabled

	return nil
}

func (tm *TransparencyManager) Disable() error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for id := range tm.processes {
		delete(tm.processes, id)
	}
	tm.mode = ModeDisabled
	return nil
}

func (tm *TransparencyManager) IsEnabled() bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.mode == ModeEnabled
}

func (tm *TransparencyManager) IsProcessHidden(pid uint32) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	state, exists := tm.processes[pid]
	return exists && state.IsHidden
}

func (tm *TransparencyManager) GetHiddenProcesses() map[uint32]TransparencyProcessState {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	result := make(map[uint32]TransparencyProcessState)
	maps.Copy(result, tm.processes)
	return result
}

func (tm *TransparencyManager) AddHook(pid uint32) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if state, ok := tm.processes[pid]; ok {
		state.HookCount++
		tm.processes[pid] = state
	}
}

func (tm *TransparencyManager) RemoveHook(pid uint32) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if state, ok := tm.processes[pid]; ok {
		if state.HookCount > 0 {
			state.HookCount--
			tm.processes[pid] = state
		}
	}
}

func makeUint64(lo, hi uint32) uint64 {
	return uint64(hi)<<32 | uint64(lo)
}

func TransparentModeRdtscDiffVmexit() uint64 {
	retLo, retHi := common.RdTsc()
	ret := makeUint64(retLo, retHi)

	common.CpuId(0x00)

	ret2Lo, ret2Hi := common.RdTsc()
	ret2 := makeUint64(ret2Lo, ret2Hi)

	return ret2 - ret
}

func TransparentModeRdtscVmexitTracing() uint64 {
	retLo, retHi := common.RdTsc()
	ret := makeUint64(retLo, retHi)

	ret2Lo, ret2Hi := common.RdTsc()
	ret2 := makeUint64(ret2Lo, ret2Hi)

	return ret2 - ret
}

type HypervisorDetectionResult struct {
	Average              uint64
	StandardDeviation    uint64
	Median               uint64
	IsHypervisorDetected bool
}

func TransparentModeCpuidTimeStampCounter() (*HypervisorDetectionResult, error) {
	var avg float64
	var measuredTime float64
	var results []float64

	for range TestCount {
		measuredTime = float64(TransparentModeRdtscDiffVmexit())
		avg = avg + measuredTime
		results = append(results, measuredTime)
	}

	gaussianResult := GuassianGenerateRandom(results)

	avg = avg / float64(TestCount)
	isHypervisorDetected := !(avg < 1000 && avg > 0)

	return &HypervisorDetectionResult{
		Average:              gaussianResult.Average,
		StandardDeviation:    gaussianResult.StandardDeviation,
		Median:               gaussianResult.Median,
		IsHypervisorDetected: isHypervisorDetected,
	}, nil
}

func TransparentModeRdtscEmulationDetection() (*HypervisorDetectionResult, error) {
	var avg float64
	var measuredTime float64
	var results []float64

	for range TestCount {
		measuredTime = float64(TransparentModeRdtscVmexitTracing())
		avg = avg + measuredTime
		results = append(results, measuredTime)
	}

	gaussianResult := GuassianGenerateRandom(results)

	avg = avg / float64(TestCount)
	isEmulationDetected := !(avg < 750 && avg > 0)

	return &HypervisorDetectionResult{
		Average:              gaussianResult.Average,
		StandardDeviation:    gaussianResult.StandardDeviation,
		Median:               gaussianResult.Median,
		IsHypervisorDetected: isEmulationDetected,
	}, nil
}

func TransparentModeCheckHypervisorPresence() (*HypervisorDetectionResult, error) {
	result, err := TransparentModeCpuidTimeStampCounter()
	if err != nil {
		return nil, err
	}

	if result.IsHypervisorDetected {
		fmt.Println("hypervisor detected")
	} else {
		fmt.Println("hypervisor not detected")
	}

	return result, nil
}

func TransparentModeCheckRdtscpVmexit() (*HypervisorDetectionResult, error) {
	result, err := TransparentModeRdtscEmulationDetection()
	if err != nil {
		return nil, err
	}

	if result.IsHypervisorDetected {
		fmt.Println("rdtsc/p emulation detected")
	} else {
		fmt.Println("rdtsc/p emulation not detected")
	}

	return result, nil
}
