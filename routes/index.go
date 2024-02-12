package routes

import (
	"echolabstack/service"
	"echolabstack/ratelimitter"
	"github.com/labstack/echo/v4"
)


func Echoroutes(e *echo.Echo) {
	
	e.Static("/static", "static")
	//ratelimitting api
	e.GET("/users", service.NewAPI, ratelimitter.CombinedRateLimiter())
	e.GET("/pdfapi", service.PdfAPI,ratelimitter.CombinedRateLimiter())
	//emailid
	e.POST("/subscribe", service.EmailIDAPI,ratelimitter.CombinedRateLimiter())

}