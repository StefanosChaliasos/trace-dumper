// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yul

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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

// YulMetaData contains all meta data concerning the Yul contract.
var YulMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x601c600d600039601c6000f3fe60003560e01c63f8a8fd6d8103601a57601560570180600052505b50",
}

// YulABI is the input ABI used to generate the binding from.
// Deprecated: Use YulMetaData.ABI instead.
var YulABI = YulMetaData.ABI

// YulBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YulMetaData.Bin instead.
var YulBin = YulMetaData.Bin

// DeployYul deploys a new Ethereum contract, binding an instance of Yul to it.
func DeployYul(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Yul, error) {
	parsed, err := YulMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YulBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Yul{YulCaller: YulCaller{contract: contract}, YulTransactor: YulTransactor{contract: contract}, YulFilterer: YulFilterer{contract: contract}}, nil
}

// Yul is an auto generated Go binding around an Ethereum contract.
type Yul struct {
	YulCaller     // Read-only binding to the contract
	YulTransactor // Write-only binding to the contract
	YulFilterer   // Log filterer for contract events
}

// YulCaller is an auto generated read-only Go binding around an Ethereum contract.
type YulCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YulTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YulTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YulFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YulFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YulSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YulSession struct {
	Contract     *Yul              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YulCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YulCallerSession struct {
	Contract *YulCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YulTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YulTransactorSession struct {
	Contract     *YulTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YulRaw is an auto generated low-level Go binding around an Ethereum contract.
type YulRaw struct {
	Contract *Yul // Generic contract binding to access the raw methods on
}

// YulCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YulCallerRaw struct {
	Contract *YulCaller // Generic read-only contract binding to access the raw methods on
}

// YulTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YulTransactorRaw struct {
	Contract *YulTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYul creates a new instance of Yul, bound to a specific deployed contract.
func NewYul(address common.Address, backend bind.ContractBackend) (*Yul, error) {
	contract, err := bindYul(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yul{YulCaller: YulCaller{contract: contract}, YulTransactor: YulTransactor{contract: contract}, YulFilterer: YulFilterer{contract: contract}}, nil
}

// NewYulCaller creates a new read-only instance of Yul, bound to a specific deployed contract.
func NewYulCaller(address common.Address, caller bind.ContractCaller) (*YulCaller, error) {
	contract, err := bindYul(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YulCaller{contract: contract}, nil
}

// NewYulTransactor creates a new write-only instance of Yul, bound to a specific deployed contract.
func NewYulTransactor(address common.Address, transactor bind.ContractTransactor) (*YulTransactor, error) {
	contract, err := bindYul(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YulTransactor{contract: contract}, nil
}

// NewYulFilterer creates a new log filterer instance of Yul, bound to a specific deployed contract.
func NewYulFilterer(address common.Address, filterer bind.ContractFilterer) (*YulFilterer, error) {
	contract, err := bindYul(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YulFilterer{contract: contract}, nil
}

// bindYul binds a generic wrapper to an already deployed contract.
func bindYul(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YulMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yul *YulRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yul.Contract.YulCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yul *YulRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yul.Contract.YulTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yul *YulRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yul.Contract.YulTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yul *YulCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yul.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yul *YulTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yul.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yul *YulTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yul.Contract.contract.Transact(opts, method, params...)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Yul *YulTransactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yul.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Yul *YulSession) Test() (*types.Transaction, error) {
	return _Yul.Contract.Test(&_Yul.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_Yul *YulTransactorSession) Test() (*types.Transaction, error) {
	return _Yul.Contract.Test(&_Yul.TransactOpts)
}
