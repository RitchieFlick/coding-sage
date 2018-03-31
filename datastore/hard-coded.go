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
	{"> The trouble with programmers is that you can never tell what a programmer is doing until it’s too late.\n- Seymour Cray"},
	{"Have you tried turning it off and on again?"},
	{"> Provide Options, Don't Make Lame Excuses\nInstead of excuses, provide options. Don't say it can't be done; explain what can be done.\n-The Pragmatic Programmer"},
	{"DRY - Don't Repeat Yourself"},
	{"> Make It Easy To Reuse\nIf it's easy to reuse, people will. Create an environment that supports reuse.\n-The Pragmatic Programmer"},
	{"> Prototype To Learn\nPrototyping is a learning experience. Its value lies not in the code you produce, but in the lessons you learn.\n-The Pragmatic Programmer"},
	{"> Don't Assume It - Prove It\nProve your assumptions in the actual environment - with real data and boundary conditions.\n-The Pragmatic Programmer"},
	{"> Learn a Text Manipulation Language\nYou spend a large part of each day working with text. Why not have the computer do some of it for you?\n-The Pragmatic Programmer"},
	{"> Write Code That Writes Code\nCode generators increase your productivity and help avoid duplication.\n-The Pragmatic Programmer"},
	{"> Put Abstractions in Code, Details in Metadata\nProgram for the general case, and put the specifics outside the compiled code base.\n-The Pragmatic Programmer\n-Editors Note: Business policies"},
	{"> Seperate Views from Models\nGain flexibility at low cost by designing your application in terms of models and views.\n-The Pragmatic Programmer"},
	{"> Refactor Early, Refactor Often\nJust as you might weed and rearrange a garden, rewrite, rework and re-architect code when it needs it. Fix the root of the problem.\n-The Pragmatic Programmer"},
	{"> Design to Test\nStart thinking about testing before you write a line of code\n-The Pragmatic Programmer"},
	{"> Test Your Software, or Your Users Will\nTest ruthlessy. Don't make your users find bugs for you.\n-The Pragmatic Programmer"},
	{"> Don't Gather Requirements - Dig for Them\nRequirements rarely lie on the surface. They're buried deep beneath layers of assumptions, misconecptions, and politics.\n-The Pragmatic Programmer"},
	{"> Work with a User to Think Like a User\nIt's the best way to gain insight into how the system will really be used.\n-The Pragmatic Programmer"},
	{"> Don't Use Manual Procedures\nA shell script or batch file will execute the same instructions, in the same order, time after time.\n-The Pragmatic Programmer"},
	{"> Test Early. Test Often. Test Automatically.\nTests that run with every build are much more effective than test plans that site on a shelf.\n-The Pragmatic Programmer"},
	{"> Test State Coverage, Not Code Coverage\nIdentify and test significant program states. Just testing lines of code isn't enough.\n-The Pragmatic Programmer"},
	{"> Find Bugs Once\nOnce a human tester finds a bug, it should be the last time a human tester finds that bug. Automatic tests should check for it from then on.\n-The Pragmatic Programmer"},
	{"> English is Just a Programming Language\nWrite documents as you would write code: honor the DRY principle, use metadata, MVC, automatic generation, and so on.\n-The Pragmatic Programmer"},
	{"> Build Documentation in, Don't Bolt It On\nDocumentation created seperately from code is less likely to be correct and up to date.\n-The Pragmatic Programmer"},
}

func (HardCoded) GetPhrases() ([]string, error) {
	var list []string
	for _, phrase := range phrases {
		list = append(list, phrase.value)
	}
	return list, nil
}
