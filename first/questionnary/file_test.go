package questionnary

import (
	"strings"
	"testing"
)

func TestParseFile(t *testing.T) {
	t.Run("Testing parseFile", func(t *testing.T) {

		linesNum := 10
		lines := ""
		for index := 0; index < linesNum; index++ {
			lines += "5+5,10\n"
		}

		fileContent := strings.NewReader(lines)
		questionnary := parseFile(fileContent)

		if len(questionnary.Questions) != linesNum {
			t.Errorf("Wanted %v, got %v", linesNum, len(questionnary.Questions))
		}
	})
}
