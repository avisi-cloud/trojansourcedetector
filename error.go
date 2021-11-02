package trojansourcedetector

import (
	"encoding/json"
	"sync"
)

// Error describes a detected error.
type Error interface {
	// Code returns the ErrorCode of the detector.
	Code() ErrorCode
	// Details contains the details of the error.
	Details() string
	// File returns the ErrorCode of the file this error was found in.
	File() string
	// Line returns the number of the line this error was found in.
	Line() uint
	// Column returns the number of the byte this error was caused by.
	Column() uint
	// JSON returns a JSON-encoded variant of this error.
	JSON() ([]byte, error)
}

// NewErrors creates a new implementation of Errors.
func NewErrors() Errors {
	return &errorsImpl{
		lock: &sync.Mutex{},
	}
}

// Errors is a container to hold multiple errors.
type Errors interface {
	// Add adds a new error.
	Add(code ErrorCode, details string, file string, line uint, column uint)
	// AddAll adds all errors from another Errors instance.
	AddAll(e Errors)
	// Get returns all errors collected.
	Get() []Error
}

type errorsImpl struct {
	lock   *sync.Mutex
	errors []Error
}

func (e *errorsImpl) AddAll(e2 Errors) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = append(e.errors, e2.Get()...)
}

func (e *errorsImpl) Add(code ErrorCode, details string, file string, line uint, column uint) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = append(e.errors, &errorImpl{
		ErrorCode:    code,
		ErrorDetails: details,
		ErrorFile:    file,
		ErrorLine:    line,
		ErrorColumn:  column,
	})
}

func (e *errorsImpl) Get() []Error {
	return e.errors
}

type errorImpl struct {
	ErrorCode    ErrorCode `json:"name"`
	ErrorFile    string    `json:"file"`
	ErrorLine    uint      `json:"line"`
	ErrorColumn  uint      `json:"column"`
	ErrorDetails string    `json:"details"`
}

func (e errorImpl) Details() string {
	return e.ErrorDetails
}

func (e errorImpl) Code() ErrorCode {
	return e.ErrorCode
}

func (e errorImpl) File() string {
	return e.ErrorFile
}

func (e errorImpl) Line() uint {
	return e.ErrorLine
}

func (e errorImpl) Column() uint {
	return e.ErrorColumn
}

func (e errorImpl) JSON() ([]byte, error) {
	return json.Marshal(e)
}
