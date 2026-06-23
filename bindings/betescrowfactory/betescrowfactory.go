// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package betescrowfactory

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

// BetEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type BetEscrowTerms struct {
	YesAgent        common.Address
	NoAgent         common.Address
	Arbiter         common.Address
	Token           common.Address
	YesStake        *big.Int
	NoStake         *big.Int
	EventTime       uint64
	ClaimDeadline   uint64
	ChallengeWindow uint64
	Nonce           *big.Int
	Statement       string
	PrimarySource   string
	FallbackSource  string
	Visibility      uint8
}

// BetEscrowFactoryMetaData contains all meta data concerning the BetEscrowFactory contract.
var BetEscrowFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauseController_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"baseFeeBps_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registry_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractStablecoinAllowlist\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"terms\",\"type\":\"tuple\",\"internalType\":\"structBetEscrow.Terms\",\"components\":[{\"name\":\"yesAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbiter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"yesStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"eventTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"claimDeadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"challengeWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"statement\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"primarySource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fallbackSource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"visibility\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executionFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEmergencyPauseController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAgentRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetCreated\",\"inputs\":[{\"name\":\"escrow\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"yesAgent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"termsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"visibility\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProtocolPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotAllowed\",\"inputs\":[]}]",
}

// BetEscrowFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BetEscrowFactoryMetaData.ABI instead.
var BetEscrowFactoryABI = BetEscrowFactoryMetaData.ABI

// BetEscrowFactory is an auto generated Go binding around an Ethereum contract.
type BetEscrowFactory struct {
	BetEscrowFactoryCaller     // Read-only binding to the contract
	BetEscrowFactoryTransactor // Write-only binding to the contract
	BetEscrowFactoryFilterer   // Log filterer for contract events
}

// BetEscrowFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BetEscrowFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BetEscrowFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BetEscrowFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BetEscrowFactorySession struct {
	Contract     *BetEscrowFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetEscrowFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BetEscrowFactoryCallerSession struct {
	Contract *BetEscrowFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BetEscrowFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BetEscrowFactoryTransactorSession struct {
	Contract     *BetEscrowFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BetEscrowFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BetEscrowFactoryRaw struct {
	Contract *BetEscrowFactory // Generic contract binding to access the raw methods on
}

// BetEscrowFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BetEscrowFactoryCallerRaw struct {
	Contract *BetEscrowFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BetEscrowFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BetEscrowFactoryTransactorRaw struct {
	Contract *BetEscrowFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBetEscrowFactory creates a new instance of BetEscrowFactory, bound to a specific deployed contract.
func NewBetEscrowFactory(address common.Address, backend bind.ContractBackend) (*BetEscrowFactory, error) {
	contract, err := bindBetEscrowFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFactory{BetEscrowFactoryCaller: BetEscrowFactoryCaller{contract: contract}, BetEscrowFactoryTransactor: BetEscrowFactoryTransactor{contract: contract}, BetEscrowFactoryFilterer: BetEscrowFactoryFilterer{contract: contract}}, nil
}

// NewBetEscrowFactoryCaller creates a new read-only instance of BetEscrowFactory, bound to a specific deployed contract.
func NewBetEscrowFactoryCaller(address common.Address, caller bind.ContractCaller) (*BetEscrowFactoryCaller, error) {
	contract, err := bindBetEscrowFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFactoryCaller{contract: contract}, nil
}

// NewBetEscrowFactoryTransactor creates a new write-only instance of BetEscrowFactory, bound to a specific deployed contract.
func NewBetEscrowFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BetEscrowFactoryTransactor, error) {
	contract, err := bindBetEscrowFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFactoryTransactor{contract: contract}, nil
}

// NewBetEscrowFactoryFilterer creates a new log filterer instance of BetEscrowFactory, bound to a specific deployed contract.
func NewBetEscrowFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BetEscrowFactoryFilterer, error) {
	contract, err := bindBetEscrowFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFactoryFilterer{contract: contract}, nil
}

// bindBetEscrowFactory binds a generic wrapper to an already deployed contract.
func bindBetEscrowFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BetEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BetEscrowFactory *BetEscrowFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BetEscrowFactory.Contract.BetEscrowFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BetEscrowFactory *BetEscrowFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.BetEscrowFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BetEscrowFactory *BetEscrowFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.BetEscrowFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BetEscrowFactory *BetEscrowFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BetEscrowFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BetEscrowFactory *BetEscrowFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BetEscrowFactory *BetEscrowFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.contract.Transact(opts, method, params...)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCaller) Allowlist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "allowlist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_BetEscrowFactory *BetEscrowFactorySession) Allowlist() (common.Address, error) {
	return _BetEscrowFactory.Contract.Allowlist(&_BetEscrowFactory.CallOpts)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) Allowlist() (common.Address, error) {
	return _BetEscrowFactory.Contract.Allowlist(&_BetEscrowFactory.CallOpts)
}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactoryCaller) BaseFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "baseFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactorySession) BaseFeeBps() (*big.Int, error) {
	return _BetEscrowFactory.Contract.BaseFeeBps(&_BetEscrowFactory.CallOpts)
}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) BaseFeeBps() (*big.Int, error) {
	return _BetEscrowFactory.Contract.BaseFeeBps(&_BetEscrowFactory.CallOpts)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactoryCaller) ExecutionFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "executionFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactorySession) ExecutionFee() (*big.Int, error) {
	return _BetEscrowFactory.Contract.ExecutionFee(&_BetEscrowFactory.CallOpts)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) ExecutionFee() (*big.Int, error) {
	return _BetEscrowFactory.Contract.ExecutionFee(&_BetEscrowFactory.CallOpts)
}

