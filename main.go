package main

import (
	"encoding/json"
	"gocrud/config"

	alamatController "gocrud/controller/alamatcontroller"
	userController "gocrud/controller/usercontroller"
	"gocrud/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
config.LoadConfig()
config.ConnectDB()
	r := mux.NewRouter()
	//user
	r.HandleFunc("/users", userController.Index).Methods("GET")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")
	r.HandleFunc("/users", userController.Create).Methods("POST")
	r.HandleFunc("/users/{id}/detail", userController.Detail).Methods("GET")
	r.HandleFunc("/users/{id}/delete", userController.Delete).Methods("DELETE")
	r.HandleFunc("/users/{id}/update", userController.Update).Methods("PUT")

	//alamat
	r.HandleFunc("/alamats", alamatController.Index).Methods("GET")
	r.HandleFunc("/alamats", alamatController.Create).Methods("POST")
	r.HandleFunc("/alamats/{id}/detail", alamatController.Detail).Methods("GET")
	r.HandleFunc("/alamats/{id}/delete", alamatController.Delete).Methods("DELETE")
	r.HandleFunc("/alamats/{id}/update", alamatController.Update).Methods("PUT")



	r.Use(middleware.LoggingMiddleware)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}





