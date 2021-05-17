package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", byte)
	fmt.Println("测试pull")
}
