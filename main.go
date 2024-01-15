package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts/types"
)

type ChainInfo struct {
	ProviderURL   string
	ContractAddrs map[string]common.Address
}

var NetProviders map[string]string = map[string]string{
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

func main() {
	pk := os.Getenv("P_K")
	if len(pk) == 0 {
		log.Fatalf("Invalid required private key for transactions: %v", pk)
	}

	chain := "eraTestnet"
	ctx := context.Background()
}

//type Transaction712 struct {
//	Nonce      *big.Int         // Nonce to use for the transaction execution.
//	GasTipCap  *big.Int         // EIP-1559 tip per gas.
//	GasFeeCap  *big.Int         // EIP-1559 fee cap per gas.
//	Gas        *big.Int         // Gas limit to set for the transaction execution.
//	To         *common.Address  // The address of the recipient.
//	Value      *big.Int         // Funds to transfer along the transaction (nil = 0 = no funds).
//	Data       hexutil.Bytes    // Input data, usually an ABI-encoded contract method invocation.
//	AccessList types.AccessList // EIP-2930 access list.
//
//	ChainID *big.Int        // Chain ID of the network.
//	From    *common.Address // The address of the sender.
//	Meta    *Eip712Meta     // EIP-712 metadata.
//}
//type Eip712Meta struct {
//	// GasPerPubdata denotes the maximum amount of gas the user is willing
//	// to pay for a single byte of pubdata.
//	GasPerPubdata *hexutil.Big `json:"gasPerPubdata,omitempty"`
//	// CustomSignature is used for the cases in which the signer's account
//	// is not an EOA.
//	CustomSignature hexutil.Bytes `json:"customSignature,omitempty"`
//	// FactoryDeps is a non-empty array of bytes. For deployment transactions,
//	// it should contain the bytecode of the contract being deployed.
//	// If the contract is a factory contract, i.e. it can deploy other contracts,
//	// the array should also contain the bytecodes of the contracts which it can deploy.
//	FactoryDeps []hexutil.Bytes `json:"factoryDeps"`
//	// PaymasterParams contains parameters for configuring the custom paymaster
//	// for the transaction.
//	PaymasterParams *PaymasterParams \`json:"paymasterParams,omitempty"`
//}

func GetTransaction(ctx context.Context, from, to common.Address, value *big.Int) *types.Transaction712 {
	//tx := &types.Transaction712{
	//	Nonce:
	//	GasTipCap  *big.Int         // EIP-1559 tip per gas.
	//	GasFeeCap  *big.Int         // EIP-1559 fee cap per gas.
	//	Gas        *big.Int         // Gas limit to set for the transaction execution.
	//	To         *common.Address  // The address of the recipient.
	//	Value      *big.Int         // Funds to transfer along the transaction (nil = 0 = no funds).
	//	Data       hexutil.Bytes    // Input data, usually an ABI-encoded contract method invocation.
	//	AccessList types.AccessList // EIP-2930 access list.
	//	ChainID *big.Int        // Chain ID of the network.
	//	From    *common.Address // The address of the sender.
	//	Meta    *types.Eip712Meta     // EIP-712 metadata.
	//}
	// getNonce from client
	//calc gas

	//

}
