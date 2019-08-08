package common

import(
        "math/big"
)

type CalculationAPI interface {
        AddString(number2 string)
        SubString(number2 string)
        MulString(number2 string)
}

type Calculation struct {
        Number1 interface{}
}

func NewCalculation(num interface{}) CalculationAPI{
        return &Calculation{
                Number1:num,
        }
}

func (C *Calculation)AddString(number2 string){
        bignumber1 := new(big.Int)
        bignumber1.SetString(C.Number1.(string), 10)
        bignumber2 := new(big.Int)
        bignumber2.SetString(number2, 10)
        C.Number1 = bignumber1.Add(bignumber1, bignumber2).String()
}

func (C *Calculation)SubString(number2 string){
        bignumber1 := new(big.Int)
        bignumber1.SetString(C.Number1.(string), 10)
        bignumber2 := new(big.Int)
        bignumber2.SetString(number2, 10)
        C.Number1 = bignumber1.Sub(bignumber1, bignumber2).String()
}

func (C *Calculation)MulString(number2 string){
        bignumber1 := new(big.Int)
        bignumber1.SetString(C.Number1.(string), 10)
        bignumber2 := new(big.Int)
        bignumber2.SetString(number2, 10)
        C.Number1 = bignumber1.Mul(bignumber1, bignumber2).String()
}

