package model_test

import (
	"github.com/lukachi/apiserver/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func () *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "With encrypt password",
			u: func () *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedPassword"

				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func () *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func () *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func () *model.User {
				u := model.TestUser(t)
				u.Password = "123"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
