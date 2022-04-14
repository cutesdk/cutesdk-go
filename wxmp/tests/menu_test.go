package tests

import "testing"

func TestGetCurrentMen(t *testing.T) {
	client := getClient()

	res, err := client.GetCurrentMenu()

	t.Error(res, err)
}
