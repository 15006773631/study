package main

import (
	"io"
	"net/http"
)

const form = `<html><body><form action= "#" method="post" name="bar">
                 <input typre="text" name="in"/>
                 <input typre="text" name="in"/>
                 <input type="submit"value="submit"/>
                 </form></body></html>`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello,world")
}
func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))

	}
}
func main() {
	http.HandleFunc("test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8888", nil); err != nil {

	}
}
