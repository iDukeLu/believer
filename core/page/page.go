package page

type Page struct {
	Curr    int         `json:"current"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
	Records interface{} `json:"records"`
}
