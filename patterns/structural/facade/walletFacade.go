package facade

import "fmt"

type WalletFacade struct {
	Account      *account
	Wallet       *wallet
	SecurityCode *securityCode
	Notification *notification
	Ledger       *ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		Account:      newAccount(accountID),
		SecurityCode: newSecurityCode(code),
		Wallet:       newWallet(),
		Notification: &notification{},
		Ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.Wallet.creditBalance(amount)
	w.Notification.sendWalletCreditNotification()
	w.Ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.Wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.Notification.sendWalletDebitNotification()
	w.Ledger.makeEntry(accountID, "credit", amount)
	return nil
}
