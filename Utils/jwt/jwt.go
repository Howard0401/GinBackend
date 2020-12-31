package jwt

import (
	md5 "VueGin/Utils/crypto"
	"VueGin/global"
	"VueGin/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id         string `json:"id"`
	NickName   string `json:"nickName"`
	BufferTime int64
	jwt.StandardClaims
}

type JWT struct {
	key []byte
}

func NewJWT() *JWT {
	return &JWT{key: []byte(global.Global_Viper.GetString("JWT.Secret"))}
}

// func GetJWTSecret() []byte {
// 	return []byte(global.Global_Viper.GetString("JWT.Secret"))
// }

//JWT => Header + Payload + Signature(對前面兩個部分檢驗Secrect)
func (j *JWT) GenerateToken(User model.User) (string, error) {
	nowTime := time.Now()
	//從config讀取設定，轉換成過期時間
	exTime := time.Duration(global.Global_Viper.GetInt("JWT.Expire")) * time.Second
	expireTime := nowTime.Add(exTime)
	//理論上，claim還可以再加其他非隱私資訊/驗證，僅示例故採用最簡單的資訊
	claims := Claims{
		//自定義認證
		Id:       User.UserId,
		NickName: md5.Md5(User.NickName),
		StandardClaims: jwt.StandardClaims{ //對應JWT中Payload的字串
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.Global_Viper.GetString("JWT.Issuer"),
		},
	}
	token, err := j.CreateToken(claims)
	return token, err
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	//根據HS256生成的加密token (可選HS384、HS512)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//簽名字串
	// token, err := tokenClaims.SignedString(GetJWTSecret()) //改成用pointer receiver，所以直接從struct取得輸入簽名的字串
	token, err := tokenClaims.SignedString(j.key)
	return token, err
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// return GetJWTSecret(), nil
		return j.key, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
