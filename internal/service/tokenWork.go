package service

import (
	"module/internal/database"
	"module/internal/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var accessKey = []byte("basic_key")
var refreshKey = []byte("super_mega_key")

// получение пары access и refresh token. передача refresh в базу данных
func TokenGetPair(GUID string) models.TokenResponser {

	var access string = createTokenAccess(GUID)
	var refresh string = createTokenRefresh(GUID, access)

	responser := models.TokenResponser{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	dbData := models.TokenDB{
		GUID:         GUID,
		RefreshToken: refresh,
	}
	database.InsertToken(dbData)

	return responser
}

// метод для проведения проверки токена
func TokenAccessValidate(bearer string) string {

	token, err := validateAccessToken(bearer)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "unathorized"
		}
		return "unathorized"
	}

	if !token.Valid {
		return "token expired"
	}

	// токен валиден. вернуть результат
	user := token.Claims.(*TokenAccessData)
	return "token useful, user = " + user.GUID

}

// получить новый рефреш токен
func RenewToken(refreshToken string) string {

	token, err := validateRefreshToken(refreshToken)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "unauthorized"
		}
		return "unauthorized"
	}

	if !token.Valid {
		return "token expired"
	}

	// токен валиден. удаляем рефреш токен из базы, получаем новый аксес токен
	refToken := token.Claims.(*TokenRefreshData)

	var DBmodel models.TokenDB
	DBmodel.GUID = refToken.GUID
	DBmodel.RefreshToken = refreshToken

	// проверка что в базе есть такой токен, и что он принадлежит этому пользователю
	retModel := database.SearchToken(DBmodel)
	if retModel.RefreshToken != DBmodel.RefreshToken {
		return "wrong user in token"
	}

	// удаление данного токена пользователя из базы
	var DBmodelDel models.TokenDB
	DBmodelDel.GUID = refToken.GUID
	DBmodelDel.RefreshToken = refreshToken
	database.DeleteToken(DBmodelDel)

	// создание нового токена
	newAccessToken := createTokenAccess(refToken.GUID)

	return newAccessToken
}

// создание аксес токена. срок жизни 1 минуты для теста
func createTokenAccess(GUID string) string {

	// создаем токен
	var tokenObj = TokenAccessData{
		GUID: GUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	}

	// method HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(accessKey)
	if err != nil {
		return ""
	}

	return tokenString
}

// создание рефреш токена. срок жини 5 минут для теста
func createTokenRefresh(GUID string, accessToken string) string {

	// создаем токен
	var tokenObj = TokenRefreshData{
		GUID:         GUID,
		AcceessToken: accessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	// method HS = HMAC + SHA 512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenObj)
	tokenString, err := token.SignedString(refreshKey)

	if err != nil {
		return ""
	}

	return tokenString
}

// проверка валидности access токена
func validateAccessToken(bearerToken string) (*jwt.Token, error) {

	tokenString := strings.Split(bearerToken, " ")[1]
	token, err := jwt.ParseWithClaims(tokenString, &TokenAccessData{}, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})
	return token, err
}

// проверка валидности refresh токена
func validateRefreshToken(bearerToken string) (*jwt.Token, error) {

	//tokenString := strings.Split(bearerToken, " ")[1]

	tokenString := bearerToken
	token, err := jwt.ParseWithClaims(tokenString, &TokenRefreshData{}, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	return token, err
}

// модели для работы с токенами
type TokenAccessData struct {
	GUID string
	jwt.StandardClaims
}

type TokenRefreshData struct {
	GUID         string
	AcceessToken string
	jwt.StandardClaims
}
