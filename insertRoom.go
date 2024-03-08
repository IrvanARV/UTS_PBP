package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	idRoom, _ := strconv.Atoi(r.Form.Get("id_Room"))
	idUser, _ := strconv.Atoi(r.Form.Get("id_User"))
	countQuery := "SELECT COUNT(*) AS player_count FROM participants WHERE id_room = ?"
   	var playerCount int
	err = db.QueryRow(countQuery, idRoom).Scan(&playerCount)
    if err != nil {
        log.Fatal(err)
    }

	var maxPlayer int
    query := "SELECT max_player FROM games WHERE ID = (SELECT id_game FROM rooms WHERE ID = ?)"
    err = db.QueryRow(query, idRoom).Scan(&maxPlayer)
    if err != nil {
        fmt.Println("Failed to join room: Player Max")
    	return
    }

    insertQuery := "INSERT INTO participants (id_room, id_account) VALUES (?, ?)"
    _, err = db.Exec(insertQuery, idRoom, idUser) // Contoh: id_account yang masuk ke dalam room
    if err != nil {
        log.Fatal(err)
	}
    fmt.Println("Successfully joined room.")
	log.Println("Joined the room", err)
}