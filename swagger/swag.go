package swagger

import (
	_ "gdemo1/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @securityDefinitions.apikey ApiKeyAuth
// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", HandleHello)
		// v1.POST("/login", HandleLogin)
		// v1Auth := r.Use(HandleAuth)
		// {
		//     v1Auth.POST("/upload", HandleUpload)
		//     v1Auth.GET("/list", HandleList)
		// }
	}
	r.Run(":8080")
	// _ = s.bar

	// dialector := mysql.Open("fage:Fage501526~@/mytest?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open(dialector)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var hookLevels []logrus.Level

	// hookLevels = append(hookLevels, 0, 1, 2, 3, 4, 5, 6)

	// h := loggerhook.New(loggergormhook.New(db),

	// 	loggerhook.SetMaxQueues(1),
	// 	loggerhook.SetMaxWorks(1),
	// 	loggerhook.SetLevels(hookLevels...),
	// )
	// fmt.Println(h.Levels())
	// logger.AddHook(h)

	// ctx1 := logger.NewTagContext(context.Background(), "__main__")
	// logger.SetFormatter("json")
	// logger.WithContext(context.Background()).Info("Trace12")
	// logger.WithContext(ctx1).Info("Trace")
	// time.Sleep(time.Second * 10)
	// fmt.Println("end") // @Security ApiKeyAuth
}

//@Security ApiKeyAuth
/// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]
func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}
