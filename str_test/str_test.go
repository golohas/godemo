package str_test

import (
	"fmt"
	"strconv"
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

func TestStrIndex(t *testing.T) {
	str1 := "this is str"
	str2 := "is"
	num1 := strings.Index(str1, str2)
	num2 := strings.LastIndex(str1, str2)
	t.Log(num1, num2)
	print(num1, num2)
}

func TestStrplus(t *testing.T) {
	var s string = " "
	s += "foo"
	t.Log("s is:", s)
}
func TestStrconv(t *testing.T) {
	str := "10"
	num, _ := strconv.ParseInt(str, 10, 64)
	t.Logf("num: %v", num)
	str2 := "3.141592654"
	num2, _ := strconv.ParseFloat(str2, 10)
	t.Logf("num2: %v, num2 type: %T", num2, num2)
}
