package cos

import (
	"math/rand"

	store "github.com/RitchieFlick/cos/datastore"
)

type datastore interface {
	GetPhrases() ([]string, error)
}

// NewAPI creates a new instance of the API of cos.
func NewAPI() API {
	var listDatastores []datastore
	listDatastores = append(listDatastores, store.HardCoded{})
	return API{datastores: listDatastores}
}

// API represents the main entry point for the project.
type API struct {
	datastores []datastore
}

// GetPhrases returns a random phrase from the list of different Datastores.
func (api API) GetRandomPhrase() (string, error) {
	phrases, err := api.getAllPhrases()
	if err != nil {
		return "", err
	}
	return phrases.chooseRandomPhrase(), nil
}

func (api API) getAllPhrases() (phrases, error) {
	var phrases []string
	for _, datastore := range api.datastores {
		datastorePhrases, err := datastore.GetPhrases()
		if err != nil {
			return nil, err
		}
		phrases = append(phrases, datastorePhrases...)
	}
	return phrases, nil
}

type phrases []string

func (i phrases) chooseRandomPhrase() string {
	numberOfRandomPhrases := len(i)
	randomPhraseIndex := rand.Intn(numberOfRandomPhrases)
	return i[randomPhraseIndex]
}
