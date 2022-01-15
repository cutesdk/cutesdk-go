package cache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	c := New()

	v, ok := c.Get("foo")
	t.Error(v, ok)
	c.Set("foo", "bar3", 3600*time.Second)

	v, ok = c.Get("foo")

	t.Error(v, ok)
}
