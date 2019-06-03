package questionnary

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

type Questions []Question

type Question struct {
	Q string
	A string
}

type Questionnary struct {
	Questions    Questions
	Reader       io.Reader
	RightAnswers uint
	TimeLength   uint
}

func (qs *Questionnary) Run() *Questionnary {

	reader := bufio.NewReader(qs.Reader)
	answerChan := make(chan string)
	timer := qs.getTimer()

questionLoop:
	for index, question := range qs.Questions {

		go func() {
			fmt.Printf("%v ? ", question.Q)
			input, _ := reader.ReadString('\n')
			answerChan <- input
		}()

		select {
		case <-timer:
			break questionLoop
		case input := <-answerChan:
			qs.checkAnswer(input, index)
			continue questionLoop
		}
	}

	fmt.Printf("Right answer : %v\n", qs.RightAnswers)
	return qs
}

func (qs *Questionnary) SetTimerDuration(timer uint) *Questionnary {
	qs.TimeLength = timer

	return qs
}

func (qs *Questionnary) getTimer() <-chan time.Time {
	if qs.TimeLength > 0 {
		timer := time.NewTimer(time.Duration(qs.TimeLength) * time.Second)
		return timer.C
	}

	infiniteChan := make(chan time.Time)
	return infiniteChan
}

func (qs *Questionnary) checkAnswer(answer string, index int) {
	if strings.TrimSpace(strings.ToLower(qs.Questions[index].A)) == strings.TrimSpace(strings.ToLower(answer)) {
		qs.RightAnswers++
	}
}

func (qs *Questionnary) Shuffle(shouldShuffle bool) *Questionnary {
	if shouldShuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(qs.Questions), func(i, j int) {
			qs.Questions[i], qs.Questions[j] = qs.Questions[j], qs.Questions[i]
		})
	}

	return qs
}
