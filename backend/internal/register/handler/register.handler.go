package register

import (
	"net/http"
	register_service "perna/internal/register/service"
	interfaces "perna/internal/swagger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	Service *register_service.RegisterService
}

func NewRegisterController(service *register_service.RegisterService) *RegisterController {
	return &RegisterController{Service: service}
}

// @Summary Create Domestic Violence BO
// @Description Create a new Domestic Violence BO record
// @Tags DomesticViolenceBO
// @Accept json
// @Produce json
// @Param DomesticViolenceBO body interfaces.DomesticViolenceBO true "Domestic Violence BO Data"
// @Success 200 {object} db.DomesticViolenceBO
// @Router /api/domestic-violence-bos [post]
func (ctrl *RegisterController) CreateRegister(c *gin.Context) {
	var domesticViolenceBO interfaces.DomesticViolenceBO

	if err := c.ShouldBindJSON(&domesticViolenceBO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.CreateRegister(&domesticViolenceBO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domesticViolenceBO)
}

// @Summary Get All Domestic Violence BOs
// @Description Get all Domestic Violence BOs
// @Tags DomesticViolenceBO
// @Produce json
// @Success 200 {object} []db.DomesticViolenceBO
// @Router /api/domestic-violence-bos [get]
func (ctrl *RegisterController) GetAllRegisters(c *gin.Context) {
	records, err := ctrl.Service.GetAllRegisters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// @Summary Get Domestic Violence BO by ID
// @Description Get a single Domestic Violence BO by ID
// @Tags DomesticViolenceBO
// @Produce json
// @Param id path int true "Domestic Violence BO ID"
// @Success 200 {object} db.DomesticViolenceBO
// @Router /api/domestic-violence-bos/{id} [get]
func (ctrl *RegisterController) GetRegisterByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	record, err := ctrl.Service.GetRegisterByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Domestic Violence BO not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Summary Update Domestic Violence BO by ID
// @Description Update a Domestic Violence BO record by ID
// @Tags DomesticViolenceBO
// @Accept json
// @Produce json
// @Param id path int true "Domestic Violence BO ID"
// @Param DomesticViolenceBO body interfaces.DomesticViolenceBO true "Updated Domestic Violence BO Data"
// @Success 200 {object} db.DomesticViolenceBO
// @Router /api/domestic-violence-bos/{id} [put]
func (ctrl *RegisterController) UpdateRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedDomesticViolenceBO interfaces.DomesticViolenceBO
	if err := c.ShouldBindJSON(&updatedDomesticViolenceBO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.UpdateRegister(id, &updatedDomesticViolenceBO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDomesticViolenceBO)
}

// @Summary Delete Domestic Violence BO by ID
// @Description Delete a Domestic Violence BO record by ID
// @Tags DomesticViolenceBO
// @Produce json
// @Param id path int true "Domestic Violence BO ID"
// @Success 200 {string} string "Deleted Successfully"
// @Router /api/domestic-violence-bos/{id} [delete]
func (ctrl *RegisterController) DeleteRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.DeleteRegister(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
