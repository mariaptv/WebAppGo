package models

type TemplateData struct{
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	//We are not sure what the data will be, so we use interface{}
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}