package register_service

import (
	"errors"
	"perna/db"
	interfaces "perna/internal/swagger"

	"gorm.io/gorm"
)

type RegisterService struct {
	DB *gorm.DB
}

func NewRegisterService(db *gorm.DB) *RegisterService {
	return &RegisterService{DB: db}
}

func (s *RegisterService) CreateRegister(data *interfaces.DomesticViolenceBO) error {
	dbRecord := db.DomesticViolenceBO{
		ReportDetails:   data.ReportDetails,
		VictimName:      data.VictimName,
		VictimContact:   data.VictimContact,
		OffenderName:    data.OffenderName,
		IncidentAddress: data.IncidentAddress,
		DateOfIncident:  data.DateOfIncident,
		PoliceStation:   data.PoliceStation,
		Status:          data.Status,
	}

	if err := s.DB.Create(&dbRecord).Error; err != nil {
		return err
	}
	return nil
}

func (s *RegisterService) GetAllRegisters() ([]db.DomesticViolenceBO, error) {
	var records []db.DomesticViolenceBO
	if err := s.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (s *RegisterService) GetRegisterByID(id int) (*db.DomesticViolenceBO, error) {
	var record db.DomesticViolenceBO
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (s *RegisterService) UpdateRegister(id int, data *interfaces.DomesticViolenceBO) error {
	var record db.DomesticViolenceBO
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	dbUpdate := db.DomesticViolenceBO{
		ReportDetails:   data.ReportDetails,
		VictimName:      data.VictimName,
		VictimContact:   data.VictimContact,
		OffenderName:    data.OffenderName,
		IncidentAddress: data.IncidentAddress,
		DateOfIncident:  data.DateOfIncident,
		PoliceStation:   data.PoliceStation,
		Status:          data.Status,
	}

	if err := s.DB.Model(&record).Updates(dbUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (s *RegisterService) DeleteRegister(id int) error {
	if err := s.DB.Delete(&db.DomesticViolenceBO{}, id).Error; err != nil {
		return err
	}
	return nil
}
