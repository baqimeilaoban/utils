package greedy

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_partitionLabels(t *testing.T) {
	convey.Convey("TestUnit_partitionLabels", t, func() {
		convey.Convey("test-01", func() {
			s := "ababcbacadefegdehijhklij"
			res := partitionLabels(s)
			convey.So(reflect.DeepEqual(res, []int{9, 7, 8}), convey.ShouldBeTrue)
		})
		convey.Convey("test-02", func() {
			s := "eccbbbbdec"
			res := partitionLabels(s)
			convey.So(reflect.DeepEqual(res, []int{10}), convey.ShouldBeTrue)
		})
	})
}
