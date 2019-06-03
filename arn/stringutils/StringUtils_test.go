package stringutils_test

import (
	"testing"

	"github.com/animenotifier/notify.moe/arn/stringutils"
	"github.com/stretchr/testify/assert"
)

func TestRemoveSpecialCharacters(t *testing.T) {
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Hello World"), "Hello World")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Aldnoah.Zero 2"), "Aldnoah Zero 2")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Working!"), "Working ")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Working!!"), "Working  ")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Working!!!"), "Working   ")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("Lucky☆Star"), "Lucky Star")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("ChäoS;Child"), "ChäoS Child")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("K-On!"), "KOn ")
	assert.Equal(t, stringutils.RemoveSpecialCharacters("僕だけがいない街"), "僕だけがいない街")
}

func TestContainsUnicodeLetters(t *testing.T) {
	assert.False(t, stringutils.ContainsUnicodeLetters("hello"))
	assert.True(t, stringutils.ContainsUnicodeLetters("こんにちは"))
	assert.True(t, stringutils.ContainsUnicodeLetters("hello こんにちは"))
}
