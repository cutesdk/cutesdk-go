package tests

import "testing"

func TestCode2Session(t *testing.T) {
	client := getClient()

	code := "091pS8Ga1yE4ZC0LtvIa1rPq9g3pS8GY"

	res, err := client.Code2Session(code)

	t.Error(res, err)
}

func TestGetUserPhone(t *testing.T) {
	client := getClient()

	code := "cda802f413a0a7c9767819742a84f3fe3e7839da6661e62f39d0290a21116dce"

	res, err := client.GetUserPhone(code)

	t.Error(res, err)
}
