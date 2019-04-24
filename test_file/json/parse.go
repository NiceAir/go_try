package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})   //断言
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv{
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is a type of I dont't know how to handle")
		}
	}
}