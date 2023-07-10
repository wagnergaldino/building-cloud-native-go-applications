package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native GO", Author: "M. L. Reiner", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native GO","author":"M. L. Reiner","isbn":"0123456789"}`, string(json), "Book JSON marshalling gone wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native GO","author":"M. L. Reiner","isbn":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native GO", Author: "M. L. Reiner", ISBN: "0123456789"}, book, "Book JSON unmarshalling gone wrong")
}
