package bytecode

import (
	"math"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// Utils provides byte-level conversions and shifting utilities.
type Utils struct{}

// IntToBytesBigEndian converts an int to a big-endian byte slice of the given size.
// Panics: if size <= 0 (not checked, will panic on slice allocation).
func (u *Utils) IntToBytesBigEndian(value int, size int) []byte {
	result := make([]byte, size)
	val := uint64(value)
	for i := size - 1; i >= 0; i-- {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// IntToBytesLittleEndian converts an int to a little-endian byte slice of the given size.
// Panics: if size <= 0 (not checked, will panic on slice allocation).
func (u *Utils) IntToBytesLittleEndian(value int, size int) []byte {
	result := make([]byte, size)
	val := uint64(value)
	for i := 0; i < size; i++ {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// IntToBytes converts an int to a byte slice with the given endianness.
// Panics: if size <= 0 (via called functions).
func (u *Utils) IntToBytes(value int, size int, endianess public.EndianType) []byte {
	if endianess == public.BigEndian {
		return u.IntToBytesBigEndian(value, size)
	}
	return u.IntToBytesLittleEndian(value, size)
}

// BytesToIntBigEndian converts a big-endian byte slice to an int.
// It handles sign extension.
func (u *Utils) BytesToIntBigEndian(bytes []byte) int {
	var val uint64 = 0
	for _, b := range bytes {
		val = (val << 8) | uint64(b)
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToIntLittleEndian converts a little-endian byte slice to an int.
// It handles sign extension.
func (u *Utils) BytesToIntLittleEndian(bytes []byte) int {
	var val uint64 = 0
	for i, b := range bytes {
		val |= uint64(b) << (8 * i)
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToInt converts a byte slice to an int using the specified endianness.
func (u *Utils) BytesToInt(bytes []byte, endianess public.EndianType) int {
	if endianess == public.BigEndian {
		return u.BytesToIntBigEndian(bytes)
	}
	return u.BytesToIntLittleEndian(bytes)
}

// Float64ToBytes converts a float64 in range [-1,1] to a byte slice of given size.
// Panics: if size <= 0 or value exceeds the representable range (clamped).
func (u *Utils) Float64ToBytes(value float64, size int, endianess public.EndianType) []byte {
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

// BytesToFloat64 converts a byte slice to a float64 in range [-1,1].
func (u *Utils) BytesToFloat64(bytes []byte, endianess public.EndianType) float64 {
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

// Float64ToBytesRange converts a float64 in [minVal, maxVal] to a byte slice.
// Panics: if size <= 0, minVal >= maxVal, or value out of range (clamped).
func (u *Utils) Float64ToBytesRange(value float64, size int, minVal, maxVal float64, endianess public.EndianType) []byte {
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

// BytesToFloat64Range converts a byte slice to a float64 in [minVal, maxVal].
// Panics: if minVal >= maxVal.
func (u *Utils) BytesToFloat64Range(bytes []byte, minVal, maxVal float64, endianess public.EndianType) float64 {
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

// Float64ToBytesBigEndian is a convenience wrapper for big-endian conversion.
func (u *Utils) Float64ToBytesBigEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, public.BigEndian)
}

// Float64ToBytesLittleEndian is a convenience wrapper for little-endian conversion.
func (u *Utils) Float64ToBytesLittleEndian(value float64, size int) []byte {
	return u.Float64ToBytes(value, size, public.LittleEndian)
}

// BytesToFloat64BigEndian is a convenience wrapper.
func (u *Utils) BytesToFloat64BigEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, public.BigEndian)
}

// BytesToFloat64LittleEndian is a convenience wrapper.
func (u *Utils) BytesToFloat64LittleEndian(bytes []byte) float64 {
	return u.BytesToFloat64(bytes, public.LittleEndian)
}

// Shift provides safe and unsafe byte reading from a buffer with an internal index.
type Shift struct {
	Code []byte
	Idx  *int
}

// NewShift creates a new Shift instance.
func NewShift(code []byte, idx *int) *Shift {
	return &Shift{
		Code: code,
		Idx:  idx,
	}
}

// ShiftError reads `length` bytes from the buffer and returns them.
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError:
//   - On unexpected end of data.
//     Meta: EMK(0, "int") – requested length,
//     EMK(1, "int") – current index,
//     EMK(2, "int") – total buffer length.
func (s *Shift) ShiftError(length int) ([]byte, core.ErrorInterface) {
	if *s.Idx+length > len(s.Code) {
		return nil, core.Err(errors.BytecodeShiftError, "Unexpected end of data").
			WithMeta(core.EMK(0, "int"), length).
			WithMeta(core.EMK(1, "int"), *s.Idx).
			WithMeta(core.EMK(2, "int"), len(s.Code))
	}
	res := s.Code[*s.Idx : *s.Idx+length]
	*s.Idx += length
	return res, nil
}

// ShiftPanic reads `length` bytes, panicking if not enough data.
// Use only in contexts where you are certain the buffer is large enough.
func (s *Shift) ShiftPanic(length int) []byte {
	bytes, err := s.ShiftError(length)
	if err != nil {
		panic("Can't continue shifting, error: " + err.Error())
	}
	return bytes
}

// ShiftFloat64Error reads `size` bytes and interprets them as a float64 in [-1,1].
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError (wrapped from ShiftError).
func (s *Shift) ShiftFloat64Error(size int, endianess public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess), nil
}

// ShiftFloat64Panic reads and converts a float64, panicking on error.
func (s *Shift) ShiftFloat64Panic(size int, endianess public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianess)
}

// ShiftFloat64RangeError reads `size` bytes and interprets them as a float64 in [minVal, maxVal].
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError (wrapped from ShiftError).
func (s *Shift) ShiftFloat64RangeError(size int, minVal, maxVal float64, endianess public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess), nil
}

// ShiftFloat64RangePanic reads and converts a float64 with range, panicking on error.
func (s *Shift) ShiftFloat64RangePanic(size int, minVal, maxVal float64, endianess public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianess)
}

// AutoIntToBytes chooses the minimal number of bytes needed to represent `value`
// (excluding sign extension) and converts it.
func (u *Utils) AutoIntToBytes(value int, endianess public.EndianType) []byte {
	size := 1
	temp := value
	if temp < 0 {
		temp = -temp
	}
	for temp > 0xFF {
		temp >>= 8
		size++
	}
	if value < 0 {
		size = 8
	}
	return u.IntToBytes(value, size, endianess)
}

// AutoFloat64ToBytes chooses an optimal byte size based on the magnitude of `value`,
// then converts it to bytes.
func (u *Utils) AutoFloat64ToBytes(value float64, endianess public.EndianType) []byte {
	size := 1
	absValue := math.Abs(value)
	if absValue > 1.0 {
		if absValue <= 2 {
			size = 2
		} else if absValue <= 4 {
			size = 3
		} else if absValue <= 8 {
			size = 4
		} else {
			size = 8
		}
	} else {
		if absValue == 0 {
			size = 1
		} else if absValue < 0.01 {
			size = 4
		} else if absValue < 0.1 {
			size = 3
		} else if absValue < 0.5 {
			size = 2
		}
	}
	return u.Float64ToBytes(value, size, endianess)
}
