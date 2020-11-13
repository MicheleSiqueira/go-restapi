package go_restapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var dao = MoviesDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", movierouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/movies", movierouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}