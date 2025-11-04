package project_repo

import (
	"errors"
	"os"
  "fmt"
	"portfolioAPI/database"
	projectModel "portfolioAPI/project/models"
	statusModels "portfolioAPI/status/models"
	tagModel "portfolioAPI/tag/models"
	"sort"

	"gorm.io/gorm"
)

type Project_Repo struct {
	db *gorm.DB
}

func NewProjectRepo() *Project_Repo {
	return &Project_Repo{
		db: database.GetDBClient(),
	}
}

func (repo *Project_Repo) GetAllProjects(includeHidden bool, languageCode string) []projectModel.ProjectDisplay {
	var selectedProjects []projectModel.ProjectDataSelect

	query := repo.db.Select("ps.Status, ps.ID as StatusID, p.ID as ProjectID, p.Name, p.About, p.GithubLink, p.DemoLink, p.LogoPath, t.Tag, t.Icon as TagIcon, t.ID as TagId, p.DevDate, p.Hidden").Table("Project as p").
		Joins("Inner join ProjectStatus as ps ON ps.ID = p.ProjectStatusID").
		Joins("inner join Project_Tags as pt ON p.ID = pt.ProjectID").
		Joins("inner join Tag as t ON t.ID = pt.TagID").
		Where("p.DeletedAt IS NULL").Order("p.DevDate ASC")

	if !includeHidden {
		query = query.Where("Hidden = false")
	}

  if(languageCode != ""){
		query = query.Where("LanguageCode = ?", languageCode)
  }
	result := query.Find(&selectedProjects)

	if result.Error != nil {
		return nil
	}

	projects := mapDataRowsToProjects(selectedProjects)

	sortProjectsByDate(projects)

	return projects
}

func (repo *Project_Repo) GetProjectById(projectId int, includeHidden bool) (*projectModel.ProjectDisplay, error) {
	var selectedProjects []projectModel.ProjectDataSelect

	query := repo.db.Select("ps.Status, ps.ID as StatusID, p.ID as ProjectID, p.Name, p.About, p.GithubLink, p.DemoLink, p.LogoPath, t.Tag, t.Icon as TagIcon, t.ID as TagId, p.DevDate, p.Hidden").Table("Project as p").
		Joins("Inner join ProjectStatus as ps ON ps.ID = p.ProjectStatusID").
		Joins("inner join Project_Tags as pt ON p.ID = pt.ProjectID").
		Joins("inner join Tag as t ON t.ID = pt.TagID").
		Where("p.DeletedAt IS NULL AND p.ID = ?", projectId).Order("p.DevDate ASC")

	if !includeHidden {
		query = query.Where("Hidden = false")
	}

	result := query.Find(&selectedProjects)

	if result.Error != nil {
		return nil, result.Error
	}

	projects := mapDataRowsToProjects(selectedProjects)

	return &projects[0], nil
}

func mapDataRowsToProjects(projects []projectModel.ProjectDataSelect) []projectModel.ProjectDisplay {
	apiURL := os.Getenv("API_ENDPOINT_URL")
	projectMap := make(map[int]*projectModel.ProjectDisplay)
	for _, project := range projects {
		_, projectExists := projectMap[project.ProjectID]

		if !projectExists {
			projectMap[project.ProjectID] = &projectModel.ProjectDisplay{
				ProjectID: project.ProjectID,
				Name:      project.Name,
				About:     project.About,
				Status: statusModels.ProjectStatusDisplay{
					StatusID: project.StatusID,
					Status:   project.Status,
				},
				Github_Link: project.GithubLink,
				Demo_Link:   project.DemoLink,
				Logo_Path:   apiURL + project.LogoPath,
				Tags:        []tagModel.JsonTag{},
				DevDate:     project.DevDate,
				Hidden:      project.Hidden,
			}
		}

		projectMap[project.ProjectID].Tags = append(projectMap[project.ProjectID].Tags, tagModel.JsonTag{
			Icon:  project.TagIcon,
			Tag:   project.Tag,
			TagId: project.TagId,
		})
	}
	var projectDisplay []projectModel.ProjectDisplay
	for _, project := range projectMap {
		projectDisplay = append(projectDisplay, *project)
	}

	return projectDisplay
}

func sortProjectsByDate(projects []projectModel.ProjectDisplay) {
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].DevDate.Before(projects[j].DevDate)
	})
}

func (repo *Project_Repo) Insert(projectToCreate *projectModel.Project) (*projectModel.Project, error) {
	result := repo.db.Create(&projectToCreate)

	if result.Error != nil {
		return nil, result.Error
	}

	return *&projectToCreate, nil
}

func (repo *Project_Repo) Update(projectToUpdate *projectModel.Project) (*projectModel.Project, error) {
	var dbProject = projectModel.Project{ID: projectToUpdate.ID}
	existingProject := repo.db.First(&dbProject)

	if existingProject.Error != nil {
		return nil, errors.New("Project not found")
	}

  projectToUpdate.LogoPath = dbProject.LogoPath
	updateProject := repo.db.Model(&dbProject).Select("*").Omit("CreatedAt").Updates(projectToUpdate)

	if updateProject.Error != nil {
		return nil, errors.New(updateProject.Error.Error())
	}

	return *&projectToUpdate, nil
}

func (repo *Project_Repo) Delete(projectID int) error {

	var dbProject = projectModel.Project{ID: projectID}
	existingProject := repo.db.First(&dbProject)

	if existingProject.Error != nil {
		return errors.New("Project not found")
	}

	result := repo.db.Delete(&dbProject)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *Project_Repo) InsertIntoProjectTags(projectId int, tagIds []int) error {
	var project projectModel.Project
	repo.db.First(&project, projectId)

	var tags []tagModel.Tag
	if err := repo.db.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
		return fmt.Errorf("failed to find tags with IDs %v: %w", tagIds, err)
	}

	if err := repo.db.Model(&project).Association("Tags").Replace(&tags); err != nil {
		return fmt.Errorf("failed to append tags to project ID %d: %w", project.ID, err)
	}

	return nil
}
