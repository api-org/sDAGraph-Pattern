package crypto

import(
	"testing"
)

func TestGenerateKey(t *testing.T){
	key := NewCrypto("123")
	t.Error(key.PrivateKeys())
	t.Error(key.PublicKeys())
	t.Error(key.Addresses())
}
