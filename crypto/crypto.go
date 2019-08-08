package crypto

type CryptoAPI interface {
        //GeneratePrivateKey()
        //SubString(number2 string)
}

type Crypto struct {
        PrivateKey interface{}
	PublicKey interface{}
	Address interface{}
}

func NewCrypto() CryptoAPI{
        return &Crypto{
                PrivateKey:"",
        }
}

func (C *Crypto)GeneratePrivateKey() CryptoAPI{
	return NewCrypto()
}


