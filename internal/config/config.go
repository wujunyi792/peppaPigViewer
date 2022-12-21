package config

type User struct {
	StaffId  string `json:"staffId"`
	Password string `json:"password"`
	//JSESSIONID string `json:"JSESSIONID"`
	//Route      string `json:"route"`
	IsAutoReAuth bool `json:"auto_auth"`
}

type Config struct {
	User User `json:"user"`

	Target []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"target"`
	//ClassNumber []string `json:"classNumber"`
	ErrTag []string `json:"errTag"`
	Rate   int      `json:"rate"`
	Ua     string   `json:"ua"`

	Interval int `json:"interval"`
}
