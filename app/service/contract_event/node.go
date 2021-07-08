/**
@author: 铁柱
@date:2021/6/29
@note:
**/
package contract_event

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"nft_service/app/model"
)

type TestUserService struct {
}

func (m *TestUserService) Generate(count int) error {
	var addrs []model.TestUser

	for i := 0; i <= count; i++ {
		key, addr := generateAddr()
		addrs = append(addrs, model.TestUser{
			Addr: addr,
			Key:  key,
		})
	}

	return model.AddUsers(&addrs)
}

func (m *TestUserService) ClearAll() error {

	return model.CleanAll()
}

func generateAddr() (key, addr string) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {

	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	key = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	addr = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}
