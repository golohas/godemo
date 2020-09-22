package str_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	var str1 = "你好"
	var str2 = "golang"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	t.Log(str3)
}

func TestStrCon(t *testing.T) {
	str1 := "test1"
	str2 := "golang"
	t.Log(str1 + str2)
}

func TestStrSplit(t *testing.T) {
	var s = "123-456-789"
	a := strings.Split(s, "-")
	t.Log(a)
	str2 := strings.Join(a, "*")
	t.Log(str2)
}

func TestJoin(t *testing.T) {
	arr := []string{"java", "php", "golang"}
	str3 := strings.Join(arr, "-")
	t.Log(str3)
}
