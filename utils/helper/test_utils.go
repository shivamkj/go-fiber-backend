package helper

import (
	"reflect"
	"testing"

	"github.com/qnify/api-server/utils/errors"
)

func TestPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Error(errors.New("code did not panic"))
	}
}

func ShouldErr(err error, t *testing.T) {
	if err == nil {
		t.Error(errors.Wrap("should err failed", err))
	}
}

func NoErr(err error, t *testing.T) {
	if err != nil {
		t.Error(errors.Wrap("found error in no err", err))
	}
}

func Check(condition bool, t *testing.T) {
	if !condition {
		t.Error(errors.New("conditional check failed"))
	}
}

func DeepEqual(x, y any, t *testing.T) {
	if !reflect.DeepEqual(x, y) {
		t.Error(errors.New("deep equal check failed"))
	}
}
