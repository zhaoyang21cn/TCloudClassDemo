package main

import (
	"classdemo/src"
	"fmt"
	"net/http"
	"runtime"
	"github.com/gin-gonic/gin"
)

func main() {

	defer func() {
		// crash保护
		if e := recover(); e != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			stack := string(buf[:n])
			fmt.Printf( "crash:%s", stack)
		}

	}()


	g := gin.Default()
	g.Use(Options)

	// demo路由
	demoV1 := g.Group("/demo/v1")
	SetRouter(demoV1)

	addr := fmt.Sprintf(":%d", 8080)
	fmt.Println(addr)

	g.Run(addr)
}
func SetRouter(rGroup *gin.RouterGroup) {
	rGroup.POST("/new_enter/create",  src.CreateNewEnterId)
	rGroup.POST("/user/register",  src.UserRegister)
	rGroup.POST("/login",  src.UserOpenLogin)
	rGroup.POST("/class/create",  src.CreateClass)
	rGroup.POST("/classroom_code/create",  src.CreateClassRoomCode)

}
func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.AbortWithStatus(http.StatusOK)
	}
}
