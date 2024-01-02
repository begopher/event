package test

import(
	"fmt"
	"testing"
	"github.com/begopher/event/dispatcher"
)

type Account struct {
	Id int
}
type Bank struct {
	Id int
	T *testing.T
	Expected int
}


func (b Bank) Disable(acc Account) error {
	if acc.Id != b.Expected {
		b.T.Errorf("Expectd account id(%d) got (%d)", b.Expected, acc.Id)
	}
	return nil
}

type bankrupt struct {
	Account Account
	Bank Bank
}

func (b bankrupt) Occur() {
	b.Bank.Disable(b.Account)
}

func (b bankrupt) Name() string {
	return fmt.Sprintf("bankrupt[%v@%v]",
		b.Account.Id,
		b.Bank.Id)
}

func Test_Dispacher(t *testing.T) {
	expected := 55
	account := Account{expected}
	bank := Bank{
		Id: 1,
		T: t,
		Expected: expected,
	}
	bankrupted := 0
	dis := dispacher.New()
	dis.Publish(bankrupted)
	registration := bankrupt{account, bank}
	dis.Bind(bankrupted, registration)
	dis.Send(bankrupted)
}
