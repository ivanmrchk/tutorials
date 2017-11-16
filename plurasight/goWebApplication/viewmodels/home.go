package viewmodels

type Home struct {
	Title  string
	Active string
}

func GetHome() Home {
	result := Home{
		Title:  "Lemondae Stand Society",
		Active: "home",
	}
	return result
}
