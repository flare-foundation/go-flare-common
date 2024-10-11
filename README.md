# Go Flare Common V0

Golang packages that are used across multiple Flare projects.

Contains the following packages:

    - database: connecting and reading the c-chain indexer database
    - events: parsing event logs as stored in c-chain indexer database
    - logger: logging solution
    - merkle: Merkle tree implementation
    - payload: working with transaction inputs to submission contract
    - policy: storing and parsing signing policies
    - queue: an implementation of a priority queue with two lanes
    - restserver
    - storage: an implementation of a cyclic storage
    - voters: working with Flare entities
