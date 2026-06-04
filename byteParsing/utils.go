package byteParsing

import "errors"

type Utils struct{}

func (u *Utils) BytesToInt(bytes []byte) int {
	result := 0
	for _, b := range bytes {
		result = (result << 8) | int(b)
	}
	return result
}

type shift struct {
	code []byte
	Idx  int
}

func (u *Utils) ShiftStruct(code []byte) *shift {
	return &shift{
		code: code,
		Idx:  0,
	}
}

func (s *shift) ShiftError(length int) ([]byte, error) {
	if s.Idx+length > len(s.code) {
		return nil, errors.New("Unexpected end of data")
	}
	res := s.code[s.Idx : s.Idx+length]
	s.Idx += length
	return res, nil
}

func (s *shift) SiftPanic(length int) []byte {
	bytes, err := s.ShiftError(length)
	if err != nil {
		panic("Can't continue shifting, error: " + err.Error())
	}
	return bytes
}
