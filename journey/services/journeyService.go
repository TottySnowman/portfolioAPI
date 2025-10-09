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

func (service *JourneyService) Insert(journey *journeyModels.JourneyDisplay) (*journeyModels.JourneyDisplay, error) {
	mappedExperience := journeyMapping.MapJourneyDisplayModelToJourney(*journey)
	experience, err := service.repository.Insert(&mappedExperience)
	if err != nil {
		return nil, err
	}

	service.taskServie.InsertTasks(experience.Tasks)
	return service.repository.GetJourney(experience.ID)
}

func (service *JourneyService) Delete(journeyId int) error {
	return service.repository.Delete(journeyId)
}

func (service *JourneyService) Update(journeyToUpdate *journeyModels.JourneyDisplay, experienceId int) (*journeyModels.JourneyDisplay, error) {
	experience := journeyMapping.MapJourneyDisplayModelToJourney(*journeyToUpdate)
	experience.ID = experienceId

	_, error := service.repository.Update(&experience)

	if error != nil {
		return nil, error
	}

	service.taskServie.DeleteTasks(experience.ID)

  var tasks []journeyModels.Task
  for _, task := range experience.Tasks {
    tasks = append(tasks, journeyModels.Task{
      Details: task.Details,
      ExperienceId: uint(experience.ID),
    })
  }

	service.taskServie.InsertTasks(tasks)
	return service.repository.GetJourney(experience.ID)
}
