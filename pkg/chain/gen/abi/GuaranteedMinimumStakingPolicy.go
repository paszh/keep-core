// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// GuaranteedMinimumStakingPolicyMetaData contains all meta data concerning the GuaranteedMinimumStakingPolicy contract.
var GuaranteedMinimumStakingPolicyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_now\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"grantedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawn\",\"type\":\"uint256\"}],\"name\":\"getStakeableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GuaranteedMinimumStakingPolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use GuaranteedMinimumStakingPolicyMetaData.ABI instead.
var GuaranteedMinimumStakingPolicyABI = GuaranteedMinimumStakingPolicyMetaData.ABI

// GuaranteedMinimumStakingPolicy is an auto generated Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicy struct {
	GuaranteedMinimumStakingPolicyCaller     // Read-only binding to the contract
	GuaranteedMinimumStakingPolicyTransactor // Write-only binding to the contract
	GuaranteedMinimumStakingPolicyFilterer   // Log filterer for contract events
}

// GuaranteedMinimumStakingPolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuaranteedMinimumStakingPolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuaranteedMinimumStakingPolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuaranteedMinimumStakingPolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuaranteedMinimumStakingPolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuaranteedMinimumStakingPolicySession struct {
	Contract     *GuaranteedMinimumStakingPolicy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GuaranteedMinimumStakingPolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuaranteedMinimumStakingPolicyCallerSession struct {
	Contract *GuaranteedMinimumStakingPolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// GuaranteedMinimumStakingPolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuaranteedMinimumStakingPolicyTransactorSession struct {
	Contract     *GuaranteedMinimumStakingPolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// GuaranteedMinimumStakingPolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicyRaw struct {
	Contract *GuaranteedMinimumStakingPolicy // Generic contract binding to access the raw methods on
}

// GuaranteedMinimumStakingPolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicyCallerRaw struct {
	Contract *GuaranteedMinimumStakingPolicyCaller // Generic read-only contract binding to access the raw methods on
}

// GuaranteedMinimumStakingPolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuaranteedMinimumStakingPolicyTransactorRaw struct {
	Contract *GuaranteedMinimumStakingPolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuaranteedMinimumStakingPolicy creates a new instance of GuaranteedMinimumStakingPolicy, bound to a specific deployed contract.
func NewGuaranteedMinimumStakingPolicy(address common.Address, backend bind.ContractBackend) (*GuaranteedMinimumStakingPolicy, error) {
	contract, err := bindGuaranteedMinimumStakingPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuaranteedMinimumStakingPolicy{GuaranteedMinimumStakingPolicyCaller: GuaranteedMinimumStakingPolicyCaller{contract: contract}, GuaranteedMinimumStakingPolicyTransactor: GuaranteedMinimumStakingPolicyTransactor{contract: contract}, GuaranteedMinimumStakingPolicyFilterer: GuaranteedMinimumStakingPolicyFilterer{contract: contract}}, nil
}

// NewGuaranteedMinimumStakingPolicyCaller creates a new read-only instance of GuaranteedMinimumStakingPolicy, bound to a specific deployed contract.
func NewGuaranteedMinimumStakingPolicyCaller(address common.Address, caller bind.ContractCaller) (*GuaranteedMinimumStakingPolicyCaller, error) {
	contract, err := bindGuaranteedMinimumStakingPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuaranteedMinimumStakingPolicyCaller{contract: contract}, nil
}

// NewGuaranteedMinimumStakingPolicyTransactor creates a new write-only instance of GuaranteedMinimumStakingPolicy, bound to a specific deployed contract.
func NewGuaranteedMinimumStakingPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*GuaranteedMinimumStakingPolicyTransactor, error) {
	contract, err := bindGuaranteedMinimumStakingPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuaranteedMinimumStakingPolicyTransactor{contract: contract}, nil
}

// NewGuaranteedMinimumStakingPolicyFilterer creates a new log filterer instance of GuaranteedMinimumStakingPolicy, bound to a specific deployed contract.
func NewGuaranteedMinimumStakingPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*GuaranteedMinimumStakingPolicyFilterer, error) {
	contract, err := bindGuaranteedMinimumStakingPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuaranteedMinimumStakingPolicyFilterer{contract: contract}, nil
}

// bindGuaranteedMinimumStakingPolicy binds a generic wrapper to an already deployed contract.
func bindGuaranteedMinimumStakingPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GuaranteedMinimumStakingPolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuaranteedMinimumStakingPolicy.Contract.GuaranteedMinimumStakingPolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.GuaranteedMinimumStakingPolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.GuaranteedMinimumStakingPolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuaranteedMinimumStakingPolicy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.contract.Transact(opts, method, params...)
}

// GetStakeableAmount is a free data retrieval call binding the contract method 0xdab40935.
//
// Solidity: function getStakeableAmount(uint256 _now, uint256 grantedAmount, uint256 duration, uint256 start, uint256 cliff, uint256 withdrawn) view returns(uint256)
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyCaller) GetStakeableAmount(opts *bind.CallOpts, _now *big.Int, grantedAmount *big.Int, duration *big.Int, start *big.Int, cliff *big.Int, withdrawn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GuaranteedMinimumStakingPolicy.contract.Call(opts, &out, "getStakeableAmount", _now, grantedAmount, duration, start, cliff, withdrawn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeableAmount is a free data retrieval call binding the contract method 0xdab40935.
//
// Solidity: function getStakeableAmount(uint256 _now, uint256 grantedAmount, uint256 duration, uint256 start, uint256 cliff, uint256 withdrawn) view returns(uint256)
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicySession) GetStakeableAmount(_now *big.Int, grantedAmount *big.Int, duration *big.Int, start *big.Int, cliff *big.Int, withdrawn *big.Int) (*big.Int, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.GetStakeableAmount(&_GuaranteedMinimumStakingPolicy.CallOpts, _now, grantedAmount, duration, start, cliff, withdrawn)
}

// GetStakeableAmount is a free data retrieval call binding the contract method 0xdab40935.
//
// Solidity: function getStakeableAmount(uint256 _now, uint256 grantedAmount, uint256 duration, uint256 start, uint256 cliff, uint256 withdrawn) view returns(uint256)
func (_GuaranteedMinimumStakingPolicy *GuaranteedMinimumStakingPolicyCallerSession) GetStakeableAmount(_now *big.Int, grantedAmount *big.Int, duration *big.Int, start *big.Int, cliff *big.Int, withdrawn *big.Int) (*big.Int, error) {
	return _GuaranteedMinimumStakingPolicy.Contract.GetStakeableAmount(&_GuaranteedMinimumStakingPolicy.CallOpts, _now, grantedAmount, duration, start, cliff, withdrawn)
}
