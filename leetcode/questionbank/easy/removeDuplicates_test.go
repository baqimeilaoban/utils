package easy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_removeDuplicates(t *testing.T) {
	convey.Convey("TestUnit_removeDuplicates", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{1, 1, 2}
			convey.So(removeDuplicates(nums), convey.ShouldEqual, 2)
		})
		convey.Convey("test-02", func() {
			nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
			convey.So(removeDuplicates(nums), convey.ShouldEqual, 5)
		})
	})
}
