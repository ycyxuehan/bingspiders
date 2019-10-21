package types

import(
	"encoding/json"
)

type Page struct {
	ID int `json:"ID" mysql:"id"`
	Books []Book `json:"Books" mysql:"books"`
	Total int `json:"Total" mysql:"total"`
}

func NewPage(){

}

func (p *Page)ToJSON()string{
	d, _ := json.MarshalIndent(p, "", "\t")
	return string(d)
}