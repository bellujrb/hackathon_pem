package helper

import (
	"net/http"
	helper_service "perna/internal/helper/service"
	interfaces "perna/internal/swagger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HelperController struct {
	Service *helper_service.HelperService
}

func NewHelperController(service *helper_service.HelperService) *HelperController {
	return &HelperController{Service: service}
}

// @Summary Create Emergency Request
// @Description Create a new Emergency Request record
// @Tags EmergencyRequest
// @Accept json
// @Produce json
// @Param EmergencyRequest body interfaces.EmergencyRequest true "Emergency Request Data"
// @Success 200 {object} db.EmergencyRequest
// @Router /api/emergency-requests [post]
func (ctrl *HelperController) CreateHelper(c *gin.Context) {
	var emergencyRequest interfaces.EmergencyRequest

	if err := c.ShouldBindJSON(&emergencyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.CreateHelper(&emergencyRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emergencyRequest)
}

// @Summary Get All Emergency Requests
// @Description Get all Emergency Requests
// @Tags EmergencyRequest
// @Produce json
// @Success 200 {object} []db.EmergencyRequest
// @Router /api/emergency-requests [get]
func (ctrl *HelperController) GetAllHelpers(c *gin.Context) {
	records, err := ctrl.Service.GetAllHelpers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// @Summary Get Emergency Request by ID
// @Description Get a single Emergency Request by ID
// @Tags EmergencyRequest
// @Produce json
// @Param id path int true "Emergency Request ID"
// @Success 200 {object} db.EmergencyRequest
// @Router /api/emergency-requests/{id} [get]
func (ctrl *HelperController) GetHelperByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	record, err := ctrl.Service.GetHelperByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Emergency Request not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Summary Update Emergency Request by ID
// @Description Update an Emergency Request record by ID
// @Tags EmergencyRequest
// @Accept json
// @Produce json
// @Param id path int true "Emergency Request ID"
// @Param EmergencyRequest body interfaces.EmergencyRequest true "Updated Emergency Request Data"
// @Success 200 {object} db.EmergencyRequest
// @Router /api/emergency-requests/{id} [put]
func (ctrl *HelperController) UpdateHelper(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedEmergencyRequest interfaces.EmergencyRequest
	if err := c.ShouldBindJSON(&updatedEmergencyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.UpdateHelper(id, &updatedEmergencyRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEmergencyRequest)
}

// @Summary Delete Emergency Request by ID
// @Description Delete an Emergency Request record by ID
// @Tags EmergencyRequest
// @Produce json
// @Param id path int true "Emergency Request ID"
// @Success 200 {string} string "Deleted Successfully"
// @Router /api/emergency-requests/{id} [delete]
func (ctrl *HelperController) DeleteHelper(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.DeleteHelper(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
