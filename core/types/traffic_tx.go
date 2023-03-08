package types

import (
	"factor-mx/common"
	"math/big"
)

type TrafficTx struct {
	ChainID        *big.Int        // destination chain ID
	Nonce          uint64          // nonce of sender account
	GasPrice       *big.Int        // wei per gas
	Gas            uint64          // gas limit
	Violator       *common.Address // violator address
	CentralAccount *common.Address // Decentralized traffic control system address
	Penalty        *big.Int        // wei amount penalty
	Data           []byte          // tx meta data
	AccessList     AccessList      // Central Manager List
	V, R, S        *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *TrafficTx) copy() TxData {
	cpy := &TrafficTx{
		Nonce:          tx.Nonce,
		Violator:       copyAddressPtr(tx.Violator),
		CentralAccount: copyAddressPtr(tx.CentralAccount),
		Data:           common.CopyBytes(tx.Data),
		Gas:            tx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		Penalty:    new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Penalty != nil {
		cpy.Penalty.Set(tx.Penalty)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *TrafficTx) txType() byte           { return AccessListTxType }
func (tx *TrafficTx) chainID() *big.Int      { return tx.ChainID }
func (tx *TrafficTx) accessList() AccessList { return tx.AccessList }
func (tx *TrafficTx) data() []byte           { return tx.Data }
func (tx *TrafficTx) gas() uint64            { return tx.Gas }
func (tx *TrafficTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *TrafficTx) gasTipCap() *big.Int    { return tx.GasPrice }
func (tx *TrafficTx) gasFeeCap() *big.Int    { return tx.GasPrice }
func (tx *TrafficTx) value() *big.Int        { return tx.Penalty }
func (tx *TrafficTx) nonce() uint64          { return tx.Nonce }
func (tx *TrafficTx) to() *common.Address    { return tx.Violator }

func (tx *TrafficTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *TrafficTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}
