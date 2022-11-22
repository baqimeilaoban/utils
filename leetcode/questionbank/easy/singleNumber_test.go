package easy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_singleNumber(t *testing.T) {
	convey.Convey("TestUnit_singleNumber", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{2, 2, 1}
			res := singleNumber(nums)
			convey.So(res, convey.ShouldEqual, 1)
		})
		convey.Convey("test-02", func() {
			nums := []int{4, 1, 2, 1, 2}
			res := singleNumberOpt(nums)
			convey.So(res, convey.ShouldEqual, 4)
		})
		convey.Convey("test-03", func() {
			nums := []int{1}
			res := singleNumber(nums)
			convey.So(res, convey.ShouldEqual, 1)
		})
	})
}
