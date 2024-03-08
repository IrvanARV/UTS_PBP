package controllers

import (
	m "uts_week6/models"
	"encoding/json"
	"log"
	"net/http"

)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ERROR : ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []m.Rooms

	for rows.Next() {
		var room m.Rooms
		err := rows.Scan(&room.ID, &room.Room_name, &room.Id_game)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "ERROR : ", http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}