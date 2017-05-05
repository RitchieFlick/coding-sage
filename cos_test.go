package cos_test

import (
	"testing"

	"github.com/RitchieFlick/cos"
)

type testHardCode struct{}

func (t testHardCode) GetPhrases() ([]string, error) {
	var list []string
	list = append(list, "Remember: YAGNI (You Ain’t Gonna Need It)")
	list = append(list, "Remember: 3-2-1 Backup Strategy")
	return list, nil
}

func TestGetRandomPhrase(t *testing.T) {
	api := cos.NewAPI()
	api.AddDatastore(testHardCode{})
	firstPhrase, err := api.GetRandomPhrase()
	if err != nil {
		t.Errorf("An error was raised!", err)
	}

	var samePhrase = true

	for i := 0; i < 4; i++ {
		api := cos.NewAPI()
		api.AddDatastore(testHardCode{})
		phrase, err := api.GetRandomPhrase()
		if err != nil {
			t.Errorf("An error was raised!", err)
		}
		t.Logf(phrase)
		if firstPhrase != phrase {
			samePhrase = false
		}
	}

	if samePhrase {
		t.Errorf("Only the same phrase was always returned!")
	}
}
