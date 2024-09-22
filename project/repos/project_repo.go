package project_repo

import (
	"portfolioAPI/database"
	projectModel "portfolioAPI/project/models"

  "os"
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
	var projects []projectModel.ProjectDataSelect

	result := repo.db.Select("Status, p.ID as ProjectID, p.Name, p.About, p.GitHubLink, p.DemoLink, p.LogoPath, t.Tag, t.Icon as TagIcon").Table("Project as p").
		Joins("Inner join ProjectStatus as ps ON ps.ID = p.ProjectStatusID").
		Joins("inner join Project_Tags as pt ON p.ID = pt.ProjectID").
		Joins("inner join Tag as t ON t.ID = pt.TagID").
		Where("Hidden = false").Order("DevDate ASC").Find(&projects)

	if result.Error != nil {
		return nil
	}

  apiURL := os.Getenv("API_ENDPOINT_URL") 

	projectMap := make(map[int]*projectModel.ProjectDisplay)
	for _, project := range projects {
    if _, exists := projectMap[project.ProjectID]; !exists{
      projectMap[project.ProjectID] = &projectModel.ProjectDisplay{
        ProjectID: project.ProjectID,
        Name: project.Name,
        About: project.About,
        Status: project.Status,
        Github_Link: project.GithubLink,
        Demo_Link: project.DemoLink,
        Logo_Path: apiURL + project.LogoPath,
        Tags: []projectModel.JsonTag{},
      }
    }

    projectMap[project.ProjectID].Tags = append(projectMap[project.ProjectID].Tags, projectModel.JsonTag{
      TagIcon: project.TagIcon,
      Tag: project.Tag,
    })
	}
  var responseProjects []projectModel.ProjectDisplay
	for _, project := range projectMap {
		responseProjects = append(responseProjects, *project)
	}
	return responseProjects
}
