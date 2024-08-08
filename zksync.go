package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/utils"
)

// bridgeToL2(w, zkClient, ethCl)
func bridgeToL2(w *accounts.Wallet, zkClient clients.Client, ethCl *ethclient.Client) {
	// Show balance before deposit
	balance, err := w.Balance(context.Background(), utils.EthAddress, nil)
	errHandler(err)
	printer(weiToEther(balance))

	// Perform deposit
	tx, err := w.Deposit(nil, accounts.DepositTransaction{
		Token:  utils.EthAddress,
		Amount: toWei(0.005, 18),
		To:     w.Address(),
	})
	errHandler(err)

	fmt.Println("L1 transaction hash: ", tx.Hash())

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
