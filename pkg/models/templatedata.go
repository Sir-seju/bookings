package models

// A struct to hold Template Data
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
  	Data      map[string]interface{}
	CRSFToken string
	Error     string
	Flash     string
	Warning   string
}