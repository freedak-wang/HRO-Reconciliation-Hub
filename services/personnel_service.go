package services

import (
	"time"
	"your-project/models"
)

type PersonnelService struct {
	db *gorm.DB
}

func NewPersonnelService(db *gorm.DB) *PersonnelService {
	return &PersonnelService{db: db}
}

func (s *PersonnelService) AddPersonnel(personnel *models.Personnel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(personnel).Error; err != nil {
			return err
		}

		change := models.PersonnelChange{
			PersonnelID: personnel.ID,
			ContractID:  personnel.ContractID,
			ChangeType:  "add",
			ChangeDate:  personnel.StartDate,
			Reason:      "New personnel added",
		}

		return tx.Create(&change).Error
	})
}

func (s *PersonnelService) RemovePersonnel(personnelID uint, endDate time.Time, reason string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		personnel := models.Personnel{}
		if err := tx.First(&personnel, personnelID).Error; err != nil {
			return err
		}

		personnel.EndDate = &endDate
		personnel.Status = "inactive"
		if err := tx.Save(&personnel).Error; err != nil {
			return err
		}

		change := models.PersonnelChange{
			PersonnelID: personnelID,
			ContractID:  personnel.ContractID,
			ChangeType:  "remove",
			ChangeDate:  endDate,
			Reason:      reason,
		}

		return tx.Create(&change).Error
	})
}

func (s *PersonnelService) ListPersonnel(contractID uint) ([]models.Personnel, error) {
	var personnel []models.Personnel
	err := s.db.Where("contract_id = ?", contractID).Find(&personnel).Error
	return personnel, err
}

func (s *PersonnelService) GetPersonnelChanges(contractID uint, startDate, endDate time.Time) ([]models.PersonnelChange, error) {
	var changes []models.PersonnelChange
	err := s.db.Where("contract_id = ? AND change_date BETWEEN ? AND ?", contractID, startDate, endDate).
		Order("change_date").
		Find(&changes).Error
	return changes, err
}