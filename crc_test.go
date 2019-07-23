package crc8

import (
	"strings"
	"testing"
)

var largeText = []byte(strings.Repeat("a", 50000))
var smallText = []byte(strings.Repeat("a", 5))

func AssetEqual(t *testing.T, expected uint8, actual uint8) {
	if expected != actual {
		t.Errorf("Expected 0x%02X got 0x%02X", expected, actual)
	}
}

func TestCrc8EmptyVector(t *testing.T) {
	AssetEqual(t, 0x00, Checksum(nil, 0x07, 0x00, 0x00))
}

func TestNormalized(t *testing.T) {

	input := []byte("abcdefgh")

	t.Log("CRC-8")
	{
		AssetEqual(t, 0xCB, Checksum(input, 0x07, 0x00, 0x00))
	}

}

func BenchmarkCrcSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		Checksum(smallText, 0x07, 0x00, 0x00)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		Checksum(largeText, 0x07, 0x00, 0x00)
	}
}
