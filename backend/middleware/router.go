package middleware

import (
	_ "perna/docs"
	helper "perna/internal/helper/handler"
	helper_service "perna/internal/helper/service"
	programs "perna/internal/programs/handler"
	programs_service "perna/internal/programs/service"
	register "perna/internal/register/handler"
	register_service "perna/internal/register/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title API Documentation
// @version 1.0
// @description API Documentation for Domestic Violence BO, Programs, and Emergency Requests
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ResponseHandler())
	// Use response, but not Token
	r.GET("/token", generateTokenHandler)

	auth := r.Group("/api")

	// Register Controllers
	registerController := register.NewRegisterController(register_service.NewRegisterService(db))
	auth.POST("/domestic-violence-bos", registerController.CreateRegister)
	auth.GET("/domestic-violence-bos", registerController.GetAllRegisters)
	auth.GET("/domestic-violence-bos/:id", registerController.GetRegisterByID)
	auth.PUT("/domestic-violence-bos/:id", registerController.UpdateRegister)
	auth.DELETE("/domestic-violence-bos/:id", registerController.DeleteRegister)

	// Programs Controllers
	programsController := programs.NewProgramsController(programs_service.NewProgramsService(db))
	auth.POST("/programs", programsController.CreateProgram)
	auth.GET("/programs", programsController.GetAllPrograms)
	auth.GET("/programs/:id", programsController.GetProgramByID)
	auth.PUT("/programs/:id", programsController.UpdateProgram)
	auth.DELETE("/programs/:id", programsController.DeleteProgram)

	// Helper Controllers (Emergency Request)
	helperController := helper.NewHelperController(helper_service.NewHelperService(db))
	auth.POST("/emergency-requests", helperController.CreateHelper)
	auth.GET("/emergency-requests", helperController.GetAllHelpers)
	auth.GET("/emergency-requests/:id", helperController.GetHelperByID)
	auth.PUT("/emergency-requests/:id", helperController.UpdateHelper)
	auth.DELETE("/emergency-requests/:id", helperController.DeleteHelper)

	return r
}
