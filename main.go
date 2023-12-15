package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/zksync-sdk/zksync2-go/accounts"

	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	//"context"
	//"fmt"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/zksync-sdk/zksync2-go/utils"
	//"math/big"
	//"os"
)

var (
	//PrivateKey        = os.Getenv("PRIVATE_KEY")
	privateKey     = "0x8c526e46c473bdc280b840c7f61ca53a12f5efc3dd0b5496d06b40c8d033cea6"
	zkSyncProvider = "https://sepolia.era.zksync.dev"
	ethProvider    = "https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
)

func main() {
	// Connect to zkSync network
	client, err := clients.Dial(zkSyncProvider)
	errHandler(err)
	defer client.Close()

	// Connect to Ethereum network
	ethClient, err := ethclient.Dial(ethProvider)
	errHandler(err)
	defer ethClient.Close()

	ctx := context.Background()
	fmt.Println(client.BridgeContracts(ctx))

	chainID, err := client.ChainID(context.Background())
	errHandler(err)
	wallet, err := accounts.NewWalletFromRawPrivateKey(common.FromHex(privateKey), chainID.Int64(), &client, ethClient)
	errHandler(err)
	printer([]interface{}{wallet, chainID.Int64()})
}
