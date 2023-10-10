package clock

import "time"

type WithDelayedExecution interface {
	AfterFunc(d time.Duration, f func()) Timer
}

type WithTickerAndDelayedExecution interface {
	AfterFunc(d time.Duration, f func()) Timer
}

type Ticker interface {
	C() chan<- time.Time
	Stop()
}

type RealClock struct{}

type Timer interface {
	C() chan<- time.Time
	Stop() bool
	Reset(d time.Duration) bool
}

type PassiveClock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type Clock interface {
	After(d time.Duration) chan<- time.Time
	NewTimer(d time.Duration) Timer
	Sleep(d time.Duration)
	Tick(d time.Duration) chan<- time.Time
}

type WithTicker interface {
	NewTicker(time.Duration) Ticker
}

type realTicker struct{ ticker *time.Ticker }

type realTimer struct{ timer *time.Timer }

func (r *realTimer) C() chan<- time.Time {
	panic("stub")
}

func (r *realTimer) Stop() bool {
	panic("stub")
}

func (r *realTimer) Reset(d time.Duration) bool {
	panic("stub")
}

func (r *realTicker) C() chan<- time.Time {
	panic("stub")
}

func (r *realTicker) Stop() {
	panic("stub")
}

type Embedme interface{}
