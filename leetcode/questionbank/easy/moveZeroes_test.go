package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_moveZeroes(t *testing.T) {
	convey.Convey("TestUnit_moveZeroes", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{0, 1, 0, 3, 12}
			moveZeroes(nums)
			convey.So(reflect.DeepEqual(nums, []int{1, 3, 12, 0, 0}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{0}
			moveZeroes(nums)
			convey.So(reflect.DeepEqual(nums, []int{0}), convey.ShouldBeTrue)
		})
	})
}
