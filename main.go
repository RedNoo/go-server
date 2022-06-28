package main

import (
	"fmt"
	"log"
	"net/http"
)


func formmHandler(w http.ResponseWriter, r *http.Request){
	 if err := r.ParseForm() ; err != nil {
		 fmt.Fprintf(w, "Parseform() err : %v", err)
		 return
	 }

	 fmt.Fprintf(w,"Form post successfull \n")
	 name := r.FormValue("name")
	 fmt.Fprintf(w,"name: %s\n", name)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"GET mothod not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello"	)

}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formmHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server starting at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}