package main

import (
	"fmt"
	"net/http"
)



func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "this is the about page")

}

func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}


