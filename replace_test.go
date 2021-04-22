package replace

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestReplace(t *testing.T) {
	var data1 struct {
		Name     string
		UserName string
		Email    string
		Password string
	}
	data1.Name = "user"
	data1.UserName = "user"
	data1.Email = "user"
	data1.Password = "password"
	var data2 struct {
		Name     *string
		UserName *string
		Email    *string
		Password *string
	}
	value := "newuser"
	data2.Name = &value
	Replace(&data1, &data2)
	assert.Equal(t, data1.Name, *data2.Name)
	assert.Equal(t, data1.UserName, "user")
	assert.Equal(t, data1.Email, "user")
	assert.Equal(t, data1.Password, "password")
}
