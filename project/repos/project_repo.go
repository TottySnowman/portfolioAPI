package project_repo

import (
	"errors"
	"os"
	"portfolioAPI/database"
	projectModel "portfolioAPI/project/models"
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

func (repo *Project_Repo) GetAllProjects() []projectModel.ProjectDisplay {
	var selectedProjects []projectModel.ProjectDataSelect

	result := repo.db.Select("Status, p.ID as ProjectID, p.Name, p.About, p.GithubLink, p.DemoLink, p.LogoPath, t.Tag, t.Icon as TagIcon, p.DevDate").Table("Project as p").
		Joins("Inner join ProjectStatus as ps ON ps.ID = p.ProjectStatusID").
		Joins("inner join Project_Tags as pt ON p.ID = pt.ProjectID").
		Joins("inner join Tag as t ON t.ID = pt.TagID").
		Where("Hidden = false").Order("p.DevDate ASC").Find(&selectedProjects)

	if result.Error != nil {
		return nil
	}

	projects := mapDataRowsToProjects(selectedProjects)

	sortProjectsByDate(projects)

	return projects
}

func mapDataRowsToProjects(projects []projectModel.ProjectDataSelect) []projectModel.ProjectDisplay {

	apiURL := os.Getenv("API_ENDPOINT_URL")

	projectMap := make(map[int]*projectModel.ProjectDisplay)
	for _, project := range projects {
		_, projectExists := projectMap[project.ProjectID]

		if !projectExists {
			projectMap[project.ProjectID] = &projectModel.ProjectDisplay{
				ProjectID:   project.ProjectID,
				Name:        project.Name,
				About:       project.About,
				Status:      project.Status,
				Github_Link: project.GithubLink,
				Demo_Link:   project.DemoLink,
				Logo_Path:   apiURL + project.LogoPath,
				Tags:        []tagModel.JsonTag{},
				DevDate:     project.DevDate,
			}
		}

		projectMap[project.ProjectID].Tags = append(projectMap[project.ProjectID].Tags, tagModel.JsonTag{
			TagIcon: project.TagIcon,
			Tag:     project.Tag,
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

func (repo *Project_Repo) Insert(projectToCreate *projectModel.Project) error {
	result := repo.db.Create(&projectToCreate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *Project_Repo) Update(projectToUpdate *projectModel.Project) error {
	var dbProject = projectModel.Project{Model: gorm.Model{ID: projectToUpdate.ID}}
	existingProject := repo.db.First(&dbProject)

	if existingProject.Error != nil {
		return errors.New("Project not found")
	}

	updateProject := repo.db.Model(&dbProject).Select("*").Omit("CreatedAt").Updates(projectToUpdate)

	if updateProject.Error != nil {
		return errors.New(updateProject.Error.Error())
	}

	return nil
}

func (repo *Project_Repo) Delete(projectID int) error {

	var dbProject = projectModel.Project{Model: gorm.Model{ID: uint(projectID)}}
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
