package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	idRoom := r.Form.Get("id_room")
	idUser := r.Form.Get("id_user")

	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	deleteQuery := "DELETE FROM participants WHERE id_room = ? AND id_account = ?"
	result, err := db.Exec(deleteQuery, idRoom, id)
	if err != nil {
		http.Error(w, "Failed to delete participant", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("User tidak ditemukan")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User tidak ditemukan"))
		return
	}

	fmt.Println("User keluar")
	log.Println("User Keluar:", err)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User keluar"))
}