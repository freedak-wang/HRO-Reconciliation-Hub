package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"your-project/models"
	"your-project/services"
)

type PersonnelHandler struct {
	personnelService *services.PersonnelService
}

func NewPersonnelHandler(personnelService *services.PersonnelService) *PersonnelHandler {
	return &PersonnelHandler{personnelService: personnelService}
}

func (h *PersonnelHandler) AddPersonnel(c *gin.Context) {
	var personnel models.Personnel
	if err := c.ShouldBindJSON(&personnel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.personnelService.AddPersonnel(&personnel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, personnel)
}

func (h *PersonnelHandler) RemovePersonnel(c *gin.Context) {
	personnelID := c.Param("id")
	var req struct {
		EndDate string `json:"end_date" binding:"required"`
		Reason  string `json:"reason" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	if err := h.personnelService.RemovePersonnel(personnelID, endDate, req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Personnel removed successfully"})
}

func (h *PersonnelHandler) ListPersonnel(c *gin.Context) {
	contractID := c.Param("contractId")
	personnel, err := h.personnelService.ListPersonnel(contractID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personnel)
}

func (h *PersonnelHandler) GetPersonnelChanges(c *gin.Context) {
	contractID := c.Param("contractId")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	changes, err := h.personnelService.GetPersonnelChanges(contractID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, changes)
}