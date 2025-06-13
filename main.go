package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pagedata struct {
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "unable to load page", http.StatusInternalServerError)
		return
	}
	data := Pagedata{Result: ""}

	if r.Method == http.MethodGet && r.URL.Query().Has("result") {
		data.Result = r.URL.Query().Get("result")
	}
	tmpl.Execute(w, data)

}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	num1 := r.FormValue("fno")
	num2 := r.FormValue("sno")
	op := r.FormValue("operation")
	var result string

	var n1, n2 float64

	fmt.Sscanf(num1, "%f", &n1)
	fmt.Sscanf(num2, "%f", &n2)
	switch op {
	case "add":
		result = fmt.Sprintf("%.2f", n1+n2)
	case "sub":
		result = fmt.Sprintf("%.2f", n1-n2)
	case "mul":
		result = fmt.Sprintf("%.2f", n1*n2)
	case "div":
		if n2 == 0 {
			result = "Cannot divide by zero"
		} else {
			result = fmt.Sprintf("%.2f", n1/n2)
		}
	default:
		result = "invalid operation"

	}
	http.Redirect(w, r, "/?result="+result, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
