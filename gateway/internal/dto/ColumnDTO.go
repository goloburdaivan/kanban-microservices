package dto

type ColumnDTO struct {
	ID     int        `json:"ID"`
	Title  string     `json:"title"`
	Order  int        `json:"order"`
	IsDone bool       `json:"is_done"`
	Tasks  *[]TaskDTO `json:"tasks"`
}
