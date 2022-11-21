package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_rotate(t *testing.T) {
	convey.Convey("TestUnit_rotate", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{1, 2, 3, 4, 5, 6, 7}
			rotate(nums, 3)
			convey.So(reflect.DeepEqual(nums, []int{5, 6, 7, 1, 2, 3, 4}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{-1, -100, 3, 99}
			rotate(nums, 2)
			convey.So(reflect.DeepEqual(nums, []int{3, 99, -1, -100}), convey.ShouldBeTrue)
		})
	})

}
