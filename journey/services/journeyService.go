package journeyService

import (
	journeyModels "portfolioAPI/journey/models"
	journeyRepo "portfolioAPI/journey/repos"
	journeyMapping "portfolioAPI/journey/services/mapping"
	taskService "portfolioAPI/tasks/services"
)

type JourneyService struct {
	repository *journeyRepo.JourneyRepo
	taskServie *taskService.TaskService
}

func NewJourneyService(journeyRepo *journeyRepo.JourneyRepo, taskService *taskService.TaskService) *JourneyService {
	return &JourneyService{
		repository: journeyRepo,
		taskServie: taskService,
	}
}

func (service *JourneyService) GetFullJourney(languageCode string) journeyModels.JourneyResponse {
	fullJourney := service.repository.GetFullJourney(languageCode)

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

func (service *JourneyService) Insert(journeysToInsert *journeyModels.JourneyUpsertModel) error {
	mappedExperiences := journeyMapping.MapJourneyDisplayModelToJourney(journeysToInsert)

	for _, experience := range mappedExperiences {
		_, err := service.repository.Insert(&experience)
		if err != nil {
			return err
		}

		service.taskServie.InsertTasks(experience.Tasks)
	}

	return nil
}

func (service *JourneyService) Delete(journeyId int) error {
	return service.repository.Delete(journeyId)
}

func (service *JourneyService) Update(journeyToUpdate *journeyModels.JourneyDisplay, experienceId int) (*journeyModels.JourneyDisplay, error) {
	experience := journeyMapping.MapJourneyDisplayModelToJourneyWithId(journeyToUpdate, experienceId)
	existingExperience, err := service.repository.GetJourney(experienceId)
	if err != nil {
		return nil, err
	}
	experience.LanguageCode = existingExperience.LanguageCode

	_, error := service.repository.Update(&experience)

	if error != nil {
		return nil, error
	}

	service.taskServie.DeleteTasks(experience.ID)
	service.taskServie.InsertTasks(experience.Tasks)

	return service.repository.GetJourney(experienceId)
}
