package datastore

type HardCoded struct{}

var phrases = []struct {
	value string
}{
	{"Remember: YAGNI (You Ain’t Gonna Need It)"},
	{"> Truth can only be found in one place: the code.\n- Robert C. Martin, Clean Code: A Handbook of Agile Software Craftsmanship"},
	{"Remember: 3-2-1 Backup Strategy"},
	{"> Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live.\n- Rick Osborne"},
	{"> Debugging is twice as hard as writing the code in the first place. Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it.\n- Brian W. Kernighan."},
	{"The trouble with programmers is that you can never tell what a programmer is doing until it’s too late.\n- Seymour Cray"},
	{"Have you tried turning it off and on again?"},
}

func (HardCoded) GetPhrases() ([]string, error) {
	var list []string
	for _, phrase := range phrases {
		list = append(list, phrase.value)
	}
	return list, nil
}
