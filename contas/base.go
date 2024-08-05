package contas

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/samucap/zksyncer/contracts/syncswap/classicPoolFactory"
)

type Network struct {
	ChainID       *big.Int
	ProviderURL   string
	ContractAddrs map[string]common.Address
	TokenAddrs    map[string]common.Address
}

var NetProviders map[string]Network = map[string]Network{
	"eraTestnet": {
		ProviderURL: "https://zksync2-testnet.zksync.dev",
		ContractAddrs: map[string]common.Address{
			"vault":              common.HexToAddress("0x6C41e61C5449B0916103d536c638E470bbabf95e"),
			"poolMaster":         common.HexToAddress("0x4bEAC45efEE4DfB45f7397C19c13a89611dE193D"),
			"classicPoolFactory": common.HexToAddress("0xeaeDD100A9e61CE0412664CE598F0a624CFB3ccB"),
			"stablePoolFactory":  common.HexToAddress("0x51372e0F0BfAa5be848762AE828D1a24bb0d7Fa0"),
			"router":             common.HexToAddress("0x4dbcd68F735e91ccBa5dff2d4DAb7B0729BBc1a4"),
		},
		TokenAddrs: map[string]common.Address{
			"weth": common.HexToAddress("0x20b28b1e4665fff290650586ad76e977eab90c5d"),
			"eth":  common.HexToAddress("0x20b28b1e4665fff290650586ad76e977eab90c5d"),
			"usdc": common.HexToAddress("0x0faF6df7054946141266420b43783387A78d82A9"),
		},
	},
}

type Account struct {
	Addr     common.Address
	PK       string
	L1Client *ethclient.Client
	L2Client *ethclient.Client
	Network  Network
}

func (acct *Account) PubKeyAddr() {
	privKey, err := crypto.HexToECDSA(acct.PK)
	errHandler(err)

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	acct.Addr = crypto.PubkeyToAddress(*pubKeyECDSA)
}

func getClient(ctx context.Context, netProvider string) *ethclient.Client {
	cl, err := ethclient.DialContext(ctx, netProvider)
	errHandler(err)

	return cl
}

func NewAccount(ctx context.Context, pk string, chain string) *Account {
	acct := &Account{
		PK:      pk,
		Network: NetProviders[chain],
	}

	acct.L2Client = getClient(ctx, acct.Network.ProviderURL)
	chainID, err := acct.L2Client.ChainID(ctx)
	errHandler(err)
	acct.Network.ChainID = chainID

	acct.PubKeyAddr()
	return acct
}

func (acct *Account) GetTransaction(ctx context.Context) *bind.TransactOpts {
	privKey, err := crypto.HexToECDSA(acct.PK)
	errHandler(err)
	tx, err := bind.NewKeyedTransactorWithChainID(privKey, acct.Network.ChainID)
	errHandler(err)
	nonce, err := acct.L2Client.PendingNonceAt(ctx, acct.Addr)
	errHandler(err)
	tx.Nonce = nonce
	pools := classicPoolFactory.NewClassicPoolFactory()
	pool := pools.GetPool(acct.Network.TokenAddrs.weth, acct.Network.TokenAddrs.usdc)

	fmt.Println(">>>>>>>>>>>", pool)
	return &tx

	//tx := &types.Transaction712{
	//	Nonce: nonce,
	//	To: to,
	//	Value: value,
	//	ChainID: acct.Network.ChainID,
	//}
	// getNonce from client
	//calc gas

	//

}

func errHandler(err error) {
	if err != nil {
		fmt.Printf("Error >>>>>>>>>>>> %+v\n", err)
		panic(err)
	}
}
