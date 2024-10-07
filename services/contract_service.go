package services

import (
	"time"
	"your-project/models"
)

type ContractService struct {
	db *gorm.DB
}

func NewContractService(db *gorm.DB) *ContractService {
	return &ContractService{db: db}
}

func (s *ContractService) CreateContract(contract *models.Contract) error {
	return s.db.Create(contract).Error
}

func (s *ContractService) GetContract(id uint) (*models.Contract, error) {
	var contract models.Contract
	err := s.db.Preload("ServiceItems").First(&contract, id).Error
	return &contract, err
}

func (s *ContractService) UpdateContract(contract *models.Contract) error {
	return s.db.Save(contract).Error
}

func (s *ContractService) ListContracts(status string) ([]models.Contract, error) {
	var contracts []models.Contract
	query := s.db.Preload("ServiceItems")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&contracts).Error
	return contracts, err
}

func (s *ContractService) TerminateContract(id uint, terminationDate time.Time) error {
	return s.db.Model(&models.Contract{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":   "terminated",
			"end_date": terminationDate,
		}).Error
}