package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {

	resp, err := http.Get("http://www.google.co.uk")
	fmt.Println("Given Error is: ", err)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Body Read error is: ", err)
	fmt.Println("Body type - ", reflect.TypeOf(body))

	//Parse html
	prs, err := html.Parse(resp.Body)
	fmt.Println("Body type after Parse - ", reflect.TypeOf(prs))

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Println(n.Data)
		}
		fmt.Println(n.Data)
	}
	f(prs)
	fmt.Println(prs.Data)

}

