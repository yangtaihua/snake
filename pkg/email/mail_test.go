package email

import (
	"testing"

	"github.com/1024casts/snake/config"
)

func TestSend(t *testing.T) {
	// init config
	cfg := "../../conf/config.sample.yaml"
	if err := config.Init(cfg); err != nil {
		panic(err)
	}

	// init log
	config.InitLog(config.Conf)

	Init()

	type args struct {
		to      string
		subject string
		body    string
	}

	subject, body := NewResetPasswordHTMLEmail("test", "http://snake.com")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test send reset mail", args{"test100@test.com", subject, body}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.to, tt.args.subject, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
