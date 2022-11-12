package algorithmproblem

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_climbStairs(t *testing.T) {
	convey.Convey("TestUnit_climbStairs", t, func() {
		convey.Convey("test-01", func() {
			res := climbStairs(2)
			convey.So(res, convey.ShouldEqual, 2)
		})
		convey.Convey("test-02", func() {
			res := climbStairs(3)
			convey.So(res, convey.ShouldEqual, 3)
		})
	})
}
