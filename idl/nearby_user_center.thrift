struct Meta{
    1: i32 code;    // 业务自定义的返回值, 0为成功, 其他情况为失败(20018为未登录)
    2: string msg;
}

struct User{
    1: i64 user_id;              // 用户id
    2: string nickname;          // 用户名
    3: string head_url;          // 用户头像
}

struct ProfileMeData {
    1: User user;
}

struct ProfileMeResponse{
    1: Meta meta;
    2: ProfileMeData data;
}


struct ProfileMeRequest{
}

struct RegisterVisitorRequest{
    1: string user_id;
    2: string verify_code;
}

struct RegisterVisitorResponse{
    1: Meta meta;
    2: RegisterVisitorData data;
}

struct RegisterVisitorData{
    1: string token;
    2: i64 token_expires_at; // token过期时间戳
}


service UserCenter{
    // 获取用户基本信息
    ProfileMeResponse ProfileMe(1: ProfileMeRequest req)(api.get="/profile/me")

    // 注册游客用户
    RegisterVisitorResponse RegisterVisitor(1: RegisterVisitorRequest req) (api.post="/get_token")
}