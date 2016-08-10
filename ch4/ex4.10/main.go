// Exercise 4.10:
// Modify the issues to report the results in age categories, say less than
// a month old, less than a year, and more than a year old.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopl.io/ch4/github"
)

var ageParam = flag.String("age", "", `An age requirement in <1y (less than a year) or >2m (more than a month) format.`)

func main() {
	flag.Parse()
	filter, err := setFilter(*ageParam)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[2:])
	filteredIssues := filterAge(result.Items, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", len(filteredIssues))
	for _, item := range filteredIssues {
		fmt.Printf("[%s]#%-5d %9.9s %.55s\n",
			item.CreatedAt.Format((time.UnixDate)),
			item.Number, item.User.Login, item.Title)
	}
}

func filterAge(result []*github.Issue, filter filterFunc) []*github.Issue {
	remains := result[:0]
	for _, issue := range result {
		if filter(issue) {
			remains = append(remains, issue)
		}
	}
	return remains
}

type filterFunc func(*github.Issue) bool

func setFilter(ageParam string) (filterFunc, error) {
	if len(ageParam) == 0 {
		return func(_ *github.Issue) bool {
			return true
		}, nil
	}

	if len(ageParam) != 3 {
		return nil, fmt.Errorf("setFilter: Age parameter '%s' in a wrong error format.", ageParam)
	}

	var isLess bool
	if ageParam[0] == '<' {
		isLess = true
	} else if ageParam[0] == '>' {
		isLess = false
	} else {
		return nil, fmt.Errorf("setFilter: Age parameter '%s' in a wrong error format.", ageParam)
	}

	t, err := strconv.Atoi(ageParam[1:2])
	if err != nil {
		return nil, fmt.Errorf("%v\nsetFilter: Age parameter '%s' in a wrong error format.", ageParam, err)
	}
	var duration time.Duration
	if ageParam[2] == 'm' {
		duration = time.Duration(t) * time.Hour * 24 * 30
	} else if ageParam[2] == 'y' {
		duration = time.Duration(t) * time.Hour * 24 * 30 * 12
	} else {
		return nil, fmt.Errorf("setFilter: Age parameter '%s' in a wrong error format.", ageParam)
	}

	return func(issue *github.Issue) bool {
		delta := time.Now().Sub(issue.CreatedAt)
		if isLess {
			return delta < duration
		} else {
			return delta >= duration
		}
	}, nil
}
