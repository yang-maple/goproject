package main

import (
	"html/template"
	"net/http"
	"strconv"
	"web/modity"
)

type Task struct {
	ID     int
	Name   string
	Status string
}

func main() {

	addr := "0.0.0.0:9999"
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		tpl := template.Must(template.ParseFiles("webhtml/list.html"))
		tpl.ExecuteTemplate(w, "list.html", modity.Findtask())

	})

	http.HandleFunc("/modity", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			r.ParseForm()
			fid, _ := strconv.Atoi(r.FormValue("fid"))
			task1 := Task{fid, r.FormValue("fname"), r.FormValue("fstatus")}
			switch r.FormValue("fsubmit") {
			case "提交":
				modity.Addtask(modity.Task(task1))
			case "删除":
				modity.Deltask(modity.Task(task1))
			case "更新":
				modity.Updatatask(modity.Task(task1))

			}

			//	http.Redirect(w, r, "/login", http.StatusFound)
		}
		tpl := template.Must(template.ParseFiles("webhtml/modity.html"))
		tpl.ExecuteTemplate(w, "modity.html", nil)

	})

	http.ListenAndServe(addr, nil)

}
