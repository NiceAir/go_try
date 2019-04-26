package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend3 struct {
	Fname string
}

type Person3 struct {
	UserName string
	Emails []string
	Friends	[]*Friend3
}

func EmailDealwith(args ... interface{}) string {
	ok := false
	var s string
	if len(args) == 1{
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	subtrs := strings.Split(s, "@")
	if len(subtrs) != 2 {
		return s
	}
	return (subtrs[0] + "at" + subtrs[1])
}

func main()  {
	f1 := Friend3{Fname: "minux.ma"}
	f2 := Friend3{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal":EmailDealwith})
	t, _ = t.Parse(`hello {{.UserName}} !
								{{range .Emails}}
									an email {{.|emailDeal}}
								{{end}}
								{{with .Friends}}
								{{range .}}
									my friend name is {{.Fname}}
								{{end}}
								{{end}}
								`)
	p := Person3{UserName: "ly",
				Emails:		[]string{"@@", "haha1@qq.com", "<script>4321</script>"},
				Friends:	[]*Friend3{&f1,&f2}}
	t.Execute(os.Stdout, p)

}
