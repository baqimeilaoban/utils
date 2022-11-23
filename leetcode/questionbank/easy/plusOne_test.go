package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_plusOne(t *testing.T) {
	convey.Convey("TestUnit_plusOne", t, func() {
		convey.Convey("test-01", func() {
			digits := []int{1, 2, 3}
			res := plusOne(digits)
			convey.So(reflect.DeepEqual(res, []int{1, 2, 4}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			digits := []int{4, 3, 2, 1}
			res := plusOne(digits)
			convey.So(reflect.DeepEqual(res, []int{4, 3, 2, 2}), convey.ShouldBeTrue)
		})
		convey.Convey("test-03", func() {
			digits := []int{0}
			res := plusOne(digits)
			convey.So(reflect.DeepEqual(res, []int{1}), convey.ShouldBeTrue)
		})
	})
}
