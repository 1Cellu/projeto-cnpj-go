package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"projeto-cnpj-go/internal/modules"
	"projeto-cnpj-go/internal/services"
)

func SetupRouter(svc services.Service) *gin.Engine {
	router := gin.Default()

	router.GET("v1/list", func(c *gin.Context) {
		fmt.Println("GET path v1/list")
		items, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, items)
	})

	router.GET("v1/list/:cnpj", func(c *gin.Context) {
		fmt.Println("GET path v1/list/:cnpj")
		cnpj := c.Param("cnpj")
		item, err := svc.Get(cnpj)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, &item)
	})

	router.POST("v1/add/:cnpj", func(c *gin.Context) {
		fmt.Println("POST path v1/add/:cnpj")
		cnpj := c.Param("cnpj")
		var companyInfo modules.CompanyInfo

		if err := c.ShouldBindJSON(&companyInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err := svc.AddRecord(cnpj, companyInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Success")
	})

	router.DELETE("v1/delete/:cnpj", func(c *gin.Context) {
		fmt.Println("GET path v1/delete/:cnpj")
		cnpj := c.Param("cnpj")
		err := svc.Delete(cnpj)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Success")
	})

	return router
}
