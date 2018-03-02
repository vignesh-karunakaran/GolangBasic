package main

import ("fmt"
		"net/http"
		"html/template")

type vicky struct{
	Title string
	Sub string
}


func index(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello cruel World...!</h1>")
}


func vickyset(w http.ResponseWriter,r *http.Request) {
  k :=vicky{Title: "summoning to my demons", Sub: "work ...Hard!"}
  t,_ :=template.ParseFiles("source.html")
  t.Execute(w, k)

	
}



func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/hell",vickyset)
	http.ListenAndServe(":8000", nil)

}