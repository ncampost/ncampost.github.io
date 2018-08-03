package main

import (
	"log"
	"net/http"

	"./routing"
	"./pages/view"
	"./pages/edit"
	"./pages/save"
	"./pages/index"
)


func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/view/", routing.MakeHandler(viewendpoints.ViewHandler))
	http.HandleFunc("/edit/", routing.MakeHandler(editendpoints.EditHandler))
	http.HandleFunc("/save/", routing.MakeHandler(saveendpoints.SaveHandler))
	http.HandleFunc("/", routing.MakeHandler(indexendpoints.Handler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
