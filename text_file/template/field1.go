package main

import "html/template"
import "os"

type Friend struct {
	Fname string
}

type person struct {
	UserName	string
	Emails		[]string
	Friends		[]*Friend
}

func main()  {
	f1 := Friend{"f1"}
	f2 := Friend{"f2"}

	t := template.New("haha")
	t, _ = t.Parse(`hello {{.UserName}}
	{{range .Emails}}
	an email {{. | html}}
	{{end}}
	{{range .Friends}}
		my friend name is {{.Fname}}
	{{end}}`)

	p := person{UserName:"ly",
				Emails:[]string{"<145>", "136@qq.com"},
				Friends:[]*Friend{&f1, &f2}}

	t.Execute(os.Stdout, p)
}
