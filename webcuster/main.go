package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"webcuster/webmodity"
)

func main() {
	// 服务器监听地址
	addr := "0.0.0.0:9999"
	// 打开保存数据的文件 如果没有就创建
	file, _ := os.OpenFile("/root/go/goproject/webcuster/custer.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	// 延迟关闭文件
	defer file.Close()
	// 初始化服务的同时读取数据内容
	webmodity.Read(*file)

	// 监听的用户管理系统主页面
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// 实现了 保存功能
		if r.FormValue("fsubmit") == "保存用户信息" {
			webmodity.Save(*file)
		}
		tpl := template.Must(template.ParseFiles("webhtml/menu.html"))
		tpl.ExecuteTemplate(w, "menu.html", nil)

	})
	// 监听的查询用户信息的页面
	http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("webhtml/find.html"))
		tpl.ExecuteTemplate(w, "find.html", webmodity.Findtask())

	})
	// 监听的新增用户信息的页面
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			// 获取请求的信息
			r.ParseForm()
			// string 转为 int
			fage, _ := strconv.Atoi(r.FormValue("fage"))
			fphone, _ := strconv.Atoi(r.FormValue("fphone"))
			// 请求的数据写入 task1 中
			task1 := webmodity.Person{Name: r.FormValue("fname"), Age: fage, Phone: fphone, Emil: r.FormValue("femil"), Addr: webmodity.Addr{City: r.FormValue("fcity")}}
			// 用户返回和提交新增的数据
			switch r.FormValue("fsubmit") {
			case "提交":
				// webmodity.Addtask 函数处理 task1 中的数据
				webmodity.Addtask(task1)
			case "返回":
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
		tpl := template.Must(template.ParseFiles("webhtml/add.html"))
		tpl.ExecuteTemplate(w, "add.html", nil)

	})
	// 监听的删除用户信息的页面
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			r.ParseForm()
			// 用户返回和提交删除的数据
			switch r.FormValue("fsubmit") {
			case "提交":
				// webmodity.Deltask 函数处理 task1 中的数据
				webmodity.Deltask(r.FormValue("fname"))
			case "返回":
				http.Redirect(w, r, "/login", http.StatusFound)
			}

		}
		tpl := template.Must(template.ParseFiles("webhtml/delete.html"))
		tpl.ExecuteTemplate(w, "delete.html", nil)

	})
	// 监听的更新用户信息的页面
	http.HandleFunc("/updata", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			r.ParseForm()
			fage, _ := strconv.Atoi(r.FormValue("fage"))
			fphone, _ := strconv.Atoi(r.FormValue("fphone"))
			task1 := webmodity.Person{Name: r.FormValue("fname"), Age: fage, Phone: fphone, Emil: r.FormValue("femil"), Addr: webmodity.Addr{City: r.FormValue("fcity")}}
			// 用户返回和提交更新的数据
			switch r.FormValue("fsubmit") {
			case "提交":
				webmodity.Updatatask(task1)
			case "返回":
				http.Redirect(w, r, "/login", http.StatusFound)
			}

		}
		tpl := template.Must(template.ParseFiles("webhtml/updata.html"))
		tpl.ExecuteTemplate(w, "updata.html", nil)

	})
	// 打开服务监听
	http.ListenAndServe(addr, nil)

}
