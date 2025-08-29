package journeyService

import (
	journeyModels "portfolioAPI/journey/models"
	journeyRepo "portfolioAPI/journey/repos"
)

type JourneyService struct {
	repository *journeyRepo.JourneyRepo
}

func NewJourneyService(journeyRepo *journeyRepo.JourneyRepo) *JourneyService {
	return &JourneyService{
		repository: journeyRepo,
	}
}

func (service *JourneyService) GetFullJourney() journeyModels.JourneyResponse {
	fullJourney := service.repository.GetFullJourney()

	return mapFullJourneyToJourneyResponse(fullJourney)
}

func mapFullJourneyToJourneyResponse(journeyDisplay []journeyModels.JourneyDisplay) journeyModels.JourneyResponse {
	var journeyResponse journeyModels.JourneyResponse
	for _, journey := range journeyDisplay {
		switch journey.ExperienceType {
		case journeyModels.Education:
			journeyResponse.Educations = append(journeyResponse.Educations, journey)
			break
		case journeyModels.Work:
			journeyResponse.Work = append(journeyResponse.Work, journey)
			break
		}
	}

	return journeyResponse
}
