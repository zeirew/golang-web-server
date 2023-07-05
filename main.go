package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request) { //star is pointer 
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter,r *http.Request) { //star is pointer 
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static")) //:= declares and defines 
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n") //for newline, use printf
	if err := http.ListenAndServe(":8080", nil); err!=nil { // listenandserve creates server
		log.Fatal(err)
	}
}
