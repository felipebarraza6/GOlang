//web server

package main

import (
"fmt"
"log"
"net/http"
)

func main(){
	http.HandleFunc("/", handler)//cada peticion manda a llamar a un handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//handler echos the Path componente of the requeest URL r.
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
