package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波拉契数列
//1,1,2,3,5,8,13
//  a,b
//    a,b
func fibonacii() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

//为函数实现接口
func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	//TODO: incorrect if p is too small!!
	return strings.NewReader(s).Read(p)
}

//打印文件
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacii()
	printFileContents(f)
}
