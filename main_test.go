package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {

	if reflect.TypeOf(version.name) != reflect.TypeOf("a string") {
		t.Errorf("Is not a string")
	}

}
