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
		UserName *string
		Name     *string
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

func catchPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}

func TestPanic1(t *testing.T) {
	defer catchPanic(t)
	var data struct{}
	Replace(data, &data)
}
func TestPanic2(t *testing.T) {
	defer catchPanic(t)
	var data struct{}
	Replace(&data, data)
}

func TestPanic3(t *testing.T) {
	defer catchPanic(t)
	var data struct{}
	var data2 []string
	Replace(&data2, &data)
}

func TestPanic4(t *testing.T) {
	defer catchPanic(t)
	var data struct{}
	var data2 []string
	Replace(&data, &data2)
}
