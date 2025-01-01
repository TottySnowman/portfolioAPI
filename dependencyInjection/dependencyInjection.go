package dependencyinjection

import (
	fileController "portfolioAPI/fileUpload/controllers"
	fileHandler "portfolioAPI/fileUpload/handler"
	fileServices "portfolioAPI/fileUpload/services"
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
	FileController    *fileController.FileController
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
	fileService    *fileServices.FileService
}

func NewAppContainer() *AppContainer {
	repos := getRepos()
	services := getServices(repos)

	return &AppContainer{
		ProjectController: projectController.NewProjectController(services.projectService),
		TagController:     tagController.NewTagController(services.tagService),
		StatusController:  statusController.NewStatusController(services.statusService),
		FileController:    fileController.NewFileController(services.fileService),
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
	uploader := &fileHandler.FileUploadHandler{}
	deleter := &fileHandler.FileDeleteHandler{}

	fileService := fileServices.NewFileService(uploader, deleter)
	return services{
		projectService: projectService.NewProjectService(repos.projectRepo, tagService, fileService),
		tagService:     tagService,
		statusService:  statusService.NewStatusService(repos.statusRepo),
		fileService:    fileService,
	}
}
