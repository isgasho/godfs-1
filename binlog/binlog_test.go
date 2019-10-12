package binlog_test

import (
	"fmt"
	"github.com/hetianyi/godfs/util"
	"github.com/hetianyi/gox"
	"github.com/hetianyi/gox/convert"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var a = gox.GetTimestamp(time.Now())
	buff := make([]byte, 8)
	convert.Length2Bytes(a, buff)
	fmt.Println(buff)

	b := 2 << 20
	convert.Length2Bytes(int64(b), buff)
	fmt.Println(buff)
}

func TestFixZeros(t *testing.T) {
	fmt.Println(util.FixZeros(0, 5))
	fmt.Println(util.FixZeros(1, 5))
	fmt.Println(util.FixZeros(11, 5))
	fmt.Println(util.FixZeros(111, 5))
	fmt.Println(util.FixZeros(1111, 5))
	fmt.Println(util.FixZeros(11111, 5))
	fmt.Println(util.FixZeros(0, 3))
	fmt.Println(util.FixZeros(1, 3))
	fmt.Println(util.FixZeros(11, 3))
	fmt.Println(util.FixZeros(111, 3))
}
