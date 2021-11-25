package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/desmos-labs/cosmos-go-wallet/types"
)

// Wallet represents a Cosmos cosmos that should be used to create and send transactions to the chain
type Wallet struct {
	privKey cryptotypes.PrivKey

	TxConfig client.TxConfig
	Client   *Client
}

// NewWallet allows to build a new Wallet instance
func NewWallet(accountCfg *types.AccountConfig, client *Client, txConfig client.TxConfig) (*Wallet, error) {
	// Get the private types
	algo := hd.Secp256k1
	derivedPriv, err := algo.Derive()(accountCfg.Mnemonic, "", accountCfg.HDPath)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		privKey:  algo.Generate()(derivedPriv),
		TxConfig: txConfig,
		Client:   client,
	}, nil
}

// AccAddress returns the address of the account that is going to be used to sign the transactions
func (w *Wallet) AccAddress() string {
	bech32Addr, err := bech32.ConvertAndEncode(w.Client.GetAccountPrefix(), w.privKey.PubKey().Address())
	if err != nil {
		panic(err)
	}
	return bech32Addr
}

// BroadcastTxAsync creates and signs a transaction with the provided messages and fees,
// then broadcasts it using the async method
func (w *Wallet) BroadcastTxAsync(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	builder, err := w.buildTx(msgs)
	if err != nil {
		return nil, err
	}

	return w.Client.BroadcastTxAsync(builder.GetTx())
}

// BroadcastTxSync creates and signs a transaction with the provided messages and fees,
// then broadcasts it using the sync method
func (w *Wallet) BroadcastTxSync(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	builder, err := w.buildTx(msgs)
	if err != nil {
		return nil, err
	}

	return w.Client.BroadcastTxSync(builder.GetTx())
}

// BroadcastTxCommit creates and signs a transaction with the provided messages and fees,
// then broadcasts it using the commit method
func (w *Wallet) BroadcastTxCommit(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	builder, err := w.buildTx(msgs)
	if err != nil {
		return nil, err
	}

	return w.Client.BroadcastTxCommit(builder.GetTx())
}

func (w *Wallet) buildTx(msgs []sdk.Msg) (client.TxBuilder, error) {
	// Get the account
	account, err := w.Client.GetAccount(w.AccAddress())
	if err != nil {
		return nil, fmt.Errorf("error while getting the account from the chain: %s", err)
	}

	builder := w.TxConfig.NewTxBuilder()
	builder.SetFeeAmount(w.Client.GetFees(200000))
	builder.SetGasLimit(200000)
	err = builder.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	// Set an empty signature first
	sigData := signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   w.privKey.PubKey(),
		Data:     &sigData,
		Sequence: account.GetSequence(),
	}

	err = builder.SetSignatures(sig)
	if err != nil {
		return nil, err
	}

	chainID, err := w.Client.GetChainID()
	if err != nil {
		return nil, err
	}

	// Sign the transaction with the private key
	sig, err = tx.SignWithPrivKey(
		signing.SignMode_SIGN_MODE_DIRECT,
		authsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: account.GetAccountNumber(),
			Sequence:      account.GetSequence(),
		},
		builder,
		w.privKey,
		w.TxConfig,
		account.GetSequence(),
	)
	if err != nil {
		return nil, err
	}

	err = builder.SetSignatures(sig)
	if err != nil {
		return nil, err
	}

	return builder, nil
}
