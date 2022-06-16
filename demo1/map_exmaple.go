package demo1

import "fmt"

func Test() {
	//var s = map[string][]interface{}
	var s1 map[string][]interface{}

	var s = map[string]string{}
	s["a"] = "abc"
	s["b"] = "nil"
	s["c"] = "null"
	s["d"] = ""
	fmt.Printf("map %s\n", s)
	fmt.Printf("map s1 %s   is nil %s \n", s1, nil == s1)

	for item, index := range s {
		fmt.Printf("map 遍历 --->>%s %s\n", item, index)
	}
}
