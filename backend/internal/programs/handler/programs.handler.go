package programs

import (
	"net/http"
	programs_service "perna/internal/programs/service"
	interfaces "perna/internal/swagger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProgramsController struct {
	Service *programs_service.ProgramsService
}

func NewProgramsController(service *programs_service.ProgramsService) *ProgramsController {
	return &ProgramsController{Service: service}
}

// @Summary Create Program
// @Description Create a new Program record
// @Tags Program
// @Accept json
// @Produce json
// @Param Program body interfaces.Program true "Program Data"
// @Success 200 {object} db.Program
// @Router /api/programs [post]
func (ctrl *ProgramsController) CreateProgram(c *gin.Context) {
	var program interfaces.Program

	if err := c.ShouldBindJSON(&program); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.CreateProgram(&program); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, program)
}

// @Summary Get All Programs
// @Description Get all Programs
// @Tags Program
// @Produce json
// @Success 200 {object} []db.Program
// @Router /api/programs [get]
func (ctrl *ProgramsController) GetAllPrograms(c *gin.Context) {
	records, err := ctrl.Service.GetAllPrograms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// @Summary Get Program by ID
// @Description Get a single Program by ID
// @Tags Program
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {object} db.Program
// @Router /api/programs/{id} [get]
func (ctrl *ProgramsController) GetProgramByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	record, err := ctrl.Service.GetProgramByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Summary Update Program by ID
// @Description Update a Program record by ID
// @Tags Program
// @Accept json
// @Produce json
// @Param id path int true "Program ID"
// @Param Program body interfaces.Program true "Updated Program Data"
// @Success 200 {object} db.Program
// @Router /api/programs/{id} [put]
func (ctrl *ProgramsController) UpdateProgram(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProgram interfaces.Program
	if err := c.ShouldBindJSON(&updatedProgram); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.UpdateProgram(id, &updatedProgram); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProgram)
}

// @Summary Delete Program by ID
// @Description Delete a Program record by ID
// @Tags Program
// @Produce json
// @Param id path int true "Program ID"
// @Success 200 {string} string "Deleted Successfully"
// @Router /api/programs/{id} [delete]
func (ctrl *ProgramsController) DeleteProgram(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.DeleteProgram(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
