package models

type Page struct {
	Name string `json:"name"`
	Lang string `json:"lang"`
	Body string `json:"body"`
}

func NewPage(name, lang, body string) *Page {
	return &Page{
		Name: name,
		Lang: lang,
		Body: body,
	}
}
