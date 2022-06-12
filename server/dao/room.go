package dao

import "chess-room/server/model"

func InsertRoom(room *model.Room) error {
	result := DB.Select("owner").Create(&room)
	return result.Error
}

func SelectRoom(room *model.Room) (string, error) {
	var id string
	rows := MDB.QueryRow("SELECT room_id FROM room WHERE owner = ?", room.RoomId)
	if rows.Err() != nil {
		return "", rows.Err()
	}
	err := rows.Scan(&id)
	return id, err
}
