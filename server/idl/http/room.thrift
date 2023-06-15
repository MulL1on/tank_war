namespace go room

include "../base/common.thrift"
include "../base/room.thrift"

struct CreateRoomReq {
    1: required string name (api.raw="name" api.vd="len($)>0 && lne($)<33")
    2: required i32 max_player (api.raw="max_player" )
}

struct CreateRoomResp {
    1: required common.BaseResponse base_resp
}

struct JoinRoomReq {
    1: required string name (api.raw="name" api.vd="len($)>0 && lne($)<33")
}
struct JoinRoomResp {
    1: required i32 room_id
    2: required string address
    3: required i32 port
}


struct GetRoomListReq {
}

struct GetRoomListResp {
    1: required list<room.Room> rooms
}

service RoomService{
     CreateRoomResp CreateRoom(1: CreateRoomReq req) (api.POST="/room")
     JoinRoomResp JoinRoom(1: JoinRoomReq req) (api.POST="/room/join")
    GetRoomListResp GetRoomList(1: GetRoomListReq req) (api.GET="/room")
}



