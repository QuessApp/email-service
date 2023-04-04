package entities

// Response is a model for each response in app.
type Response struct {
	Ok      bool `json:"ok"`
	Error   bool `json:"error"`
	Message any  `json:"message"`
	Data    any  `json:"data"`
}
