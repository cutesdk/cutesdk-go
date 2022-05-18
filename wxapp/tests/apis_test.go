package tests

import "testing"

func TestCode2Session(t *testing.T) {
	client := getClient()

	code := "051kk0nl2jfTd94EqBnl2cp1IS3kk0nO"

	res, err := client.Code2Session(code)

	t.Error(res, err)
}

func TestGetUserPhone(t *testing.T) {
	client := getClient()

	code := "cda802f413a0a7c9767819742a84f3fe3e7839da6661e62f39d0290a21116dce"

	res, err := client.GetUserPhone(code)

	t.Error(res, err)
}
