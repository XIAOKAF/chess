package serverService

import (
	"chess-room/server/dao"
	"chess-room/server/model"
	"chess-room/service"
	"github.com/go-redis/redis"
)

func InsertRoom(room *model.Room) error {
	err := dao.InsertRoom(room)
	if err != nil {
		return err
	}
	fields := make(map[string]interface{})
	fields["one"] = room.Owner
	fields["one_status"] = "wait"
	return dao.HashSet(room.RoomId, fields)
}

func RecordRoom(room *model.Room) error {
	fields := make(map[string]interface{})
	fields[room.RoomId] = room.Owner
	return dao.HashSet("room", fields)
}

func SelectRoom(room *model.Room) (string, error) {
	return dao.SelectRoom(room)
}

func GetPlayer(player *model.Player, playerId string) (error, bool) {
	_, err := dao.HashGet(player.RoomId, playerId)
	if err == redis.Nil {
		return nil, false
	}
	return err, true
}

func GetStatus(player *model.Player, playerId string) (string, error) {
	status, err := dao.HashGet(player.RoomId, playerId)
	return status, err
}

func Join(player *model.Player) error {
	fields := make(map[string]interface{})
	fields["two"] = player.Mobile
	fields["two_status"] = "wait"
	return dao.HashSet(player.RoomId, fields)
}

func SetStatus(player *model.Player, request *service.UpdateRequest) error {
	fields := make(map[string]interface{})
	fields[player.PlayerId] = request.Status
	return dao.HashSet(player.RoomId, fields)
}
