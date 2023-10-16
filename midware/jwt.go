package midware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	orm "liuliang/model"
	"net/http"
	"time"
)

var jwtkey = []byte("i_really_like_you")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Gettoken(username1 string) (string, error) {
	expirationtime := time.Now().Add(1 * time.Hour)
	claims := Claims{
		username1,
		jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "zzw",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtkey)
	return tokenstring, err
}

func parsetoken(tokenstring string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效token")
}

func Sendtoken() func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		passwd := c.PostForm("passwd")
		if orm.Userlogin(username, passwd) {
			tokenstring, _ := Gettoken(username)
			c.JSON(http.StatusOK, gin.H{
				"token": tokenstring,
			})
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "密码错误",
			})
			c.Abort()
		}
	}
}

func Twjparse() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenstring := c.PostForm("token")
		if tokenstring == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "token为空",
			})
			c.Abort()
		}
		_, err := parsetoken(tokenstring)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "无效的token",
			})
			c.Abort()
		}
		c.Next()
	}
}
