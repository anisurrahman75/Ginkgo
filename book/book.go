package books

import (
	"strings"
)

type Book struct {
	Title  string
	Author string
	Pages  int64
}

func (b *Book) IsValid() bool {
	return b.Title != "" && b.Author != "" && b.Pages > 0
}

func (b *Book) AuthorLastName() string {
	names := strings.Fields(b.Author)
	if len(names) > 0 {
		return names[len(names)-1]
	}
	return ""
}

func (b *Book) AuthorFirstName() string {
	names := strings.Fields(b.Author)
	if len(names) > 1 {
		return names[0]
	}
	return ""
}
