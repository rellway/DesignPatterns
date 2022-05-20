package singleton

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetInstance(t *testing.T) {
	convey.Convey("normal", t, func() {
		GetInstance()
	})
	convey.Convey("normal", t, func() {
		GetInstance()
	})
}
