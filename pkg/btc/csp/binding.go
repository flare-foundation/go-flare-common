package csp

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Binding is the wallet's multisig configuration, provisioned to each machine
// so it can derive the addresses it signs for.
//
// It exists as its own instruction because of an ordering constraint that
// cannot be worked around: each machine generates its key independently, so the
// complete xpub set only exists once EVERY key has been confirmed on chain.
// There is therefore no moment during key generation at which a machine could
// be told who its co-signers are.
//
// A machine must not take this from a proposal. It derives the witness script
// for each input from these xpubs, and that derivation is what proves the
// wallet owns the input being spent — if the set were substitutable per
// proposal, whoever assembled one could point the machine at a wallet it does
// not own and use it as a signing oracle.
type Binding struct {
	WalletID [32]byte
	// Network selects the chain parameters. Part of the key's identity: the
	// same seed yields different addresses on regtest and mainnet.
	Network string
	// Threshold is k in k-of-n.
	Threshold uint32
	// ParentXpubs are the n published wallet-level keys at m/87'/coin', in
	// confirmation order. Order does not affect the derived address — BIP-67
	// sorts the leaf pubkeys — but it is preserved so a machine can locate its
	// own entry.
	ParentXpubs []string
}

var bindingArgs abi.Arguments

func init() {
	ty, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "walletId", Type: "bytes32"},
		{Name: "network", Type: "string"},
		{Name: "threshold", Type: "uint32"},
		{Name: "parentXpubs", Type: "string[]"},
	})
	if err != nil {
		panic("building binding ABI type: " + err.Error())
	}
	bindingArgs = abi.Arguments{{Type: ty}}
}

type abiBinding struct {
	WalletID    [32]byte `abi:"walletId"`
	Network     string   `abi:"network"`
	Threshold   uint32   `abi:"threshold"`
	ParentXpubs []string `abi:"parentXpubs"`
}

// Validate rejects a binding that cannot produce a signable address.
func (b Binding) Validate() error {
	if len(b.ParentXpubs) == 0 {
		return errors.New("binding has no xpubs")
	}
	if b.Threshold < 1 || int(b.Threshold) > len(b.ParentXpubs) {
		return fmt.Errorf("threshold %d is out of range for n=%d", b.Threshold, len(b.ParentXpubs))
	}
	if b.Network == "" {
		return errors.New("binding has no network")
	}
	// Duplicate xpubs collapse signer independence: one holder could satisfy
	// several CHECKMULTISIG slots, so a k-of-n with repeats is not a genuine
	// k-of-n.
	seen := make(map[string]struct{}, len(b.ParentXpubs))
	for i, x := range b.ParentXpubs {
		if x == "" {
			return fmt.Errorf("xpub %d is empty", i)
		}
		if _, dup := seen[x]; dup {
			return fmt.Errorf("xpub %d is a duplicate; every signer must be distinct", i)
		}
		seen[x] = struct{}{}
	}
	return nil
}

// EncodeBinding packs the binding for the instruction payload.
func EncodeBinding(b Binding) ([]byte, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	out, err := bindingArgs.Pack(abiBinding(b))
	if err != nil {
		return nil, fmt.Errorf("encoding binding: %w", err)
	}
	return out, nil
}

// DecodeBinding parses a provisioned binding, rejecting anything unusable
// before it can be stored.
func DecodeBinding(raw []byte) (Binding, error) {
	vals, err := bindingArgs.Unpack(raw)
	if err != nil {
		return Binding{}, fmt.Errorf("decoding binding: %w", err)
	}
	if len(vals) != 1 {
		return Binding{}, errors.New("binding decoded to more than one value")
	}
	rawB, ok := abi.ConvertType(vals[0], new(abiBinding)).(*abiBinding)
	if !ok || rawB == nil {
		return Binding{}, errors.New("binding did not convert to the expected shape")
	}
	b := Binding{
		WalletID:    rawB.WalletID,
		Network:     rawB.Network,
		Threshold:   rawB.Threshold,
		ParentXpubs: rawB.ParentXpubs,
	}
	if err := b.Validate(); err != nil {
		return Binding{}, err
	}
	return b, nil
}

// Contains reports whether xpub is one of the binding's signers, which is how a
// machine confirms it is being provisioned for a wallet it actually holds a key
// for.
func (b Binding) Contains(xpub string) bool {
	return slices.Contains(b.ParentXpubs, xpub)
}

// BindingWalletID reads just the wallet id, for callers routing before decode.
func BindingWalletID(raw []byte) (common.Hash, error) {
	b, err := DecodeBinding(raw)
	if err != nil {
		return common.Hash{}, err
	}
	return b.WalletID, nil
}
