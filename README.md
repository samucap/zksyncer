#### zkSync Transacions

1. Bridge ETH to zkSync bridge
2. swap syncswap and other swaps
3. Transfer to liquidity pool

#### Dependencies
    * go >= go version in go.mod

#### Install
go install

#### Running
go run *.go

#### TODO
    - make new accounts
        - 
    - fund new accounts
    - abstract functions to use go-ethereum only


######## steps
centralized exchange -> eth account -> bridge to zksync -> transactions
    * Account manager with Clef

############ ethRPCProviders
    * Infura
        * URL: https://sepolia.infura.io/v3/de4d9283b6f04c07a5d4d5d16ca31291
        * Limit: 100,000 requests/day (1.15requests/s); 25,000 Ethereum Mainnet Archive Requests/Day
    * Alchemy
        * URL: https://eth-sepolia.g.alchemy.com/v2/CjnjKlh6O1s0NOi9qP86MsCcwULTKQqM
        * Limit: 330 compute units/s, 40 req/min ??
            * https://docs.alchemy.com/reference/throughput
    * Ankr
        * URL: https://rpc.ankr.com/eth_sepolia
        * Limit: 1800 req/min
            * https://www.ankr.com/docs/rpc-service/service-plans/#rate-limits
    * Quicknode
        * URL: https://restless-delicate-tent.ethereum-sepolia.quiknode.pro/a12d29f2c3d0cd9ce3cdd905638ec8b1fb7b42f9
        * Liimit: 330 credit/s; 50M / month
            * https://www.quicknode.com/pricing

    seems like response 429 is too many requests error
