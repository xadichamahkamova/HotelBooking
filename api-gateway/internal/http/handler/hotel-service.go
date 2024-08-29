package handler

import (
	pb "api-gateway/genproto/hotelpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/hotels [get]
// @Summary LIST OF HOTELS
// @Security  		BearerAuth
// @Description This method lists all hotels
// @Tags HOTEL
// @Accept json
// @Produce json
// @Success 200 {object} models.ListOfHotelResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListOfHotel(c *gin.Context) {

	resp, err := h.Service.ListOfHotel(ctx, &pb.ListOfHotelRequest{})
	if err != nil {
		logger.Error("Error retrieving list of hotels: ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("Successfully retrieved list of hotels")
	c.JSON(200, resp)
}

// @Router /api/hotels/{id} [get]
// @Summary GET DETAILS OF HOTEL
// @Security  		BearerAuth
// @Description This method retrieves details of a hotel by its ID
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} models.GetDetailsOfHotelResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDetailsOfHotel(c *gin.Context) {

	req := pb.GetDetailsOfHotelRequest{}
	req.HotelId = c.Param("id")
	resp, err := h.Service.GetDetailsOfHotel(ctx, &req)
	if err != nil {
		logger.Error("Error retrieving details of hotel (ID: ", req.HotelId, "): ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Successfully retrieved details of hotel (ID: ", req.HotelId, ")")
	c.JSON(200, resp)
}

// @Router /api/hotels/{id}/rooms/availability [get]
// @Summary GET AVAILABILITY OF ROOMS
// @Security  		BearerAuth
// @Description This method retrieves room availability by hotel ID
// @Tags HOTEL
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} models.GetAvailabilityRoomsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetAvailabilityRooms(c *gin.Context) {

	req := pb.GetAvailabilityRoomsRequest{}
	req.HotelId = c.Param("id")
	resp, err := h.Service.GetAvailabilityRooms(ctx, &req)
	if err != nil {
		logger.Error("Error retrieving room availability for hotel (ID: ", req.HotelId, "): ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("Successfully retrieved room availability for hotel (ID: ", req.HotelId, ")")
	c.JSON(200, resp)
}
