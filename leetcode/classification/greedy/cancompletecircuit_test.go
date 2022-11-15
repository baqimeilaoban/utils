package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_canCompleteCircuit(t *testing.T) {
	convey.Convey("TestUnit_canCompleteCircuit", t, func() {
		convey.Convey("test-01", func() {
			g := []int{1, 2, 3, 4, 5}
			c := []int{3, 4, 5, 1, 2}
			res := canCompleteCircuit(g, c)
			convey.So(res, convey.ShouldEqual, 3)
		})
		convey.Convey("test-02", func() {
			g := []int{2, 3, 4}
			c := []int{3, 4, 3}
			res := canCompleteCircuit(g, c)
			convey.So(res, convey.ShouldEqual, -1)
		})
	})
}

func TestUnit_canCompleteCircuitOpt(t *testing.T) {
	convey.Convey("TestUnit_canCompleteCircuitOpt", t, func() {
		convey.Convey("test-01", func() {
			g := []int{1, 2, 3, 4, 5}
			c := []int{3, 4, 5, 1, 2}
			res := canCompleteCircuitOpt(g, c)
			convey.So(res, convey.ShouldEqual, 3)
		})
		convey.Convey("test-02", func() {
			g := []int{2, 3, 4}
			c := []int{3, 4, 3}
			res := canCompleteCircuitOpt(g, c)
			convey.So(res, convey.ShouldEqual, -1)
		})
	})
}

func TestUnit_canCompleteCircuitOpt02(t *testing.T) {
	convey.Convey("TestUnit_canCompleteCircuitOpt02", t, func() {
		convey.Convey("test-01", func() {
			g := []int{1, 2, 3, 4, 5}
			c := []int{3, 4, 5, 1, 2}
			res := canCompleteCircuitOpt02(g, c)
			convey.So(res, convey.ShouldEqual, 3)
		})
		convey.Convey("test-02", func() {
			g := []int{2, 3, 4}
			c := []int{3, 4, 3}
			res := canCompleteCircuitOpt02(g, c)
			convey.So(res, convey.ShouldEqual, -1)
		})
		convey.Convey("test-03", func() {
			g := []int{4, 5, 2, 6, 5, 3}
			c := []int{3, 2, 7, 3, 2, 9}
			res := canCompleteCircuitOpt02(g, c)
			convey.So(res, convey.ShouldEqual, -1)
		})
	})
}
