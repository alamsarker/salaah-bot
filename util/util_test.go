package util

import (
	"reflect"
	"testing"
)

func TestGetTimeByHourMin(t *testing.T) {
	respTime := GetTimeByHourMin("10:11")
	if reflect.TypeOf(respTime) == nil {
		t.Error("Something went wrong to time")
	}
}

func TestGetCronTime(t *testing.T) {
	cronTime := GetCronTime("10:10")
	if cronTime != "10:00" {
		t.Error("Something went wrong to time")
	}
}
