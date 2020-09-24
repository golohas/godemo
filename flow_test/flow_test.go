package flow_test

import (
	"testing"
)

func TestIf(t *testing.T) {
	var num = 10
	if num == 10 {
		t.Log("hello 10")
	} else if num > 10 {
		t.Log("hello>10")
	} else {
		t.Log("hello<10")
	}
}

func TestFor(t *testing.T) {
	var str = "你好golang"
	for key, value := range str {
		t.Logf("%v-%c", key, value)
	}
}

func TestFor2(t *testing.T) {
	var arr = []string{"php", "java", "node", "golang"}
	for index, value := range arr {
		t.Logf("%v %s", index, value)
	}
}

func TestSwitch(t *testing.T) {
	extname := ".a"
	switch extname {
	case ".html":
		{
			t.Log(".html")
			break
		}
	case ".doc":
		{
			t.Log(".doc")
		}
	case ".js":
		{
			t.Log(".js")
		}
	default:
		{
			t.Log("other suffix")
		}
	}
}

func TestSwitchFallthrough(t *testing.T) {
	extname := ".txt"
	switch extname {
	case ".html":
		t.Log(".html")
		fallthrough
	case ".txt", ".doc":
		t.Log("传递来的是文档")
		fallthrough
	case ".js":
		t.Log(".js")
		fallthrough
	default:
		t.Log("其它后缀")
	}
}
