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

	t.Log("CRC8/CRC-8")
	{
		AssetEqual(t, 0xCB, Checksum(input, 0x07, 0x00, 0x00))
	}

	t.Log("CRC8/SAE-J1850")
	{
		AssetEqual(t, 0xD7, Checksum(input, 0x1D, 0xFF, 0xFF))
	}

	t.Log("CRC8/SAE-J1850-ZERO")
	{
		AssetEqual(t, 0x3E, Checksum(input, 0x1D, 0x00, 0x00))
	}

	t.Log("CRC8/8H2F")
	{
		AssetEqual(t, 0x54, Checksum(input, 0x2F, 0xFF, 0xFF))
	}

	t.Log("CRC8/CDMA2000")
	{
		AssetEqual(t, 0xF7, Checksum(input, 0x9B, 0xFF, 0x00))
	}

	//t.Log("CRC8/DARC")
	//{
	//AssetEqual(t, 0x62, Checksum(input, 0x39, 0x00, 0x00, true, true))
	//}

	t.Log("CRC8/DVB-S2")
	{
		AssetEqual(t, 0x62, Checksum(input, 0xD5, 0x00, 0x00))
	}

	//t.Log("CRC8/EBU")
	//{
	//AssetEqual(t, 0x41, Checksum(input, 0x1D, 0xFF, 0x00, true, true))
	//}

	t.Log("CRC8/ICODE")
	{
		AssetEqual(t, 0x96, Checksum(input, 0x1D, 0XFD, 0x00))
	}

	t.Log("CRC8/ITU")
	{
		AssetEqual(t, 0x9E, Checksum(input, 0x7, 0X00, 0x55))
	}

	//t.Log("CRC8/MAXIM")
	//{
	//AssetEqual(t, 0x92, Checksum(input, 0x31, 0X00, 0x00, true, true))
	//}

	//t.Log("CRC8/ROHC")
	//{
	//AssetEqual(t, 0x15, Checksum(input, 0x7, 0XFF, 0x0))
	//}

	//t.Log("CRC8/ITU")
	//{
	//AssetEqual(t, 0x3D, Checksum(input, 0x9B, 0X00, 0x00))
	//}
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
