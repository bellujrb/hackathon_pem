package programs_service

import (
	"errors"
	"perna/db"
	interfaces "perna/internal/swagger"

	"gorm.io/gorm"
)

type ProgramsService struct {
	DB *gorm.DB
}

func NewProgramsService(db *gorm.DB) *ProgramsService {
	return &ProgramsService{DB: db}
}

func (s *ProgramsService) CreateProgram(data *interfaces.Program) error {
	dbProgram := db.Program{
		ProgramName:  data.ProgramName,
		Description:  data.Description,
		WebsiteURL:   data.WebsiteURL,
		ContactEmail: data.ContactEmail,
		ContactPhone: data.ContactPhone,
	}

	if err := s.DB.Create(&dbProgram).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProgramsService) GetAllPrograms() ([]db.Program, error) {
	var records []db.Program
	if err := s.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (s *ProgramsService) GetProgramByID(id int) (*db.Program, error) {
	var record db.Program
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func (s *ProgramsService) UpdateProgram(id int, data *interfaces.Program) error {
	var record db.Program
	if err := s.DB.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	dbUpdate := db.Program{
		ProgramName:  data.ProgramName,
		Description:  data.Description,
		WebsiteURL:   data.WebsiteURL,
		ContactEmail: data.ContactEmail,
		ContactPhone: data.ContactPhone,
	}

	if err := s.DB.Model(&record).Updates(dbUpdate).Error; err != nil {
		return err
	}

	return nil
}

func (s *ProgramsService) DeleteProgram(id int) error {
	if err := s.DB.Delete(&db.Program{}, id).Error; err != nil {
		return err
	}
	return nil
}
