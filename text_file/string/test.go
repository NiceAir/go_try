package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}

func main()  {
	fmt.Printf("[%q]\n", strings.Trim("   !!! Achtung !!! ", "! "))
	fmt.Printf("Fields are: %q\n", strings.Fields("       foo bar  baz   "))

	str := make([]byte, 0, 100)
	fmt.Println(string(str))
	str = strconv.AppendInt(str, 4567, 10)  //base: 进制
	fmt.Println(string(str))
	str = strconv.AppendBool(str, false)
	fmt.Println(string(str))
	str = strconv.AppendQuote(str, "abcdefg")
	fmt.Println(string(str))
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)


	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)

}
