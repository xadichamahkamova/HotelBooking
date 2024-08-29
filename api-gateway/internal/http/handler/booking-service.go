package handler

import (
	pb "api-gateway/genproto/bookingpb"
	pbNotif "api-gateway/genproto/notificationpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/bookings [post]
// @Summary CREATE BOOKINGS
// @Security  		BearerAuth
// @Description This method creates a new booking
// @Tags BOOKING
// @Accept json
// @Produce json
// @Param booking body models.CreateBookingRequest true "Booking"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateBooking(c *gin.Context) {

	req := pb.CreateBookingRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("Failed to bind CreateBookingRequest: ", err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
	}

	logger.Info("Creating booking for user:", req.UserId)
	_, err := h.Service.CreateBooking(ctx, &req)
	if err != nil {
		logger.Error("Failed to create booking: ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("Booking created successfully for user:", req.UserId)

	email, exists := c.Get("email")
	if !exists {
		logger.Error("Email not found in context")
		c.AbortWithStatusJSON(500, gin.H{
			"error": "email not found",
		})
	}

	logger.Info("Sending email to:", email.(string))
	requestToNotifSrv := pbNotif.SendEmailRequest{
		To:        email.(string),
		HotelName: req.RoomType,
	}
	respFromNotifSrv, err := h.Service.SendEmail(ctx, &requestToNotifSrv)
	if err != nil {
		logger.Error("Failed to send email:", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("Email sent successfully to:", email.(string))

	err = h.Producer.ProduceMessage(email.(string), []byte(respFromNotifSrv.Message))
	if err != nil {
		logger.Error("Failed to produce message:", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	
	c.JSON(200, gin.H{"message": "The message was successfully sent to the WebSocket service"})
}

// @Router /api/bookings/{id} [get]
// @Summary GET BOOKING DETAILS
// @Security  		BearerAuth
// @Description This method retrieves the details of a booking
// @Tags BOOKING
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} models.Booking
// @Failure 500 {object} string
func (h *HandlerST) GetDetailsOfBooking(c *gin.Context) {

	req := pb.GetDetailsOfBookingRequest{}
	req.BookingId = c.Param("id")

	logger.Info("Fetching booking details for booking ID:", req.BookingId)
	resp, err := h.Service.GetDetailsOfBooking(ctx, &req)
	if err != nil {
		logger.Error("Failed to fetch booking details for booking ID ", req.BookingId, ": ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("Successfully fetched booking details for booking ID:", req.BookingId)
	c.JSON(200, resp)
}

// @Router /api/bookings/{id} [put]
// @Summary UPDATE BOOKINGS
// @Security  		BearerAuth
// @Description This method updates an existing booking
// @Tags BOOKING
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Param booking body models.UpdateBookingRequest true "Booking"
// @Success 200 {object} models.Booking
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateBooking(c *gin.Context) {

	req := pb.UpdateBookingRequest{}
	req.BookingId = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("Failed to bind UpdateBookingRequest for booking ID ", req.BookingId, ": ", err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
	}
	resp, err := h.Service.UpdateBooking(ctx, &req)
	if err != nil {
		logger.Error("Failed to update booking with booking ID ", req.BookingId, ": ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("Successfully updated booking with booking ID:", req.BookingId)
	c.JSON(200, resp)
}

// @Router /api/bookings/{id} [delete]
// @Summary CANCEL BOOKING
// @Security  		BearerAuth
// @Description This method cancels an existing booking
// @Tags BOOKING
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} models.CancelBookingResponse
// @Failure 500 {object} string
func (h *HandlerST) CancelBooking(c *gin.Context) {

	req := pb.CancelBookingRequest{}
	req.BookingId = c.Param("id")
	logger.Info("Cancelling booking with booking ID:", req.BookingId)

	resp, err := h.Service.CancelBooking(ctx, &req)
	if err != nil {
		logger.Error("Failed to cancel booking with booking ID ", req.BookingId, ": ", err)
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}

	logger.Info("Successfully cancelled booking with booking ID:", req.BookingId)
	c.JSON(200, resp)
}

// @Router /api/users/{id}/bookings [get]
// @Summary LIST USER BOOKINGS
// @Security  		BearerAuth
// @Description This method lists all bookings of a user
// @Tags BOOKING
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.ListBookingOfUserResponse
// @Failure 500 {object} string
func (h *HandlerST) ListBookingOfUsers(c *gin.Context) {

	req := pb.ListBookingOfUserRequest{}
	req.UserId = c.Param("id")
	logger.Info("Listing bookings for user ID:", req.UserId)

	resp, err := h.Service.ListBookingOfUsers(ctx, &req)
	if err != nil {
		logger.Error("Failed to list bookings for user ID ", req.UserId, ": ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("Successfully listed bookings for user ID:", req.UserId)
	c.JSON(200, resp)
}
