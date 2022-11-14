package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_jump(t *testing.T) {
	convey.Convey("TestUnit_jump", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{2, 3, 1, 1, 4}
			res := jump(nums)
			convey.So(res, convey.ShouldEqual, 2)
		})
		convey.Convey("test-02", func() {
			nums := []int{2, 3, 0, 1, 4}
			res := jump(nums)
			convey.So(res, convey.ShouldEqual, 2)
		})
	})
}
