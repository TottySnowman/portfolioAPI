package chatService

import (
	"fmt"
	chatModel "portfolioAPI/chat/models"
	vectorRepo "portfolioAPI/chat/repos"
	projectModel "portfolioAPI/project/models"
	projectService "portfolioAPI/project/services"
	tagModel "portfolioAPI/tag/models"
	"strings"
)

type VectorService struct {
	vectorRepo       *vectorRepo.VectorRepo
	projectService   *projectService.ProjectService
	embeddingService *EmbeddingService
}

func NewVectorService(vectorRepo *vectorRepo.VectorRepo,
	projectService *projectService.ProjectService,
	embeddingService *EmbeddingService) *VectorService {
	if !vectorRepo.DoesCollectionExist() {
		vectorRepo.CreateCollection()
	}

	return &VectorService{
		vectorRepo:       vectorRepo,
		projectService:   projectService,
		embeddingService: embeddingService,
	}
}

func (service *VectorService) ResetDatabase(syncModel *chatModel.SyncModel) error {
	if !syncModel.ResetProject && !syncModel.ResetPersonal {
		return nil
	}

	if syncModel.ResetProject && syncModel.ResetPersonal {
    err := service.vectorRepo.FullResetDatabase()
    if err != nil{
      return err
    }

    service.vectorRepo.CreateCollection()
	}

	if syncModel.ResetProject {
		return service.vectorRepo.ResetProject()
	}

	if syncModel.ResetPersonal {
		return service.vectorRepo.ResetPersonal()
	}

	return nil
}

func (service *VectorService) InsertProjectsAsync() {
	go func() {
		projects := service.projectService.GetAllProjects(false)
		for _, project := range projects {
			go service.insertProject(project)
		}
	}()
}

func (service *VectorService) insertProject(project projectModel.ProjectDisplay) {
	tags := extractTags(project.Tags)
	embeddingInput := fmt.Sprintf("%s %s %s", project.Name, project.About, strings.Join(tags, " "))
  vector, err := service.embeddingService.GetVectorByText(embeddingInput)
  if err != nil{
    // TODO logging
  }

  modifyProjectModel := &chatModel.ModifyProjectModel{
    ProjectPayload: project,
    Vector: vector,
    ProjectTags: tags,
  }

  service.vectorRepo.UpsertProject(*modifyProjectModel)
}

func extractTags(tags []tagModel.JsonTag) []string {
	concatTags := make([]string, 0)

	for _, tag := range tags {
		concatTags = append(concatTags, tag.Tag)
	}

	return concatTags
}
