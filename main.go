package main

import (
	"controller"
	"fmt"
	"github.com/13283339616/apollo"
	"github.com/13283339616/eureka"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
	"vo"
)

func main() {

	fmt.Print("董桂君的框架开启")

	router := gin.Default()
	sep := string(os.PathSeparator)
	dir, _ := os.Getwd()
	path := dir + sep + "resource"
	fmt.Print(path)
	apollo.InitAll(path)

	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://39.97.194.236:8081/eureka/",
		App:                   "five",
		Port:                  9090,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
	// start client, register、heartbeat、refresh
	client.Start()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello eureka!")
	})

	router.POST("/user/info", func(c *gin.Context) {
		u := new(vo.GetUserInfoResponsetVo)
		requestVo := vo.GetUserInfoRequestVo{}
		if errA := c.ShouldBind(&requestVo); errA != nil {
			c.String(http.StatusOK, `the body should be formA`)
			// Always an error is occurred by this because c.Request.Body is EOF now.
		}
		fmt.Println("xiaowu")
		fmt.Println(requestVo.Name)
		var wg sync.WaitGroup
		wg.Add(1)
		poolObj.Submit(func() {
			defer wg.Done()
			user := new(conroller.UserController)
			u = user.GetInfo(c)
		})
		wg.Wait()
		c.JSON(200, u)

	})
	router.Use(Authorize())
	// 指定地址和端口号
	router.Run("127.0.0.1:9090")

}
