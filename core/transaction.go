package core

const (
	result = true
)

type Transaction interface {
	EncodeRawTransaction() bool
	Sign() bool
}

func NewTransactionAPI() Transaction {
	return &TransactionData{}
}

type TransactionData struct{
	Result bool
}

func (T *TransactionData) EncodeRawTransaction() bool{
	return T.Result
}

func (T *TransactionData) EncodeForSign() bool{
	return T.Result
}

func (T *TransactionData) Sign() bool{
	return true
}

func (T *TransactionData) VerifySig() bool{
	return true
}

func (T *TransactionData) VerifyAccount() bool{
	return true
}


