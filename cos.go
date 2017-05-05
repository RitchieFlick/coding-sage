package cos

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type datastore interface {
	GetPhrases() ([]string, error)
}

// NewAPI creates a new instance of the API of cos.
func NewAPI() API {
	return API{}
}

// API represents the main entry point for the project.
type API struct {
	datastores []datastore
}

// AddDatastore adds a new datastore to the list of existing datastores
func (api *API) AddDatastore(datastore datastore) error {
	api.datastores = append(api.datastores, datastore)
	return nil
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
