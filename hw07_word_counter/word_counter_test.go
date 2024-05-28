package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "should return empty map when input is empty",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "should return correct map of words when input contains one word",
			input: "42",
			want:  map[string]int{"42": 1},
		},
		{
			name:  "should return correct map of words when input contains multiple words with one space between",
			input: "hello world",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "should return correct map of words when input contains multiple words with one many space between",
			input: "hello    world      welcome",
			want:  map[string]int{"hello": 1, "world": 1, "welcome": 1},
		},
		{
			name:  "should return correct map of words when input contains spaces at the beginning and end of input",
			input: "	hello world welcome			  ",
			want:  map[string]int{"hello": 1, "world": 1, "welcome": 1},
		},
		{
			name:  "should return correct map of words when input contains the same words several times",
			input: "two three one three two three",
			want:  map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name:  "should return correct map of words when input contains punctuation marks",
			input: "two: three ! ? one ,  three.   two ;three",
			want:  map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name:  "should return correct map of words when input contains punctuation marks",
			input: "two: three ! ? one ,  three.   two ;three",
			want:  map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name:  "should return correct map of words when input contains non-alphabetic and non-numeric characters",
			input: " - two № three $$ @ one **  three^   #two ( ) _three",
			want:  map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name:  "should return correct map of words when input contains non-english alphabetic characters #1",
			input: " - два № три $$ @ один **  три^   #два ( ) _три",
			want:  map[string]int{"один": 1, "два": 2, "три": 3},
		},
		{
			name:  "should return correct map of words when input contains non-english alphabetic characters #1",
			input: " - 二 № 三 $$ @ ワン **  三^   #二 ( ) _三",
			want:  map[string]int{"ワン": 1, "二": 2, "三": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.input)
			require.Equal(t, len(tt.want), len(got), "Returned map of words has the wrong length")
			for k, v := range tt.want {
				require.Equal(t, v, got[k], "Returned map contains the wrong count of word: %s", v)
			}
		})
	}
}
