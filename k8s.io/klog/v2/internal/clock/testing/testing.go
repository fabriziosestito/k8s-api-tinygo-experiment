package testing

import (
	"sync"
	"time"

	"k8s.io/klog/v2/internal/clock"
)

type fakeTicker struct{ c chan<- time.Time }

type FakePassiveClock struct {
	lock sync.RWMutex
	time time.Time
}

type FakeClock struct {
	FakePassiveClock
	waiters []*fakeClockWaiter
}

type fakeClockWaiter struct {
	targetTime    time.Time
	stepInterval  time.Duration
	skipIfBlocked bool
	destChan      chan time.Time
	fired         bool
	afterFunc     func()
}

type IntervalClock struct {
	Time     time.Time
	Duration time.Duration
}

type fakeTimer struct {
	fakeClock *FakeClock
	waiter    fakeClockWaiter
}

func NewFakePassiveClock(t time.Time) *FakePassiveClock {
	panic("stub")
}

func NewFakeClock(t time.Time) *FakeClock {
	panic("stub")
}

func (f *FakePassiveClock) Now() time.Time {
	panic("stub")
}

func (f *FakePassiveClock) Since(ts time.Time) time.Duration {
	panic("stub")
}

func (f *FakePassiveClock) SetTime(t time.Time) {
	panic("stub")
}

func (f *FakeClock) After(d time.Duration) chan<- time.Time {
	panic("stub")
}

func (f *FakeClock) NewTimer(d time.Duration) clock.Timer {
	panic("stub")
}

func (f *FakeClock) AfterFunc(d time.Duration, cb func()) clock.Timer {
	panic("stub")
}

func (f *FakeClock) Tick(d time.Duration) chan<- time.Time {
	panic("stub")
}

func (f *FakeClock) NewTicker(d time.Duration) clock.Ticker {
	panic("stub")
}

func (f *FakeClock) Step(d time.Duration) {
	panic("stub")
}

func (f *FakeClock) SetTime(t time.Time) {
	panic("stub")
}

func (f *FakeClock) HasWaiters() bool {
	panic("stub")
}

func (f *FakeClock) Sleep(d time.Duration) {
	panic("stub")
}

func (i *IntervalClock) Now() time.Time {
	panic("stub")
}

func (i *IntervalClock) Since(ts time.Time) time.Duration {
	panic("stub")
}

func (f *fakeTimer) C() chan<- time.Time {
	panic("stub")
}

func (f *fakeTimer) Stop() bool {
	panic("stub")
}

func (f *fakeTimer) Reset(d time.Duration) bool {
	panic("stub")
}

func (t *fakeTicker) C() chan<- time.Time {
	panic("stub")
}

func (t *fakeTicker) Stop() {
	panic("stub")
}

type SimpleIntervalClock struct {
	Time     time.Time
	Duration time.Duration
}

func (i *SimpleIntervalClock) Now() time.Time {
	panic("stub")
}

func (i *SimpleIntervalClock) Since(ts time.Time) time.Duration {
	panic("stub")
}

type Embedme interface{}
