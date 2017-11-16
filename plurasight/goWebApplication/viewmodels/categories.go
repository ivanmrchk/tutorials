package viewmodels

type Categories struct {
	Title, Active string
	Categories    []Category
}

type Category struct {
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
	Id            int
}

func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "Shop",
	}

	juiceCategory := Category{
		Id:            1,
		ImageUrl:      "lemon.png",
		Title:         "Juices and Mixes",
		Description:   `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		IsOrientRight: false,
	}

	supplyCategory := Category{
		Id:            2,
		ImageUrl:      "kiwi.png",
		Title:         "Cups, Straws, and Other Supplies",
		Description:   `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		IsOrientRight: true,
	}

	advertisingCategory := Category{
		Id:            3,
		ImageUrl:      "pineapple.png",
		Title:         "Signs and Advertising",
		Description:   `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		IsOrientRight: false,
	}

	result.Categories = []Category{
		juiceCategory,
		supplyCategory,
		advertisingCategory,
	}

	return result
}
