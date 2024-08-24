package routes

import (
	"restaurent-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
	// incomingRoutes.DELETE("/invoices/:invoice_id", controllers.DeleteInvoice())
}