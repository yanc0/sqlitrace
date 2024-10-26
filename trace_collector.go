package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Stat struct {
	min   time.Duration
	max   time.Duration
	mean  time.Duration
	count int64
}

func (s *Stat) String() string {
	return fmt.Sprintf("mean: %s | min: %s | max: %s | count: %d", s.mean, s.min, s.max, s.count)
}

func (stat *Stat) AddPoint(duration time.Duration) *Stat {
	stat.count++

	if duration < stat.min {
		stat.min = duration
	}

	if duration > stat.max {
		stat.max = duration
	}

	stat.mean = stat.mean + (duration-stat.mean)/time.Duration(stat.count)

	return stat
}

type TraceCollector struct {
	mu        *sync.Mutex
	selfstat  *Stat
	stmtstats map[string]*Stat
}

func NewTraceCollector(ctx context.Context) *TraceCollector {
	collector := &TraceCollector{
		mu:        new(sync.Mutex),
		selfstat:  new(Stat),
		stmtstats: make(map[string]*Stat),
	}

	go collector.printStats()

	return collector
}

func (collector *TraceCollector) printStats() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		collector.mu.Lock()
		for stmt, stat := range collector.stmtstats {
			fmt.Println(collector.selfstat, "[ SELF STATS ]")
			fmt.Println(stat, "[", stmt, "]")
		}
		collector.mu.Unlock()
	}
}

func (collector *TraceCollector) SaveProbe(stmt string, duration time.Duration) {
	start := time.Now()
	collector.mu.Lock()
	defer collector.mu.Unlock()

	stat, found := collector.stmtstats[stmt]
	if !found {
		collector.stmtstats[stmt] = &Stat{
			min:   duration,
			max:   duration,
			mean:  duration,
			count: 1,
		}
		collector.selfstat = collector.selfstat.AddPoint(time.Since(start))
		return
	}
	stat.AddPoint(duration)
	collector.selfstat = collector.selfstat.AddPoint(time.Since(start))
}
