package database

import "time"

// Abstract entity, all other entities should be derived from it.
type BaseEntity struct {
	ID uint64 `gorm:"primaryKey"`
}

// Should be synchronized with the the c-chain indexer.
type State struct {
	BaseEntity
	Name           string `gorm:"type:varchar(50);index"` // first_database_block", “last_database_block”, or “last_chain_block”
	Index          uint64 // blockNumber
	BlockTimestamp uint64
	Updated        time.Time
}

// Should be synchronized with the c-chain indexer.
type Transaction struct {
	BaseEntity
	Hash             string `gorm:"type:varchar(64);index;unique"`
	FunctionSig      string `gorm:"type:varchar(50);index"` //function selector
	Input            string `gorm:"type:string"`
	BlockNumber      uint64
	BlockHash        string `gorm:"type:varchar(64)"`
	TransactionIndex uint64
	FromAddress      string `gorm:"type:varchar(40);index"`
	ToAddress        string `gorm:"type:varchar(40);index"`
	Status           uint64
	Value            string `gorm:"type:string"`
	GasPrice         string `gorm:"type:string"`
	Gas              uint64
	Timestamp        uint64 `gorm:"index"`
}

// Should be synchronized with the the c-chain indexer.
type Log struct {
	BaseEntity
	TransactionID   uint64       `gorm:"default:null"` // database ID of the transaction, should not be confused wit the hash of the transaction
	Transaction     *Transaction `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE"`
	Address         string       `gorm:"type:varchar(40);index"`
	Data            string       `gorm:"type:string"`
	Topic0          string       `gorm:"type:varchar(64);index"`
	Topic1          string       `gorm:"type:varchar(64);index"`
	Topic2          string       `gorm:"type:varchar(64);index"`
	Topic3          string       `gorm:"type:varchar(64);index"`
	TransactionHash string       `gorm:"type:varchar(64);uniqueIndex:hash_index_unique"`
	LogIndex        uint64       `gorm:"uniqueIndex:hash_index_unique"`
	Timestamp       uint64       `gorm:"index"`
	BlockNumber     uint64       `gorm:"index"`
}
