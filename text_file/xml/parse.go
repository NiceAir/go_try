package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	/*
	字段叫做MLName，且类型为xml.Name字段，
	那么在解析的时候就会保存这个element的名字到该字段,
	 */
	XMLName		xml.Name 	`xml:"servers"`

	/*
	tag定义了中含有",attr"，
	那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段
	 */
	Version		string		`xml:"version,attr"`
	Svs 		[]server 	`xml:"server"`

	/*
	string或者[]byte类型且它的tag含有",innerxml"，
	Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上
	 */
	Description string 		`xml:",innerxml"`
}

type server struct {
	XMLName 	xml.Name 	`xml:"server"`

	/*
	tag定义中含有XML结构中element的名称，
	那么解析的时候就会把相应的element值赋值给该字段
	 */
	ServerName	string 		`xml:"serverName"`
	ServerIP 	string		`xml:"serverIP"`
}

func main()  {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)  //解析xml
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}