namespace go thriftAPI
struct UserInfo{
    1:i32 id;
    2:string username;
    3:string password;
    4:string realName;
    5:string mobile;
    6:string email;
}

service UserInfoService{
    list<UserInfo> getUserByName(1:string username)
    list<UserInfo> getUserByNameWait(1:string  username)
}