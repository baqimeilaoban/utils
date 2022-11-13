package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_maxSubArray(t *testing.T) {
	convey.Convey("TestUnit_maxSubArray", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			res := maxSubArray(nums)
			convey.So(res, convey.ShouldEqual, 6)
		})
		convey.Convey("test-02", func() {
			nums := []int{1}
			res := maxSubArray(nums)
			convey.So(res, convey.ShouldEqual, 1)
		})
	})
}

func TestUnit_maxSubArrayOpt(t *testing.T) {
	convey.Convey("TestUnit_maxSubArrayOpt", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			res := maxSubArrayOpt(nums)
			convey.So(res, convey.ShouldEqual, 6)
		})
		convey.Convey("test-02", func() {
			nums := []int{1}
			res := maxSubArrayOpt(nums)
			convey.So(res, convey.ShouldEqual, 1)
		})
	})
}

func TestUnit_maxSubArrayOpt02(t *testing.T) {
	convey.Convey("TestUnit_maxSubArrayOpt02", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			res := maxSubArrayOpt2(nums)
			convey.So(res, convey.ShouldEqual, 6)
		})
		convey.Convey("test-02", func() {
			nums := []int{1}
			res := maxSubArrayOpt2(nums)
			convey.So(res, convey.ShouldEqual, 1)
		})
	})
}
