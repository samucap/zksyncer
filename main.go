package main

import (
	"context"
	"log"
	"os"

	"github.com/samucap/zksyncer/contas"
)

func main() {
	pk := os.Getenv("P_K")
	if len(pk) == 0 {
		log.Fatalf("Invalid required private key for transactions: %v", pk)
	}

	chain := "eraTestnet"
	ctx := context.Background()
	acct := contas.NewAccount(ctx, pk, chain)
	acct.GetTransaction(ctx)
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
