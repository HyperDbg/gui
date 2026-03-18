package transparency

import (
	"fmt"
	"math"
	"sync"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/ux/widget/treetable"
)

type Measurement struct {
	Timestamp   int64
	Rdtsc       uint64
	VmexitCount uint64
}

type Provider struct {
	mu                 sync.RWMutex
	enabled            bool
	rdtscMeasurements  []uint64
	vmexitMeasurements []uint64
	maxMeasurements    int
	table              *treetable.TreeTable[Measurement]
}

func New() api.Interface {
	return &Provider{
		enabled:            false,
		rdtscMeasurements:  make([]uint64, 0),
		vmexitMeasurements: make([]uint64, 0),
		maxMeasurements:    1000,
	}
}

func (tm *Provider) Self() any {
	return tm
}

func (tm *Provider) Layout() layout.Widget {
	return tm.table.AirTable.Layout
}

func (tm *Provider) Update() error {
	return nil
}

func (tm *Provider) Clear() {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.rdtscMeasurements = make([]uint64, 0)
	tm.vmexitMeasurements = make([]uint64, 0)
}

func (tm *Provider) Enable() error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.enabled {
		return fmt.Errorf("transparency already enabled")
	}

	tm.enabled = true
	return nil
}

func (tm *Provider) Disable() error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if !tm.enabled {
		return fmt.Errorf("transparency already disabled")
	}

	tm.enabled = false
	return nil
}

func (tm *Provider) IsEnabled() bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.enabled
}

func (tm *Provider) AddRdtscMeasurement(rdtsc uint64) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if len(tm.rdtscMeasurements) >= tm.maxMeasurements {
		tm.rdtscMeasurements = tm.rdtscMeasurements[1:]
	}
	tm.rdtscMeasurements = append(tm.rdtscMeasurements, rdtsc)
}

func (tm *Provider) AddVmexitMeasurement(vmexit uint64) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if len(tm.vmexitMeasurements) >= tm.maxMeasurements {
		tm.vmexitMeasurements = tm.vmexitMeasurements[1:]
	}
	tm.vmexitMeasurements = append(tm.vmexitMeasurements, vmexit)
}

func (tm *Provider) GetRdtscMeasurements() []uint64 {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	result := make([]uint64, len(tm.rdtscMeasurements))
	copy(result, tm.rdtscMeasurements)
	return result
}

func (tm *Provider) GetVmexitMeasurements() []uint64 {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	result := make([]uint64, len(tm.vmexitMeasurements))
	copy(result, tm.vmexitMeasurements)
	return result
}

func (tm *Provider) CalculateRdtscVariation() float64 {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if len(tm.rdtscMeasurements) < 2 {
		return 0
	}

	var sum float64
	for _, v := range tm.rdtscMeasurements {
		sum += float64(v)
	}
	mean := sum / float64(len(tm.rdtscMeasurements))

	var variance float64
	for _, v := range tm.rdtscMeasurements {
		diff := float64(v) - mean
		variance += diff * diff
	}
	variance /= float64(len(tm.rdtscMeasurements))

	return math.Sqrt(variance)
}

func (tm *Provider) CalculateVmexitFrequency() float64 {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if len(tm.vmexitMeasurements) < 2 {
		return 0
	}

	var totalTime float64
	for i := 1; i < len(tm.vmexitMeasurements); i++ {
		diff := int64(tm.vmexitMeasurements[i]) - int64(tm.vmexitMeasurements[i-1])
		if diff > 0 {
			totalTime += float64(diff)
		}
	}

	avgInterval := totalTime / float64(len(tm.vmexitMeasurements)-1)
	if avgInterval == 0 {
		return 0
	}

	return 1.0 / avgInterval * 1e9
}

func (tm *Provider) ResetMeasurements() {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.rdtscMeasurements = make([]uint64, 0)
	tm.vmexitMeasurements = make([]uint64, 0)
}
