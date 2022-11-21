package easy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_maxProfit(t *testing.T) {
	convey.Convey("TestUnit_maxProfit", t, func() {
		convey.Convey("test-01", func() {
			prices := []int{7, 1, 5, 3, 6, 4}
			convey.So(maxProfit(prices), convey.ShouldEqual, 7)
		})
		convey.Convey("test-02", func() {
			prices := []int{1, 2, 3, 4, 5}
			convey.So(maxProfit(prices), convey.ShouldEqual, 4)
		})
		convey.Convey("test-03", func() {
			prices := []int{7, 6, 4, 3, 1}
			convey.So(maxProfit(prices), convey.ShouldEqual, 0)
		})
	})
}
