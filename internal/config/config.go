package config

type Config struct {
	User struct {
		StaffId    string `json:"staffId"`
		Password   string `json:"password"`
		JSESSIONID string `json:"JSESSIONID"`
		Route      string `json:"route"`
	} `json:"user"`

	Target      []string `json:"target"`
	ClassNumber []string `json:"classNumber"`
	ErrTag      []string `json:"errTag"`
	BucketFull  int      `json:"bucketFull"`
	Rate        int      `json:"rate"`
	Ua          string   `json:"ua"`

	Interval int `json:"interval"`
}
