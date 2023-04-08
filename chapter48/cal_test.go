package chapter48

import (
	"testing"
)

// 单元测试
func TestAddNum(t *testing.T) {
	res := addNum(10, 2)
	if res != 7 {
		t.Fatalf("addNum(10, 3)执行错误。期望值=%v, 实际值=%v", 7, res)
	} else {
		t.Logf("addNum(10, 3)执行正确。")
	}
}
