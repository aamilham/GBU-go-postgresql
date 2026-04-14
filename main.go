package main

import (
	"fmt"
	"log"
	"net/http"
	
	"gbu-go-postgresql/config"
	"gbu-go-postgresql/controllers"
	"github.com/gorilla/mux"
)

func main() {
	//database init
	config.ConnectDB()

	//router mux init
	r := mux.NewRouter()

	//handling routes dengan methodsnya
	//route saya buat sama dengan route php slim sebelumnya untuk kemudahan testing
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	//run server, port saya samakan dengan port php slim sebelumnya untuk kemudahan testing
	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}