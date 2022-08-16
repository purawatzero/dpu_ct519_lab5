package main

import (
    "net/http"
)

// func fncHello(p string){
// 	http.ServeFile(w ResponseWriter, r *Request, p)
// }

func main() {
	// http.HandleFunc("/", fncHello("./www/index.html") )
	http.Handle("/", http.FileServer(http.Dir("./www")))
    http.ListenAndServe(":8080", nil)
}