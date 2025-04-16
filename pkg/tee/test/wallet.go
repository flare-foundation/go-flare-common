package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teeregistry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teewalletkeymanager"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teewalletmanager"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/teewalletprojectmanager"
)

type Client struct {
	prv                     *ecdsa.PrivateKey
	teeRegistry             *teeregistry.TeeRegistry
	teeWalletManager        *teewalletmanager.TeeWalletManager
	teeWalletKeyManager     *teewalletkeymanager.TeeWalletKeyManager
	teeWalletProjectManager *teewalletprojectmanager.TeeWalletProjectManager
	chainClient             *ethclient.Client
	chainID                 *big.Int
}

func NewClient(prv *ecdsa.PrivateKey, chainClient *ethclient.Client, teeRegistryAddr, teeWalletManagerAddr, teeWalletKeyManagerAddr, teeWalletProjectManagerAddr common.Address) (*Client, error) {
	teeRegistry, err := teeregistry.NewTeeRegistry(teeRegistryAddr, chainClient)
	if err != nil {
		return nil, err
	}

	teeWalletManager, err := teewalletmanager.NewTeeWalletManager(teeWalletManagerAddr, chainClient)
	if err != nil {
		return nil, err
	}

	teeWalletKeyManager, err := teewalletkeymanager.NewTeeWalletKeyManager(teeWalletKeyManagerAddr, chainClient)
	if err != nil {
		return nil, err
	}

	teeWalletProjectManager, err := teewalletprojectmanager.NewTeeWalletProjectManager(teeWalletProjectManagerAddr, chainClient)
	if err != nil {
		return nil, err
	}

	chainID, err := chainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{
		prv:                     prv,
		teeRegistry:             teeRegistry,
		teeWalletManager:        teeWalletManager,
		teeWalletProjectManager: teeWalletProjectManager,
		teeWalletKeyManager:     teeWalletKeyManager,
		chainClient:             chainClient,
		chainID:                 chainID,
	}, nil
}

func (c *Client) InitializeWallet(opType common.Hash, submitAddress common.Address) (common.Hash, common.Hash, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.prv, c.chainID)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}

	tx, err := c.teeWalletProjectManager.CreateProject(opts, opType, submitAddress)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}

	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, c.chainClient, tx)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Hash{}, common.Hash{}, fmt.Errorf("transaction CreateProject %s not successful", tx.Hash())
	}

	logProject, err := c.teeWalletProjectManager.ParseProjectCreated(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	projectID := logProject.ProjectId

	tx, err = c.teeWalletManager.CreateWallet(opts, projectID)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}

	receipt, err = bind.WaitMined(ctx, c.chainClient, tx)
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Hash{}, common.Hash{}, fmt.Errorf("transaction CreateWallet %s not successful", tx.Hash())
	}

	logWallet, err := c.teeWalletManager.ParseWalletCreated(*receipt.Logs[0])
	if err != nil {
		return common.Hash{}, common.Hash{}, err
	}
	walletID := logWallet.WalletId

	return projectID, walletID, nil
}

func (c *Client) AddKey(teeID common.Address, walletID common.Hash) (uint64, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(c.prv, c.chainID)
	if err != nil {
		return 0, err
	}

	tx, err := c.teeWalletKeyManager.AddKey(opts, teeID, walletID)
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

	log, err := c.teeWalletKeyManager.ParseWalletKeyAdded(*receipt.Logs[0])
	if err != nil {
		return 0, err
	}

	return log.KeyId.Uint64(), nil
}

func (c *Client) InitializeAndAddKeys() ([]uint64, error) {
	x := [32]byte{}
	copy(x[:], "XRP")

	submitAddress := common.BytesToAddress([]byte("randomAddress")) // unused so far

	_, walletID, err := c.InitializeWallet(x, submitAddress)
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
