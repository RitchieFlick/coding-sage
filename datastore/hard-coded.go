package datastore

type HardCoded struct{}

func (HardCoded) GetPhrases() ([]string, error) {
	var list []string
	list = append(list, "Remember: YAGNI (You Ain’t Gonna Need It)")
	return list, nil
}
