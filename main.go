package main

import (
	"uts_week6/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/rooms/1", controllers.GetDetailRooms).Methods("GET")
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/rooms", controllers.LeaveRoom).Methods("DELETE")

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
