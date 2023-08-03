package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// TransactionData contains all the data about a transaction
type TransactionData struct {
	Messages   []sdk.Msg
	Memo       string
	GasLimit   uint64
	GasAuto    bool
	FeeAmount  sdk.Coins
	FeeAuto    bool
	FeeGranter sdk.AccAddress
	Sequence   *uint64
}

// NewTransactionData builds a new TransactionData instance
func NewTransactionData(msgs ...sdk.Msg) *TransactionData {
	return &TransactionData{
		Messages: msgs,
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

// WithGasAuto allows to automatically compute the amount of gas to be used when broadcasting the transaction
func (t *TransactionData) WithGasAuto() *TransactionData {
	t.GasAuto = true
	return t
}

// WithFeeAmount allows to set the given fee amount
func (t *TransactionData) WithFeeAmount(amount sdk.Coins) *TransactionData {
	t.FeeAmount = amount
	return t
}

// WithFeeAuto allows to automatically compute the fee amount to be used when broadcasting the transaction
func (t *TransactionData) WithFeeAuto() *TransactionData {
	t.FeeAuto = true
	return t
}

// WithFeeGranter allows to set the given fee granter that will pay for fees.
// To work properly, a fee grant must exist from the granter towards the transaction signer.
func (t *TransactionData) WithFeeGranter(granter sdk.AccAddress) *TransactionData {
	t.FeeGranter = granter
	return t
}

// WithSequence allows to set the given sequence
func (t *TransactionData) WithSequence(sequence uint64) *TransactionData {
	t.Sequence = &sequence
	return t
}
