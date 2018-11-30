// 单元测试文件
//
// 1、 文件名 *_test.go
// 2、 import testing
// 3、 测试函数必须以 Test开头，如： TestXxxx 或 Test_+xxxx
// 4、 测试用例会按照源代码中写的顺序执行
// 5、 测试函数TestXxxx() 的参数是testing.T,可以使用该类型来记录错误或者是测试
// 6、 测试格式：func TestXxx (t *testing.T), Xxx部分可以为任意的字符数字的组合，但是首字母不能是小写字母


package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没有通过...")
	} else {
		t.Log("通过测试赛")
	}
}

func Test_Division_2(t *testing.T)  {
	//t.Error("就是不通过！")
	if _, e := Division(6, 0); e == nil {
		t.Error("Division did not work as expected.")// 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) // 记录一些你期望记录的信息
	}
}























