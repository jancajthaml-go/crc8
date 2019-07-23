package crc8

func Checksum(data []byte, poly uint8, init uint8, xorout uint8) uint8 {
	var crc uint8 = init
	var bit uint8

	for _, item := range data {
		for j := uint16(0x0080); j != 0; j >>= 1 {
			if (uint16(item) & j) != 0 {
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
