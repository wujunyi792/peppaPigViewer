package parse2cron_test

import (
	"fmt"
	"newJwCourseHelper/internal/util/parse2cron"
	"testing"
)

func TestFromSeconds(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "In one minute",
			args: args{
				n: 59,
			},
			want: "*/59 */0 */0 * * *",
		},
		{
			name: "Beyond one minute",
			args: args{
				n: 85,
			},
			want: "*/25 */1 */0 * * *",
		},
		{
			name: "In an hour",
			args: args{
				n: 60 * 60 * 3,
			},
			want: "*/0 */0 */3 * * *",
		},
		{
			name: "An hour with some minutes",
			args: args{
				n: 60 * 85 * 3,
			},
			want: "*/0 */15 */4 * * *",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse2cron.FromSeconds(tt.args.n); got != tt.want {
				t.Errorf("FromSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCronEntity(t *testing.T) {
	fmt.Printf("parse2cron.Seconds(10).Minutes().Hours().End(): %v\n", parse2cron.Seconds(80).Minutes(20).Hours(1).End())
}
