package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_findMinArrowShots(t *testing.T) {
	convey.Convey("TestUnit_findMinArrowShots", t, func() {
		convey.Convey("test-01", func() {
			points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
			res := findMinArrowShots(points)
			convey.So(res, convey.ShouldEqual, 2)
		})
		convey.Convey("test-02", func() {
			points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
			res := findMinArrowShots(points)
			convey.So(res, convey.ShouldEqual, 4)
		})
		convey.Convey("test-03", func() {
			points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
			res := findMinArrowShots(points)
			convey.So(res, convey.ShouldEqual, 2)
		})
	})
}
