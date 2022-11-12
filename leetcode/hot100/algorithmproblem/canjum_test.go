package algorithmproblem

import (
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAlgorithmProblem_CanJump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	Convey("TestAlgorithmProblem_CanJump", t, func() {
		Convey("true", func() {
			nums := []int{2, 3, 1, 1, 4}
			res := CanJump(nums)
			So(res, ShouldBeTrue)
		})
		Convey("false", func() {
			nums := []int{3, 2, 1, 0, 4}
			res := CanJump(nums)
			So(res, ShouldBeFalse)
		})
	})
}
