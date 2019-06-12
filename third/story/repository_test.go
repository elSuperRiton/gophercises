package story

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	testRepository Repository
	testDataFile   = "./testdata/testgopher.json"
)

func setupStubTest(t *testing.T, dataType string) func(t *testing.T) {
	var err error
	testRepository, err = NewRepositoryFromFile("./testdata/testgopher.json")
	if err != nil {
		t.Fatal(err)
	}

	return func(t *testing.T) {
		testRepository = &repository{}
	}
}

func Test_Current(t *testing.T) {
	t.Run("Testing repository.Current() with empty current", func(t *testing.T) {
		repository := &repository{}
		if curr := repository.Current(); curr != "intro" {
			t.Errorf("Wanted current to be \"intro\", got %v", curr)
		}
	})
	t.Run("Testing repository.Current() with test current", func(t *testing.T) {
		repository := &repository{
			current: "fakecurrent",
		}
		if curr := repository.Current(); curr != "fakecurrent" {
			t.Errorf("Wanted current to be \"fakecurrent\", got %v", curr)
		}
	})
}

func Test_CurrentChapter(t *testing.T) {
	t.Run("Testing repository.CurrentChapter() with empty current", func(t *testing.T) {
		repository := &repository{
			story: Story{
				ChapterTitle("intro"): Chapter{
					Title: "intro",
				},
				ChapterTitle("testchapter"): Chapter{
					Title: "testchapter",
				},
			},
		}

		if chapter := repository.CurrentChapter(); chapter.Title != "intro" {
			t.Errorf("Wanted chapter title to be \"intro\", got %v", chapter.Title)
		}
	})
}

func Test_NewRepositoryFromFile(t *testing.T) {
	t.Run("Testing NewRepositoryFromFile with existing file", func(t *testing.T) {
		_, err := NewRepositoryFromFile(testDataFile)
		if err != nil {
			t.Errorf("expected err to be nil, got %v", err)
		}
	})
	t.Run("Testing NewRepositoryFromFile with non existing file", func(t *testing.T) {
		_, err := NewRepositoryFromFile("./testdata/idontexists.json")
		if err == nil {
			t.Error("expected err not to be nil")
		}
	})
}

func Test_LoadFromFile(t *testing.T) {
	t.Run("Testing LoadFromFile with proper file", func(t *testing.T) {
		repo := &repository{}

		if err := repo.LoadFromFile(testDataFile); err != nil {
			t.Errorf("Wanted err to be nil, got %v", err)
		}

		if repo.story == nil {
			t.Error("Wanted story not be nil")
		}
	})
}

func Test_ChapterFromRequest(t *testing.T) {
	t.Run("Testing repository.ChapterFromRequest(*http.Request) without params", func(t *testing.T) {
		repository := &repository{}

		req := httptest.NewRequest("GET", "/", nil)
		repository.ChapterFromRequest(req)
		if repository.current != "intro" {
			t.Errorf("Wanted current to be \"intro\", got %v", repository.current)
		}
	})

	t.Run("Testing repository.ChapterFromRequest(*http.Request) with params", func(t *testing.T) {
		repository := &repository{}
		param := make(url.Values)
		param["arc"] = []string{"testarc"}

		req := httptest.NewRequest("GET", "/?"+param.Encode(), nil)

		repository.ChapterFromRequest(req)
		if repository.current != "testarc" {
			t.Errorf("Wanted current to be \"testarc\", got %v", repository.current)
		}
	})
}
