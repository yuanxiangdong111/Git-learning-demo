package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimeUnix(t *testing.T) {
	timeUnix := time.Now().Unix()
	fmt.Println(GetTimeUnixToTime(timeUnix))
}

func TestGetStringTimeByTimezone(t *testing.T) {
	timeUnix := time.Now().Unix()
	r := FormatTimeByTimezone(GetTimeUnixToTime(timeUnix), TimeFormat, -7)
	fmt.Println(r)
}
