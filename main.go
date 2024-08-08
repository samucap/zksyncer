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
