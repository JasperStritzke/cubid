package model

type ServerGroup struct {
	Name     string                  `json:"name"`
	Template ServerGroupTemplateData `json:"template"`
}

type ServerGroupTemplateData struct {
	TemplateGroup string `json:"templateGroup"`
	TemplateName  string `json:"templateName"`
}
