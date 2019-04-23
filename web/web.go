package main

import (
	"crypto/md5"
	"fmt"
	"go_try/web/session"
	_ "go_try/web/session/providers/memory"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var gloalsessions *session.Manager

func init()  {
	gloalsessions,_ = session.NewManger("memory", "gosessionid", 60)
	go gloalsessions.GC()
}



func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	http.Redirect(w, r, "/login", http.StatusFound)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cur_time := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cur_time, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("login.html")
		if err != nil {
			log.Println(err)
			return
		}
		//w.Header().Set("Content-Type", "text/html")
		log.Println(t.Execute(w, token))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			fmt.Fprintln(w,"error!")
			return
		} else {
			//todo 验证token的合法性、验证用户名密码。验证通过后：
			if false {
				sess := gloalsessions.SessionStart(w, r)
				sess.Set("username", r.Form["username"])
				http.Redirect(w, r, "/upload", http.StatusFound)
			}
		}
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form.Get("password"))
		http.Redirect(w, r, "/upload", http.StatusFound)     //目前没有做登陆校验，所以只要点击登陆按钮就会重定向到"/upload"
	}
}

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		cur_time := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cur_time, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("upload.html")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(t.Execute(w, token))
	} else if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {

	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func count(w http.ResponseWriter, r *http.Request)  {
	sess := gloalsessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	}else if (createtime.(int64)+60) < (time.Now().Unix()) {
		gloalsessions.SessionDestory(w, r)
		sess = gloalsessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	}else {
		sess.Set("countnum", ct.(int) + 1)
	}
	t, _ := template.ParseFiles("count.html")
	w.Header().Set("content-type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}