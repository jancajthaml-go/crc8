## 8Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc8)](https://goreportcard.com/report/jancajthaml-go/crc8)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation is tableless with variable 8bit polynomial.

### Performance ###

```
BenchmarkCrcSmall  58.47 MB/s  0 B/op  0 allocs/op
BenchmarkCrcLarge  47.11 MB/s  0 B/op  0 allocs/op
```

### Usage ###

```
import "github.com/jancajthaml-go/crc16"

data := []byte("abcdefgh")
poly := 0x07
init := 0x00
xorout := 0x00

crc8.Checksum(data, poly, init, xorout) // 0xCB
```
