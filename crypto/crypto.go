package crypto
import(
)

type CryptoAPI interface {
	PrivateKeys()(interface{})
	PublicKeys()(interface{})
	Addresses()(interface{})
}

type Crypto struct {
        PrivateKey interface{}
	PublicKey interface{}
	Address interface{}
}

func NewCrypto(key string) CryptoAPI{
        return &Crypto{
		PrivateKey:key,
        }
}

func GenerateKey() CryptoAPI{
	return &Crypto{
		PrivateKey:"",
		PublicKey:"",
		Address:"",
	}
}

func (C *Crypto) PrivateKeys()(interface{}){
	return C.PrivateKey
}

func (C *Crypto) PublicKeys()(interface{}){
	return C.PublicKey
}

func (C *Crypto) Addresses()(interface{}){
	return C.Address
}
