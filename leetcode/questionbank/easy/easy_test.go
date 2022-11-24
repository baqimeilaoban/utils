package easy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_twoSum(t *testing.T) {
	convey.Convey("TestUnit_twoSum", t, func() {
		convey.Convey("test-01", func() {
			nums := []int{2, 7, 11, 15}
			target := 9
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{0, 1}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			nums := []int{3, 2, 4}
			target := 6
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{1, 2}), convey.ShouldBeTrue)
		})
		convey.Convey("test-03", func() {
			nums := []int{3, 3}
			target := 6
			res := twoSum(nums, target)
			convey.So(reflect.DeepEqual(res, []int{0, 1}), convey.ShouldBeTrue)
		})
	})
}

func TestUnit_isValidSudoku(t *testing.T) {
	convey.Convey("TestUnit_isValidSudoku", t, func() {
		convey.Convey("test-01", func() {
			board := [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
			res := isValidSudoku(board)
			convey.So(res, convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			board := [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
			res := isValidSudoku(board)
			convey.So(res, convey.ShouldBeFalse)
		})
	})
}

func TestUnit_rotateImage(t *testing.T) {
	convey.Convey("TestUnit_rotateImage", t, func() {
		convey.Convey("test-01", func() {
			matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
			rotateImage(matrix)
			convey.So(reflect.DeepEqual(matrix, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}), convey.ShouldBeTrue)
		})
	})
}
