package chirps

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Chirp{})

	if len(feed.Chirps) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Chirp{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}