package helper

import "hash/fnv"

func Fnv32(in []byte) uint32 {
	h := fnv.New32()
	_, _ = h.Write(in)
	return h.Sum32()
}

func Fnv32Str(s string) uint32 {
	return Fnv32([]byte(s))
}
