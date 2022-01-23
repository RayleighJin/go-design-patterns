package facade

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkSecurityCode(code int) error {
	if s.code != code {
		return fmt.Errorf("securityCode is incorrect")
	}
	fmt.Println("security code verified")
	return nil
}
