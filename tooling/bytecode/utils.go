package bytecode

import (
	"math"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// Utils provides byte-level conversions and shifting utilities.
type Utils struct{}

// IntToBytesBigEndian converts an int to a big-endian byte slice of the given
// size. A size below one panics on the slice allocation.
func (u *Utils) IntToBytesBigEndian(value int, size int) []byte {
	result := make([]byte, size)
	val := uint64(value)
	for i := size - 1; i >= 0; i-- {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// IntToBytesLittleEndian converts an int to a little-endian byte slice of the
// given size. A size below one panics on the slice allocation.
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
// A size below one panics.
func (u *Utils) IntToBytes(value int, size int, endianness public.EndianType) []byte {
	if endianness == public.BigEndian {
		return u.IntToBytesBigEndian(value, size)
	}
	return u.IntToBytesLittleEndian(value, size)
}

// BytesToIntBigEndian converts a big-endian byte slice to a sign-extended int.
func (u *Utils) BytesToIntBigEndian(bytes []byte) int {
	var val uint64
	for _, b := range bytes {
		val = (val << 8) | uint64(b)
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToIntLittleEndian converts a little-endian byte slice to a
// sign-extended int.
func (u *Utils) BytesToIntLittleEndian(bytes []byte) int {
	var val uint64
	for i, b := range bytes {
		val |= uint64(b) << (8 * uint(i))
	}
	bits := uint(len(bytes) * 8)
	if bits > 0 && (val>>(bits-1))&1 == 1 {
		val |= ^uint64(0) << bits
	}
	return int(val)
}

// BytesToInt converts a byte slice to an int using the specified endianness.
func (u *Utils) BytesToInt(bytes []byte, endianness public.EndianType) int {
	if endianness == public.BigEndian {
		return u.BytesToIntBigEndian(bytes)
	}
	return u.BytesToIntLittleEndian(bytes)
}

// unsignedMaxValue is the largest unsigned value representable in size bytes.
// size above 8 does not fit in uint64, so it is capped at 8.
func unsignedMaxValue(size int) uint64 {
	if size >= 8 {
		return math.MaxUint64
	}
	return 1<<(uint(size)*8) - 1
}

// bytesToUint decodes an unsigned integer, so the top bit is not sign-extended.
func (u *Utils) bytesToUint(bytes []byte, endianness public.EndianType) uint64 {
	var val uint64
	if endianness == public.BigEndian {
		for _, b := range bytes {
			val = (val << 8) | uint64(b)
		}
		return val
	}
	for i, b := range bytes {
		val |= uint64(b) << (8 * uint(i))
	}
	return val
}

func (u *Utils) uintToBytes(value uint64, size int, endianness public.EndianType) []byte {
	result := make([]byte, size)
	if endianness == public.BigEndian {
		for i := size - 1; i >= 0; i-- {
			result[i] = byte(value)
			value >>= 8
		}
		return result
	}
	for i := 0; i < size; i++ {
		result[i] = byte(value)
		value >>= 8
	}
	return result
}

// Float64ToBytes converts a float64 in range [-1,1] to a byte slice of given size.
// Panics: if size <= 0 or value exceeds the representable range (clamped).
func (u *Utils) Float64ToBytes(value float64, size int, endianness public.EndianType) []byte {
	if size <= 0 {
		panic("Float64ToBytes: size must be positive")
	}
	if value < -1.0 {
		value = -1.0
	}
	if value > 1.0 {
		value = 1.0
	}
	// The code is (value+1)/2*max, so a mid-grid value lands exactly between
	// two codes. Rounding (not truncating) keeps the error within one LSB and
	// makes BytesToFloat64(Float64ToBytes(x)) a fixed point.
	scaled := math.Round((value + 1.0) / 2.0 * float64(unsignedMaxValue(size)))
	return u.uintToBytes(uint64(scaled), size, endianness)
}

// BytesToFloat64 converts a byte slice to a float64 in range [-1,1].
func (u *Utils) BytesToFloat64(bytes []byte, endianness public.EndianType) float64 {
	size := len(bytes)
	if size == 0 {
		return 0.0
	}
	intValue := u.bytesToUint(bytes, endianness)
	maxValue := unsignedMaxValue(size)
	if maxValue == 0 {
		return 0.0
	}
	return float64(intValue)/float64(maxValue)*2.0 - 1.0
}

// Float64ToBytesRange converts a float64 in [minVal, maxVal] to a byte slice.
// Panics: if size <= 0, minVal >= maxVal, or value out of range (clamped).
func (u *Utils) Float64ToBytesRange(value float64, size int, minVal, maxVal float64, endianness public.EndianType) []byte {
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
	scaled := math.Round(normalized * float64(unsignedMaxValue(size)))
	return u.uintToBytes(uint64(scaled), size, endianness)
}

// BytesToFloat64Range converts a byte slice to a float64 in [minVal, maxVal].
// Panics: if minVal >= maxVal.
func (u *Utils) BytesToFloat64Range(bytes []byte, minVal, maxVal float64, endianness public.EndianType) float64 {
	size := len(bytes)
	if size == 0 {
		return minVal
	}
	if minVal >= maxVal {
		panic("BytesToFloat64Range: minVal must be less than maxVal")
	}
	intValue := u.bytesToUint(bytes, endianness)
	maxValue := unsignedMaxValue(size)
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

// Shift reads bytes from a buffer and moves an index through it, either
// reporting an error or panicking when the buffer is exhausted.
type Shift struct {
	Code []byte
	Idx  *int
}

// NewShift creates a Shift over the given buffer and index.
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
//     Meta: EMK(0, "int") - requested length,
//     EMK(1, "int") - current index,
//     EMK(2, "int") - total buffer length.
func (s *Shift) ShiftError(length int) ([]byte, core.ErrorInterface) {
	// A negative length passes an "end of data" check and then panics on the
	// slice, so it has to be rejected up front. A nil Idx is a setup error,
	// not a parsing one, and would nil-deref below.
	if s.Idx == nil {
		return nil, core.Err(errors.BytecodeShiftError, "Shift index is not initialized").
			WithMeta(core.EMK(0, "int"), length).
			WithMeta(core.EMK(1, "int"), 0).
			WithMeta(core.EMK(2, "int"), len(s.Code))
	}
	if length < 0 {
		return nil, core.Err(errors.BytecodeShiftError, "Negative length requested: %d", length).
			WithMeta(core.EMK(0, "int"), length).
			WithMeta(core.EMK(1, "int"), *s.Idx).
			WithMeta(core.EMK(2, "int"), len(s.Code))
	}
	if *s.Idx < 0 || *s.Idx+length > len(s.Code) {
		return nil, core.Err(errors.BytecodeShiftError, "Unexpected end of data").
			WithMeta(core.EMK(0, "int"), length).
			WithMeta(core.EMK(1, "int"), *s.Idx).
			WithMeta(core.EMK(2, "int"), len(s.Code))
	}
	res := s.Code[*s.Idx : *s.Idx+length]
	*s.Idx += length
	return res, nil
}

// ShiftPanic reads length bytes, panicking when the buffer is exhausted.
// Use it only where the buffer is known to be large enough.
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
func (s *Shift) ShiftFloat64Error(size int, endianness public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianness), nil
}

// ShiftFloat64Panic reads and converts a float64, panicking on error.
func (s *Shift) ShiftFloat64Panic(size int, endianness public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64(bytes, endianness)
}

// ShiftFloat64RangeError reads `size` bytes and interprets them as a float64 in [minVal, maxVal].
// If there is not enough data, returns a core.Error.
//
// Err errors.BytecodeShiftError (wrapped from ShiftError).
func (s *Shift) ShiftFloat64RangeError(size int, minVal, maxVal float64, endianness public.EndianType) (float64, core.ErrorInterface) {
	bytes, err := s.ShiftError(size)
	if err != nil {
		return 0, err
	}
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianness), nil
}

// ShiftFloat64RangePanic reads and converts a float64 with range, panicking on error.
func (s *Shift) ShiftFloat64RangePanic(size int, minVal, maxVal float64, endianness public.EndianType) float64 {
	bytes := s.ShiftPanic(size)
	utils := &Utils{}
	return utils.BytesToFloat64Range(bytes, minVal, maxVal, endianness)
}

// AutoIntToBytes chooses the minimal number of bytes needed to represent `value`
// (excluding sign extension) and converts it.
func (u *Utils) AutoIntToBytes(value int, endianness public.EndianType) []byte {
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
	return u.IntToBytes(value, size, endianness)
}

// AutoFloat64ToBytes chooses an optimal byte size based on the magnitude of `value`,
// then converts it to bytes.
func (u *Utils) AutoFloat64ToBytes(value float64, endianness public.EndianType) []byte {
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
	return u.Float64ToBytes(value, size, endianness)
}
