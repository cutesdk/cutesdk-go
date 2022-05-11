package tests

import (
	"testing"
	"time"
)

func TestSetComponentVerifyTicket(t *testing.T) {
	c := getClient()

	ticket := "ticket@@@BDJg_t76rRbAT-riW9cfI04C0v4eUR93de82xXAPf3abXgfCim4u8tMIRZY35v8iK20LGhiVvBqA9d9AjqAZgA"
	expire := 720 * time.Minute

	err := c.SetComponentVerifyTicket(ticket, expire)

	t.Error(err)
}
