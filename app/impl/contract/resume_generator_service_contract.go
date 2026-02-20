package contract

type ResumeGeneratorServiceContract interface {
	GenerateBusinessResume(data string) (text []byte, err error)
}
