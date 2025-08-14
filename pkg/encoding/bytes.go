package encoding

import "encoding/binary"

func Uint16toBytes(i uint16) (arr [2]byte) {
	binary.BigEndian.PutUint16(arr[0:2], i)
	return
}
