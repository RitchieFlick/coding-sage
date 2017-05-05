package main

import (
	"fmt"
	"os"

	cowsay "github.com/Code-Hex/Neo-cowsay"
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

	say, err := cowsay.Say(&cowsay.Cow{
		Phrase:      phrase,
		Type:        "sage",
		BallonWidth: 75,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
	os.Exit(0)
}
