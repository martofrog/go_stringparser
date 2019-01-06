package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
 *
 * We receive our odds from many different sources, we receive them as strings and manipulate them.
 *
 * A correct odds string is for example "5/2" which corresponds to 3.5 decimal. / is the only valid separator.
 * Unfortunately our feeds aren't always reliable and the string can contain literally anything.
 * Aim of this task is to provide a small and solid module which parses a string and returns a valid odds
 * if possible.
 *
 * Examples of invalid strings: “0/1”, “SP/1”, “2/3/4”,  “23*4”
 *
 **/

func TestOddsNotNull(t *testing.T) {
	oddsToTest := Odds("1/2")
	assert.NotNil(t, oddsToTest, "is Nil")
}