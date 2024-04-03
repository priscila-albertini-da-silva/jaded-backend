package infra

import (
	"github.com/priscila-albertini-silva/jaded-backend/internal/models"
	"github.com/priscila-albertini-silva/jaded-backend/internal/schemas"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DividendRepositoryAdapter interface{}

type DividendRepository struct{ db *gorm.DB }

func NewDividendRepository(db *gorm.DB) DividendRepositoryAdapter {
	return DividendRepository{db: db}
}

func (r DividendRepository) FindDividends(filter schemas.DividendFilter) ([]models.Dividend, error) {
	var Dividends []models.Dividend

	query := r.db.Model(&models.Dividend{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", filter.Name)
	}
	if filter.Code != "" {
		query = query.Where("code LIKE ?", filter.Code)
	}

	result := query.Find(&Dividends)
	if result.Error != nil {
		log.Error("Error while querying Dividends:", result.Error)

		return nil, result.Error
	}

	return Dividends, nil
}

func (r DividendRepository) FindDividendByID(id int64) (*models.Dividend, error) {
	var Dividend models.Dividend

	result := r.db.First(&Dividend, id)
	if result.Error != nil {
		log.Error("Error getting Dividend by id:", result.Error)

		return nil, result.Error
	}

	return &Dividend, nil
}

func (r DividendRepository) Create(Dividend models.Dividend) (*models.Dividend, error) {
	tx := r.db.Begin()

	result := r.db.Create(&Dividend)
	if result.Error != nil {
		log.Error("Error creating Dividend: ", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &Dividend, nil
}

func (r DividendRepository) Update(Dividend models.Dividend) (*models.Dividend, error) {
	tx := r.db.Begin()

	result := r.db.Save(&Dividend)

	if result.Error != nil {
		log.Error("Error updating Dividend:", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &Dividend, nil
}

func (r DividendRepository) Delete(Dividend models.Dividend) error {

	result := r.db.Delete(&Dividend)

	if result.Error != nil {
		log.Error("Error deleting Dividend:", result.Error)

		return result.Error
	}

	return nil
}
