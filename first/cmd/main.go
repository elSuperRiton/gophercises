package main

import (
	"flag"

	"github.com/elSuperRiton/gophercices/first/questionnary"
)

const (
	defaultFilename             = "problems.csv"
	filenameFlag                = "file"
	questionnaryDurationFlag    = "duration"
	defaultQuestionnaryDuration = 5
	questionnaryShuffleFlag     = "shuffle"
	defaultQuestionnaryShuffle  = false
)

var (
	filename             string
	questionnaryDuration uint
	questionnaryShuffle  bool
)

func init() {
	flag.StringVar(&filename, filenameFlag, defaultFilename, "the filename to load")
	flag.UintVar(&questionnaryDuration, questionnaryDurationFlag, defaultQuestionnaryDuration, "the questionnary timer")
	flag.BoolVar(&questionnaryShuffle, questionnaryShuffleFlag, defaultQuestionnaryShuffle, "the questionnary timer")
	flag.Parse()
}

func main() {
	questionary := questionnary.LoadFile(filename)
	questionary.SetTimerDuration(questionnaryDuration).Shuffle(questionnaryShuffle).Run()
}
