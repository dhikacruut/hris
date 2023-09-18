package servers

import (
	"github.com/dhikacruut/hris/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutesCompliance(g *gin.RouterGroup) {
	{
		g.GET("/GenerateSlip", controllers.GenerateSlip)
		g.GET("/ListSlip/:period", controllers.SalarySlipShow)
		g.GET("/ListSlipDetail/:period/:id", controllers.SalarySlipDetailShow)
		g.POST("/CreateSlip/", controllers.SalarySlipCreate)
		g.POST("/CreateSlipDetail/", controllers.SalarySlipDetailCreate)
		g.PUT("/UpdateSlip/:id", controllers.SalarySlipUpdate)
		g.PUT("/UpdateSlipDetail/:id", controllers.SalarySlipDetailUpdate)
		g.DELETE("/DeleteSlip/:id", controllers.SalarySlipDelete)
		g.DELETE("/DeleteSlipDetail/:id", controllers.SalarySlipDetailDelete)
	}
}
