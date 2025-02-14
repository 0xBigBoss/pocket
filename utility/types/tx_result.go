package types

import (
	"github.com/pokt-network/pocket/shared/codec"
	"github.com/pokt-network/pocket/shared/crypto"
	"github.com/pokt-network/pocket/shared/modules"
)

// INVESTIGATE: Look into a way of removing this type altogether or from shared interfaces.

var _ modules.TxResult = &DefaultTxResult{}

func (txr *DefaultTxResult) Bytes() ([]byte, error) {
	return codec.GetCodec().Marshal(txr)
}

func (*DefaultTxResult) FromBytes(bz []byte) (modules.TxResult, error) {
	result := new(DefaultTxResult)
	if err := codec.GetCodec().Unmarshal(bz, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (txr *DefaultTxResult) Hash() ([]byte, error) {
	bz, err := txr.Bytes()
	if err != nil {
		return nil, err
	}
	return txr.HashFromBytes(bz)
}

func (txr *DefaultTxResult) HashFromBytes(bz []byte) ([]byte, error) {
	return crypto.SHA3Hash(bz), nil
}

func (tx *Transaction) ToTxResult(height int64, index int, signer, recipient, msgType string, err Error) (*DefaultTxResult, Error) {
	txBytes, er := tx.Bytes()
	if er != nil {
		return nil, ErrProtoMarshal(er)
	}
	code, errString := int32(0), ""
	return &DefaultTxResult{
		Tx:            txBytes,
		Height:        height,
		Index:         int32(index),
		ResultCode:    code,
		Error:         errString,
		SignerAddr:    signer,
		RecipientAddr: recipient,
		MessageType:   msgType,
	}, nil
}
