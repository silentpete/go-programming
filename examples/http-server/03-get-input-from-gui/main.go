package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("cmd.tmpl")
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		r.ParseForm() // need to parse the form before interacting with form data
		cmd := r.Form["text"]
		switch cmd[0] {
		case "pg --blog":
			http.Redirect(w, r, "http://blog.pg-h.io/", http.StatusFound)
		case "pg --blog-metrics":
			http.Redirect(w, r, "http://blog.pg-h.io/metrics", http.StatusFound)
		case "pg --github":
			http.Redirect(w, r, "https://github.com/silentpete", http.StatusFound)
		case "pg --grafana":
			http.Redirect(w, r, "http://grafana.pg-h.io/", http.StatusFound)
		case "pg --grafana-metrics":
			http.Redirect(w, r, "http://grafana.pg-h.io/metrics", http.StatusFound)
		case "pg --help":
			http.Redirect(w, r, "/", http.StatusFound)
		case "pg -h":
			http.Redirect(w, r, "/", http.StatusFound)
		case "pg --prometheus":
			http.Redirect(w, r, "http://prometheus.pg-h.io/", http.StatusFound)
		case "pg --prometheus-metrics":
			http.Redirect(w, r, "http://prometheus.pg-h.io/metrics", http.StatusFound)
		case "pg --resume":
			http.Redirect(w, r, "https://www.linkedin.com/in/petegallerani/", http.StatusFound)
		case "durp":
			fmt.Fprintf(w, "durp")
		default:
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func main() {
	http.HandleFunc("/", rootHandler)        // setting router rule
	err := http.ListenAndServe(":9090", nil) // starting http server with listening port 9090
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
