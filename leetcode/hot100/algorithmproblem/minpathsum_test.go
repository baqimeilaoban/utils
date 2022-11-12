package algorithmproblem

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAlgorithmProblem_MinPaths(t *testing.T) {
	convey.Convey("TestAlgorithmProblem_MinPaths", t, func() {
		convey.Convey("test-01", func() {
			grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
			res := MinPathSum(grid)
			convey.So(res, convey.ShouldEqual, 7)
		})
	})
}
