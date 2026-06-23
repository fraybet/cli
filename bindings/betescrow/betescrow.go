// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package betescrow

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

// BetEscrowMetaData contains all meta data concerning the BetEscrow contract.
var BetEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"t\",\"type\":\"tuple\",\"internalType\":\"structBetEscrow.Terms\",\"components\":[{\"name\":\"yesAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbiter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"yesStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"eventTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"claimDeadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"challengeWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"statement\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"primarySource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fallbackSource\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"visibility\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"baseFeeBps_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accept\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agreeOutcome\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"arbiter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengeDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimedOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evidenceURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fallbackSource\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"noAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"primarySource\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBondRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"evidenceURI_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revoke\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"statement\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"status\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"termsHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"visibility\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voidUnclaimed\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yesAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Accepted\",\"inputs\":[{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ArbiterResolved\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"evidenceURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Challenged\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"challengeDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Funded\",\"inputs\":[{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutcomeAgreed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Revoked\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settled\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WentLive\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadTerms\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcome\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotContested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFunding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRefundable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooEarly\",\"inputs\":[]}]",
}

// BetEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use BetEscrowMetaData.ABI instead.
var BetEscrowABI = BetEscrowMetaData.ABI

// BetEscrow is an auto generated Go binding around an Ethereum contract.
type BetEscrow struct {
	BetEscrowCaller     // Read-only binding to the contract
	BetEscrowTransactor // Write-only binding to the contract
	BetEscrowFilterer   // Log filterer for contract events
}

// BetEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type BetEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BetEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BetEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BetEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BetEscrowSession struct {
	Contract     *BetEscrow        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BetEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BetEscrowCallerSession struct {
	Contract *BetEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BetEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BetEscrowTransactorSession struct {
	Contract     *BetEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BetEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type BetEscrowRaw struct {
	Contract *BetEscrow // Generic contract binding to access the raw methods on
}

// BetEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BetEscrowCallerRaw struct {
	Contract *BetEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// BetEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BetEscrowTransactorRaw struct {
	Contract *BetEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBetEscrow creates a new instance of BetEscrow, bound to a specific deployed contract.
func NewBetEscrow(address common.Address, backend bind.ContractBackend) (*BetEscrow, error) {
	contract, err := bindBetEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BetEscrow{BetEscrowCaller: BetEscrowCaller{contract: contract}, BetEscrowTransactor: BetEscrowTransactor{contract: contract}, BetEscrowFilterer: BetEscrowFilterer{contract: contract}}, nil
}

// NewBetEscrowCaller creates a new read-only instance of BetEscrow, bound to a specific deployed contract.
func NewBetEscrowCaller(address common.Address, caller bind.ContractCaller) (*BetEscrowCaller, error) {
	contract, err := bindBetEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BetEscrowCaller{contract: contract}, nil
}

// NewBetEscrowTransactor creates a new write-only instance of BetEscrow, bound to a specific deployed contract.
func NewBetEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*BetEscrowTransactor, error) {
	contract, err := bindBetEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BetEscrowTransactor{contract: contract}, nil
}

// NewBetEscrowFilterer creates a new log filterer instance of BetEscrow, bound to a specific deployed contract.
func NewBetEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*BetEscrowFilterer, error) {
	contract, err := bindBetEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFilterer{contract: contract}, nil
}

// bindBetEscrow binds a generic wrapper to an already deployed contract.
func bindBetEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BetEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BetEscrow *BetEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BetEscrow.Contract.BetEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BetEscrow *BetEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.Contract.BetEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BetEscrow *BetEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BetEscrow.Contract.BetEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BetEscrow *BetEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BetEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BetEscrow *BetEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BetEscrow *BetEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BetEscrow.Contract.contract.Transact(opts, method, params...)
}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_BetEscrow *BetEscrowCaller) Arbiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "arbiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_BetEscrow *BetEscrowSession) Arbiter() (common.Address, error) {
	return _BetEscrow.Contract.Arbiter(&_BetEscrow.CallOpts)
}

// Arbiter is a free data retrieval call binding the contract method 0xfe25e00a.
//
// Solidity: function arbiter() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) Arbiter() (common.Address, error) {
	return _BetEscrow.Contract.Arbiter(&_BetEscrow.CallOpts)
}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrow *BetEscrowCaller) BaseFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "baseFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrow *BetEscrowSession) BaseFeeBps() (*big.Int, error) {
	return _BetEscrow.Contract.BaseFeeBps(&_BetEscrow.CallOpts)
}

// BaseFeeBps is a free data retrieval call binding the contract method 0xbf5a5940.
//
// Solidity: function baseFeeBps() view returns(uint256)
func (_BetEscrow *BetEscrowCallerSession) BaseFeeBps() (*big.Int, error) {
	return _BetEscrow.Contract.BaseFeeBps(&_BetEscrow.CallOpts)
}

// ChallengeDeadline is a free data retrieval call binding the contract method 0xca593dad.
//
// Solidity: function challengeDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowCaller) ChallengeDeadline(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "challengeDeadline")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChallengeDeadline is a free data retrieval call binding the contract method 0xca593dad.
//
// Solidity: function challengeDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowSession) ChallengeDeadline() (uint64, error) {
	return _BetEscrow.Contract.ChallengeDeadline(&_BetEscrow.CallOpts)
}

// ChallengeDeadline is a free data retrieval call binding the contract method 0xca593dad.
//
// Solidity: function challengeDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowCallerSession) ChallengeDeadline() (uint64, error) {
	return _BetEscrow.Contract.ChallengeDeadline(&_BetEscrow.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint64)
func (_BetEscrow *BetEscrowCaller) ChallengeWindow(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "challengeWindow")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint64)
func (_BetEscrow *BetEscrowSession) ChallengeWindow() (uint64, error) {
	return _BetEscrow.Contract.ChallengeWindow(&_BetEscrow.CallOpts)
}

