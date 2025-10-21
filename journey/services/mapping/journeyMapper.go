package journeyMapping

import journeyModels "portfolioAPI/journey/models"

func MapJourneyDisplayModelToJourney(journeyUpsertModel *journeyModels.JourneyUpsertModel) []journeyModels.Experience {
	var journeys []journeyModels.Experience
	var tasks []journeyModels.Task

	for _, translation := range journeyUpsertModel.JourneyTransaltions {
		for _, task := range translation.Tasks {
			tasks = append(tasks, MapTaskDisplayModelToTask(task))
		}

		experience := journeyModels.Experience{
			Title:            translation.Title,
			From:             journeyUpsertModel.From,
			To:               journeyUpsertModel.To,
			Diploma:          translation.Diploma,
			ExperienceTypeId: uint(journeyUpsertModel.ExperienceType),
			Tasks:            tasks,
			LanguageCode:     translation.LanguageCode,
		}

		journeys = append(journeys, experience)
	}

	return journeys
}

func MapTaskDisplayModelToTask(taskDisplay journeyModels.TaskDisplay) journeyModels.Task {
	return journeyModels.Task{
		Details: taskDisplay.Details,
	}
}

func MapJourneyDisplayModelToJourneyWithId(journeyDisplay *journeyModels.JourneyDisplay, experienceId int) journeyModels.Experience {
	var tasks []journeyModels.Task
	for _, task := range journeyDisplay.Tasks {
		tasks = append(tasks, MapTaskDisplayModelToTaskWithId(task, experienceId))
	}
	return journeyModels.Experience{
		Title:            journeyDisplay.Title,
		From:             journeyDisplay.From,
		To:               journeyDisplay.To,
		Diploma:          journeyDisplay.Diploma,
		ExperienceTypeId: uint(journeyDisplay.ExperienceType),
		Tasks:            tasks,
		ID:               experienceId,
	}
}

func MapTaskDisplayModelToTaskWithId(taskDisplay journeyModels.TaskDisplay, experienceId int) journeyModels.Task {
	return journeyModels.Task{
		Details:      taskDisplay.Details,
		ExperienceId: uint(experienceId),
	}
}
