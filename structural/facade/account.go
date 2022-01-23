package facade

import "fmt"

type account struct {
	id string
}

func newAccount(accountID string) *account {
	return &account{
		id: accountID,
	}
}

func(a *account) checkAccount(accountID string) error {
	if a.id != accountID {
		return fmt.Errorf("accountID is incorrect")
	}
	fmt.Println("account verified")
	return nil
}