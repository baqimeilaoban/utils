package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_intersect(t *testing.T) {
	convey.Convey("TestUnit_intersect", t, func() {
		convey.Convey("test-01", func() {
			nums1 := []int{1, 2, 2, 1}
			nums2 := []int{2, 2}
			res := intersect(nums1, nums2)
			convey.So(reflect.DeepEqual(res, []int{2, 2}), convey.ShouldBeTrue)
		})
	})
}
