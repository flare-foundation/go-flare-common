<div align="center">
  <a href="https://flare.network/" target="blank">
    <img src="https://content.flare.network/Flare-2.svg" width="300" alt="Flare Logo" />
  </a>
  <br />
  <a href="CONTRIBUTING.md">Contributing</a>
  ·
  <a href="SECURITY.md">Security</a>
  ·
  <a href="CHANGELOG.md">Changelog</a>
</div>

# Go Flare Common V1

Golang packages that are used across multiple Flare projects.

Contains the following packages:

- call: rest api calls
- contracts: [abigen](https://geth.ethereum.org/docs/tools/abigen) bindings for Flare contracts
- database: connecting to and reading from the Flare's c-chain indexer database
- events: parsing event logs as stored in Flare's c-chain indexer database
- heapt: generically typed heap implementation based on "container/heap"
- logger: logging solution
- merkle: Merkle tree implementation as used by Flare's protocols
- payload: working with transaction inputs for Submission smart contract
- policy: storing and parsing signing policies
- priority: heap implementation of a priority queue with two lanes
- queue: channel implementation of a priority queue with two lanes
- restserver
- retry: retried execution of a function
- storage: an implementation of a cyclic storage
- tee
  - attestation: working with jwt tokens carrying Google Cloud tee attestations
  - instructions: working with tee instructions
  - op: opTypes and opCommands
  - signer: signer server for remote ETH messages signing and decrypting
  - structs: encoding and decoding messages from tee instructions events, FDC2 requests and attestations
- toml: reading toml files
- voters: working with Flare entities
- xrpl
  - address: XRPL addresses, account ID, and public keys
  - aggregator: working with XRPL multisig
  - base58: base58 encoding with xrpl or custom alphabet
  - encoding: encoding and decoding XRPL transactions
  - hash: hashes used in XRPL
  - signing: signing of XRPL transactions
    - ed25519
    - secp256k1
    - signer: working with xrpl signer items
