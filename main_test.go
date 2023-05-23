package main

import (
    "gitlab.com/nyarla/go-crypt"
    "testing"
)


func TestGimmeTime(t *testing.T) {
    want := "01"
    msg, err := gimmeTime()
    if msg != want || err != nil {
	t.Fatalf(`Want: %v --> return value: %v`, want, msg)
    }
}

func TestCrypt(t *testing.T) {
    want := "esDRYJnY4VaGM"
    msg := crypt.Crypt("testtest", "es")
    if msg != want {
	t.Fatalf(`Want: %v --> return value: %v`, want, msg)
    }
}
