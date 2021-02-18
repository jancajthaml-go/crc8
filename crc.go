package crc8

// CRC holds crc8 parameters and precalculated table
type CRC struct {
	table  []uint8
	poly   uint8
	xorout uint8
	init   uint8
}

// New returns CRC8 instance with precalculated table
func New(poly uint8, init uint8, xorout uint8) CRC {
	return CRC{
		poly:   poly,
		xorout: xorout,
		init:   init,
		table:  createTable(poly, init, xorout),
	}
}

func createTable(poly uint8, init uint8, xorout uint8) []uint8 {
	result := make([]uint8, 256)
	var bit uint8
	for divident := 0; divident < 256; divident++ {
		var current uint8 = 0x00
		for j := byte(0x0080); j != 0; j >>= 1 {
			if (uint8(divident) & j) != 0 {
				bit = (current & 0x80) ^ 0x80
			} else {
				bit = current & 0x80
			}
			switch bit {
			case 0:
				current <<= 1
			default:
				current = (current << 1) ^ poly
			}
		}
		result[uint8(divident)] = current & 0xFF
	}
	return result
}

// Checksum returns CRC8 checksum of given CRC instance
func (crc *CRC) Checksum(data []byte) uint8 {
	var (
		result uint16 = uint16(crc.init)
		pos uint8
	)
	for _, item := range data {
		result = result ^ uint16(item)
		pos = uint8(result & 0xFF)
		result = (result << 8) & 0xFF
		result = (result ^ uint16(crc.table[pos])) & 0xFF
	}
	return uint8((result ^ uint16(crc.xorout)) & 0xFF)
}

// Checksum returns CRC8 checksum for given parameters
func Checksum(data []byte, poly uint8, init uint8, xorout uint8) uint8 {
	var (
		crc uint8 = init
		bit uint8
	)
	for _, item := range data {
		for j := byte(0x0080); j != 0; j >>= 1 {
			if (item & j) != 0 {
				bit = (crc & 0x80) ^ 0x80
			} else {
				bit = crc & 0x80
			}
			switch bit {
			case 0:
				crc <<= 1
			default:
				crc = (crc << 1) ^ poly
			}
		}
	}
	return crc ^ xorout
}
