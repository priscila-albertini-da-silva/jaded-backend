package infra

import (
	"github.com/priscila-albertini-silva/jaded-backend/internal/models"
	"github.com/priscila-albertini-silva/jaded-backend/internal/schemas"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StockRepositoryAdapter interface{}

type StockRepository struct{ db *gorm.DB }

func NewStockRepository(db *gorm.DB) StockRepositoryAdapter {
	return StockRepository{db: db}
}

func (r StockRepository) FindStocks(filter schemas.StockFilter) ([]models.Stock, error) {
	var stocks []models.Stock

	query := r.db.Model(&models.Stock{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", filter.Name)
	}
	if filter.Code != "" {
		query = query.Where("code LIKE ?", filter.Code)
	}

	result := query.Find(&stocks)
	if result.Error != nil {
		log.Error("Error while querying stocks:", result.Error)

		return nil, result.Error
	}

	return stocks, nil
}

func (r StockRepository) FindStockByID(id int64) (*models.Stock, error) {
	var stock models.Stock

	result := r.db.First(&stock, id)
	if result.Error != nil {
		log.Error("Error getting stock by id:", result.Error)

		return nil, result.Error
	}

	return &stock, nil
}

func (r StockRepository) Create(stock models.Stock) (*models.Stock, error) {
	tx := r.db.Begin()

	result := r.db.Create(&stock)
	if result.Error != nil {
		log.Error("Error creating stock: ", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &stock, nil
}

func (r StockRepository) Update(stock models.Stock) (*models.Stock, error) {
	tx := r.db.Begin()

	result := r.db.Save(&stock)

	if result.Error != nil {
		log.Error("Error updating stock:", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &stock, nil
}

func (r StockRepository) Delete(stock models.Stock) error {

	result := r.db.Delete(&stock)

	if result.Error != nil {
		log.Error("Error deleting stock:", result.Error)

		return result.Error
	}

	return nil
}
