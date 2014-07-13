package hello

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type Score struct {
	Score int
	Date  time.Time
}

func scoreKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Scores", "default_scoreboard", 0, nil)
}

func populate(c appengine.Context) error {
	score := Score{
		Score: rand.Intn(1234),
		Date:  time.Now(),
	}
	key := datastore.NewIncompleteKey(c, "Score", scoreKey(c))
	_, err := datastore.Put(c, key, &score)
	return err
}

func getScore(c appengine.Context) (int, error) {
	query := datastore.NewQuery("Score").Ancestor(scoreKey(c)).Order("-Date").Limit(1)
	for t := query.Run(c); ; {
		var score Score
		if _, err := t.Next(&score); err != nil {
			return -1, err
		}

		return score.Score, nil
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := populate(c); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	score, err := getScore(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprintf(w, "Hello, world %d!", score)
	}
}
