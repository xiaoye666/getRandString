// getRandString
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World!")
}

func getRandString(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//进行切片处理，变成数组
	bytes := []byte(str)
	//声明空数组
	result := []byte{}
	//
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
