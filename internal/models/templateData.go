package models

// TemplateData host data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // when you are not sure what kind of type it can be, use interface
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
