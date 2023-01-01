package entity

type State struct {
	Username    string `json:"username" validate:"required"`
	Folder      string `json:"folder" validate:"required"`
	FileName    string `json:"file_name" validate:"required"`
	VirtualName string `json:"virtual_name" validate:"required"`
	FileSize    int    `json:"file_size" validate:"required"`
	State       int    `json:"state" validate:"required,eq=200|eq=300|eq=400"`
}
