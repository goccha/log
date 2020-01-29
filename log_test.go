package log

import (
	"errors"
	"testing"
)

func TestWarn(t *testing.T) {
	err := errors.New("warning")
	Warn("%+v", err)
}

func TestError(t *testing.T) {
	err := errors.New("error")
	Error("%+v", err)
}
