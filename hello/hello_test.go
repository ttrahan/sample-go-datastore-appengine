package hello

import (
	"testing"

	"appengine/aetest"
)

func TestStorage(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if err := populate(c); err != nil {
		t.Fatal(err)
	}

	score, err := getScore(c)
	if err != nil {
		t.Fatal(err)
	}
	if score < 0 || score > 1023 {
		t.Errorf("Score outside of expected range: %d", score)
	}
}



