//package main
//
//import (
//	"fmt"
//)
//
//var MsgFlags = map[int]string{
//	SUCCESS:                        "ok",
//	ERROR:                          "fail",
//	INVALID_PARAMS:                 "请求参数错误",
//	ERROR_EXIST_TAG:                "已存在该标签名称",
//	ERROR_NOT_EXIST_TAG:            "该标签不存在",
//	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
//	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
//	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
//	ERROR_AUTH_TOKEN:               "Token生成失败",
//	ERROR_AUTH:                     "Token错误",
//}
//
//func GetMsg(code int) string {
//	msg, ok := MsgFlags[code]
//	if ok {
//		return msg
//	}
//	return MsgFlags[ERROR]
//}
//
//func main() {
//	fmt.Println(GetMsg(400))
//}

package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
