package dependencyinjection

import (
	projectController "portfolioAPI/project/controllers"
	project_repo "portfolioAPI/project/repos"
	projectService "portfolioAPI/project/services"
	statusController "portfolioAPI/status/controllers"
	statusRepo "portfolioAPI/status/repos"
	statusService "portfolioAPI/status/services"
	tagController "portfolioAPI/tag/controllers"
	tagRepo "portfolioAPI/tag/repos"
	tagService "portfolioAPI/tag/services"
)

type AppContainer struct {
	ProjectController *projectController.ProjectController
	TagController     *tagController.TagController
	StatusController  *statusController.StatusController
}

type repos struct {
	projectRepo *project_repo.Project_Repo
	tagRepo     *tagRepo.TagRepo
	statusRepo  *statusRepo.StatusRepo
}

type services struct {
	projectService *projectService.ProjectService
	tagService     *tagService.TagService
	statusService  *statusService.StatusService
}

func NewAppContainer() *AppContainer {
	repos := getRepos()
	services := getServices(repos)

	return &AppContainer{
		ProjectController: projectController.NewProjectController(services.projectService),
		TagController:     tagController.NewTagController(services.tagService),
		StatusController:  statusController.NewStatusController(services.statusService),
	}
}

func getRepos() repos {
	return repos{
		projectRepo: project_repo.NewProjectRepo(),
		tagRepo:     tagRepo.NewTagRepo(),
		statusRepo:  statusRepo.NewStatusRepo(),
	}
}

func getServices(repos repos) services {
	tagService := tagService.NewTagService(repos.tagRepo)

	return services{
		projectService: projectService.NewProjectService(repos.projectRepo, tagService),
		tagService:     tagService,
		statusService:  statusService.NewStatusService(repos.statusRepo),
	}
}
