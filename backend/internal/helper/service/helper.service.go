package helper_service

import (
	"errors"
	"perna/db"
	interfaces "perna/internal/swagger"

	"gorm.io/gorm"
)

type HelperService struct {
	DB *gorm.DB
}

func NewHelperService(db *gorm.DB) *HelperService {
	return &HelperService{DB: db}
}

func (s *HelperService) CreateHelper(data *interfaces.EmergencyRequest) error {
	dbEmergencyRequest := db.EmergencyRequest{
		VictimName:    data.VictimName,
		VictimContact: data.VictimContact,
		Location:      data.Location,
	}

	if err := s.DB.Create(&dbEmergencyRequest).Error; err != nil {
		return err
	}
	return nil
}

func (s *HelperService) GetAllHelpers() ([]db.EmergencyRequest, error) {
	var records []db.EmergencyRequest
	if err := s.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (s *HelperService) GetHelperByID(id int) (*db.EmergencyRequest, error) {
	var record db.EmergencyRequest
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (s *HelperService) UpdateHelper(id int, data *interfaces.EmergencyRequest) error {
	var record db.EmergencyRequest
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	dbUpdate := db.EmergencyRequest{
		VictimName:    data.VictimName,
		VictimContact: data.VictimContact,
		Location:      data.Location,
	}

	if err := s.DB.Model(&record).Updates(dbUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (s *HelperService) DeleteHelper(id int) error {
	if err := s.DB.Delete(&db.EmergencyRequest{}, id).Error; err != nil {
		return err
	}
	return nil
}
