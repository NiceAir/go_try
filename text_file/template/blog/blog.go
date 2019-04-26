package main

import (
	"html/template"
	"net/http"
)

func show(w http.ResponseWriter, r *http.Request)  {
	s1, _ := template.ParseFiles("file/index.html",
		"file/header.html",
		"file/content.html",
		"file/footer.html")
	s1.Execute(w, s1)
}

func main()  {
	http.HandleFunc("/", show)
	http.ListenAndServe(":8080", nil)
}
