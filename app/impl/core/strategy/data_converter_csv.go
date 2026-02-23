package strategy

import (
	"encoding/csv"
	"gitcrawler/app/impl/core/entity"
	"os"
	"strings"
)

type ConverterCsv struct{}

func NewConverterCsv() *ConverterCsv {
	return &ConverterCsv{}
}

func (c *ConverterCsv) Convert(data *entity.RepositoryData) (err error) {

	w, _ := c.createPath()
	writer := csv.NewWriter(w)
	defer writer.Flush()

	headers := []string{"Name", "Data", "Path"}

	writer.Write(headers)

	for i := 0; i < len(data.Files); i++ {
		trimmedData := strings.ReplaceAll(data.Files[i].Data, "\n", "\\n")
		record := []string{
			data.Name,
			trimmedData,
			data.Files[i].Path,
		}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
func (c *ConverterCsv) createPath() (w *os.File, err error) {
	homeDir, _ := os.UserHomeDir()

	dirPath := homeDir + "/Downloads"

	os.MkdirAll(dirPath, os.ModePerm)

	filePath := dirPath + "/output.csv"

	w, err = os.Create(filePath)
	if err != nil {
		return w, err
	}
	return w, nil
}
