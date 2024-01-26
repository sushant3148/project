package routes

import (
    "github.com/gin-gonic/gin"
    "dhobi_ghat/controllers"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
)
func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Use(RequestLogger())

    route := router.Group("/", ginBodyLogMiddleware())

    route.GET("/users/product_view", controllers.Products())
    route.GET("/users/search", controllers.ProductByQuery())

    dhobi := route.Group("/dhobi")
    {
        dhobi.POST("/users/singup", controllers.UserSingup())
        dhobi.POST("/users/login", controllers.UserLogin())
        dhobi.POST("/admin/add_product", controllers.AddProduct())
        
        
    }

    return router
}