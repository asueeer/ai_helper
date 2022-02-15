package middleware

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/domain/val_obj"
	"ai_helper/biz/model"
	"github.com/spf13/cast"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var notLogin = model.Response{
	Meta: model.Meta{
		Code: common.LoginFailErrCode,
		Msg:  "user not login",
	},
}

// Auth 校验用户是否登陆
func Auth(jwtClient *JwtClient, noAuth map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if noAuth[c.FullPath()] {
			c.Next()
			return
		}
		isLogin := isLoginAndDo(c, jwtClient, func(token *jwt.Token, claims *val_obj.UserClaims) {
			c.Set(
				common.UserProfile,
				&val_obj.UserClaims{
					UserID:   claims.UserID,
					Nickname: claims.Nickname,
					HeadURL:  claims.HeadURL,
					IsHelper: common.IsHelper(cast.ToString(claims.UserID)),
				},
			)
		})
		// 模拟用户登陆, 测试用, 生产环境应失效!
		if config.IsDebugMode() && c.Query("mock_login") != "" {
			c.Set(
				common.UserProfile,
				&val_obj.UserClaims{
					UserID:   435737,
					Nickname: "admin",
					HeadURL:  "https://images.gitee.com/uploads/images/2021/0731/134131_a864d20c_7809561.jpeg",
					IsHelper: true,
				},
			)
			isLogin = true
		}
		if !isLogin {
			c.JSON(200, notLogin)
			c.Abort()
			return
		}
		c.Next()
	}
}

func isLoginAndDo(c *gin.Context, client *JwtClient, do func(token *jwt.Token, claims *val_obj.UserClaims)) bool {
	tokenString := c.GetHeader(common.AuthTokenName)
	if tokenString == "" {
		tokenString, _ = c.GetQuery(common.AuthTokenName)
	}
	token, claims, err := client.ParseToken(tokenString)
	if err != nil {
		log.Printf("ParseToken fail, err: %+v", err)
		return false
	}
	do(token, claims)
	return true
}
