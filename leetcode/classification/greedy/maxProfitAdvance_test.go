package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUint_maxProfitAd(t *testing.T) {
	convey.Convey("TestUint_maxProfitAd", t, func() {
		convey.Convey("test-01", func() {
			prices := []int{1, 3, 2, 8, 4, 9}
			res := maxProfitAd(prices, 2)
			convey.So(res, convey.ShouldEqual, 8)
		})
		convey.Convey("test-03", func() {
			prices := []int{1, 3, 7, 5, 10, 3}
			res := maxProfitAd(prices, 3)
			convey.So(res, convey.ShouldEqual, 6)
		})
	})
}
