// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IMailboxL2CanonicalTransaction is an auto generated low-level Go binding around an user-defined struct.
type IMailboxL2CanonicalTransaction struct {
	TxType                 *big.Int
	From                   *big.Int
	To                     *big.Int
	GasLimit               *big.Int
	GasPerPubdataByteLimit *big.Int
	MaxFeePerGas           *big.Int
	MaxPriorityFeePerGas   *big.Int
	Paymaster              *big.Int
	Nonce                  *big.Int
	Value                  *big.Int
	Reserved               [4]*big.Int
	Data                   []byte
	Signature              []byte
	FactoryDeps            []*big.Int
	PaymasterInput         []byte
	ReservedDynamic        []byte
}

// L2Log is an auto generated low-level Go binding around an user-defined struct.
type L2Log struct {
	L2ShardId       uint8
	IsService       bool
	TxNumberInBatch uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// L2Message is an auto generated low-level Go binding around an user-defined struct.
type L2Message struct {
	TxNumberInBatch uint16
	Sender          common.Address
	Data            []byte
}

// IZkSyncMetaData contains all meta data concerning the IZkSync contract.
var IZkSyncMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"factoryDeps\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIMailbox.L2CanonicalTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeEthWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"requestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IZkSyncABI is the input ABI used to generate the binding from.
// Deprecated: Use IZkSyncMetaData.ABI instead.
var IZkSyncABI = IZkSyncMetaData.ABI

// IZkSync is an auto generated Go binding around an Ethereum contract.
type IZkSync struct {
	IZkSyncCaller     // Read-only binding to the contract
	IZkSyncTransactor // Write-only binding to the contract
	IZkSyncFilterer   // Log filterer for contract events
}

// IZkSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZkSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZkSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZkSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZkSyncSession struct {
	Contract     *IZkSync          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZkSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZkSyncCallerSession struct {
	Contract *IZkSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IZkSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZkSyncTransactorSession struct {
	Contract     *IZkSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IZkSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZkSyncRaw struct {
	Contract *IZkSync // Generic contract binding to access the raw methods on
}

// IZkSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZkSyncCallerRaw struct {
	Contract *IZkSyncCaller // Generic read-only contract binding to access the raw methods on
}

// IZkSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZkSyncTransactorRaw struct {
	Contract *IZkSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZkSync creates a new instance of IZkSync, bound to a specific deployed contract.
func NewIZkSync(address common.Address, backend bind.ContractBackend) (*IZkSync, error) {
	contract, err := bindIZkSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZkSync{IZkSyncCaller: IZkSyncCaller{contract: contract}, IZkSyncTransactor: IZkSyncTransactor{contract: contract}, IZkSyncFilterer: IZkSyncFilterer{contract: contract}}, nil
}

// NewIZkSyncCaller creates a new read-only instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncCaller(address common.Address, caller bind.ContractCaller) (*IZkSyncCaller, error) {
	contract, err := bindIZkSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncCaller{contract: contract}, nil
}

// NewIZkSyncTransactor creates a new write-only instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*IZkSyncTransactor, error) {
	contract, err := bindIZkSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncTransactor{contract: contract}, nil
}

// NewIZkSyncFilterer creates a new log filterer instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*IZkSyncFilterer, error) {
	contract, err := bindIZkSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZkSyncFilterer{contract: contract}, nil
}

// bindIZkSync binds a generic wrapper to an already deployed contract.
func bindIZkSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZkSyncMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSync *IZkSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSync.Contract.IZkSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSync *IZkSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.Contract.IZkSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSync *IZkSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSync.Contract.IZkSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSync *IZkSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSync *IZkSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSync *IZkSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSync.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncSession) GetName() (string, error) {
	return _IZkSync.Contract.GetName(&_IZkSync.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncCallerSession) GetName() (string, error) {
	return _IZkSync.Contract.GetName(&_IZkSync.CallOpts)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) pure returns(uint256)
func (_IZkSync *IZkSyncCaller) L2TransactionBaseCost(opts *bind.CallOpts, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "l2TransactionBaseCost", _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) pure returns(uint256)
func (_IZkSync *IZkSyncSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSync.Contract.L2TransactionBaseCost(&_IZkSync.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) pure returns(uint256)
func (_IZkSync *IZkSyncCallerSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSync.Contract.L2TransactionBaseCost(&_IZkSync.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2LogInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2LogInclusion", _batchNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2MessageInclusion", _batchNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _batchNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _batchNumber, _index, _message, _proof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactor) FinalizeEthWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "finalizeEthWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactorSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncTransactor) RequestL2Transaction(opts *bind.TransactOpts, _contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "requestL2Transaction", _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.RequestL2Transaction(&_IZkSync.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncTransactorSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.RequestL2Transaction(&_IZkSync.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// IZkSyncEthWithdrawalFinalizedIterator is returned from FilterEthWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for EthWithdrawalFinalized events raised by the IZkSync contract.
type IZkSyncEthWithdrawalFinalizedIterator struct {
	Event *IZkSyncEthWithdrawalFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IZkSyncEthWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncEthWithdrawalFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IZkSyncEthWithdrawalFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IZkSyncEthWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncEthWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncEthWithdrawalFinalized represents a EthWithdrawalFinalized event raised by the IZkSync contract.
type IZkSyncEthWithdrawalFinalized struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawalFinalized is a free log retrieval operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSync *IZkSyncFilterer) FilterEthWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address) (*IZkSyncEthWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncEthWithdrawalFinalizedIterator{contract: _IZkSync.contract, event: "EthWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawalFinalized is a free log subscription operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSync *IZkSyncFilterer) WatchEthWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IZkSyncEthWithdrawalFinalized, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncEthWithdrawalFinalized)
				if err := _IZkSync.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthWithdrawalFinalized is a log parse operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSync *IZkSyncFilterer) ParseEthWithdrawalFinalized(log types.Log) (*IZkSyncEthWithdrawalFinalized, error) {
	event := new(IZkSyncEthWithdrawalFinalized)
	if err := _IZkSync.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the IZkSync contract.
type IZkSyncNewPriorityRequestIterator struct {
	Event *IZkSyncNewPriorityRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IZkSyncNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewPriorityRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IZkSyncNewPriorityRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IZkSyncNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewPriorityRequest represents a NewPriorityRequest event raised by the IZkSync contract.
type IZkSyncNewPriorityRequest struct {
	TxId                *big.Int
	TxHash              [32]byte
	ExpirationTimestamp uint64
	Transaction         IMailboxL2CanonicalTransaction
	FactoryDeps         [][]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSync *IZkSyncFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*IZkSyncNewPriorityRequestIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewPriorityRequestIterator{contract: _IZkSync.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSync *IZkSyncFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *IZkSyncNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewPriorityRequest)
				if err := _IZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriorityRequest is a log parse operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSync *IZkSyncFilterer) ParseNewPriorityRequest(log types.Log) (*IZkSyncNewPriorityRequest, error) {
	event := new(IZkSyncNewPriorityRequest)
	if err := _IZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
