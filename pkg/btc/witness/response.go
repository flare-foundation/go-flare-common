package witness

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// SignResponse is what a TEE machine returns for a BTC batch: one signature per
// input, per key it holds.
//
// The wire form of what Assemble consumes, defined here so the machine that
// produces it and the facilitator that consumes it cannot drift — the two live
// in different repos and a silent field rename between them would surface only
// as "not enough verified signatures".
//
// Note the shape: XRPL returns ONE signed transaction per fee tier, BTC returns
// N signatures for one transaction. The facilitator's job differs accordingly —
// it assembles rather than selects.
type SignResponse struct {
	Txid       common.Hash         `json:"txid"`
	Signatures []SignatureResponse `json:"signatures"`
}

// SignatureResponse is one key's signature over one input.
type SignatureResponse struct {
	// KeyID is the wallet key that produced it, as the chain numbers them.
	KeyID uint64 `json:"keyId"`
	// Index is the input's position in the transaction.
	Index int `json:"index"`
	// Signature is DER without the trailing sighash-type byte.
	Signature hexutil.Bytes `json:"signature"`
	// PubKey is the compressed key, which decides the signature's slot.
	PubKey hexutil.Bytes `json:"pubKey"`
}

// Answer converts a machine's response into the form Assemble takes.
func (r SignResponse) Answer(machine int) Answer {
	a := Answer{Machine: machine, Signatures: make([]InputSignature, 0, len(r.Signatures))}
	for _, s := range r.Signatures {
		a.Signatures = append(a.Signatures, InputSignature{
			Index: s.Index, Signature: s.Signature, PubKey: s.PubKey,
		})
	}
	return a
}
