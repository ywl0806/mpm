package project

type Project struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Usage       int         `json:"usage"`
	Last_use_at string      `json:"last_use_at"`
	Created_at  string      `json:"created_at"`
	Directories []Directory `json:"directories"`
	DefaultCmd  string      `json:"default_cmd"`
}

type Directory struct {
	Path    string `json:"path"`
	Cmd     string `json:"cmd"`
	Options string `json:"options"`
}
