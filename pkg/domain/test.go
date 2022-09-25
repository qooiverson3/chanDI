package domain

type Info struct {
	Data []int  `json:"data"`
	Name string `json:"name"`
}

type Service interface {
	GetData() Info
}
