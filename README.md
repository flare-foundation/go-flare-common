<p align="left">
  <a href="https://flare.network/" target="blank"><img src="https://flare.network/wp-content/uploads/Artboard-1-1.svg" width="400" height="300" alt="Flare Logo" /></a>
</p>

# Go Flare Common V1

Golang packages that are used across multiple Flare projects.

Contains the following packages:

    - database: connecting and reading the c-chain indexer database
    - events: parsing event logs as stored in Flare's c-chain indexer database
    - heapt: generically typed heap implementation based on based on "container/heap"
    - logger: logging solution
    - merkle: Merkle tree implementation
    - payload: working with transaction inputs for Submission smart contract
    - policy: storing and parsing signing policies
    - queue: an implementation of a priority queue with two lanes
    - restserver
    - storage: an implementation of a cyclic storage
    - voters: working with Flare entities
    - abigen binding for Flare contracts (generated using [abigen](https://geth.ethereum.org/docs/tools/abigen))
