package journeyRepo

import (
	"errors"
	"portfolioAPI/database"
	journeyModels "portfolioAPI/journey/models"

	"gorm.io/gorm"
)

type JourneyRepo struct {
	db *gorm.DB
}

func NewJourneyRepo() *JourneyRepo {
	return &JourneyRepo{
		db: database.GetDBClient(),
	}
}

func (repo *JourneyRepo) GetFullJourney() []journeyModels.JourneyDisplay {
	var selectedJourney []journeyModels.ExperienceSelect

	query := repo.db.
		Model(&journeyModels.Experience{}).
		Select("Experience.ID as ExperienceId, Title, Company, `From`, `To`, Diploma, et.Name, t.Details, Experience.ExperienceTypeId").
		Joins("INNER JOIN ExperienceType as et on et.ID = Experience.ExperienceTypeId").
		Joins("LEFT JOIN Task as t on t.ExperienceId = Experience.ID")

	result := query.Find(&selectedJourney)
	if result.Error != nil {
		return nil
	}

	return mapDataRowToExperiences(selectedJourney)
}

func (repo *JourneyRepo) Insert(experienceToCreate *journeyModels.Experience) (*journeyModels.Experience, error) {
	result := repo.db.Create(&experienceToCreate)

	if result.Error != nil {
		return nil, result.Error
	}

	return *&experienceToCreate, nil
}

func (repo *JourneyRepo) GetJourney(experienceId int) (*journeyModels.JourneyDisplay, error) {
	var selectedJourney []journeyModels.ExperienceSelect

	query := repo.db.
		Model(&journeyModels.Experience{}).
		Select("Experience.ID as ExperienceId, Title, Company, `From`, `To`, Diploma, et.Name, t.Details, Experience.ExperienceTypeId").
		Joins("INNER JOIN ExperienceType as et on et.ID = Experience.ExperienceTypeId").
		Joins("LEFT JOIN Task as t on t.ExperienceId = Experience.ID").
		Where("Experience.ID = ?", experienceId)

	result := query.Find(&selectedJourney)

	if result.Error != nil {
		return nil, result.Error
	}

	experiences := mapDataRowToExperiences(selectedJourney)

	return &experiences[0], nil
}

func (repo *JourneyRepo) Update() *journeyModels.Experience {
	return nil
}

func mapDataRowToExperiences(journeys []journeyModels.ExperienceSelect) []journeyModels.JourneyDisplay {
	experienceMap := make(map[int]*journeyModels.JourneyDisplay)

	for _, experience := range journeys {
		println(experience.ExperienceId)
		_, experienceExists := experienceMap[int(experience.ExperienceId)]

		if !experienceExists {
			experienceMap[int(experience.ExperienceId)] = &journeyModels.JourneyDisplay{
				Id:             int(experience.ExperienceId),
				Title:          experience.Title,
				Diploma:        experience.Diploma,
				From:           experience.From,
				To:             experience.To,
				ExperienceType: journeyModels.ExperienceTypeEnum(experience.ExperienceTypeId),
			}
		}

		if experience.Details != "" {
			existingExperience, _ := experienceMap[int(experience.ExperienceId)]
			existingExperience.Tasks = append(experienceMap[existingExperience.Id].Tasks, journeyModels.TaskDisplay{
				Details: experience.Details,
			})
		}
	}

	var journeyDisplay []journeyModels.JourneyDisplay
	for _, journey := range experienceMap {
		journeyDisplay = append(journeyDisplay, *journey)
	}

	return journeyDisplay

}

func (repo *JourneyRepo) Delete(journeyId int) error {
	dbJourney := journeyModels.Experience{ID: journeyId}

	existingExperience := repo.db.First(&dbJourney)

	if existingExperience.Error != nil {
		return errors.New("Experience not found")
	}

	result := repo.db.Delete(&dbJourney)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
