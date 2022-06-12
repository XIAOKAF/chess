package model

type Room struct {
	RoomId string
	Owner  string
}

type Player struct {
	Mobile   string
	RoomId   string
	PlayerId string
}
