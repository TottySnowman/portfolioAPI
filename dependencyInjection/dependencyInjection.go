package dependencyinjection

import (
	"portfolioAPI/apiClients"
	chatController "portfolioAPI/chat/controllers"
	vectorRepo "portfolioAPI/chat/repos"
	chatService "portfolioAPI/chat/services"
	contactController "portfolioAPI/contact/controllers"
	contactService "portfolioAPI/contact/services"
	fileController "portfolioAPI/fileUpload/controllers"
	fileHandler "portfolioAPI/fileUpload/handler"
	fileServices "portfolioAPI/fileUpload/services"
	journeyController "portfolioAPI/journey/controllers"
	journeyRepo "portfolioAPI/journey/repos"
	journeyService "portfolioAPI/journey/services"
	knowledgeController "portfolioAPI/knowledge/controllers"
	projectController "portfolioAPI/project/controllers"
	project_repo "portfolioAPI/project/repos"
	projectService "portfolioAPI/project/services"
	statusController "portfolioAPI/status/controllers"
	statusRepo "portfolioAPI/status/repos"
	statusService "portfolioAPI/status/services"
	tagController "portfolioAPI/tag/controllers"
	tagRepo "portfolioAPI/tag/repos"
	tagService "portfolioAPI/tag/services"
	taskRepo "portfolioAPI/tasks/repos"
	taskService "portfolioAPI/tasks/services"
)

type AppContainer struct {
	ProjectController   *projectController.ProjectController
	TagController       *tagController.TagController
	StatusController    *statusController.StatusController
	FileController      *fileController.FileController
	ChatController      *chatController.ChatController
	KnowledgeController *knowledgeController.KnowledgeController
	ContactController   *contactController.ContactController
	JourneyController   *journeyController.JourneyController
}

type repos struct {
	projectRepo *project_repo.Project_Repo
	tagRepo     *tagRepo.TagRepo
	statusRepo  *statusRepo.StatusRepo
	vectorRepo  *vectorRepo.VectorRepo
	journeyRepo *journeyRepo.JourneyRepo
	taskRepo    *taskRepo.TaskRepo
}

type services struct {
	projectService   *projectService.ProjectService
	tagService       *tagService.TagService
	statusService    *statusService.StatusService
	fileService      *fileServices.FileService
	embeddingService *chatService.EmbeddingService
	vectorService    *chatService.VectorService
	wsService        *chatService.WsService
	contactService   *contactService.ContactService
	journeyService   *journeyService.JourneyService
	taskService      *taskService.TaskService
}

func NewAppContainer() *AppContainer {
	repos := getRepos()
	services := getServices(repos)

	return &AppContainer{
		ProjectController:   projectController.NewProjectController(services.projectService),
		TagController:       tagController.NewTagController(services.tagService),
		StatusController:    statusController.NewStatusController(services.statusService),
		FileController:      fileController.NewFileController(services.fileService),
		ChatController:      chatController.NewChatController(services.embeddingService, services.vectorService, services.wsService),
		KnowledgeController: knowledgeController.NewKnowledgeController(services.vectorService, services.embeddingService),
		ContactController:   contactController.NewContactController(services.contactService),
		JourneyController:   journeyController.NewJourneyController(services.journeyService),
	}
}

func getRepos() repos {
	return repos{
		projectRepo: project_repo.NewProjectRepo(),
		tagRepo:     tagRepo.NewTagRepo(),
		statusRepo:  statusRepo.NewStatusRepo(),
		journeyRepo: journeyRepo.NewJourneyRepo(),
		taskRepo:    taskRepo.NewTaskRepo(),
	}
}

func getServices(repos repos) services {
	uploader := &fileHandler.FileUploadHandler{}
	deleter := &fileHandler.FileDeleteHandler{}

	fileService := fileServices.NewFileService(uploader, deleter)
	tagService := tagService.NewTagService(repos.tagRepo)
	projectService := projectService.NewProjectService(repos.projectRepo, tagService, fileService)
	embeddingService := chatService.NewEmbeddingService(apiClients.NewHuggingFaceClient())
	responseService := chatService.NewResponseService(apiClients.NewOpenAiClient())
	repos.vectorRepo = vectorRepo.NewVectorRepo(projectService)
	taskService := taskService.NewTaskService(repos.taskRepo)

	return services{
		projectService:   projectService,
		tagService:       tagService,
		statusService:    statusService.NewStatusService(repos.statusRepo),
		fileService:      fileService,
		embeddingService: embeddingService,
		vectorService:    chatService.NewVectorService(repos.vectorRepo, embeddingService, projectService, responseService),
		wsService:        chatService.NewWsService(),
		contactService:   contactService.NewContactService(),
		journeyService:   journeyService.NewJourneyService(repos.journeyRepo, taskService),
	}
}
