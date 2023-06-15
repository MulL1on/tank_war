namespace go user

include "../base/common.thrift"
include "../base/user.thrift"


struct RegisterReq{
  1: string username
  2: string password
}

struct RegisterResp{
    1:common.BaseResponse base_resp
}

struct LoginReq{
  1: string username
    2: string password
}

struct LoginResp{
    1:common.BaseResponse base_resp
    2:string token
}

struct GetUserInfoReq{
    1: i64 user_id
}
struct GetUserInfoResp{
    1:common.BaseResponse base_resp
    2:user.User user
}


service UserService{
    RegisterResp Register(1:RegisterReq req)
    LoginResp Login(1:LoginReq req)
    GetUserInfoResp GetUserInfo(1:GetUserInfoReq req)
}