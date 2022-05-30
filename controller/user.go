package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var usersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user := common.User{}

	Db.AutoMigrate(&common.User{})

	if err := Db.First(&user, "user_name=?", username).Error; err == nil {
		c.JSON(http.StatusOK,
			common.UserRegisterResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  "失败，账户已经存在了",
				},
				Token:  token,
				UserID: 0})
	} else {
		//创建users数据库
		//Db.AutoMigrate(&User{})
		//创建记录
		user = common.User{Id: 0,
			UserName:      username,
			UserPwd:       password,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Token:         token,
		}
		//向数据库插入数据
		err := Db.Create(&user).Error
		if err != nil {
			fmt.Println("插入失败")
		} else {
			c.JSON(http.StatusOK,
				common.UserRegisterResponse{
					Response: common.Response{
						StatusCode: 0,
						StatusMsg:  "成功",
					},
					Token:  token,
					UserID: 0,
				})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user := &common.User{}

	if err := Db.First(&user, "user_name=?", username).Error; err == nil && user.UserPwd == password {
		c.JSON(http.StatusOK, common.UserLoginResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "登录成功"},
			Token:  token,
			UserID: user.Id})
	} else {
		c.JSON(http.StatusOK, common.UserLoginResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "登录失败"},
			Token:  token,
			UserID: user.Id,
		})
	}

}

//这里的是否关注有点问题，回来再修改
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user := &common.User{}

	//没有报错，表示查询成功，有数据
	if err := Db.First(&user, "token=?", token).Error; err == nil {
		c.JSON(http.StatusOK,
			common.UserInfoResponse{
				Response: common.Response{
					StatusCode: 0,
					StatusMsg:  "查询信息成功",
				},
				User: common.UserFroInfo{
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					ID:            user.Id,
					IsFollow:      user.IsFollow,
					Name:          user.Name,
				},
			})
	} else { //err不为空，说明有错误
		c.JSON(http.StatusOK,
			common.UserInfoResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  "查询信息失败",
				},
				User: common.UserFroInfo{
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					ID:            user.Id,
					IsFollow:      user.IsFollow,
					Name:          user.Name,
				},
			})
	}

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
