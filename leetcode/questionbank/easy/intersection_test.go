package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_intersection(t *testing.T) {
	convey.Convey("TestUnit_intersection", t, func() {
		convey.Convey("test-01", func() {
			nums1 := []int{1, 2, 2, 1}
			nums2 := []int{2, 2}
			res := intersection(nums1, nums2)
			convey.So(reflect.DeepEqual(res, []int{2}), convey.ShouldBeTrue)
		})
	})
}
