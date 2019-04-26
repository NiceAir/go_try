package main

import (
	"html/template"
	"os"
)

func main()  {
	t := template.New("t")
	t, _ = t.Parse(`{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
`)
	t.Execute(os.Stdout, nil)

	t1 := template.New("t")
	t, _ = t1.Parse(`{{with $x := "output"}}{{printf "%q" $x}}{{end}}
`)
	t.Execute(os.Stdout, nil)

	t2 := template.New("t")
	t, _ = t2.Parse(`{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
`)
	t.Execute(os.Stdout, nil)
}
