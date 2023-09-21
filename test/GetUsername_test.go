package test

import (
	"fmt"
	apisdk "github.com/quanxiangsheng/go-api-sdk"
	"testing"
)

var client = apisdk.NewApiClient("ABC", "abc")

func TestGetUsername(t *testing.T) {

	id := 1
	username := client.GetUsername(id)
	fmt.Println(username)
}