// ChallengeWindow is a free data retrieval call binding the contract method 0x861a1412.
//
// Solidity: function challengeWindow() view returns(uint64)
func (_BetEscrow *BetEscrowCallerSession) ChallengeWindow() (uint64, error) {
	return _BetEscrow.Contract.ChallengeWindow(&_BetEscrow.CallOpts)
}

// ClaimDeadline is a free data retrieval call binding the contract method 0x3ba86c44.
//
// Solidity: function claimDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowCaller) ClaimDeadline(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "claimDeadline")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClaimDeadline is a free data retrieval call binding the contract method 0x3ba86c44.
//
// Solidity: function claimDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowSession) ClaimDeadline() (uint64, error) {
	return _BetEscrow.Contract.ClaimDeadline(&_BetEscrow.CallOpts)
}

// ClaimDeadline is a free data retrieval call binding the contract method 0x3ba86c44.
//
// Solidity: function claimDeadline() view returns(uint64)
func (_BetEscrow *BetEscrowCallerSession) ClaimDeadline() (uint64, error) {
	return _BetEscrow.Contract.ClaimDeadline(&_BetEscrow.CallOpts)
}

// ClaimedOutcome is a free data retrieval call binding the contract method 0x0f383d11.
//
// Solidity: function claimedOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) ClaimedOutcome(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "claimedOutcome")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ClaimedOutcome is a free data retrieval call binding the contract method 0x0f383d11.
//
// Solidity: function claimedOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowSession) ClaimedOutcome() (uint8, error) {
	return _BetEscrow.Contract.ClaimedOutcome(&_BetEscrow.CallOpts)
}

// ClaimedOutcome is a free data retrieval call binding the contract method 0x0f383d11.
//
// Solidity: function claimedOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) ClaimedOutcome() (uint8, error) {
	return _BetEscrow.Contract.ClaimedOutcome(&_BetEscrow.CallOpts)
}

// Disputed is a free data retrieval call binding the contract method 0x0695c46c.
//
// Solidity: function disputed() view returns(bool)
func (_BetEscrow *BetEscrowCaller) Disputed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "disputed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Disputed is a free data retrieval call binding the contract method 0x0695c46c.
//
// Solidity: function disputed() view returns(bool)
func (_BetEscrow *BetEscrowSession) Disputed() (bool, error) {
	return _BetEscrow.Contract.Disputed(&_BetEscrow.CallOpts)
}

// Disputed is a free data retrieval call binding the contract method 0x0695c46c.
//
// Solidity: function disputed() view returns(bool)
func (_BetEscrow *BetEscrowCallerSession) Disputed() (bool, error) {
	return _BetEscrow.Contract.Disputed(&_BetEscrow.CallOpts)
}

// EventTime is a free data retrieval call binding the contract method 0x0c317e7b.
//
// Solidity: function eventTime() view returns(uint64)
func (_BetEscrow *BetEscrowCaller) EventTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "eventTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EventTime is a free data retrieval call binding the contract method 0x0c317e7b.
//
// Solidity: function eventTime() view returns(uint64)
func (_BetEscrow *BetEscrowSession) EventTime() (uint64, error) {
	return _BetEscrow.Contract.EventTime(&_BetEscrow.CallOpts)
}

// EventTime is a free data retrieval call binding the contract method 0x0c317e7b.
//
// Solidity: function eventTime() view returns(uint64)
func (_BetEscrow *BetEscrowCallerSession) EventTime() (uint64, error) {
	return _BetEscrow.Contract.EventTime(&_BetEscrow.CallOpts)
}

// EvidenceURI is a free data retrieval call binding the contract method 0xca6f3d28.
//
// Solidity: function evidenceURI() view returns(string)
func (_BetEscrow *BetEscrowCaller) EvidenceURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "evidenceURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EvidenceURI is a free data retrieval call binding the contract method 0xca6f3d28.
//
// Solidity: function evidenceURI() view returns(string)
func (_BetEscrow *BetEscrowSession) EvidenceURI() (string, error) {
	return _BetEscrow.Contract.EvidenceURI(&_BetEscrow.CallOpts)
}

// EvidenceURI is a free data retrieval call binding the contract method 0xca6f3d28.
//
// Solidity: function evidenceURI() view returns(string)
func (_BetEscrow *BetEscrowCallerSession) EvidenceURI() (string, error) {
	return _BetEscrow.Contract.EvidenceURI(&_BetEscrow.CallOpts)
}

// FallbackSource is a free data retrieval call binding the contract method 0xf698113f.
//
// Solidity: function fallbackSource() view returns(string)
func (_BetEscrow *BetEscrowCaller) FallbackSource(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "fallbackSource")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FallbackSource is a free data retrieval call binding the contract method 0xf698113f.
//
// Solidity: function fallbackSource() view returns(string)
func (_BetEscrow *BetEscrowSession) FallbackSource() (string, error) {
	return _BetEscrow.Contract.FallbackSource(&_BetEscrow.CallOpts)
}

// FallbackSource is a free data retrieval call binding the contract method 0xf698113f.
//
// Solidity: function fallbackSource() view returns(string)
func (_BetEscrow *BetEscrowCallerSession) FallbackSource() (string, error) {
	return _BetEscrow.Contract.FallbackSource(&_BetEscrow.CallOpts)
}

// FinalOutcome is a free data retrieval call binding the contract method 0x404002a6.
//
// Solidity: function finalOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) FinalOutcome(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "finalOutcome")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FinalOutcome is a free data retrieval call binding the contract method 0x404002a6.
//
// Solidity: function finalOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowSession) FinalOutcome() (uint8, error) {
	return _BetEscrow.Contract.FinalOutcome(&_BetEscrow.CallOpts)
}

