# vincentinttsh/replace

[![Build Status](https://travis-ci.com/vincentinttsh/replace.svg?branch=master)](https://travis-ci.com/vincentinttsh/replace)
[![codecov](https://codecov.io/gh/vincentinttsh/replace/branch/master/graph/badge.svg?token=NATJW3S1UO)](https://codecov.io/gh/vincentinttsh/replacd)
[![Go Report Card](https://goreportcard.com/badge/github.com/vincentinttsh/replace)](https://goreportcard.com/report/github.com/vincentinttsh/replace)
[![GoDoc](https://godoc.org/github.com/vincentinttsh/replace?status.svg)](https://godoc.org/github.com/vincentinttsh/replace)

replace struct value by another struct

``` go
package main

import (
	"fmt"

	"github.com/vincentinttsh/replace"
)

func main() {
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
	fmt.Println(data1)
	replace.replace(&data1, &data2)
	fmt.Println(data1)
}
```
