package greedy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_reconstructQueue(t *testing.T) {
	convey.Convey("TestUnit_reconstructQueue", t, func() {
		convey.Convey("test-01", func() {
			people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
			res := reconstructQueue(people)
			expect := [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
			convey.So(reflect.DeepEqual(res, expect), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			people := [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
			res := reconstructQueue(people)
			expect := [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}
			convey.So(reflect.DeepEqual(res, expect), convey.ShouldBeTrue)
		})
	})
}
