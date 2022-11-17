package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_eraseOverlapIntervals(t *testing.T) {
	convey.Convey("TestUnit_eraseOverlapIntervals", t, func() {
		convey.Convey("test-01", func() {
			intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
			res := eraseOverlapIntervals(intervals)
			convey.So(res, convey.ShouldEqual, 1)
		})
		convey.Convey("test-02", func() {
			intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
			res := eraseOverlapIntervals(intervals)
			convey.So(res, convey.ShouldEqual, 2)
		})
		convey.Convey("test-03", func() {
			intervals := [][]int{{1, 2}, {2, 3}}
			res := eraseOverlapIntervals(intervals)
			convey.So(res, convey.ShouldEqual, 0)
		})
	})
}
