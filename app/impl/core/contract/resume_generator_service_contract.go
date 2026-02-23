package contract

type ResumeGeneratorServiceContract interface {
	GenerateBusinessResume(data string) (text string, err error)
}
