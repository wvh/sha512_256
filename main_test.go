package main

import (
	"strings"
	"testing"

	"encoding/hex"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		// https://en.wikipedia.org/wiki/SHA-2
		{"", "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a"},
		// https://csrc.nist.gov/CSRC/media/Projects/Cryptographic-Standards-and-Guidelines/documents/examples/SHA512_256.pdf
		{"abc", "53048e2681941ef99b2e29b76b4c7dabe4c2d0c634fc6d46e0e2f13107e7af23"},
		{"abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu", "3928e184fb8690f840da3988121d31be65cb9d3ef83ee6146feac861e19b563a"},
	}

	for _, test := range tests {
		t.Run("hash for "+test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			res := hex.EncodeToString(sum512_256(r))

			if res != test.out {
				t.Errorf("invalid hash for %q, expected %s, got %s\n", test.in, test.out, res)
			}
		})
	}
}
