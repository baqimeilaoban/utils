package algorithmproblem

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAlgorithmProblem_UniquePaths(t *testing.T) {
	convey.Convey("TestAlgorithmProblem_UniquePaths", t, func() {
		convey.Convey("test-01", func() {
			res := UniquePaths(3, 2)
			convey.So(res, convey.ShouldEqual, 3)
		})
		convey.Convey("test-02", func() {
			res := UniquePaths(3, 7)
			convey.So(res, convey.ShouldEqual, 28)
		})
	})
}
