package config

type Config struct {
	User struct {
		StaffId  string `json:"staffId"`
		Password string `json:"password"`
	} `json:"user"`

	Target     []string `json:"target"`
	ErrTag     []string `json:"errTag"`
	BucketFull int      `json:"bucketFull"`
	Rate       int      `json:"rate"`
	Ua         string   `json:"ua"`

	Interval int `json:"interval"`
}
