package models

type Accounts struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type Rooms struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	Id_game   int    `json:"id_game"`
}

type RoomDetails struct {
	ID       int     `json:"id"`
	RoomName string  `json:"room_name"`
	Rooms    []Rooms `json:"data"`
}

type RoomDetail struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type DetailRoom struct {
	ID       int      `json:"id"`
	User     Accounts     `json:"user"`
	Quantity int      `json:"quantity"`
}

type DetailRoomsResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []DetailRoom `json:"data"`
}

type Participant struct {
	ID         int `json:"id"`
	Id_room    int `json:"id_room"`
	Id_account int `json:"id_account"`
}


