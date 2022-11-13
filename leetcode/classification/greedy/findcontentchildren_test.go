package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_findContentChildren(t *testing.T) {
	convey.Convey("TestUnit_findContentChildren", t, func() {
		convey.Convey("test-01", func() {
			g := []int{1, 2, 3}
			s := []int{1, 1}
			res := findContentChildren(g, s)
			convey.So(res, convey.ShouldEqual, 1)
		})
		convey.Convey("test-02", func() {
			g := []int{10,9,8,7}
			s := []int{5,6,7,8}
			res := findContentChildren(g, s)
			convey.So(res, convey.ShouldEqual, 2)
		})
	})
}
