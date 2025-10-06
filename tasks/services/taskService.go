package taskService

import (
	journeyModels "portfolioAPI/journey/models"
	taskRepo "portfolioAPI/tasks/repos"
)

type TaskService struct {
	repo taskRepo.TaskRepo
}

func NewTaskService(taskRepo *taskRepo.TaskRepo) *TaskService {
	return &TaskService{
		repo: *taskRepo,
	}
}

func(taskService *TaskService) InsertTasks(tasks []journeyModels.Task) error{
  for _, task := range tasks {
    err := taskService.repo.Insert(task)
    if err != nil{
      return err
    }
  }

  return nil
}

func(taskService *TaskService) DeleteTasks(experienceId uint) error{
  return taskService.DeleteTasks(experienceId)
}
