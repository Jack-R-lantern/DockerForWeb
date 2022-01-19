package service

import (
	"testing"
)

func TestGetContainerList(t *testing.T) {
	containers := GetContainerList()
	if len(containers) != 1 {
		t.Error(containers)
	}
}
