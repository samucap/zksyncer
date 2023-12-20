package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/utils"
)

var (
	//PrivateKey        = os.Getenv("PRIVATE_KEY")
	privateKey     = "da00d59541d59b4b257ea8a2702d221f7deecc2025a86d3b9e66bb43d54f3532"
	zkSyncProvider = "https://sepolia.era.zksync.dev"
	//ethProvider    = "https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
	ethProvider = "https://eth-sepolia.g.alchemy.com/v2/CjnjKlh6O1s0NOi9qP86MsCcwULTKQqM"
)

type RequestExecuteTransaction struct {
	ContractAddress common.Address // The L2 receiver address.
	Calldata        []byte         // The input of the L2 transaction.
	L2GasLimit      *big.Int       // Maximum amount of L2 gas that transaction can consume during execution on L2.
	L2Value         *big.Int       // `msg.value` of L2 transaction.
	FactoryDeps     [][]byte       // An array of L2 bytecodes that will be marked as known on L2.

	// If the ETH value passed with the transaction is not explicitly stated Auth.Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address
}

func main() {
	//// Connect to zkSync network
	zkClient, err := clients.DialContext(context.Background(), zkSyncProvider)
	errHandler(err)
	defer zkClient.Close()

	// Connect to Ethereum network
	ethCl, err := ethclient.DialContext(context.Background(), ethProvider)
	errHandler(err)
	defer ethCl.Close()

	// Connect to zksync network
	//zkCl, err := rpc.DialContext(context.Background(), zkSyncProvider)
	//errHandler(err)

	//gasPrice, err := ethCl.SuggestGasPrice(context.Background())
	//errHandler(err)

	l2GasLimit := big.NewInt(90000)

	//func (_IZkSync *IZkSyncTransactor) RequestL2Transaction(opts *bind.TransactOpts, _contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	//zks, err := NewIZkSync(common.HexToAddress("0x9a6de0f62Aa270A8bCB1e2610078650D539B1Ef9"), zkCl)
	//ethCl.SuggestGasPrice(context.Background())
	pk, err := crypto.HexToECDSA(privateKey)
	errHandler(err)

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		errHandler(errors.New("error casting public key to ECDSA"))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethCl.PendingNonceAt(context.Background(), fromAddress)
	errHandler(err)

	//gasPrice, err := ethCl.SuggestGasPrice(context.Background())
	//errHandler(err)

	//chainID, err := ethCl.ChainID(context.Background())
	//errHandler(err)
	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(300))
	errHandler(err)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = toWei(0.001, 18)  // in wei
	auth.GasLimit = uint64(700000) // in units
	auth.GasPrice = big.NewInt(10000000000)
	requestTx := RequestExecuteTransaction{
		ContractAddress:   common.HexToAddress("0x9a6de0f62Aa270A8bCB1e2610078650D539B1Ef9"),
		L2Value:           toWei(0.01, 18),
		L2GasLimit:        l2GasLimit,
		GasPerPubdataByte: big.NewInt(800),
		OperatorTip:       big.NewInt(800),
		RefundRecipient:   fromAddress,
	}

	printer("sending to " + fromAddress.String())
	zksync, err := NewIZkSync(common.HexToAddress("0x9a6de0f62Aa270A8bCB1e2610078650D539B1Ef9"), zkClient)
	errHandler(err)
	trxx, err := zksync.RequestL2Transaction(auth, fromAddress, requestTx.L2Value, nil, l2GasLimit, big.NewInt(80000), nil, fromAddress)
	errHandler(err)
	printer(trxx.Hash())
	printer("waiting for mined....")
	_, err = bind.WaitMined(context.Background(), ethCl, trxx)
	if err != nil {
		log.Panic(err)
	}

	// Get transaction receipt for deposit transaction on L1 network
	l1Receipt, err := ethCl.TransactionReceipt(context.Background(), trxx.Hash())
	if err != nil {
		log.Panic(err)
	}
	printer(l1Receipt)
	// 0x681A1AFdC2e06776816386500D2D461a6C96cB45
	//0x9a6de0f62Aa270A8bCB1e2610078650D539B1Ef9
	//a := "0x2470ad8A20BECc539b20D026e750D9434335E351"
	////a := "0x84DbCC0B82124bee38e3Ce9a92CdE2f943bab60D"
	//contractAddr := common.HexToAddress(a)
	//l1Bridge, err := NewL1Bridge(contractAddr, ethCl)
	//errHandler(err)
	////, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int
	//tx, err := l1Bridge.Deposit(auth, fromAddress, fromAddress, toWei(0.001, 18), big.NewInt(700000), big.NewInt(50_000))
	//errHandler(err)
	//printer(tx.Hash())

	//w, err := accounts.NewWallet(common.Hex2Bytes(privateKey), &zkClient, ethCl)
	//errHandler(err)

	//bridgeToL2(w, zkClient, ethCl)
}

func bridgeToL2(w *accounts.Wallet, zkClient clients.Client, ethCl *ethclient.Client) {
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
