package greedy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUint_monotoneIncreasingDigits(t *testing.T) {
	convey.Convey("TestUint_monotoneIncreasingDigits", t, func() {
		convey.Convey("test-01", func() {
			res := monotoneIncreasingDigits(10)
			convey.So(res, convey.ShouldEqual, 9)
		})
		convey.Convey("test-02", func() {
			res := monotoneIncreasingDigits(1234)
			convey.So(res, convey.ShouldEqual, 1234)
		})
		convey.Convey("test-03", func() {
			res := monotoneIncreasingDigits(332)
			convey.So(res, convey.ShouldEqual, 299)
		})
	})
}
