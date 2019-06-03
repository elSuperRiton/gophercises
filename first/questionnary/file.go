package questionnary

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/elSuperRiton/gophercices/first/utils"
)

func LoadFile(filename string) *Questionnary {
	os.Chdir("../")
	f, err := os.Open(filename)
	utils.PanicIfErr(err, fmt.Sprintf("cannot read file %v", filename))

	return parseFile(f)
}

func parseFile(content io.Reader) *Questionnary {

	var questionnary Questionnary

	reader := csv.NewReader(bufio.NewReader(content))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		utils.PanicIfErr(err, fmt.Sprintf("error parsing line %v", line))

		questionnary.Reader = os.Stdin
		questionnary.Questions = append(questionnary.Questions, Question{
			Q: line[0],
			A: line[1],
		})
	}

	return &questionnary
}
