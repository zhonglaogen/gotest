package main

import (
	"testing"
	"zlx.com/laogen/gotest/basic/basictest"
)

/**
表格驱动测试，一个位置错了，后面的还可以继续测试，并且得到错误出现的位置
命令的方式： cd到测试文件的包  cgo test .
 */

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3,4,5},
		{3,3,5},
		{5,12,13},
	}

	for _, tt := range tests{
		if actual := basictest.CalculateTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calculate(%d,%d); " +
				"got %d; expected %d",
					tt.a, tt.b, actual, tt.c)
		}
	}

}

func BenchmarkSubstr(b *testing.B)  {
//	性能测试

b.Logf("log")
//从这里开始计算时间
b.ResetTimer()
	 s1, s2, c := 3,4,5
	// 系统自动算法算出次数
	for  i := 0; i < b.N ; i++  {
		if actual := basictest.CalculateTriangle(s1, s2); actual != c {
			b.Errorf("calculate(%d,%d); " +
				"got %d; expected %d",
				s1, s2, actual, c)
		}
	}




}
