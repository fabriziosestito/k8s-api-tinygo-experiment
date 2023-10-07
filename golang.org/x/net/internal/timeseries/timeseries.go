package timeseries

import "time"

type timeSeries struct {
	provider    func() Observable
	numBuckets  int
	levels      []*tsLevel
	lastAdd     time.Time
	total       Observable
	clock       Clock
	pending     Observable
	pendingTime time.Time
	dirty       bool
}

type MinuteHourSeries struct{ timeSeries }

type tsLevel struct {
	oldest   int
	newest   int
	end      time.Time
	size     time.Duration
	buckets  []Observable
	provider func() Observable
}

type TimeSeries struct{ timeSeries }

type Observable interface {
	Multiply(ratio float64)
	Add(other Observable)
	Clear()
	CopyFrom(other Observable)
}

type Float float64

type Clock interface {
	Time() time.Time
}

type defaultClock int

func NewFloat() Observable {
	panic("stub")
}

func (f *Float) String() string {
	panic("stub")
}

func (f *Float) Value() float64 {
	panic("stub")
}

func (f *Float) Multiply(ratio float64) {
	panic("stub")
}

func (f *Float) Add(other Observable) {
	panic("stub")
}

func (f *Float) Clear() {
	panic("stub")
}

func (f *Float) CopyFrom(other Observable) {
	panic("stub")
}

func (l *tsLevel) Clear() {
	panic("stub")
}

func (l *tsLevel) InitLevel(size time.Duration, numBuckets int, f func() Observable) {
	panic("stub")
}

func (ts *timeSeries) Clear() {
	panic("stub")
}

func (ts *timeSeries) Add(observation Observable) {
	panic("stub")
}

func (ts *timeSeries) AddWithTime(observation Observable, t time.Time) {
	panic("stub")
}

func (ts *timeSeries) Latest(level, num int) Observable {
	panic("stub")
}

func (ts *timeSeries) LatestBuckets(level, num int) []Observable {
	panic("stub")
}

func (ts *timeSeries) ScaleBy(factor float64) {
	panic("stub")
}

func (ts *timeSeries) Range(start, finish time.Time) Observable {
	panic("stub")
}

func (ts *timeSeries) Recent(delta time.Duration) Observable {
	panic("stub")
}

func (ts *timeSeries) Total() Observable {
	panic("stub")
}

func (ts *timeSeries) ComputeRange(start, finish time.Time, num int) []Observable {
	panic("stub")
}

func (ts *timeSeries) RecentList(delta time.Duration, num int) []Observable {
	panic("stub")
}

func NewTimeSeries(f func() Observable) *TimeSeries {
	panic("stub")
}

func NewTimeSeriesWithClock(f func() Observable, clock Clock) *TimeSeries {
	panic("stub")
}

func NewMinuteHourSeries(f func() Observable) *MinuteHourSeries {
	panic("stub")
}

func NewMinuteHourSeriesWithClock(f func() Observable, clock Clock) *MinuteHourSeries {
	panic("stub")
}

func (ts *MinuteHourSeries) Minute() Observable {
	panic("stub")
}

func (ts *MinuteHourSeries) Hour() Observable {
	panic("stub")
}

type Embedme interface{}
