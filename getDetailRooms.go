package controllers

import (
	m "uts_week6/models"
	"net/http"
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, s string) {
	panic("unimplemented")
}

func GetDetailRooms (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.id, r.room_name, p.id, a.id, a.usename,
			FROM rooms r JOIN participant p ON r.ID = p.ID_Room
			JOIN accounts a ON p.ID_Account = a.ID`
	detailRoomsRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w, "Invalid query")
		return
	}

	var detailRoom m.Rooms
	var detailRooms []m.RoomDetails
	for detailRoomsRow.Next() {
		if err := detailRoomsRow.Scan(&detailRoom.ID, &detailRoom.Id_game, &detailRoom.Room_name); 
		err != nil {
			print(err.Error())
			sendErrorResponse(w, "Failed to scan data")
			return
		} else {
			detailRooms = append(detailRooms, detailRoom)
		}
	}
	var response m.DetailRoomsResponse
	response.Status = 200
	response.Message = "Success!"
	response.Data = detailRooms
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}	