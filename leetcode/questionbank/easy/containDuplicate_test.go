package easy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_containsDuplicate(t *testing.T) {
	convey.Convey("TestUnit_containsDuplicate", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{1, 2, 3, 1}
			res := containsDuplicate(nums)
			convey.So(res, convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{1, 2, 3, 4}
			res := containsDuplicate(nums)
			convey.So(res, convey.ShouldBeFalse)
		})
		convey.Convey("test-03", func() {
			nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
			res := containsDuplicate(nums)
			convey.So(res, convey.ShouldBeTrue)
		})
	})
}
