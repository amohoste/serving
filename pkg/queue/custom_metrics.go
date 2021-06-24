package queue

import "sync"

type CustomMetricStore struct {
	mu sync.Mutex
	totalExecutionTime float64
	totalExecutions int32
}

type CustomMetricReport struct {
	totalExecutionTime float64
	totalExecutions int32
}

// Adds an execution to the metric store
func (s *CustomMetricStore) LogExecution(executionTime float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalExecutionTime += executionTime
	s.totalExecutions++
}

// Returns current store state
func (s *CustomMetricStore) Get() CustomMetricReport {
	s.mu.Lock()
	defer s.mu.Unlock()

	report := CustomMetricReport{
		totalExecutionTime: s.totalExecutionTime,
		totalExecutions: s.totalExecutions,
	}

	return report
}

// Returns and resets current store state
func (s *CustomMetricStore) Report() CustomMetricReport {
	s.mu.Lock()
	defer s.mu.Unlock()

	report := CustomMetricReport{
		totalExecutionTime: s.totalExecutionTime,
		totalExecutions: s.totalExecutions,
	}

	s.totalExecutionTime = 0
	s.totalExecutions = 0

	return report
}

// Resets metric store state
func (s *CustomMetricStore) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalExecutionTime = 0
	s.totalExecutions = 0

}

var CustomMetrics *CustomMetricStore

func init() {
	// use package init to make sure path is always instantiated
	CustomMetrics = new(CustomMetricStore)
	CustomMetrics.totalExecutionTime = 0
	CustomMetrics.totalExecutions = 0
}