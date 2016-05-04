package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type indexCode int

const (
	indexName indexCode = iota
	indexLink
	indexDesc
	indexTags
	indexAuthor
)

var (
	indexNames map[string]bool
	indexUrls  map[string]bool
	lastName   string
	urlMatcher *regexp.Regexp
)

func init() {
	log.SetFlags(0) // Set logging to have no prefixes

	indexNames = make(map[string]bool)
	indexUrls = make(map[string]bool)
	urlMatcher = regexp.MustCompile(`^http(?:s)?://(?:github.com|gitlab.com|bitbucket.org|gist.github.com)/(?:\S+)/(?:\S+)$`)
}

type indexLineError struct {
	what       interface{}
	lineText   string
	lineNumber int
}

func (e *indexLineError) Error() string {
	what := ""

	switch reflect.TypeOf(e.what).Kind() {
	case reflect.Slice:
		errs := reflect.ValueOf(e.what)
		for i := 0; i < errs.Len(); i++ {
			what += "   " + errs.Index(i).String() + "\n"
		}
	default:
		what, _ = e.what.(string)
		what = "   " + what + "\n"
	}

	return fmt.Sprintf("%d %s\n   ---\n%s   ...", e.lineNumber+1, e.lineText, what)
}

type indexLine struct {
	no   int
	code indexCode
	text string
}

func (l *indexLine) verify() error {
	// Checks if fisher index is correctly formatted: 5 lines + 1 blank
	if l.text == "" && l.no%6 != 5 {
		log.Fatalf("line %d: index is incorrectly formatted\n", l.no)
	}

	switch l.code {
	case indexName:
		// Check if name is unique
		if indexNames[l.text] {
			return &indexLineError{"Name is not unique", l.text, l.no}
		}

		if lastName > l.text {
			return &indexLineError{"Index is not sorted", l.text, l.no}
		}

		lastName = l.text

		indexNames[l.text] = true
	case indexLink:
		// Check if URL is Unique
		if indexUrls[l.text] {
			return &indexLineError{"Url is not unique", l.text, l.no}
		}

		indexUrls[l.text] = true

		if urlMatcher.MatchString(l.text) != true {
			return &indexLineError{"Url in not valid", l.text, l.no}
		}
	case indexDesc:
		// Check Description
		// if _, ok := regexp.MatchString(`^[\w\s]+$`, l.text); ok != nil {
		// return &indexLineError{"Description is not valid", l.text, l.no}
		// }
	case indexTags:
		// Check duplicate and invalid tags
		encountered := map[string]bool{}
		duplicates := []string{}
		invalids := []string{}
		longs := []string{}
		errors := []string{}

		tags := strings.Fields(l.text)
		for j := range tags {
			if ok, _ := regexp.MatchString(`^[a-zA-z0-9_-]+$`, tags[j]); !ok {
				invalids = append(invalids, tags[j])
			}

			if len(tags[j]) > 15 {
				longs = append(longs, tags[j])
			}

			if encountered[tags[j]] {
				duplicates = append(duplicates, tags[j])
			} else {
				encountered[tags[j]] = true
			}
		}

		if len(tags) > 4 {
			errors = append(errors, "Too many tags")
		}

		if len(longs) != 0 {
			errors = append(errors, "Long tags: "+strings.Join(longs, ", "))
		}

		if len(invalids) != 0 {
			errors = append(errors, "Invalid tags: "+strings.Join(invalids, ", "))
		}

		if len(duplicates) != 0 {
			errors = append(errors, "Duplicate tags: "+strings.Join(duplicates, ", "))
		}

		if len(errors) != 0 {
			return &indexLineError{errors, l.text, l.no}
		}
	case indexAuthor:
		// Check Author
	}

	return nil
}
