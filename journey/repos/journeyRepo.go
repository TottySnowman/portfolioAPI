package journeyRepo

import (
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

func (repo *JourneyRepo) GetFullJourney() {

}

func (repo *JourneyRepo) Insert() *journeyModels.Experience {
	return nil
}

func (repo *JourneyRepo) Update() *journeyModels.Experience {
	return nil
}

func (repo *JourneyRepo) Delete() {

}
