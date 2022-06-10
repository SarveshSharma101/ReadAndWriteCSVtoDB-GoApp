package shops

/**
Structures representing the data model of each of the shops
*/
type (
	Shop interface{}

	SuperMarket struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Gender  string `json:"gender"`
		Email   string `json:"email"`
		PhoneNo string `json:"phoneNo"`
	}

	Electronic struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		PhoneNo  string `json:"phoneNo"`
		Address  string `json:"address"`
		Location string `json:"location"`
	}

	Mobile struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Gender   string `json:"gender"`
		Email    string `json:"email"`
		PhoneNo  string `json:"phoneNo"`
		Address  string `json:"address"`
		Location string `json:"location"`
		Hobby    string `json:"Hobby"`
		Interest string `json:"interest"`
	}

	Bakery struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Gender  string `json:"gender"`
		Email   string `json:"email"`
		PhoneNo string `json:"phoneNo"`
	}
)
