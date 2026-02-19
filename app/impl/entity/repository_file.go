package entity

type RepositoryFile struct {
	Data string
	Path string
}

func (r *RepositoryFile) String() string {
	return r.Data + r.Path
}
