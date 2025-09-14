package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/http_handler/user"
)

func InitUserGroup(r *gin.RouterGroup, app userhandler.UserHandlerRepo) {

	r.POST("/users", app.RegisterUser)
	// r.GET("/allTasks", app.GetTasks)
	// r.GET("/taskById/:id", app.GetTask)
	// r.PUT("/updateTask/:id", app.UpdateTask)
	// r.PUT("/setTaskDone/:id", app.SetTaskDone)
	// r.DELETE("/deleteTask/:id", app.DeleteTask)

	//TODO: Swagger Doc
	//docs.SwaggerInfo.BasePath = "/"
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
