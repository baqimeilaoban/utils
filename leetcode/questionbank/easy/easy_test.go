package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_twoSum(t *testing.T) {
	convey.Convey("TestUnit_twoSum", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{2, 7, 11, 15}
			target := 9
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{0, 1}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{3, 2, 4}
			target := 6
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{1, 2}), convey.ShouldBeTrue)
		})
		convey.Convey("test-03", func() {
			nums := []int{3, 3}
			target := 6
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{0, 1}), convey.ShouldBeTrue)
		})
	})
}
