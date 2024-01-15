package accounts

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network struct {
	Chain         string
	ProviderURL   string
	ContractAdrrs map[string]common.Address
	TokenAddrs    map[string]common.Address
}

type Account struct {
	Addr     common.Address
	PK       string
	L1Client *ethclient.Client
	L2Client *ethclient.Client
	Network  Network
}

func (acct *Account) PubKeyAddr() common.Address {
	privKey, err := crypto.HexToECDSA(acct.PK)
	errHandler(err)

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*pubKeyECDSA)
}

func getClient(ctx context.Context, netProvider string) *ethclient.Client {
	cl, err := ethclient.DialContext(ctx, netProvider)
	errHandler(err)

	return cl
}

func NewAccount(ctx context.Context, pk string, chain string) *Account {
	acct := &Account{
		Addr: getAddr(),
		PK:   pk,
		Network: &Network{
			Chain: chain,
		},
	}

	return acct
}