// PauseController is a free data retrieval call binding the contract method 0x60b47789.
//
// Solidity: function pauseController() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCaller) PauseController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "pauseController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseController is a free data retrieval call binding the contract method 0x60b47789.
//
// Solidity: function pauseController() view returns(address)
func (_BetEscrowFactory *BetEscrowFactorySession) PauseController() (common.Address, error) {
	return _BetEscrowFactory.Contract.PauseController(&_BetEscrowFactory.CallOpts)
}

// PauseController is a free data retrieval call binding the contract method 0x60b47789.
//
// Solidity: function pauseController() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) PauseController() (common.Address, error) {
	return _BetEscrowFactory.Contract.PauseController(&_BetEscrowFactory.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrowFactory *BetEscrowFactorySession) Registry() (common.Address, error) {
	return _BetEscrowFactory.Contract.Registry(&_BetEscrowFactory.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) Registry() (common.Address, error) {
	return _BetEscrowFactory.Contract.Registry(&_BetEscrowFactory.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCaller) RevenueWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrowFactory.contract.Call(opts, &out, "revenueWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrowFactory *BetEscrowFactorySession) RevenueWallet() (common.Address, error) {
	return _BetEscrowFactory.Contract.RevenueWallet(&_BetEscrowFactory.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrowFactory *BetEscrowFactoryCallerSession) RevenueWallet() (common.Address, error) {
	return _BetEscrowFactory.Contract.RevenueWallet(&_BetEscrowFactory.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x0c39ed7c.
//
// Solidity: function create((address,address,address,address,uint256,uint256,uint64,uint64,uint64,uint256,string,string,string,uint8) terms) returns(address escrow)
func (_BetEscrowFactory *BetEscrowFactoryTransactor) Create(opts *bind.TransactOpts, terms BetEscrowTerms) (*types.Transaction, error) {
	return _BetEscrowFactory.contract.Transact(opts, "create", terms)
}

// Create is a paid mutator transaction binding the contract method 0x0c39ed7c.
//
// Solidity: function create((address,address,address,address,uint256,uint256,uint64,uint64,uint64,uint256,string,string,string,uint8) terms) returns(address escrow)
func (_BetEscrowFactory *BetEscrowFactorySession) Create(terms BetEscrowTerms) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.Create(&_BetEscrowFactory.TransactOpts, terms)
}

// Create is a paid mutator transaction binding the contract method 0x0c39ed7c.
//
// Solidity: function create((address,address,address,address,uint256,uint256,uint64,uint64,uint64,uint256,string,string,string,uint8) terms) returns(address escrow)
func (_BetEscrowFactory *BetEscrowFactoryTransactorSession) Create(terms BetEscrowTerms) (*types.Transaction, error) {
	return _BetEscrowFactory.Contract.Create(&_BetEscrowFactory.TransactOpts, terms)
}

// BetEscrowFactoryBetCreatedIterator is returned from FilterBetCreated and is used to iterate over the raw logs and unpacked data for BetCreated events raised by the BetEscrowFactory contract.
type BetEscrowFactoryBetCreatedIterator struct {
	Event *BetEscrowFactoryBetCreated // Event containing the contract specifics and raw log

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
func (it *BetEscrowFactoryBetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowFactoryBetCreated)
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
		it.Event = new(BetEscrowFactoryBetCreated)
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
func (it *BetEscrowFactoryBetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowFactoryBetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowFactoryBetCreated represents a BetCreated event raised by the BetEscrowFactory contract.
type BetEscrowFactoryBetCreated struct {
	Escrow     common.Address
	YesAgent   common.Address
	NoAgent    common.Address
	TermsHash  [32]byte
	Visibility uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBetCreated is a free log retrieval operation binding the contract event 0x551dbb51d8f106c1d3998f4786c5fc3c4bae02b1481f5b407331c566e71d99a2.
//
// Solidity: event BetCreated(address indexed escrow, address indexed yesAgent, address indexed noAgent, bytes32 termsHash, uint8 visibility)
func (_BetEscrowFactory *BetEscrowFactoryFilterer) FilterBetCreated(opts *bind.FilterOpts, escrow []common.Address, yesAgent []common.Address, noAgent []common.Address) (*BetEscrowFactoryBetCreatedIterator, error) {

	var escrowRule []interface{}
	for _, escrowItem := range escrow {
		escrowRule = append(escrowRule, escrowItem)
	}
	var yesAgentRule []interface{}
	for _, yesAgentItem := range yesAgent {
		yesAgentRule = append(yesAgentRule, yesAgentItem)
	}
	var noAgentRule []interface{}
	for _, noAgentItem := range noAgent {
		noAgentRule = append(noAgentRule, noAgentItem)
	}

	logs, sub, err := _BetEscrowFactory.contract.FilterLogs(opts, "BetCreated", escrowRule, yesAgentRule, noAgentRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFactoryBetCreatedIterator{contract: _BetEscrowFactory.contract, event: "BetCreated", logs: logs, sub: sub}, nil
}

// WatchBetCreated is a free log subscription operation binding the contract event 0x551dbb51d8f106c1d3998f4786c5fc3c4bae02b1481f5b407331c566e71d99a2.
//
// Solidity: event BetCreated(address indexed escrow, address indexed yesAgent, address indexed noAgent, bytes32 termsHash, uint8 visibility)
func (_BetEscrowFactory *BetEscrowFactoryFilterer) WatchBetCreated(opts *bind.WatchOpts, sink chan<- *BetEscrowFactoryBetCreated, escrow []common.Address, yesAgent []common.Address, noAgent []common.Address) (event.Subscription, error) {

	var escrowRule []interface{}
	for _, escrowItem := range escrow {
		escrowRule = append(escrowRule, escrowItem)
	}
	var yesAgentRule []interface{}
	for _, yesAgentItem := range yesAgent {
		yesAgentRule = append(yesAgentRule, yesAgentItem)
	}
	var noAgentRule []interface{}
	for _, noAgentItem := range noAgent {
		noAgentRule = append(noAgentRule, noAgentItem)
	}

	logs, sub, err := _BetEscrowFactory.contract.WatchLogs(opts, "BetCreated", escrowRule, yesAgentRule, noAgentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowFactoryBetCreated)
				if err := _BetEscrowFactory.contract.UnpackLog(event, "BetCreated", log); err != nil {
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

// ParseBetCreated is a log parse operation binding the contract event 0x551dbb51d8f106c1d3998f4786c5fc3c4bae02b1481f5b407331c566e71d99a2.
//
// Solidity: event BetCreated(address indexed escrow, address indexed yesAgent, address indexed noAgent, bytes32 termsHash, uint8 visibility)
func (_BetEscrowFactory *BetEscrowFactoryFilterer) ParseBetCreated(log types.Log) (*BetEscrowFactoryBetCreated, error) {
	event := new(BetEscrowFactoryBetCreated)
	if err := _BetEscrowFactory.contract.UnpackLog(event, "BetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
