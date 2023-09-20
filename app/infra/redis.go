package infra

import (
	
	"fmt"

	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func NewSession(c echo.Context, cookieKey, redisValue string) error {
	//b := make([]byte, 64)
	redisKey := uuid.New().String()
	// if _, err := io.ReadFull(uuid, b); err != nil {
	// 	panic("ランダムな文字作成時にエラーが発生しました。")
	// }
	if err := conn.Set(conn.Context(), redisKey ,redisValue,0).Err(); err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = cookieKey
	cookie.Value = redisKey
	cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)
	return nil
}
//0f018422-652a-4bb3-b1b2-463c142c0803

func GetSession(c echo.Context, CookieKey string)(string, error) {
	
	 cookie, err := c.Cookie(CookieKey); 
	 if err != nil {
		return "", err
	 }
	redisKey := cookie.Value
	

	
	fmt.Println(redisKey)
	redisValue, err := conn.Get(conn.Context(), redisKey).Result()

	if err != nil {
		return "",err
	}
	// switch {
	// case err == redis.Nil:
	// 	fmt.Println("SessionKeyが登録されていません。")
	// 	return nil
	// case err != nil:
	// 	fmt.Println("Session取得時にエラー発生：" + err.Error())
	// 	return nil
	// }
	
	return redisValue, err

	
		
}


func DeleteSession(c echo.Context, cookieKey string)   {

cookie, _ := c.Cookie(cookieKey)

redisId := cookie.Value

	conn.Del(conn.Context(), redisId)
	//c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)

	cookie = new(http.Cookie)
	cookie.Name = cookieKey
	cookie.Value = ""
	cookie.Expires = time.Now()
 c.SetCookie(cookie)
}