package project

type UpdateProject struct {
	Path        *string
	Usage       *int
	Last_use_at *string
	Directories *[]Directory
}
