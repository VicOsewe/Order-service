package rest

import (
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case:valid email",
			args: args{
				email: gofakeit.Email(),
			},
			wantErr: false,
		},
		{
			name: "sad case:invalid email",
			args: args{
				email: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case:valid phone number",
			args: args{
				phoneNumber: gofakeit.Phone(),
			},
			wantErr: false,
		},
		{
			name: "sad case:invalid phone number with alpha numerics",
			args: args{
				phoneNumber: "invalid",
			},
			wantErr: true,
		},
		{
			name: "sad case:invalid phone number with less than ten characters",
			args: args{
				phoneNumber: "0715893",
			},
			wantErr: true,
		},
		{
			name: "sad case:invalid phone number with both alpha numerics and numbers",
			args: args{
				phoneNumber: "0715893invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePhoneNumber(tt.args.phoneNumber); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
