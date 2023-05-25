package main

import (
    "github.com/dustin/go-humanize"
    "testing"
)


func TestGimmeTime(t *testing.T) {
    want := "16 years ago"
    msg := humanize.Time(gimmeTime())
    if msg != want {
	t.Fatalf(`Want: %v --> return value: %v`, want, msg)
    }
}

