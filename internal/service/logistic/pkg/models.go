package pkg

var allEntities = []Package{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Package struct {
	Title string
}

func (p Package) String() string {
	return p.Title
}
