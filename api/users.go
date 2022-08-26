package api

import (
	"ECHO-GORM/api/helpers"
	"ECHO-GORM/db"
	"ECHO-GORM/model"
	"ECHO-GORM/redis"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c echo.Context) error {
	db := db.DbManager()
	note := model.Users{}
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	username := c.QueryParam("username")

	response := helpers.Response{}

	var kirim error

	err := db.Raw("SELECT id, username, email, password, name, is_login, refresh_token FROM users WHERE email = ? OR username = ? AND password = ?", email, username, password).Scan(&note).Error

	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusNotFound,
			Message:    "failed",
			Token:      nil,
		}
		fmt.Println(err)
	}

	value, _ := Generate(c, &note)
	token := value[0:20]
	fmt.Println(token)
	dataJson, _ := json.Marshal(note)
	fmt.Println("ini panjang token: ", len(token))

	Redis(token, dataJson)

	refreshtoken, _ := RefreshToken(c, &note)
	refresh := refreshtoken[0:20]

	updatedb := db.Model(&model.Users{}).Where("email = ? or username = ? AND password = ?", note.Email, note.Username, note.Password).Updates(map[string]interface{}{"is_login": true, "refresh_token": refresh})

	if updatedb != nil {
		fmt.Println("update login gagal")
	}
	fmt.Println(refresh)
	fmt.Println("panjang refresh token: ", len(refresh))

	response = helpers.Response{
		StatusCode:   http.StatusOK,
		Message:      "ok",
		Token:        token,
		RefreshToken: refresh,
	}

	kirim = c.JSON(http.StatusOK, response)

	return kirim

}

func Generate(c echo.Context, note *model.Users) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(note.ID), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(hash[0:20]), err
}

func RefreshToken(c echo.Context, note *model.Users) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(note.ID), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(hash[0:20]), err
}

func Redis(token string, userdata []byte) {

	rdb := redis.NewRedisClient()
	fmt.Println("redis client initialized")

	key := token
	data := userdata
	ttl := time.Duration(1800) * time.Second

	// store data using SET command
	op1 := rdb.Set(key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

}

func CheckHealth(c echo.Context) error {
	tokenParam := c.Request().Header.Get("Authorization")
	tokenKey := c.Request().Header.Get("X-App-Key")
	tokenSecret := c.Request().Header.Get("X-App-Secret")
	fmt.Println("healt", tokenParam)
	fmt.Println(tokenKey)
	fmt.Println(tokenSecret)
	tokenParamsAuth := strings.Split(tokenParam, " ")

	keyS := "training"
	secret := "raya"

	var err error
	if tokenKey != keyS {
		return err
	}
	if tokenSecret != secret {
		return err
	}

	fmt.Println("ini token params:", tokenParam)
	val, _ := GetRedis(tokenParamsAuth[1])
	fmt.Println("check -----", val)

	kirim := c.JSON(http.StatusOK, val)

	return kirim
}

func GetRedis(token string) (string, error) {

	rdb := redis.NewRedisClient()
	fmt.Println("redis client initialized")

	op2 := rdb.Get(token)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return "", nil
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return "", nil
	}
	log.Println("get operation success. result:", res)

	return res, nil
}

func Logout(c echo.Context) error {

	rdb := redis.NewRedisClient()
	tokenSecret := c.Request().Header.Get("Refresh-Token")
	fmt.Println("ini refresh token", tokenSecret)

	tokenParam := c.Request().Header.Get("Authorization")
	tokenParamAuthLogout := strings.Split(tokenParam, " ")

	db := db.DbManager()
	updaterefresh := db.Model(&model.Users{}).Where("refresh_token = ?", tokenSecret).Update("is_login", false)
	if updaterefresh != nil {
		fmt.Println("update error", updaterefresh)
	}

	err := rdb.Del(tokenParamAuthLogout...)
	if err != nil {
		return nil
	}
	return errors.New("anda berhasil logout")
}
