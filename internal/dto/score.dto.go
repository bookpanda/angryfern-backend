package dto

type IncrementScoreRequest struct {
	Code   string `json:"code" example:"TH"`
	Amount uint   `json:"amount" example:"1"`
}

type IncrementScoreResponse struct {
	Success bool `json:"success"`
}

type GetAllScoreResponse struct {
	Scores []Score `json:"scores"`
}

type Score struct {
	Code        string `example:"TH" json:"code"`
	CountryName string `example:"Thailand" json:"country_name"`
	ClickCount  uint   `example:"100" json:"click_count"`
}
