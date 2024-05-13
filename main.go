package main

import (
	"log"
	"net/http"

	controllers "github.com/dikalasenjadatang/web_app/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Static File Serve
	r.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Controllers
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/contact", controllers.Contact)

	//Create a New User
	r.HandleFunc("/register", controllers.Register)
	r.HandleFunc("/save", controllers.Save).Methods("POST")
	r.HandleFunc("/show/{id}", controllers.Show)
	//Update User
	r.HandleFunc("/edit/{id}", controllers.Edit)
	r.HandleFunc("/update", controllers.Update).Methods("POST")
	//Delete User
	r.HandleFunc("/delete/{id}", controllers.Delete)

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
