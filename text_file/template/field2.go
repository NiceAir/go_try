package main

import (
	"html/template"
	"os"
)

func main()  {
	tEmpty := template.New(" template test")
	tEmpty, _ = tEmpty.Parse("空 pipeline if demo: {{if `1`}} 不会输出. {{end}}\n")
	tEmpty.Execute(os.Stdout, tEmpty)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	i := 1
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `{{.i}}`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, i)
}

