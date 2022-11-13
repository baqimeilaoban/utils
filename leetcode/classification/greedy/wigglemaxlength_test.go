package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_wiggleMaxLength(t *testing.T) {
	convey.Convey("TestUnit_wiggleMaxLength", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{1, 7, 4, 9, 2, 5}
			res := wiggleMaxLength(nums)
			convey.So(res, convey.ShouldEqual, 6)
		})
		convey.Convey("test-02", func() {
			nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
			res := wiggleMaxLength(nums)
			convey.So(res, convey.ShouldEqual, 7)
		})
	})
}
