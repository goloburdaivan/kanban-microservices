package dto

type BoardDTO struct {
	project *ProjectDTO  `json:"project,omitempty"`
	columns []*ColumnDTO `json:"columns,omitempty"`
}
