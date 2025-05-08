package circuitbreaker

import (
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type OperationType int

const (
	Read OperationType = iota
	Write
)

type CircuitBreaker struct {
	state           State
	failureCount    int
	lastFailureTime time.Time
	mutex           sync.RWMutex
	maxFailures     int
	resetTimeout    time.Duration
}

func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        Closed,
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
	}
}

func (cb *CircuitBreaker) Execute(operation func() error, opType OperationType) error {
	if !cb.allowRequest() {
		if opType == Read {
			return ErrCircuitNotFound
		}
		return ErrCircuitOpen
	}

	err := operation()
	cb.recordResult(err)
	return err
}

func (cb *CircuitBreaker) allowRequest() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()

	switch cb.state {
	case Closed:
		return true
	case Open:
		if time.Since(cb.lastFailureTime) >= cb.resetTimeout {
			cb.mutex.RUnlock()
			cb.mutex.Lock()
			cb.state = HalfOpen
			cb.mutex.Unlock()
			cb.mutex.RLock()
			return true
		}
		return false
	case HalfOpen:
		return true
	default:
		return false
	}
}

func (cb *CircuitBreaker) recordResult(err error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()

		if cb.failureCount >= cb.maxFailures {
			cb.state = Open
		}
	} else {
		if cb.state == HalfOpen {
			cb.state = Closed
			cb.failureCount = 0
		}
	}
}

var ErrCircuitOpen = &CircuitBreakerError{message: "circuit breaker is open"}
var ErrCircuitNotFound = &CircuitBreakerError{message: "resource not found"}

type CircuitBreakerError struct {
	message string
}

func (e *CircuitBreakerError) Error() string {
	return e.message
}
