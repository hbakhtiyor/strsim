package strsim

import "fmt"

var testCases1 = []struct {
	a        string
	b        string
	expected float64
}{
	{"french", "quebec", 0},
	{"france", "france", 1},
	{"fRaNce", "france", 0.2},
	{"healed", "sealed", 0.8},
	{"web applications", "applications of the web", 0.7878787878787878},
	{"this will have a typo somewhere", "this will huve a typo somewhere", 0.92},
	{"Olive-green table for sale, in extremely good condition.", "For sale: table in very good  condition, olive green in colour.", 0.6060606060606061},
	{"Olive-green table for sale, in extremely good condition.", "For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721},
	{"Olive-green table for sale, in extremely good condition.", "Wanted: mountain bike with at least 21 gears.", 0.1411764705882353},
	{"this has one extra word", "this has one word", 0.7741935483870968},
	{"a", "a", 1},
	{"a", "b", 0},
	{"", "", 1},
	{"a", "", 0},
	{"", "a", 0},
	{"apple event", "apple    event", 1},
	{"iphone", "iphone x", 0.9090909090909091},
}

var testCases2 = []struct {
	s        string
	targets  []string
	expected *MatchResult
	err      error
}{
	{"", nil, nil, fmt.Errorf("targets parameter need at least one element")},
	{"", []string{}, nil, fmt.Errorf("targets parameter need at least one element")},
	{
		"healed",
		[]string{"mailed", "edward", "sealed", "theatre"},
		&MatchResult{
			[]*Match{
				{"mailed", 0.4},
				{"edward", 0.2},
				{"sealed", 0.8},
				{"theatre", 0.36363636363636365},
			},
			&Match{"sealed", 0.8},
			2,
		},
		nil,
	},
	{
		"Olive-green table for sale, in extremely good condition.",
		[]string{
			"For sale: green Subaru Impreza, 210,000 miles",
			"For sale: table in very good condition, olive green in colour.",
			"Wanted: mountain bike with at least 21 gears.",
		},
		&MatchResult{
			[]*Match{
				{"For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721},
				{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
				{"Wanted: mountain bike with at least 21 gears.", 0.1411764705882353},
			},
			&Match{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
			1,
		},
		nil,
	},
}

var testCases3 = []struct {
	actual   *MatchResult
	expected *MatchResult
}{
	{
		&MatchResult{
			[]*Match{
				{"mailed", 0.4},
				{"edward", 0.2},
				{"sealed", 0.8},
				{"theatre", 0.36363636363636365},
			},
			&Match{"sealed", 0.8},
			2,
		},
		&MatchResult{
			[]*Match{
				{"sealed", 0.8},
				{"mailed", 0.4},
				{"theatre", 0.36363636363636365},
				{"edward", 0.2},
			},
			&Match{"sealed", 0.8},
			0,
		},
	},
	{
		&MatchResult{
			[]*Match{
				{"For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721},
				{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
				{"Wanted: mountain bike with at least 21 gears.", 0.1411764705882353},
			},
			&Match{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
			1,
		},
		&MatchResult{
			[]*Match{
				{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
				{"For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721},
				{"Wanted: mountain bike with at least 21 gears.", 0.1411764705882353},
			},
			&Match{"For sale: table in very good condition, olive green in colour.", 0.6060606060606061},
			0,
		},
	},
}
