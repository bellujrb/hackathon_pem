package db

import "time"

type DomesticViolenceBO struct {
	BaseModel
	ReportDetails   string    `json:"report_details"`   // Detalhes do boletim de ocorrência
	VictimName      string    `json:"victim_name"`      // Nome da vítima
	VictimContact   string    `json:"victim_contact"`   // Contato da vítima
	OffenderName    string    `json:"offender_name"`    // Nome do agressor
	IncidentAddress string    `json:"incident_address"` // Local do incidente
	DateOfIncident  time.Time `json:"date_of_incident"` // Data do incidente
	PoliceStation   string    `json:"police_station"`   // Delegacia responsável
	Status          string    `json:"status"`           // Status do boletim (ex: Em andamento, Concluído)
}

type Program struct {
	BaseModel
	ProgramName  string `json:"program_name"`  // Nome do programa
	Description  string `json:"description"`   // Descrição do programa
	WebsiteURL   string `json:"website_url"`   // URL do site do programa
	ContactEmail string `json:"contact_email"` // Email de contato
	ContactPhone string `json:"contact_phone"` // Telefone de contato
}

type EmergencyRequest struct {
	BaseModel
	VictimName    string `json:"victim_name"`    // Nome da vítima solicitando socorro
	VictimContact string `json:"victim_contact"` // Contato da vítima
	Location      string `json:"location"`       // Localização geográfica da vítima (ex: coordenadas GPS)
}

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
