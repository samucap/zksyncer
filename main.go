package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/utils"
)

var (
	//PrivateKey        = os.Getenv("PRIVATE_KEY")
	privateKey     = "da00d59541d59b4b257ea8a2702d221f7deecc2025a86d3b9e66bb43d54f3532"
	zkSyncProvider = "https://sepolia.era.zksync.dev"
	ethProvider    = "https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
)

func main() {
	// Connect to zkSync network
	zkClient, err := clients.DialContext(context.Background(), zkSyncProvider)
	errHandler(err)
	defer zkClient.Close()

	// Connect to Ethereum network
	ethCl, err := ethclient.DialContext(context.Background(), ethProvider)
	errHandler(err)
	defer ethCl.Close()

	acct := common.HexToAddress(privateKey)
	bal, err := ethCl.BalanceAt(context.Background(), acct, nil)
	errHandler(err)
	printer(weiToEther(bal))

	w, err := accounts.NewWallet(common.Hex2Bytes(privateKey), &zkClient, ethCl)
	errHandler(err)
	// Show balance before deposit
	balance, err := w.Balance(context.Background(), utils.EthAddress, nil)
	errHandler(err)
	printer(weiToEther(balance))

	// Perform deposit
	tx, err := w.Deposit(nil, accounts.DepositTransaction{
		Token:  utils.EthAddress,
		Amount: toWei(0.001, 18),
		To:     w.Address(),
	})
	errHandler(err)

	fmt.Println("L1 transaction: ", tx.Hash())

	// Wait for deposit transaction to be finalized on L1 network
	fmt.Println("Waiting for deposit transaction to be finalized on L1 network")
	_, err = bind.WaitMined(context.Background(), ethCl, tx)
	errHandler(err)

	// Get transaction receipt for deposit transaction on L1 network
	l1Receipt, err := ethCl.TransactionReceipt(context.Background(), tx.Hash())
	errHandler(err)

	// Get deposit transaction on L2 network
	l2Tx, err := zkClient.L2TransactionFromPriorityOp(context.Background(), l1Receipt)
	errHandler(err)

	fmt.Println("L2 transaction", l2Tx.Hash)

	// Wait for deposit transaction to be finalized on L2 network (5-7 minutes)
	fmt.Println("Waiting for deposit transaction to be finalized on L2 network (5-7 minutes)")
	_, err = zkClient.WaitMined(context.Background(), l2Tx.Hash)
	errHandler(err)

	balance, err = w.Balance(context.Background(), utils.EthAddress, nil)
	errHandler(err)

	fmt.Println("Balance after deposit: ", weiToEther(balance))
}
