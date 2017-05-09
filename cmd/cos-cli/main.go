package main

import (
	"fmt"
	"os"

	"github.com/RitchieFlick/cos"
	"github.com/RitchieFlick/cos/datastore"
)

func main() {
	api := cos.NewAPI()
	api.AddDatastore(datastore.HardCoded{})
	phrase, err := api.GetRandomPhrase()
	if err != nil {
		fmt.Errorf("There was an error retrieving a phrase!", err)
		os.Exit(1)
	}
	fmt.Printf(phrase)
	os.Exit(0)
}
