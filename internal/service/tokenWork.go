package service

import (
	"module/internal/database"
	"module/internal/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
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
func RenewToken(refreshToken string, accessToken string) models.TokenResponser {

	var result models.TokenResponser

	// проверка рефреш токена
	token, err := validateRefreshToken(refreshToken)
	if err != nil {
		result.RefreshToken = "refresh token unauthorized"
		if err == jwt.ErrSignatureInvalid {
			result.RefreshToken = "refresh token sign unknown"
			return result
		}
		return result
	}

	if !token.Valid {
		result.RefreshToken = "refresh token expired"
		return result
	}

	refToken := token.Claims.(*TokenRefreshData)

	if refToken.AcceessToken != accessToken {
		result.RefreshToken = "access token missmatch"
		return result
	}

	// токен валиден. удаляем рефреш токен из базы, получаем новый аксес токен
	var DBmodel models.TokenDB
	DBmodel.GUID = refToken.GUID
	DBmodel.RefreshToken = refreshToken

	// проверка что в базе есть такой токен, и что он принадлежит этому пользователю
	retModel := database.SearchToken(DBmodel)
	if retModel.RefreshToken != DBmodel.RefreshToken {
		result.RefreshToken = "wrong user in token"
		return result
	}

	// удаление всех токенов (если как-то получилось много) у данного GUID из базы (возможно не надо)
	// !!!!!!!!!!!!!!!!!!!!!!
	var DBmodelDel models.TokenDB
	DBmodelDel.GUID = refToken.GUID
	database.DeleteToken(DBmodelDel)
	// !!!!!!!!!!!!!!!!!!!!!!

	// создание нового рефреш и аксес токена
	result = TokenGetPair(refToken.GUID)

	return result
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
		log.Error("token dont signed")
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
		log.Error("token dont signed")
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
