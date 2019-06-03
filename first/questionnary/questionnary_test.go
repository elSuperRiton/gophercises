package questionnary

import (
	"os"
	"strings"
	"testing"
	"time"
)

func getTestQuestionnary() Questionnary {
	return Questionnary{
		Reader: os.Stdin,
		Questions: Questions{
			Question{
				Q: "1 + 1",
				A: "2",
			},
			Question{
				Q: "2 + 1",
				A: "3",
			},
			Question{
				Q: "3 + 1",
				A: "4",
			},
			Question{
				Q: "4 + 8",
				A: "12",
			},
			Question{
				Q: "5 + 2",
				A: "7",
			},
		},
	}
}

func TestQuestionnaryRun(t *testing.T) {

	t.Run("Test questionnary Run without error", func(t *testing.T) {
		testQuestionnary := getTestQuestionnary()
		var userInput string
		for _, q := range testQuestionnary.Questions {
			userInput += q.A + "\n"
		}

		testQuestionnary.Reader = strings.NewReader(userInput)
		testQuestionnary.Run()

		if testQuestionnary.RightAnswers != uint(len(testQuestionnary.Questions)) {
			t.Errorf("Wanted %v, got %v", len(testQuestionnary.Questions), testQuestionnary.RightAnswers)
		}
	})

	t.Run("Test questionnary Run with errors", func(t *testing.T) {
		testQuestionnary := getTestQuestionnary()
		numberOfErrors := 2

		var userInput string
		for index, q := range testQuestionnary.Questions {
			if index < numberOfErrors {
				userInput += "wrong\n"
			} else {
				userInput += q.A + "\n"
			}
		}

		testQuestionnary.Reader = strings.NewReader(userInput)
		testQuestionnary.Run()

		if testQuestionnary.RightAnswers != uint(len(testQuestionnary.Questions)-numberOfErrors) {
			t.Errorf("Wanted %v, got %v", len(testQuestionnary.Questions)-numberOfErrors, testQuestionnary.RightAnswers)
		}
	})

	t.Run("Test questionnary Run with timer set", func(t *testing.T) {

		testQuestionnary := getTestQuestionnary()
		testQuestionnary.SetTimerDuration(5)
		timer := testQuestionnary.getTimer()

		var userInput string
		questionsAnswered := 2

		for index, q := range testQuestionnary.Questions {
			if index+1 > questionsAnswered {
				break
			}
			userInput += q.A + "\n"
		}

		testQuestionnary.Reader = strings.NewReader(userInput)
		// now := time.Now()
		testQuestionnary.Run()

		if testQuestionnary.RightAnswers != uint(questionsAnswered) {
			t.Errorf("Wanted %v, got %v right answers", questionsAnswered, testQuestionnary.RightAnswers)
		}

		after := <-timer
		if after.Sub(time.Now()).Seconds() != 5 {
			t.Errorf("Wanted 5, got %v right answers", after.Sub(time.Now()).Seconds())
		}
	})
}

func TestQuestionnaryCheckAnswer(t *testing.T) {
	t.Run("Test questionnary checkAnswer", func(t *testing.T) {
		testQuestionnary := getTestQuestionnary()
		for index, q := range testQuestionnary.Questions {
			testQuestionnary.checkAnswer(q.A, index)
		}

		if testQuestionnary.RightAnswers != uint(len(testQuestionnary.Questions)) {
			t.Errorf("Wanted %v, got %v", len(testQuestionnary.Questions), testQuestionnary.RightAnswers)
		}

		testQuestionnary = getTestQuestionnary()
		for index, q := range testQuestionnary.Questions {
			testQuestionnary.checkAnswer(q.A+"wrong", index)
		}

		if testQuestionnary.RightAnswers == uint(len(testQuestionnary.Questions)) {
			t.Errorf("Wanted 0, got %v", testQuestionnary.RightAnswers)
		}
	})
}

func TestGetTimer(t *testing.T) {
	t.Run("Testing getTimer()", func(t *testing.T) {
		questionnary := getTestQuestionnary()

		timerWithEmptyDuration := questionnary.getTimer()
		timerWithDuration := questionnary.SetTimerDuration(10).getTimer()

		if timerWithEmptyDuration == nil || timerWithDuration == nil {
			t.Error("Wanted timer not to be nil")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("Testing questionnary no shuffle", func(t *testing.T) {
		questionnary := getTestQuestionnary()
		originalQuestionnary := getTestQuestionnary()

		questionnary.Shuffle(false)

		for index, question := range questionnary.Questions {
			if originalQuestionnary.Questions[index] != question {
				t.Errorf("Wanted %#v, got %#v", originalQuestionnary.Questions[index], question)
			}
		}
	})

	t.Run("Testing questionnary should shuffle", func(t *testing.T) {
		questionnary := getTestQuestionnary()
		originalQuestionnary := getTestQuestionnary()

		questionnary.Shuffle(true)

		iSimilar := true
		for index, question := range questionnary.Questions {
			if originalQuestionnary.Questions[index] != question {
				iSimilar = false
			}
		}

		if iSimilar {
			t.Error("Wanted questionnary to be shuffled")
		}
	})
}
