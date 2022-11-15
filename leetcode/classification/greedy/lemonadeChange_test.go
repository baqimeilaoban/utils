package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_lemonadeChange(t *testing.T) {
	convey.Convey("TestUnit_lemonadeChange", t, func() {
		convey.Convey("test-01", func() {
			bills := []int{5, 5, 5, 10, 20}
			res := lemonadeChange(bills)
			convey.So(res, convey.ShouldEqual, true)
		})
		convey.Convey("test-02", func() {
			bills := []int{5, 5, 10, 10, 20}
			res := lemonadeChange(bills)
			convey.So(res, convey.ShouldEqual, false)
		})
	})
}
