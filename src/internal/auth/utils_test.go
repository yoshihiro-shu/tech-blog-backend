package auth_test

import (
	"testing"

	"github.com/yoshihiro-shu/tech-blog-backend/src/internal/auth"
)

func TestIsVerfyPassword(t *testing.T) {
	tests := []struct {
		Key            string
		PlainPassword  string
		HashedPassword string
		Expected       bool
	}{
		{
			Key:           "test1",
			PlainPassword: "password",
			Expected:      true,
		},
		{
			Key:            "test2",
			PlainPassword:  "password",
			HashedPassword: "hogehogehgoe",
			Expected:       false,
		},
	}
	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			if test.HashedPassword == "" {
				hash, err := auth.GenerateBcryptPassword(test.PlainPassword)
				if err != nil {
					t.Errorf("failed, generate bcrypt password %v", err.Error())
				}
				test.HashedPassword = hash
			}
			expected := auth.IsVerifyPassword(test.PlainPassword, string(test.HashedPassword))
			if test.Expected != expected {
				t.Errorf("IsVerifyPassword() failed, expected %v, got %v", test.Expected, expected)
			}
		})
	}
}
