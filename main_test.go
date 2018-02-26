package main

import (
	"strings"
	"testing"

	"encoding/hex"
)

func TestMain(t *testing.T) {
	r := strings.NewReader("")

	res := hex.EncodeToString(sum512_256(r))
	expected := "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a" // https://en.wikipedia.org/wiki/SHA-2

	if res != expected {
		t.Errorf("invalid hash, expected %s, got %s\n", expected, res)
	}
}
