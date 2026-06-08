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

func (u *Utils) Float64ToBytes(value float64, size int, endianess int) []byte {
	if size <= 0 {
		panic("Float64ToBytes: size must be positive")
	}

	maxValue := uint64(1<<(uint(size)*8) - 1)

	var scaledValue uint64

	if value >= 0 {

		if value > 1.0 {
			scaledValue = maxValue
		} else {
			scaledValue = uint64(value * float64(maxValue))
		}
	} else {

		if value < -1.0 {
			scaledValue = 0
		} else {
			scaledValue = uint64((value + 1.0) * float64(maxValue))
		}
	}
	if scaledValue > maxValue {
		scaledValue = maxValue
	}
	return u.IntToBytes(int(scaledValue), size, endianess)
}

func (u *Utils) BytesToFloat64(bytes []byte, endianess int) float64 {
	size := len(bytes)
	if size == 0 {
		return 0.0
	}

	intValue := uint64(u.BytesToInt(bytes, endianess))

	maxValue := uint64(1<<(uint(size)*8) - 1)

	if maxValue == 0 {
		return 0.0
	}

	return float64(intValue)/float64(maxValue)*2.0 - 1.0
}

func (u *Utils) Float64ToBytesRange(value float64, size int, minVal, maxVal float64, endianess int) []byte {
	if size <= 0 {
		panic("Float64ToBytesRange: size must be positive")
	}
	if minVal >= maxVal {
		panic("Float64ToBytesRange: minVal must be less than maxVal")
	}

	if value < minVal {
		value = minVal
	}
	if value > maxVal {
		value = maxVal
	}

	normalized := (value - minVal) / (maxVal - minVal)

	maxValue := uint64(1<<(uint(size)*8) - 1)

	scaledValue := uint64(normalized * float64(maxValue))

	return u.IntToBytes(int(scaledValue), size, endianess)
}

func (u *Utils) BytesToFloat64Range(bytes []byte, minVal, maxVal float64, endianess int) float64 {
	size := len(bytes)
	if size == 0 {
		return minVal
	}
	if minVal >= maxVal {
		panic("BytesToFloat64Range: minVal must be less than maxVal")
	}

	intValue := uint64(u.BytesToInt(bytes, endianess))

	maxValue := uint64(1<<(uint(size)*8) - 1)

	if maxValue == 0 {
		return minVal
	}

	normalized := float64(intValue) / float64(maxValue)

	return minVal + normalized*(maxVal-minVal)
}

func (u *Utils) Float64ToBytesBigEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, BigEndian)
}

func (u *Utils) Float64ToBytesLittleEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, LittleEndian)
}

func (u *Utils) BytesToFloat64BigEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, BigEndian)
}

func (u *Utils) BytesToFloat64LittleEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, LittleEndian)
}

type Shift struct {
	Code []byte
	Idx  *int
}

func NewShift(code []byte, idx *int) *Shift {
	return &Shift{
		Code: code,
		Idx:  idx,
	}
}

func (s *Shift) ShiftError(length int) ([]byte, error) {
	if *s.Idx+length > len(s.Code) {
		return nil, errors.New("Shift error: Unexpected end of data")
	}
	res := s.Code[*s.Idx : *s.Idx+length]
	*s.Idx += length
	return res, nil
}

func (s *Shift) ShiftPanic(length int) []byte {
	bytes, err := s.ShiftError(length)
	if err != nil {
		panic("Can't continue shifting, error: " + err.Error())
	}
	return bytes
}

func (s *Shift) ShiftFloat64Error(size int, endianess int) (float64, error) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess), nil
}

func (s *Shift) ShiftFloat64Panic(size int, endianess int) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess)
}

func (s *Shift) ShiftFloat64RangeError(size int, minVal, maxVal float64, endianess int) (float64, error) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess), nil
}

func (s *Shift) ShiftFloat64RangePanic(size int, minVal, maxVal float64, endianess int) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess)
}
