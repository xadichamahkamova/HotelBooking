package http

import (
	"api-gateway/internal/http/handler"
	auth "api-gateway/internal/http/middleware/authorization"
	rlimit "api-gateway/internal/http/middleware/rate-limiting"
	"api-gateway/internal/service"
	"crypto/tls"
	"net/http"

	producer "api-gateway/internal/kafka/producer"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(service service.ServiceRepositoryClient, producer producer.ProducerInit) *http.Server {

	r := gin.Default()

	apiHandler := handler.NewApiHandler(service, producer)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(auth.MiddleWare())
	r.Use(rlimit.RateLimit())

	api := r.Group("/api")
	{
		api.POST("/users", apiHandler.RegisterUser)
		api.POST("/users/login", apiHandler.LoginUser)

		api.GET("/hotels", apiHandler.ListOfHotel)
		api.GET("/hotels/:id", apiHandler.GetDetailsOfHotel)
		api.GET("/hotels/:id/rooms/availability", apiHandler.GetAvailabilityRooms)

		api.POST("/bookings", apiHandler.CreateBooking)
		api.GET("/bookings/:id", apiHandler.GetDetailsOfBooking)
		api.PUT("/bookings/:id", apiHandler.UpdateBooking)
		api.DELETE("/bookings/:id", apiHandler.CancelBooking)
		api.GET("/users/:id/bookings", apiHandler.ListBookingOfUsers)

	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      ":9000",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}
