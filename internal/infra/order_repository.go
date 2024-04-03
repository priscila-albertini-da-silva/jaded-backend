package infra

import (
	"github.com/priscila-albertini-silva/jaded-backend/internal/models"
	"github.com/priscila-albertini-silva/jaded-backend/internal/ui/schemas"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderRepositoryAdapter interface{}

type OrderRepository struct{ db *gorm.DB }

func NewOrderRepository(db *gorm.DB) OrderRepositoryAdapter {
	return OrderRepository{db: db}
}

func (r OrderRepository) FindOrders(filter schemas.OrderFilter) ([]models.Order, error) {
	var Orders []models.Order

	query := r.db.Model(&models.Order{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", filter.Name)
	}
	if filter.Code != "" {
		query = query.Where("code LIKE ?", filter.Code)
	}

	result := query.Find(&Orders)
	if result.Error != nil {
		log.Error("Error while querying Orders:", result.Error)

		return nil, result.Error
	}

	return Orders, nil
}

func (r OrderRepository) FindOrderByID(id int64) (*models.Order, error) {
	var Order models.Order

	result := r.db.First(&Order, id)
	if result.Error != nil {
		log.Error("Error getting Order by id:", result.Error)

		return nil, result.Error
	}

	return &Order, nil
}

func (r OrderRepository) Create(Order models.Order) (*models.Order, error) {
	tx := r.db.Begin()

	result := r.db.Create(&Order)
	if result.Error != nil {
		log.Error("Error creating Order: ", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &Order, nil
}

func (r OrderRepository) Update(Order models.Order) (*models.Order, error) {
	tx := r.db.Begin()

	result := r.db.Save(&Order)

	if result.Error != nil {
		log.Error("Error updating Order:", result.Error)

		tx.Rollback()

		return nil, result.Error
	}

	tx.Commit()

	return &Order, nil
}

func (r OrderRepository) Delete(Order models.Order) error {

	result := r.db.Delete(&Order)

	if result.Error != nil {
		log.Error("Error deleting Order:", result.Error)

		return result.Error
	}

	return nil
}
