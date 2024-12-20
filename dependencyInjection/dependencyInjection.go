package dependencyinjection

import (
	projectController "portfolioAPI/project/controllers"
	project_repo "portfolioAPI/project/repos"
	projectService "portfolioAPI/project/services"
	tagController "portfolioAPI/tag/controllers"
	tagRepo "portfolioAPI/tag/repos"
	tagService "portfolioAPI/tag/services"
)

type AppContainer struct {
	ProjectController *projectController.ProjectController
	TagController     *tagController.TagController
}

type repos struct {
	projectRepo *project_repo.Project_Repo
	tagRepo     *tagRepo.TagRepo
}

type services struct {
	projectService *projectService.ProjectService
	tagService     *tagService.TagService
}

func NewAppContainer() *AppContainer {
	repos := getRepos()
	services := getServices(repos)

	return &AppContainer{
		ProjectController: projectController.NewProjectController(services.projectService),
    TagController: tagController.NewTagController(services.tagService),
	}
}

func getRepos() repos {
	return repos{
		projectRepo: project_repo.NewProjectRepo(),
		tagRepo:     tagRepo.NewTagRepo(),
	}
}

func getServices(repos repos) services {
  tagService := tagService.NewTagService(repos.tagRepo)

	return services{
		projectService: projectService.NewProjectService(repos.projectRepo, tagService),
    tagService: tagService,
	}
}
