package main

type cmdList struct {
	From        string `long:"from" description:"from date"`
	To          string `long:"to" description:"to date"`
	ShowAllTags bool   `short:"t" long:"tags" description:"show all tags"`
}

func (cl *cmdList) Execute(args []string) error {
	dir, err := DirFromOptsOrCurrent()
	if err != nil {
		return err
	}

	for _, file := range dir {
		lines, err := ReadFile(file)
		if err != nil {
			return err
		}

		elem, err := parse(lines)
		if err != nil {
			return err
		}

	}

	return nil
}
