package main

import ( 
	
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("../FrontEnd/src/index.js")))

  	http.ListenAndServe(":3000", r)



  	r.Handle("/login", LoginHandler).method("POST")
}


var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("LoginSuccess"))

})