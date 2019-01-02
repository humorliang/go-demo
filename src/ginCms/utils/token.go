package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("JWTSCERYT")

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

//JWT 是由header,payload,signature三部分组成的

//创建一个Token
func CreateToken(key []byte, m map[string]int) (string, error) {
	//token:=
	userId := m["userId"]
	//token 生效时间
	now := time.Now()
	//token 过期时间
	expireTime := now.Add(time.Hour * time.Duration(1))

	//claims 设置
	claims := Claims{
		userId,
		jwt.StandardClaims{
			//过期时间使用unix时间戳
			ExpiresAt: expireTime.Unix(),
			//发行人
			Issuer: "blogApp",
		},
	}

	//token claim :第一个参数为加密方式  claims项
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//根据秘钥进行加密
	token, err := tokenClaims.SignedString(key)
	return token, err
}

//解析一个Token
func ParseToken(token string) (*Claims, error) {
	//解析得到一个tokenClaims项
	tokenClaims, err := jwt.ParseWithClaims(token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			//该函数返回对应token秘钥
			return JwtSecret, nil
		})
	//log.Println(tokenClaims,err)
	//对claims进行检验
	if tokenClaims != nil {
		//对tokenClaims进行验证
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
