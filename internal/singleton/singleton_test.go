package singleton

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetInstance(t *testing.T) {
	Convey("normal", t, func() {
		GetInstance()
	})
	Convey("normal", t, func() {
		GetInstance()
	})
}
