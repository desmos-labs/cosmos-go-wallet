package wallet_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/desmos-labs/cosmos-go-wallet/client"
	"github.com/desmos-labs/cosmos-go-wallet/types"
	"github.com/desmos-labs/cosmos-go-wallet/wallet"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

type WalletTestSuite struct {
	suite.Suite

	wallet *wallet.Wallet
	client *client.Client
}

func (suite *WalletTestSuite) SetupSuite() {
	chainCfg := types.ChainConfig{
		Bech32Prefix:  "desmos",
		RPCAddr:       "https://rpc.morpheus.desmos.network:443",
		GRPCAddr:      "https://grpc.morpheus.desmos.network:443",
		GasPrice:      "0.01udaric",
		GasAdjustment: 1.5,
	}
	accountCfg := types.AccountConfig{
		Mnemonic: "forward service profit benefit punch catch fan chief jealous steel harvest column spell rude warm home melody hat broccoli pulse say garlic you firm",
		HDPath:   "m/44'/852'/0'/0/0",
	}

	// Set up the SDK config with the proper bech32 prefixes
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(chainCfg.Bech32Prefix, fmt.Sprintf("%spub", chainCfg.Bech32Prefix))

	encodingCfg := simapp.MakeTestEncodingConfig()

	c, err := client.NewClient(&chainCfg, encodingCfg.Marshaler)
	suite.Require().NoError(err)
	suite.client = c

	w, err := wallet.NewWallet(&accountCfg, c, encodingCfg.TxConfig)
	suite.Require().NoError(err)
	suite.wallet = w
}

func (suite *WalletTestSuite) TestBuildTx() {
	sender, err := sdk.AccAddressFromBech32(suite.wallet.AccAddress())
	suite.Require().NoError(err)

	receiver, err := sdk.AccAddressFromBech32("desmos1q62k9kvjy7v2wh0yt9jqaepnzezz3s49j9gnpk")
	suite.Require().NoError(err)

	data := types.NewTransactionData(
		banktypes.NewMsgSend(
			sender,
			receiver,
			sdk.NewCoins(sdk.NewCoin("udaric", sdk.NewInt(10000))),
		),
	).WithGasAuto().WithFeeAuto().WithMemo("Custom memo")

	builder, err := suite.wallet.BuildTx(data)
	suite.Require().NoError(err)

	tx := builder.GetTx()
	suite.Require().Lessf(tx.GetGas(), uint64(200_000), "MsgSend should take less than 200.000 gas")
	suite.Require().NotEmptyf(tx.GetFee(), "Fees should not be empty")
}
