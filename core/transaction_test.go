package core

import(
	"testing"
)

func TestTransaction(t *testing.T){
	exampleTransaction := NewTransactionAPI()
	if(exampleTransaction.EncodeRawTransaction()!= false){
		t.Error("EncodeRawTransaction is not match:", exampleTransaction.EncodeRawTransaction())
	}
}
