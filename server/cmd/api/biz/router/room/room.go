// Code generated by hertz generator. DO NOT EDIT.

package room

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	room "tank_war/server/cmd/api/biz/handler/room"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/room", append(_getroomlistMw(), room.GetRoomList)...)
	root.POST("/room", append(_createroomMw(), room.CreateRoom)...)
	_room := root.Group("/room", _roomMw()...)
	_room.POST("/join", append(_joinroomMw(), room.JoinRoom)...)
}
