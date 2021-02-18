## zero-alloc 8Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc8)](https://goreportcard.com/report/jancajthaml-go/crc8)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation provides both tableless and tabular checksum functions with variable 8bit polynomial.

### Supported standards ###

- CRC-8/CRC-8
- CRC-8/SAE-J1850
- CRC-8/SAE-J1850-ZERO
- CRC-8/8H2F
- CRC-8/CDMA2000
- CRC-8/DVB-S2
- CRC-8/ICODE
- CRC-8/ITU

### Usage ###

```
import "github.com/jancajthaml-go/crc16"

data := []byte("abcdefgh")
poly := 0x07
init := 0x00
xorout := 0x00

// for tableless
crc8.Checksum(data, poly, init, xorout) // 0xCB

// for precalculated tabular
instance = crc8.New(poly, init, xorout)
instance.Checksum(data) // 0xCB
```
