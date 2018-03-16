package main

type cmdList struct {
	// From        string `long:"from" description:"from date"`
	// To          string `long:"to" description:"to date"`
	ShowAllTags bool     `short:"a" long:"all" description:"show all tags (local and inherited)"`
	Tags        []string `short:"t" long:"tag" description:"tag to filter"`
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

		elements, err := parse(lines)
		if err != nil {
			return err
		}

		filterVar := func(elem element) bool {
			for _, tag := range cl.Tags {
				if elem.HasTag(tag) {
					return true
				}
			}
			return false
		}
		elements = filterElements(elements, filterVar)

	}

	return nil
}
