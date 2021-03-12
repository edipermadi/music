package midi

import "io"

func marshallVarInt(writer io.Writer, v uint64) error {
	var bytes []byte
	for {
		if v > 0x7f {
			bytes = append(bytes, byte(v)&0x7f)
			v >>= 7
		} else {
			bytes = append(bytes, byte(v)&0x7f)
			break
		}
	}

	for i := 0; i < len(bytes); i++ {
		if i > 0 {
			bytes[i] = 0x80 | bytes[i]
		}
	}

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	_, err := writer.Write(bytes)
	return err
}
