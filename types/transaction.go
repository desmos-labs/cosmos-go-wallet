package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// TransactionData contains all the data about a transaction
type TransactionData struct {
	Messages  []sdk.Msg
	Memo      string
	GasLimit  uint64
	FeeAmount sdk.Coins
}

// NewTransactionData builds a new TransactionData instance
func NewTransactionData(msg sdk.Msg, msgs ...sdk.Msg) *TransactionData {
	return &TransactionData{
		Messages: append(msgs, msg),
	}
}

// WithMemo allows to set the given memo
func (t *TransactionData) WithMemo(memo string) *TransactionData {
	t.Memo = memo
	return t
}

// WithGasLimit allows to set the given gas limit
func (t *TransactionData) WithGasLimit(limit uint64) *TransactionData {
	t.GasLimit = limit
	return t
}

// WithFeeAmount allows to set the given fee amount
func (t *TransactionData) WithFeeAmount(amount sdk.Coins) *TransactionData {
	t.FeeAmount = amount
	return t
}