// FinalOutcome is a free data retrieval call binding the contract method 0x404002a6.
//
// Solidity: function finalOutcome() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) FinalOutcome() (uint8, error) {
	return _BetEscrow.Contract.FinalOutcome(&_BetEscrow.CallOpts)
}

// NoAgent is a free data retrieval call binding the contract method 0x5c3c10d7.
//
// Solidity: function noAgent() view returns(address)
func (_BetEscrow *BetEscrowCaller) NoAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "noAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NoAgent is a free data retrieval call binding the contract method 0x5c3c10d7.
//
// Solidity: function noAgent() view returns(address)
func (_BetEscrow *BetEscrowSession) NoAgent() (common.Address, error) {
	return _BetEscrow.Contract.NoAgent(&_BetEscrow.CallOpts)
}

// NoAgent is a free data retrieval call binding the contract method 0x5c3c10d7.
//
// Solidity: function noAgent() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) NoAgent() (common.Address, error) {
	return _BetEscrow.Contract.NoAgent(&_BetEscrow.CallOpts)
}

// NoAgentAgreed is a free data retrieval call binding the contract method 0x6dc87dac.
//
// Solidity: function noAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) NoAgentAgreed(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "noAgentAgreed")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NoAgentAgreed is a free data retrieval call binding the contract method 0x6dc87dac.
//
// Solidity: function noAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowSession) NoAgentAgreed() (uint8, error) {
	return _BetEscrow.Contract.NoAgentAgreed(&_BetEscrow.CallOpts)
}

// NoAgentAgreed is a free data retrieval call binding the contract method 0x6dc87dac.
//
// Solidity: function noAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) NoAgentAgreed() (uint8, error) {
	return _BetEscrow.Contract.NoAgentAgreed(&_BetEscrow.CallOpts)
}

// NoFunded is a free data retrieval call binding the contract method 0x69523a20.
//
// Solidity: function noFunded() view returns(bool)
func (_BetEscrow *BetEscrowCaller) NoFunded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "noFunded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NoFunded is a free data retrieval call binding the contract method 0x69523a20.
//
// Solidity: function noFunded() view returns(bool)
func (_BetEscrow *BetEscrowSession) NoFunded() (bool, error) {
	return _BetEscrow.Contract.NoFunded(&_BetEscrow.CallOpts)
}

// NoFunded is a free data retrieval call binding the contract method 0x69523a20.
//
// Solidity: function noFunded() view returns(bool)
func (_BetEscrow *BetEscrowCallerSession) NoFunded() (bool, error) {
	return _BetEscrow.Contract.NoFunded(&_BetEscrow.CallOpts)
}

// NoStake is a free data retrieval call binding the contract method 0x14bc116f.
//
// Solidity: function noStake() view returns(uint256)
func (_BetEscrow *BetEscrowCaller) NoStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "noStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoStake is a free data retrieval call binding the contract method 0x14bc116f.
//
// Solidity: function noStake() view returns(uint256)
func (_BetEscrow *BetEscrowSession) NoStake() (*big.Int, error) {
	return _BetEscrow.Contract.NoStake(&_BetEscrow.CallOpts)
}

