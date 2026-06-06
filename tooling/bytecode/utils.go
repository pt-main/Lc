package bytecode

import (
	"errors"
)

type Utils struct{}

const (
	BigEndian = iota
	LittleEndian
)

func (u *Utils) IntToBytesBigEndian(value int, size int) []byte {
	result := make([]byte, size)
	for i := size - 1; i >= 0; i-- {
		result[i] = byte(value & 0xFF)
		value >>= 8
	}
	return result
}

func (u *Utils) IntToBytesLittleEndian(value int, size int) []byte {
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = byte(value & 0xFF)
		value >>= 8
	}
	return result
}

func (u *Utils) IntToBytes(value int, size int, endianess int) []byte {
	if endianess == BigEndian {
		return u.IntToBytesBigEndian(value, size)
	}
	return u.IntToBytesLittleEndian(value, size)
}

func (u *Utils) BytesToIntBigEndian(bytes []byte) int {
	result := 0
	for _, b := range bytes {
		result = (result << 8) | int(b)
	}
	return result
}

func (u *Utils) BytesToIntLittleEndian(bytes []byte) int {
	result := 0
	for i, b := range bytes {
		result |= int(b) << (8 * i)
	}
	return result
}

func (u *Utils) BytesToInt(bytes []byte, endianess int) int {
	if endianess == BigEndian {
		return u.BytesToIntBigEndian(bytes)
	}
	return u.BytesToIntLittleEndian(bytes)
}

type shift struct {
	code []byte
	Idx  *int
}

func (u *Utils) ShiftStruct(code []byte, idx *int) *shift {
	return &shift{
		code: code,
		Idx:  idx,
	}
}

func (s *shift) ShiftError(length int) ([]byte, error) {
	if *s.Idx+length > len(s.code) {
		return nil, errors.New("Unexpected end of data")
	}
	res := s.code[*s.Idx : *s.Idx+length]
	*s.Idx += length
	return res, nil
}

func (s *shift) ShiftPanic(length int) []byte {
	bytes, err := s.ShiftError(length)
	if err != nil {
		panic("Can't continue shifting, error: " + err.Error())
	}
	return bytes
}
