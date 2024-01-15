package main

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	//zkSyncProvider = "https://sepolia.era.zksync.dev"
	//ethProvider    = "https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
	//ethProvider = "https://eth-sepolia.g.alchemy.com/v2/CjnjKlh6O1s0NOi9qP86MsCcwULTKQqM"
	//ethProvider = "https://goerli.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
	//ethProvider = "https://eth-goerli.g.alchemy.com/v2/TwC0w-mQvNL4ANOpLgEp1uTEXJJBNNRn"
	//ethProvider = "https://rpc.ankr.com/eth_sepolia"
	//"https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291"
	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ethAddr  = common.HexToAddress("0x000000000000000000000000000000000000800A")
	weth     = "0x20b28b1e4665fff290650586ad76e977eab90c5d"
	eth      = "0x20b28b1e4665fff290650586ad76e977eab90c5d"
)

//// get token pair pool
//poolFactory, err := classicPoolFactory.NewClassicPoolFactory(common.HexToAddress(), zkCl)
//errHandler(err)

//poolAddr, err := poolFactory.GetPool(&bind.CallOpts{Pending: false}, acct.TokenAddrs.eth, acct.TokenAddrs.usdc)
//errHandler(err)

//if poolAddr == zeroAddr {
//	log.Fatalf("Unable to retrieve swap pool for %v -> %v")
//}

//Pending: In the mempool but not yet included in a block.
//Included: Included in a block but the batch containing the block has not yet been committed.
//Verified: Included in a block and verified. Verified means the transaction has been committed, proven, and executed on the Ethereum L1 network.
//Failed: Unverified/failed transaction.
