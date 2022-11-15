package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_candy(t *testing.T) {
	convey.Convey("TestUnit_candy", t, func() {
		convey.Convey("test-01", func() {
			ratings := []int{1, 0, 2}
			res := candy(ratings)
			convey.So(res, convey.ShouldEqual, 5)
		})
		convey.Convey("test-02", func() {
			ratings := []int{1, 2, 2}
			res := candy(ratings)
			convey.So(res, convey.ShouldEqual, 4)
		})
	})
}
