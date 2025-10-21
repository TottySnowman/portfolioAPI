package journeyController

import (
	"net/http"
	journeyModels "portfolioAPI/journey/models"
	journeyService "portfolioAPI/journey/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JourneyController struct {
	journeyService *journeyService.JourneyService
}

func NewJourneyController(journeyService *journeyService.JourneyService) *JourneyController {
	return &JourneyController{
		journeyService: journeyService,
	}
}

func (con *JourneyController) GetFullJourney(context *gin.Context) {
	acceptLanguage := context.GetHeader("Accept-Language")

	fullJourney := con.journeyService.GetFullJourney(acceptLanguage)
	context.IndentedJSON(http.StatusOK, fullJourney)
}

func (con *JourneyController) InsertJourney(context *gin.Context) {
	var experience *journeyModels.JourneyUpsertModel

	if err := context.ShouldBindJSON(&experience); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}
	print(experience.ExperienceType)

	err := con.journeyService.Insert(experience)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.Status(http.StatusCreated)
}

func (con *JourneyController) DeleteExperience(context *gin.Context) {
	experienceId, err := strconv.Atoi(context.Param("ID"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := con.journeyService.Delete(experienceId); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}

func (con *JourneyController) UpdateExperience(context *gin.Context) {
	var experience *journeyModels.JourneyDisplay

	experienceId, err := strconv.Atoi(context.Param("ID"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := context.ShouldBindJSON(&experience); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	updatedProject, err := con.journeyService.Update(experience, experienceId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}
	context.IndentedJSON(http.StatusOK, updatedProject)
}
