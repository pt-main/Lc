package bytecode

import (
	"math"
	"testing"

	"github.com/pt-main/lc/v2/public"
)

var endianness = []public.EndianType{public.BigEndian, public.LittleEndian}

var sizes = []int{1, 2, 3, 4, 8}

// Float64ToBytes stored an unsigned code through int, and BytesToFloat64 read
// it back with sign extension, so any code with its top bit set decoded as a
// huge negative number. Every value broke at every size and both byte orders.
func TestFloat64RoundTrip(t *testing.T) {
	u := &Utils{}
	values := []float64{0, 0.25, 0.5, 0.75, 0.9999, 1, -1, -0.5, -0.25}
	for _, e := range endianness {
		for _, size := range sizes {
			tolerance := 2.0/float64(uint64(1)<<(uint(size)*8)) + 1e-9
			for _, v := range values {
				got := u.BytesToFloat64(u.Float64ToBytes(v, size, e), e)
				if math.Abs(got-v) > tolerance {
					t.Errorf("endianness=%v size=%d value=%v got=%v", e, size, v, got)
				}
				// The codec must also be a fixed point, otherwise a value
				// drifts every time it is read and written again.
				if again := u.BytesToFloat64(u.Float64ToBytes(got, size, e), e); again != got {
					t.Errorf("endianness=%v size=%d value=%v is not a fixed point: %v then %v",
						e, size, v, got, again)
				}
			}
		}
	}
}

// The range codec shares the same encoding, so it broke as soon as the
// normalized value reached 0.5 and set the top bit.
func TestFloat64RangeRoundTrip(t *testing.T) {
	u := &Utils{}
	for _, e := range endianness {
		for _, size := range sizes {
			tolerance := 10.0/float64(uint64(1)<<(uint(size)*8)) + 1e-9
			for _, v := range []float64{1, 2.5, 4.9, 5, 5.1, 7.5, 9} {
				got := u.BytesToFloat64Range(u.Float64ToBytesRange(v, size, 0, 10, e), 0, 10, e)
				if math.Abs(got-v) > tolerance {
					t.Errorf("endianness=%v size=%d value=%v got=%v", e, size, v, got)
				}
			}
		}
	}
}

func TestFloat64RangeEndpoints(t *testing.T) {
	u := &Utils{}
	for _, e := range endianness {
		for _, size := range []int{1, 2, 4} {
			for _, v := range []float64{0, 10} {
				got := u.BytesToFloat64Range(u.Float64ToBytesRange(v, size, 0, 10, e), 0, 10, e)
				if got != v {
					t.Errorf("endianness=%v size=%d endpoint=%v got=%v", e, size, v, got)
				}
			}
		}
	}
}

// maxValue used to be computed as 1<<(size*8)-1 in a uint64, which overflows
// for size above 8 and made the whole range unsatisfiable.
func TestLargeSizeDoesNotBreak(t *testing.T) {
	u := &Utils{}
	for _, size := range []int{9, 16} {
		if got := u.BytesToFloat64(u.Float64ToBytes(0, size, public.LittleEndian), public.LittleEndian); math.Abs(got) > 1 {
			t.Errorf("size=%d zero did not survive: %v", size, got)
		}
	}
}

func TestIntRoundTrip(t *testing.T) {
	u := &Utils{}
	for _, e := range endianness {
		for _, size := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
			// the largest value a signed field of this width can hold
			value := int(uint64(1)<<(uint(size)*8-1) - 1)
			if size == 8 {
				value = 1234
			}
			if got := u.BytesToInt(u.IntToBytes(value, size, e), e); got != value {
				t.Errorf("endianness=%v size=%d value=%d got=%d", e, size, value, got)
			}
		}
	}
}

// A negative length passed the end-of-data check and then panicked while
// slicing, and a nil index dereferenced.
func TestShiftRejectsInvalidInput(t *testing.T) {
	idx := 0
	s := NewShift([]byte{1, 2, 3}, &idx)
	for _, length := range []int{-1, -100} {
		if _, err := s.ShiftError(length); err == nil {
			t.Errorf("ShiftError(%d) must fail", length)
		}
	}
	if idx != 0 {
		t.Errorf("a rejected shift must not move the index, got %d", idx)
	}
	if _, err := (&Shift{Code: []byte{1}}).ShiftError(1); err == nil {
		t.Error("ShiftError with a nil index must fail")
	}
	past := 99
	if _, err := NewShift([]byte{1}, &past).ShiftError(1); err == nil {
		t.Error("ShiftError past the end must fail")
	}
}

func TestShiftPanicOnInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("ShiftPanic must panic on a negative length")
		}
	}()
	idx := 0
	NewShift([]byte{1, 2, 3}, &idx).ShiftPanic(-1)
}
