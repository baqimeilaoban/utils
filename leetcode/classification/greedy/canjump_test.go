package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_canJump(t *testing.T) {
	convey.Convey("TestUnit_canJump", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{2, 3, 1, 1, 4}
			res := canJump(nums)
			convey.So(res, convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{3, 2, 1, 0, 4}
			res := canJump(nums)
			convey.So(res, convey.ShouldBeFalse)
		})
	})
}
