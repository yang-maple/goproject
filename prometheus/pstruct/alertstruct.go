package pstruct

type AlertMessage struct {
	Alerts []Alerts `json:"alerts"`
}

type Alerts struct {
	Annotations `json:"annotations"`
	Labels      `json:"labels"`
	Status      string `json:"status"`
}
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type Labels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Ip        string `json:"ip"`
	Job       string `json:"job"`
	Region    string `json:"region"`
	Severity  string `json:"severity"`
	Port      string `json:"port,omitempty"`
}
