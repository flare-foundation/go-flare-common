package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teeregistry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teewalletmanager"
)

var uint64Type, _ = abi.NewType("uint64", "uint64", nil)

type Client struct {
	prv              *ecdsa.PrivateKey
	teeRegistry      *teeregistry.TeeRegistry
	teeWalletManager *teewalletmanager.TeeWalletManager
	chainClient      *ethclient.Client
	chainID          *big.Int
}

func NewClient(prv *ecdsa.PrivateKey, chainClient *ethclient.Client, teeRegistryAddr, teeWalletManagerAddr common.Address) (*Client, error) {
	teeRegistry, err := teeregistry.NewTeeRegistry(teeRegistryAddr, chainClient)
	if err != nil {
		return nil, err
	}

	teeWalletManager, err := teewalletmanager.NewTeeWalletManager(teeWalletManagerAddr, chainClient)
	if err != nil {
		return nil, err
	}

	chainID, err := chainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{
		prv:              prv,
		teeRegistry:      teeRegistry,
		teeWalletManager: teeWalletManager,
		chainClient:      chainClient,
		chainID:          chainID,
	}, nil
}

func (c *Client) InitializeWallet(opType common.Hash, multisigThreshold uint64) (common.Hash, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.prv, c.chainID)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.teeWalletManager.InitializeWallet(opts, opType, multisigThreshold)
	if err != nil {
		return common.Hash{}, err
	}

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, c.chainClient, tx)
	if err != nil {
		return common.Hash{}, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Hash{}, fmt.Errorf("transaction InitializeWallet %s not successful", tx.Hash())
	}

	return common.Hash(tx.Data()), nil
}

func (c *Client) AddKey(teeID common.Address, walletID common.Hash) (uint64, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.prv, c.chainID)
	if err != nil {
		return 0, err
	}

	tx, err := c.teeWalletManager.AddKey(opts, teeID, walletID)
	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, c.chainClient, tx)
	if err != nil {
		return 0, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return 0, fmt.Errorf("transaction AddKey %s not successful", tx.Hash())
	}

	returnArgs := abi.Arguments{{Type: uint64Type}}

	values, err := returnArgs.Unpack(tx.Data())
	if err != nil {
		return 0, err
	}

	out := abi.ConvertType(values[0], new(uint64)).(uint64)

	return out, nil
}

func (c *Client) InitializeAndAddKeys() ([]uint64, error) {
	x := [32]byte{}
	copy(x[:], "XRP")

	walletID, err := c.InitializeWallet(x, 1)
	if err != nil {
		return nil, err
	}

	ids, err := c.teeRegistry.GetActiveTeeIds(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	keyIDs := make([]uint64, len(ids))

	for i := range ids {
		keyID, err := c.AddKey(ids[i], walletID)
		if err != nil {
			return nil, err
		}
		keyIDs[i] = keyID
	}

	return keyIDs, nil
}
