package main

import (
	"fmt"
	"learngo/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

//?:something that can get
type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
