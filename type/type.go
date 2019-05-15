package types

import(
	"github.com/syndtr/goleveldb/leveldb"
)

type CoreData struct {
	PendingEvent PendingEvents
	LevelDb *leveldb.DB
	Config *Configs
	PendingBlock []Blocks
	Model string
}

type Configs struct {
        Id int64
        Hash string
        V string
        Port string
	Datadir string
	ChainName string
	Peers []string
}

type PendingEvents struct {
	pendingTransaction []Transactions
}

type Transactions struct{
        BlockHash string `json:"blockHash"`
        Tx string `json:"tx"`
	From string `json:"from"`
        To string `json:"to"`
	Balance string `json:"balance"`
        Nonce int `json:"nonce"`
        Fee string `json:"fee"`
        Type string `json:"type"`
        Input string `json:"input"`
        Sign string `json:"sign"`
        PublicKey string `json:"publicKey"`
}

type Blocks struct{
        BlockNumber int `json:"blockNumber"`
        ParentHash string `json:"parentHash"`
        Transaction []Transactions `json:"transaction"`
        Timestamp int64 `json:"timestamp"`
        Hash string `json:"hash"`
	Txs int `json:"txs"`
        Verifier []Verifiers `json:"verifier"`
}

type Verifiers struct{
        Verifier string
        Sign string
}
