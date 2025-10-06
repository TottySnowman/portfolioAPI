package journeyMapping

import journeyModels "portfolioAPI/journey/models"

func MapJourneyDisplayModelToJourney(journeyDisplay journeyModels.JourneyDisplay) journeyModels.Experience{
  var tasks []journeyModels.Task
  for _, task := range journeyDisplay.Tasks {
    tasks = append(tasks, MapTaskDisplayModelToTask(task))
  }
  return journeyModels.Experience{
    Title: journeyDisplay.Title,
    From: journeyDisplay.From,
    To: journeyDisplay.To,
    Diploma: journeyDisplay.Diploma,
    ExperienceTypeId: uint(journeyDisplay.ExperienceType),
    Tasks: tasks,
  }
}

func MapTaskDisplayModelToTask(taskDisplay journeyModels.TaskDisplay) journeyModels.Task{
  return journeyModels.Task{
    Details: taskDisplay.Details,
  }
}