// NoStake is a free data retrieval call binding the contract method 0x14bc116f.
//
// Solidity: function noStake() view returns(uint256)
func (_BetEscrow *BetEscrowCallerSession) NoStake() (*big.Int, error) {
	return _BetEscrow.Contract.NoStake(&_BetEscrow.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BetEscrow *BetEscrowCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BetEscrow *BetEscrowSession) Nonce() (*big.Int, error) {
	return _BetEscrow.Contract.Nonce(&_BetEscrow.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BetEscrow *BetEscrowCallerSession) Nonce() (*big.Int, error) {
	return _BetEscrow.Contract.Nonce(&_BetEscrow.CallOpts)
}

// PrimarySource is a free data retrieval call binding the contract method 0x7461d7de.
//
// Solidity: function primarySource() view returns(string)
func (_BetEscrow *BetEscrowCaller) PrimarySource(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "primarySource")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrimarySource is a free data retrieval call binding the contract method 0x7461d7de.
//
// Solidity: function primarySource() view returns(string)
func (_BetEscrow *BetEscrowSession) PrimarySource() (string, error) {
	return _BetEscrow.Contract.PrimarySource(&_BetEscrow.CallOpts)
}

// PrimarySource is a free data retrieval call binding the contract method 0x7461d7de.
//
// Solidity: function primarySource() view returns(string)
func (_BetEscrow *BetEscrowCallerSession) PrimarySource() (string, error) {
	return _BetEscrow.Contract.PrimarySource(&_BetEscrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrow *BetEscrowCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrow *BetEscrowSession) Registry() (common.Address, error) {
	return _BetEscrow.Contract.Registry(&_BetEscrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) Registry() (common.Address, error) {
	return _BetEscrow.Contract.Registry(&_BetEscrow.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrow *BetEscrowCaller) RevenueWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "revenueWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrow *BetEscrowSession) RevenueWallet() (common.Address, error) {
	return _BetEscrow.Contract.RevenueWallet(&_BetEscrow.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) RevenueWallet() (common.Address, error) {
	return _BetEscrow.Contract.RevenueWallet(&_BetEscrow.CallOpts)
}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_BetEscrow *BetEscrowCaller) Settled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "settled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_BetEscrow *BetEscrowSession) Settled() (bool, error) {
	return _BetEscrow.Contract.Settled(&_BetEscrow.CallOpts)
}

// Settled is a free data retrieval call binding the contract method 0x8f775839.
//
// Solidity: function settled() view returns(bool)
func (_BetEscrow *BetEscrowCallerSession) Settled() (bool, error) {
	return _BetEscrow.Contract.Settled(&_BetEscrow.CallOpts)
}

// Statement is a free data retrieval call binding the contract method 0xe42cb9f3.
//
// Solidity: function statement() view returns(string)
func (_BetEscrow *BetEscrowCaller) Statement(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "statement")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Statement is a free data retrieval call binding the contract method 0xe42cb9f3.
//
// Solidity: function statement() view returns(string)
func (_BetEscrow *BetEscrowSession) Statement() (string, error) {
	return _BetEscrow.Contract.Statement(&_BetEscrow.CallOpts)
}

// Statement is a free data retrieval call binding the contract method 0xe42cb9f3.
//
// Solidity: function statement() view returns(string)
func (_BetEscrow *BetEscrowCallerSession) Statement() (string, error) {
	return _BetEscrow.Contract.Statement(&_BetEscrow.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_BetEscrow *BetEscrowSession) Status() (uint8, error) {
	return _BetEscrow.Contract.Status(&_BetEscrow.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) Status() (uint8, error) {
	return _BetEscrow.Contract.Status(&_BetEscrow.CallOpts)
}

// TermsHash is a free data retrieval call binding the contract method 0xb311d9fd.
//
// Solidity: function termsHash() view returns(bytes32)
func (_BetEscrow *BetEscrowCaller) TermsHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "termsHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TermsHash is a free data retrieval call binding the contract method 0xb311d9fd.
//
// Solidity: function termsHash() view returns(bytes32)
func (_BetEscrow *BetEscrowSession) TermsHash() ([32]byte, error) {
	return _BetEscrow.Contract.TermsHash(&_BetEscrow.CallOpts)
}

// TermsHash is a free data retrieval call binding the contract method 0xb311d9fd.
//
// Solidity: function termsHash() view returns(bytes32)
func (_BetEscrow *BetEscrowCallerSession) TermsHash() ([32]byte, error) {
	return _BetEscrow.Contract.TermsHash(&_BetEscrow.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BetEscrow *BetEscrowCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BetEscrow *BetEscrowSession) Token() (common.Address, error) {
	return _BetEscrow.Contract.Token(&_BetEscrow.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) Token() (common.Address, error) {
	return _BetEscrow.Contract.Token(&_BetEscrow.CallOpts)
}

// Visibility is a free data retrieval call binding the contract method 0x29adec14.
//
// Solidity: function visibility() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) Visibility(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "visibility")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Visibility is a free data retrieval call binding the contract method 0x29adec14.
//
// Solidity: function visibility() view returns(uint8)
func (_BetEscrow *BetEscrowSession) Visibility() (uint8, error) {
	return _BetEscrow.Contract.Visibility(&_BetEscrow.CallOpts)
}

// Visibility is a free data retrieval call binding the contract method 0x29adec14.
//
// Solidity: function visibility() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) Visibility() (uint8, error) {
	return _BetEscrow.Contract.Visibility(&_BetEscrow.CallOpts)
}

// YesAgent is a free data retrieval call binding the contract method 0x62897e20.
//
// Solidity: function yesAgent() view returns(address)
func (_BetEscrow *BetEscrowCaller) YesAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "yesAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YesAgent is a free data retrieval call binding the contract method 0x62897e20.
//
// Solidity: function yesAgent() view returns(address)
func (_BetEscrow *BetEscrowSession) YesAgent() (common.Address, error) {
	return _BetEscrow.Contract.YesAgent(&_BetEscrow.CallOpts)
}

// YesAgent is a free data retrieval call binding the contract method 0x62897e20.
//
// Solidity: function yesAgent() view returns(address)
func (_BetEscrow *BetEscrowCallerSession) YesAgent() (common.Address, error) {
	return _BetEscrow.Contract.YesAgent(&_BetEscrow.CallOpts)
}

// YesAgentAgreed is a free data retrieval call binding the contract method 0x62ccd3b8.
//
// Solidity: function yesAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowCaller) YesAgentAgreed(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "yesAgentAgreed")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// YesAgentAgreed is a free data retrieval call binding the contract method 0x62ccd3b8.
//
// Solidity: function yesAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowSession) YesAgentAgreed() (uint8, error) {
	return _BetEscrow.Contract.YesAgentAgreed(&_BetEscrow.CallOpts)
}

// YesAgentAgreed is a free data retrieval call binding the contract method 0x62ccd3b8.
//
// Solidity: function yesAgentAgreed() view returns(uint8)
func (_BetEscrow *BetEscrowCallerSession) YesAgentAgreed() (uint8, error) {
	return _BetEscrow.Contract.YesAgentAgreed(&_BetEscrow.CallOpts)
}

// YesFunded is a free data retrieval call binding the contract method 0x67cce5e1.
//
// Solidity: function yesFunded() view returns(bool)
func (_BetEscrow *BetEscrowCaller) YesFunded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "yesFunded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// YesFunded is a free data retrieval call binding the contract method 0x67cce5e1.
//
// Solidity: function yesFunded() view returns(bool)
func (_BetEscrow *BetEscrowSession) YesFunded() (bool, error) {
	return _BetEscrow.Contract.YesFunded(&_BetEscrow.CallOpts)
}

// YesFunded is a free data retrieval call binding the contract method 0x67cce5e1.
//
// Solidity: function yesFunded() view returns(bool)
func (_BetEscrow *BetEscrowCallerSession) YesFunded() (bool, error) {
	return _BetEscrow.Contract.YesFunded(&_BetEscrow.CallOpts)
}

// YesStake is a free data retrieval call binding the contract method 0x28ea6dda.
//
// Solidity: function yesStake() view returns(uint256)
func (_BetEscrow *BetEscrowCaller) YesStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "yesStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YesStake is a free data retrieval call binding the contract method 0x28ea6dda.
//
// Solidity: function yesStake() view returns(uint256)
func (_BetEscrow *BetEscrowSession) YesStake() (*big.Int, error) {
	return _BetEscrow.Contract.YesStake(&_BetEscrow.CallOpts)
}

// YesStake is a free data retrieval call binding the contract method 0x28ea6dda.
//
// Solidity: function yesStake() view returns(uint256)
func (_BetEscrow *BetEscrowCallerSession) YesStake() (*big.Int, error) {
	return _BetEscrow.Contract.YesStake(&_BetEscrow.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_BetEscrow *BetEscrowTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_BetEscrow *BetEscrowSession) Accept() (*types.Transaction, error) {
	return _BetEscrow.Contract.Accept(&_BetEscrow.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_BetEscrow *BetEscrowTransactorSession) Accept() (*types.Transaction, error) {
	return _BetEscrow.Contract.Accept(&_BetEscrow.TransactOpts)
}

// AgreeOutcome is a paid mutator transaction binding the contract method 0xc94e8e1d.
//
// Solidity: function agreeOutcome(uint8 outcome) returns()
func (_BetEscrow *BetEscrowTransactor) AgreeOutcome(opts *bind.TransactOpts, outcome uint8) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "agreeOutcome", outcome)
}

// AgreeOutcome is a paid mutator transaction binding the contract method 0xc94e8e1d.
//
// Solidity: function agreeOutcome(uint8 outcome) returns()
func (_BetEscrow *BetEscrowSession) AgreeOutcome(outcome uint8) (*types.Transaction, error) {
	return _BetEscrow.Contract.AgreeOutcome(&_BetEscrow.TransactOpts, outcome)
}

// AgreeOutcome is a paid mutator transaction binding the contract method 0xc94e8e1d.
//
// Solidity: function agreeOutcome(uint8 outcome) returns()
func (_BetEscrow *BetEscrowTransactorSession) AgreeOutcome(outcome uint8) (*types.Transaction, error) {
	return _BetEscrow.Contract.AgreeOutcome(&_BetEscrow.TransactOpts, outcome)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_BetEscrow *BetEscrowTransactor) Challenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "challenge")
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_BetEscrow *BetEscrowSession) Challenge() (*types.Transaction, error) {
	return _BetEscrow.Contract.Challenge(&_BetEscrow.TransactOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_BetEscrow *BetEscrowTransactorSession) Challenge() (*types.Transaction, error) {
	return _BetEscrow.Contract.Challenge(&_BetEscrow.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xfb4b66eb.
//
// Solidity: function claim(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowTransactor) Claim(opts *bind.TransactOpts, outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "claim", outcome, evidenceHash)
}

// Claim is a paid mutator transaction binding the contract method 0xfb4b66eb.
//
// Solidity: function claim(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowSession) Claim(outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.Contract.Claim(&_BetEscrow.TransactOpts, outcome, evidenceHash)
}

// Claim is a paid mutator transaction binding the contract method 0xfb4b66eb.
//
// Solidity: function claim(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowTransactorSession) Claim(outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.Contract.Claim(&_BetEscrow.TransactOpts, outcome, evidenceHash)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BetEscrow *BetEscrowTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BetEscrow *BetEscrowSession) Finalize() (*types.Transaction, error) {
	return _BetEscrow.Contract.Finalize(&_BetEscrow.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BetEscrow *BetEscrowTransactorSession) Finalize() (*types.Transaction, error) {
	return _BetEscrow.Contract.Finalize(&_BetEscrow.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_BetEscrow *BetEscrowTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_BetEscrow *BetEscrowSession) Fund() (*types.Transaction, error) {
	return _BetEscrow.Contract.Fund(&_BetEscrow.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_BetEscrow *BetEscrowTransactorSession) Fund() (*types.Transaction, error) {
	return _BetEscrow.Contract.Fund(&_BetEscrow.TransactOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0xb0c80c0a.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash, string evidenceURI_) returns()
func (_BetEscrow *BetEscrowTransactor) Resolve(opts *bind.TransactOpts, outcome uint8, evidenceHash [32]byte, evidenceURI_ string) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "resolve", outcome, evidenceHash, evidenceURI_)
}

// Resolve is a paid mutator transaction binding the contract method 0xb0c80c0a.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash, string evidenceURI_) returns()
func (_BetEscrow *BetEscrowSession) Resolve(outcome uint8, evidenceHash [32]byte, evidenceURI_ string) (*types.Transaction, error) {
	return _BetEscrow.Contract.Resolve(&_BetEscrow.TransactOpts, outcome, evidenceHash, evidenceURI_)
}

// Resolve is a paid mutator transaction binding the contract method 0xb0c80c0a.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash, string evidenceURI_) returns()
func (_BetEscrow *BetEscrowTransactorSession) Resolve(outcome uint8, evidenceHash [32]byte, evidenceURI_ string) (*types.Transaction, error) {
	return _BetEscrow.Contract.Resolve(&_BetEscrow.TransactOpts, outcome, evidenceHash, evidenceURI_)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_BetEscrow *BetEscrowTransactor) Revoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "revoke")
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_BetEscrow *BetEscrowSession) Revoke() (*types.Transaction, error) {
	return _BetEscrow.Contract.Revoke(&_BetEscrow.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_BetEscrow *BetEscrowTransactorSession) Revoke() (*types.Transaction, error) {
	return _BetEscrow.Contract.Revoke(&_BetEscrow.TransactOpts)
}

// VoidUnclaimed is a paid mutator transaction binding the contract method 0x13888774.
//
// Solidity: function voidUnclaimed() returns()
func (_BetEscrow *BetEscrowTransactor) VoidUnclaimed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "voidUnclaimed")
}

// VoidUnclaimed is a paid mutator transaction binding the contract method 0x13888774.
//
// Solidity: function voidUnclaimed() returns()
func (_BetEscrow *BetEscrowSession) VoidUnclaimed() (*types.Transaction, error) {
	return _BetEscrow.Contract.VoidUnclaimed(&_BetEscrow.TransactOpts)
}

// VoidUnclaimed is a paid mutator transaction binding the contract method 0x13888774.
//
// Solidity: function voidUnclaimed() returns()
func (_BetEscrow *BetEscrowTransactorSession) VoidUnclaimed() (*types.Transaction, error) {
	return _BetEscrow.Contract.VoidUnclaimed(&_BetEscrow.TransactOpts)
}

// BetEscrowAcceptedIterator is returned from FilterAccepted and is used to iterate over the raw logs and unpacked data for Accepted events raised by the BetEscrow contract.
type BetEscrowAcceptedIterator struct {
	Event *BetEscrowAccepted // Event containing the contract specifics and raw log

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
func (it *BetEscrowAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowAccepted)
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
		it.Event = new(BetEscrowAccepted)
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
func (it *BetEscrowAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowAccepted represents a Accepted event raised by the BetEscrow contract.
type BetEscrowAccepted struct {
	Taker common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccepted is a free log retrieval operation binding the contract event 0x2e7bcc1a1a1c1bea0ff9a677c29b43662d6560afbebd99c213389fc15a6279d2.
//
// Solidity: event Accepted(address indexed taker)
func (_BetEscrow *BetEscrowFilterer) FilterAccepted(opts *bind.FilterOpts, taker []common.Address) (*BetEscrowAcceptedIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Accepted", takerRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowAcceptedIterator{contract: _BetEscrow.contract, event: "Accepted", logs: logs, sub: sub}, nil
}

// WatchAccepted is a free log subscription operation binding the contract event 0x2e7bcc1a1a1c1bea0ff9a677c29b43662d6560afbebd99c213389fc15a6279d2.
//
// Solidity: event Accepted(address indexed taker)
func (_BetEscrow *BetEscrowFilterer) WatchAccepted(opts *bind.WatchOpts, sink chan<- *BetEscrowAccepted, taker []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Accepted", takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowAccepted)
				if err := _BetEscrow.contract.UnpackLog(event, "Accepted", log); err != nil {
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

// ParseAccepted is a log parse operation binding the contract event 0x2e7bcc1a1a1c1bea0ff9a677c29b43662d6560afbebd99c213389fc15a6279d2.
//
// Solidity: event Accepted(address indexed taker)
func (_BetEscrow *BetEscrowFilterer) ParseAccepted(log types.Log) (*BetEscrowAccepted, error) {
	event := new(BetEscrowAccepted)
	if err := _BetEscrow.contract.UnpackLog(event, "Accepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowArbiterResolvedIterator is returned from FilterArbiterResolved and is used to iterate over the raw logs and unpacked data for ArbiterResolved events raised by the BetEscrow contract.
type BetEscrowArbiterResolvedIterator struct {
	Event *BetEscrowArbiterResolved // Event containing the contract specifics and raw log

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
func (it *BetEscrowArbiterResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowArbiterResolved)
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
		it.Event = new(BetEscrowArbiterResolved)
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
func (it *BetEscrowArbiterResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowArbiterResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowArbiterResolved represents a ArbiterResolved event raised by the BetEscrow contract.
type BetEscrowArbiterResolved struct {
	Outcome      uint8
	EvidenceHash [32]byte
	EvidenceURI  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterArbiterResolved is a free log retrieval operation binding the contract event 0x565df8aa55948feffff27bcda8d240698e6a278c311f8bb1f1c067057667154e.
//
// Solidity: event ArbiterResolved(uint8 outcome, bytes32 evidenceHash, string evidenceURI)
func (_BetEscrow *BetEscrowFilterer) FilterArbiterResolved(opts *bind.FilterOpts) (*BetEscrowArbiterResolvedIterator, error) {

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "ArbiterResolved")
	if err != nil {
		return nil, err
	}
	return &BetEscrowArbiterResolvedIterator{contract: _BetEscrow.contract, event: "ArbiterResolved", logs: logs, sub: sub}, nil
}

// WatchArbiterResolved is a free log subscription operation binding the contract event 0x565df8aa55948feffff27bcda8d240698e6a278c311f8bb1f1c067057667154e.
//
// Solidity: event ArbiterResolved(uint8 outcome, bytes32 evidenceHash, string evidenceURI)
func (_BetEscrow *BetEscrowFilterer) WatchArbiterResolved(opts *bind.WatchOpts, sink chan<- *BetEscrowArbiterResolved) (event.Subscription, error) {

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "ArbiterResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowArbiterResolved)
				if err := _BetEscrow.contract.UnpackLog(event, "ArbiterResolved", log); err != nil {
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

// ParseArbiterResolved is a log parse operation binding the contract event 0x565df8aa55948feffff27bcda8d240698e6a278c311f8bb1f1c067057667154e.
//
// Solidity: event ArbiterResolved(uint8 outcome, bytes32 evidenceHash, string evidenceURI)
func (_BetEscrow *BetEscrowFilterer) ParseArbiterResolved(log types.Log) (*BetEscrowArbiterResolved, error) {
	event := new(BetEscrowArbiterResolved)
	if err := _BetEscrow.contract.UnpackLog(event, "ArbiterResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowChallengedIterator is returned from FilterChallenged and is used to iterate over the raw logs and unpacked data for Challenged events raised by the BetEscrow contract.
type BetEscrowChallengedIterator struct {
	Event *BetEscrowChallenged // Event containing the contract specifics and raw log

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
func (it *BetEscrowChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowChallenged)
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
		it.Event = new(BetEscrowChallenged)
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
func (it *BetEscrowChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowChallenged represents a Challenged event raised by the BetEscrow contract.
type BetEscrowChallenged struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallenged is a free log retrieval operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed by)
func (_BetEscrow *BetEscrowFilterer) FilterChallenged(opts *bind.FilterOpts, by []common.Address) (*BetEscrowChallengedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Challenged", byRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowChallengedIterator{contract: _BetEscrow.contract, event: "Challenged", logs: logs, sub: sub}, nil
}

// WatchChallenged is a free log subscription operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed by)
func (_BetEscrow *BetEscrowFilterer) WatchChallenged(opts *bind.WatchOpts, sink chan<- *BetEscrowChallenged, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Challenged", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowChallenged)
				if err := _BetEscrow.contract.UnpackLog(event, "Challenged", log); err != nil {
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

// ParseChallenged is a log parse operation binding the contract event 0x98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c.
//
// Solidity: event Challenged(address indexed by)
func (_BetEscrow *BetEscrowFilterer) ParseChallenged(log types.Log) (*BetEscrowChallenged, error) {
	event := new(BetEscrowChallenged)
	if err := _BetEscrow.contract.UnpackLog(event, "Challenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BetEscrow contract.
type BetEscrowClaimedIterator struct {
	Event *BetEscrowClaimed // Event containing the contract specifics and raw log

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
func (it *BetEscrowClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowClaimed)
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
		it.Event = new(BetEscrowClaimed)
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
func (it *BetEscrowClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowClaimed represents a Claimed event raised by the BetEscrow contract.
type BetEscrowClaimed struct {
	By                common.Address
	Outcome           uint8
	EvidenceHash      [32]byte
	ChallengeDeadline uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd44113183.
//
// Solidity: event Claimed(address indexed by, uint8 outcome, bytes32 evidenceHash, uint64 challengeDeadline)
func (_BetEscrow *BetEscrowFilterer) FilterClaimed(opts *bind.FilterOpts, by []common.Address) (*BetEscrowClaimedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Claimed", byRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowClaimedIterator{contract: _BetEscrow.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd44113183.
//
// Solidity: event Claimed(address indexed by, uint8 outcome, bytes32 evidenceHash, uint64 challengeDeadline)
func (_BetEscrow *BetEscrowFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BetEscrowClaimed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Claimed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowClaimed)
				if err := _BetEscrow.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd44113183.
//
// Solidity: event Claimed(address indexed by, uint8 outcome, bytes32 evidenceHash, uint64 challengeDeadline)
func (_BetEscrow *BetEscrowFilterer) ParseClaimed(log types.Log) (*BetEscrowClaimed, error) {
	event := new(BetEscrowClaimed)
	if err := _BetEscrow.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowFundedIterator is returned from FilterFunded and is used to iterate over the raw logs and unpacked data for Funded events raised by the BetEscrow contract.
type BetEscrowFundedIterator struct {
	Event *BetEscrowFunded // Event containing the contract specifics and raw log

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
func (it *BetEscrowFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowFunded)
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
		it.Event = new(BetEscrowFunded)
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
func (it *BetEscrowFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowFunded represents a Funded event raised by the BetEscrow contract.
type BetEscrowFunded struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFunded is a free log retrieval operation binding the contract event 0xb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c.
//
// Solidity: event Funded(address indexed agent)
func (_BetEscrow *BetEscrowFilterer) FilterFunded(opts *bind.FilterOpts, agent []common.Address) (*BetEscrowFundedIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Funded", agentRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowFundedIterator{contract: _BetEscrow.contract, event: "Funded", logs: logs, sub: sub}, nil
}

// WatchFunded is a free log subscription operation binding the contract event 0xb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c.
//
// Solidity: event Funded(address indexed agent)
func (_BetEscrow *BetEscrowFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *BetEscrowFunded, agent []common.Address) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Funded", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowFunded)
				if err := _BetEscrow.contract.UnpackLog(event, "Funded", log); err != nil {
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

// ParseFunded is a log parse operation binding the contract event 0xb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c.
//
// Solidity: event Funded(address indexed agent)
func (_BetEscrow *BetEscrowFilterer) ParseFunded(log types.Log) (*BetEscrowFunded, error) {
	event := new(BetEscrowFunded)
	if err := _BetEscrow.contract.UnpackLog(event, "Funded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowOutcomeAgreedIterator is returned from FilterOutcomeAgreed and is used to iterate over the raw logs and unpacked data for OutcomeAgreed events raised by the BetEscrow contract.
type BetEscrowOutcomeAgreedIterator struct {
	Event *BetEscrowOutcomeAgreed // Event containing the contract specifics and raw log

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
func (it *BetEscrowOutcomeAgreedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowOutcomeAgreed)
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
		it.Event = new(BetEscrowOutcomeAgreed)
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
func (it *BetEscrowOutcomeAgreedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowOutcomeAgreedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowOutcomeAgreed represents a OutcomeAgreed event raised by the BetEscrow contract.
type BetEscrowOutcomeAgreed struct {
	By      common.Address
	Outcome uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutcomeAgreed is a free log retrieval operation binding the contract event 0x79a012f7d2faaa3daa2d4a00d83acc1a9334b76556f42bd053d2d45bf5a5b6a8.
//
// Solidity: event OutcomeAgreed(address indexed by, uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) FilterOutcomeAgreed(opts *bind.FilterOpts, by []common.Address) (*BetEscrowOutcomeAgreedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "OutcomeAgreed", byRule)
	if err != nil {
		return nil, err
	}
	return &BetEscrowOutcomeAgreedIterator{contract: _BetEscrow.contract, event: "OutcomeAgreed", logs: logs, sub: sub}, nil
}

// WatchOutcomeAgreed is a free log subscription operation binding the contract event 0x79a012f7d2faaa3daa2d4a00d83acc1a9334b76556f42bd053d2d45bf5a5b6a8.
//
// Solidity: event OutcomeAgreed(address indexed by, uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) WatchOutcomeAgreed(opts *bind.WatchOpts, sink chan<- *BetEscrowOutcomeAgreed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "OutcomeAgreed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowOutcomeAgreed)
				if err := _BetEscrow.contract.UnpackLog(event, "OutcomeAgreed", log); err != nil {
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

// ParseOutcomeAgreed is a log parse operation binding the contract event 0x79a012f7d2faaa3daa2d4a00d83acc1a9334b76556f42bd053d2d45bf5a5b6a8.
//
// Solidity: event OutcomeAgreed(address indexed by, uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) ParseOutcomeAgreed(log types.Log) (*BetEscrowOutcomeAgreed, error) {
	event := new(BetEscrowOutcomeAgreed)
	if err := _BetEscrow.contract.UnpackLog(event, "OutcomeAgreed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the BetEscrow contract.
type BetEscrowRevokedIterator struct {
	Event *BetEscrowRevoked // Event containing the contract specifics and raw log

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
func (it *BetEscrowRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowRevoked)
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
		it.Event = new(BetEscrowRevoked)
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
func (it *BetEscrowRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowRevoked represents a Revoked event raised by the BetEscrow contract.
type BetEscrowRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_BetEscrow *BetEscrowFilterer) FilterRevoked(opts *bind.FilterOpts) (*BetEscrowRevokedIterator, error) {

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &BetEscrowRevokedIterator{contract: _BetEscrow.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_BetEscrow *BetEscrowFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *BetEscrowRevoked) (event.Subscription, error) {

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowRevoked)
				if err := _BetEscrow.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_BetEscrow *BetEscrowFilterer) ParseRevoked(log types.Log) (*BetEscrowRevoked, error) {
	event := new(BetEscrowRevoked)
	if err := _BetEscrow.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowSettledIterator is returned from FilterSettled and is used to iterate over the raw logs and unpacked data for Settled events raised by the BetEscrow contract.
type BetEscrowSettledIterator struct {
	Event *BetEscrowSettled // Event containing the contract specifics and raw log

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
func (it *BetEscrowSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowSettled)
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
		it.Event = new(BetEscrowSettled)
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
func (it *BetEscrowSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowSettled represents a Settled event raised by the BetEscrow contract.
type BetEscrowSettled struct {
	Outcome uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettled is a free log retrieval operation binding the contract event 0x56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f521.
//
// Solidity: event Settled(uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) FilterSettled(opts *bind.FilterOpts) (*BetEscrowSettledIterator, error) {

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "Settled")
	if err != nil {
		return nil, err
	}
	return &BetEscrowSettledIterator{contract: _BetEscrow.contract, event: "Settled", logs: logs, sub: sub}, nil
}

// WatchSettled is a free log subscription operation binding the contract event 0x56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f521.
//
// Solidity: event Settled(uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) WatchSettled(opts *bind.WatchOpts, sink chan<- *BetEscrowSettled) (event.Subscription, error) {

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "Settled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowSettled)
				if err := _BetEscrow.contract.UnpackLog(event, "Settled", log); err != nil {
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

// ParseSettled is a log parse operation binding the contract event 0x56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f521.
//
// Solidity: event Settled(uint8 outcome)
func (_BetEscrow *BetEscrowFilterer) ParseSettled(log types.Log) (*BetEscrowSettled, error) {
	event := new(BetEscrowSettled)
	if err := _BetEscrow.contract.UnpackLog(event, "Settled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BetEscrowWentLiveIterator is returned from FilterWentLive and is used to iterate over the raw logs and unpacked data for WentLive events raised by the BetEscrow contract.
type BetEscrowWentLiveIterator struct {
	Event *BetEscrowWentLive // Event containing the contract specifics and raw log

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
func (it *BetEscrowWentLiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BetEscrowWentLive)
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
		it.Event = new(BetEscrowWentLive)
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
func (it *BetEscrowWentLiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BetEscrowWentLiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BetEscrowWentLive represents a WentLive event raised by the BetEscrow contract.
type BetEscrowWentLive struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWentLive is a free log retrieval operation binding the contract event 0x12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb.
//
// Solidity: event WentLive()
func (_BetEscrow *BetEscrowFilterer) FilterWentLive(opts *bind.FilterOpts) (*BetEscrowWentLiveIterator, error) {

	logs, sub, err := _BetEscrow.contract.FilterLogs(opts, "WentLive")
	if err != nil {
		return nil, err
	}
	return &BetEscrowWentLiveIterator{contract: _BetEscrow.contract, event: "WentLive", logs: logs, sub: sub}, nil
}

// WatchWentLive is a free log subscription operation binding the contract event 0x12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb.
//
// Solidity: event WentLive()
func (_BetEscrow *BetEscrowFilterer) WatchWentLive(opts *bind.WatchOpts, sink chan<- *BetEscrowWentLive) (event.Subscription, error) {

	logs, sub, err := _BetEscrow.contract.WatchLogs(opts, "WentLive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BetEscrowWentLive)
				if err := _BetEscrow.contract.UnpackLog(event, "WentLive", log); err != nil {
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

// ParseWentLive is a log parse operation binding the contract event 0x12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb.
//
// Solidity: event WentLive()
func (_BetEscrow *BetEscrowFilterer) ParseWentLive(log types.Log) (*BetEscrowWentLive, error) {
	event := new(BetEscrowWentLive)
	if err := _BetEscrow.contract.UnpackLog(event, "WentLive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
