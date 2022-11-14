package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_largestSumAfterKNegations(t *testing.T) {
	convey.Convey("TestUnit_largestSumAfterKNegations", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{4, 2, 3}
			k := 1
			res := largestSumAfterKNegations(nums, k)
			convey.So(res, convey.ShouldEqual, 5)
		})
		convey.Convey("test-02", func() {
			nums := []int{3, -1, 0, 2}
			k := 3
			convey.So(largestSumAfterKNegations(nums, k), convey.ShouldEqual, 6)
		})
		convey.Convey("test-03", func() {
			nums := []int{2, -3, -1, 5, -4}
			k := 2
			convey.So(largestSumAfterKNegations(nums, k), convey.ShouldEqual, 13)
		})
	})
}
