package contract

type CloneServiceContract interface {
	CloneRepository(repositoryUrl string) (string, error)
}
