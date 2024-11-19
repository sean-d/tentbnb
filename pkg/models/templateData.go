package models

// TemplateData is used for passing data to templates from handlers
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int64
	FloatMap   map[string]float64
	Data       map[string]any
	CSRFToken  string
	FlashMsg   string
	WarningMsg string
	ErrorMsg   string
}
