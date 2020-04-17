package document

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmCommonTx = "tx_common"
	TxStatusSuccess      = "success"
	TxStatusFail         = "fail"

	Tx_Field_Hash   = "tx_hash"
	Tx_Field_Type   = "type"
	Tx_Field_Status = "status"
	Tx_Field_Height = "height"
)

type CommonTx struct {
	Time   time.Time `bson:"time"`
	Height int64     `bson:"height"`
	TxHash string    `bson:"tx_hash"`
	From   string    `bson:"from"`
	To     string    `bson:"to"`
	// Amount     store.Coins     `bson:"amount"`
	Type string `bson:"type"`
	// Fee        store.Fee       `bson:"fee"`
	Memo      string  `bson:"memo"`
	Status    string  `bson:"status"`
	Code      uint32  `bson:"code"`
	Log       string  `bson:"log"`
	GasUsed   int64   `bson:"gas_used"`
	GasWanted int64   `bson:"gas_wanted"`
	GasPrice  float64 `bson:"gas_price"`
	// ActualFee  store.ActualFee `bson:"actual_fee"`
	ProposalId uint64 `bson:"proposal_id"`
	// Events     []Event `bson:"events"`

	Signers []Signer `bson:"signers"`

	Msgs []DocTxMsg `bson:"msgs"`
}

// type Event struct {
// 	Type       string   `bson:"type"`
// 	Attributes []KvPair `bson:"attributes"`
// }

type DocTxMsg struct {
	Type string `bson:"type"`
	Msg  Msg    `bson:"msg"`
}

type Msg interface {
	Type() string
	BuildMsg(msg interface{})
}

// Description
type ValDescription struct {
	Moniker  string `bson:"moniker"`
	Identity string `bson:"identity"`
	Website  string `bson:"website"`
	Details  string `bson:"details"`
}

type CommissionMsg struct {
	Rate          string `bson:"rate"`            // the commission rate charged to delegators
	MaxRate       string `bson:"max_rate"`        // maximum commission rate which validator can ever charge
	MaxChangeRate string `bson:"max_change_rate"` // maximum daily increase of the validator commission
}

type Signer struct {
	AddrHex    string `bson:"addr_hex"`
	AddrBech32 string `bson:"addr_bech32"`
}

func (d CommonTx) Name() string {
	return CollectionNmCommonTx
}

func (d CommonTx) PkKvPair() map[string]interface{} {
	return bson.M{Tx_Field_Hash: d.TxHash}
}

// func (d CommonTx) Query(query, fields bson.M, sort []string, skip, limit int) (
// 	results []CommonTx, err error) {
// 	exop := func(c *mgo.Collection) error {
// 		return c.Find(query).Sort(sort...).Select(fields).Skip(skip).Limit(limit).All(&results)
// 	}
// 	return results, store.ExecCollection(d.Name(), exop)
// }

// func (d CommonTx) CalculateTxGasAndGasPrice(txType string, limit int) (
// 	[]CommonTx, error) {
// 	query := bson.M{
// 		Tx_Field_Type:   txType,
// 		Tx_Field_Status: TxStatusSuccess,
// 	}
// 	fields := bson.M{}
// 	sort := []string{"-height"}
// 	skip := 0

// 	return d.Query(query, fields, sort, skip, limit)
// }
