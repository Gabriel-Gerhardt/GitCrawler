package entity

import "strings"

type RepositoryData struct {
	Files []*RepositoryFile
	Name  string
}

func (r *RepositoryData) String() string {
	var filesData []string
	for _, file := range r.Files {
		filesData = append(filesData, file.String())
	}

	return r.Name + " [" + strings.Join(filesData, ", ") + "]"

}
