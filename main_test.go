package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	sql := "SELECT character FROM cartoons WHERE phrase = 'zoinks' OR phrase = 'ruh roh' LIMIT 10"
	expected := "SELECT character FROM cartoons WHERE phrase = ? OR phrase = ? LIMIT ?"

	normalize(t, sql, expected)

	sql = "INSERT INTO cartoons (character, phrase) VALUES ('Shaggy', 'zoinks')"
	expected = "INSERT INTO cartoons (character, phrase) VALUES (?, ?)"

	normalize(t, sql, expected)

	sql = "INSERT INTO cartoons (character, phrase) VALUES ('Velma', 'My glasses!'), ('Scooby', 'ruh roh')"
	expected = "INSERT INTO cartoons (character, phrase) VALUES (?, ?), (?, ?)"

	normalize(t, sql, expected)

	sql = "UPDATE cartoons SET character = 'Fred' WHERE character = 'fred'"
	expected = "UPDATE cartoons SET character = ? WHERE character = ?"

	normalize(t, sql, expected)
}

func normalize(t *testing.T, sql string, expected string) {
	output, err := normalizeQuery(sql)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(output))
}
