package algorithmproblem

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAlgorithmProblem_Merge(t *testing.T) {
	convey.Convey("TestAlgorithmProblem", t, func() {
		convey.Convey("test-1", func() {
			intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
			res := Merge(intervals)
			expect := [][]int{{1, 6}, {8, 10}, {15, 18}}
			convey.So(reflect.DeepEqual(res, expect), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			intervals := [][]int{{1, 4}, {4, 5}}
			res := Merge(intervals)
			expect := [][]int{{1, 5}}
			convey.So(reflect.DeepEqual(res, expect), convey.ShouldBeTrue)
		})
	})
}
