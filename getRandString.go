// getRandString
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//测试
	rand_string := getRandString(32)
	fmt.Println(rand_string)

	text, _ := getSubString("***oioioio***", "*o", "o*")
	fmt.Println(text)
}

//取指定长度随机字符串
func getRandString(lenght int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//进行切片转换
	bytes := []byte(str)
	//声明空数组
	result := []byte{}
	//
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//取字符串中间内容
func getSubString(str string, left string, right string) (string, bool) {
	//搜索字符串位置
	left_num := strings.Index(str, left)

	if left_num == -1 {
		return "", false
	} else {
		left_num += len(left)
	}

	//string转换切片
	str_bytes := []byte(str)
	str_lenght := len(str_bytes)

	fanw := string(str_bytes[left_num:str_lenght])
	right_num := strings.Index(fanw, right)

	if right_num == -1 {
		return "", false
	}

	return string(str_bytes[left_num:(left_num + right_num)]), true
}
