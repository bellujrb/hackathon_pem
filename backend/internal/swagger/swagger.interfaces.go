package interfaces

import "time"

// DomesticViolenceBO representa um registro de boletim de ocorrência de violência doméstica.
// @Description Registro de B.O de Violência Doméstica contra Mulher.
// @property ReportDetails - Detalhes do boletim de ocorrência.
// @property VictimName - Nome da vítima.
// @property VictimContact - Contato da vítima.
// @property OffenderName - Nome do agressor.
// @property IncidentAddress - Local do incidente.
// @property DateOfIncident - Data do incidente.
// @property PoliceStation - Delegacia responsável.
// @property Status - Status do boletim.
type DomesticViolenceBO struct {
	ReportDetails   string    `json:"report_details"`   // Detalhes do boletim de ocorrência
	VictimName      string    `json:"victim_name"`      // Nome da vítima
	VictimContact   string    `json:"victim_contact"`   // Contato da vítima
	OffenderName    string    `json:"offender_name"`    // Nome do agressor
	IncidentAddress string    `json:"incident_address"` // Local do incidente
	DateOfIncident  time.Time `json:"date_of_incident"` // Data do incidente
	PoliceStation   string    `json:"police_station"`   // Delegacia responsável
	Status          string    `json:"status"`           // Status do boletim (ex: Em andamento, Concluído)
}

// Program representa um programa que pode ser vinculado à assistência.
// @Description Programas disponíveis para assistência.
// @property ProgramName - Nome do programa.
// @property Description - Descrição do programa.
// @property WebsiteURL - URL do site do programa.
// @property ContactEmail - Email de contato.
// @property ContactPhone - Telefone de contato.
type Program struct {
	ProgramName  string `json:"program_name"`  // Nome do programa
	Description  string `json:"description"`   // Descrição do programa
	WebsiteURL   string `json:"website_url"`   // URL do site do programa
	ContactEmail string `json:"contact_email"` // Email de contato
	ContactPhone string `json:"contact_phone"` // Telefone de contato
}

// EmergencyRequest representa uma solicitação de socorro em caso de emergência.
// @Description Solicitação de socorro com localização da vítima.
// @property VictimName - Nome da vítima.
// @property VictimContact - Contato da vítima.
// @property Location - Localização da vítima (coordenadas GPS ou endereço).
type EmergencyRequest struct {
	VictimName    string `json:"victim_name"`    // Nome da vítima solicitando socorro
	VictimContact string `json:"victim_contact"` // Contato da vítima
	Location      string `json:"location"`       // Localização geográfica da vítima (ex: coordenadas GPS)
}
