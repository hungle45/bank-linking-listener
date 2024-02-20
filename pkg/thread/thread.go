package thread

import "sync"

type RoutineGroup struct {
	wg *sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup {
	return &RoutineGroup{wg: &sync.WaitGroup{}}
}

func (rg *RoutineGroup) Run(fn func()) {
	rg.wg.Add(1)
	go func() {
		defer rg.wg.Done()
		fn()
	}()
}

func (rg *RoutineGroup) Wait() {
	rg.wg.Wait()
}
