package cos_test

import (
	"testing"

	"github.com/RitchieFlick/cos"
)

var ExistinghardcodedValues = []struct {
	value string
}{
	{"Remember: YAGNI (You Ain’t Gonna Need It)"},
}

var NonExistinghardcodedValues = []struct {
	value string
}{
	{"Olla"},
}

func TestExistingHardCodedValues(t *testing.T) {
	var included bool = false
	api := cos.NewAPI()
	randomPhrase, _ := api.GetRandomPhrase()
	for _, value := range ExistinghardcodedValues {
		if value.value == randomPhrase {
			included = true
		}
	}
	if !included {
		t.Errorf("The phrase wasn't found although it should!", ExistinghardcodedValues, randomPhrase)
	}
}

func TestNonExistingHardCodedValues(t *testing.T) {
	var included bool = false
	api := cos.NewAPI()
	randomPhrase, _ := api.GetRandomPhrase()
	for _, value := range NonExistinghardcodedValues {
		if value.value == randomPhrase {
			included = true
		}
	}
	if included {
		t.Errorf("The phrase was found although it shouldn't!", randomPhrase, NonExistinghardcodedValues)
	}
}
