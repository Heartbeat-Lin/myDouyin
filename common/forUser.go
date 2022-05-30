package common

type User struct {
	Id            int64  `json:"id,omitempty"`
	UserName      string `json:"user_name"`
	UserPwd       string `json:"user_pwd"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Token         string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User UserFroInfo `json:"user"` // 用户信息
}

//用户信息专属user
type UserFroInfo struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type UserLoginResponse struct {
	//StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	//StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

type UserRegisterResponse struct {
	Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}
