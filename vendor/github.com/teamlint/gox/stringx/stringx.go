package stringx

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Text 文本方法,如枚举文本,只要实现Text()方法即可
func Text(v interface{}) string {
	val := reflect.ValueOf(v)
	t := val.Type()
	nums := val.NumMethod()
	// Text 方法
	for i := 0; i < nums; i++ {
		method := t.Method(i).Name
		if method == "Text" {
			result := val.Method(i).Call(nil)
			return result[0].Interface().(string)
		}
	}
	return fmt.Sprintf("%v", val)
}

// Replace 字符串替换
func Replace(origin string, search string, replace string) string {
	return strings.Replace(origin, search, replace, -1)
}

// ReplaceByMap 使用map进行字符串替换
func ReplaceByMap(origin string, replaces map[string]string) string {
	result := origin
	for k, v := range replaces {
		result = strings.Replace(result, k, v, -1)
	}
	return result
}

// SearchArray 数组查找字符串索引位置，如果不存在则返回-1，使用完整遍历查找
func SearchArray(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return -1
}

// IsInArray 判断字符串是否在数组中
func IsInArray(a []string, s string) bool {
	return SearchArray(a, s) != -1
}

// IsLowerChar 判断给定字符是否小写
func IsLowerChar(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsUpperChar 判断给定字符是否大写
func IsUpperChar(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsNumeric 判断锁给字符串是否为数字
func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < byte('0') || s[i] > byte('9') {
			return false
		}
	}
	return true
}

// Left 返回左起n个字符,如果指定第三个参数,则使用其值作为省略符号,否则为"…"
func Left(s string, n int, ellipsis ...string) string {
	if n < 0 {
		n = 0
	}
	if len(s) <= n {
		return s
	}

	var (
		chari int
		bytei int
	)
	for bytei = range s {
		if chari >= n {
			break
		}
		chari++
	}
	if len(ellipsis) > 0 {
		return s[:bytei] + ellipsis[0]
	}

	return s[:bytei] + "…"
}

// UpperFirst 字符串首字母转换为大写
func UpperFirst(s string) string {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}
	for _, c := range s {
		sc := string(c)
		return strings.ToUpper(sc) + s[len(sc):]
	}
	return ""
}

// LowerFirst 字符串首字母转换为小写
func LowerFirst(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	for _, c := range s {
		sc := string(c)
		return strings.ToLower(sc) + s[len(sc):]
	}
	return ""
}

var reUnprintable = regexp.MustCompile("[\x00-\x1F\u200e\u200f]")

// RemoveUnprintable 删除字符串中不可打印字符 (0 to 31 ASCII)
func RemoveUnprintable(s string) string {
	return reUnprintable.ReplaceAllString(s, "")
}

// GetLine 获取指定行字符串
// gets the nth line \n-denoted line from a string.
func GetLine(in string, n int) string {
	// Would probably be faster to use []byte and find the Nth \n character, but
	// this is "fast enough"™ for now.
	arr := strings.SplitN(in, "\n", n+1)
	if len(arr) <= n-1 {
		return ""
	}
	return arr[n-1]
}
