// Package custerrors implements functions to generate custom errors in your app
//
// The New func returns an instance of custom error which has stack of trace in it
//
// New() error
//
// After call you can implement this new type of error
package custerror

import (
"fmt"
"runtime/debug"
)

//структура кастомной ошибки
type ErrorWithTrace struct {
	Text string
	Trace string
}
// функция создания для кастомной ошибки которая еще печатает стектрейс
func New(text string) error {
	return &ErrorWithTrace{
		Text: text,
		Trace: string(debug.Stack()),
	}
}
// метод Error для интерфейса error использует структуру ErrorWithTrace
func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("cust_error: %s\ntrace:\n%s", e.Text, e.Trace)
}
