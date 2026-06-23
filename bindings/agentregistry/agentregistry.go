// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agentregistry

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

// AgentStorageProfile is an auto generated low-level Go binding around an user-defined struct.
type AgentStorageProfile struct {
	Owner        common.Address
	Wallet       common.Address
	Signer       common.Address
	PolicyHash   [32]byte
	MetadataHash [32]byte
	Bond         *big.Int
	Reserved     *big.Int
	CreatedAt    uint64
	Active       bool
}

// AgentRegistryMetaData contains all meta data concerning the AgentRegistry contract.
var AgentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"agent\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAgentStorage.Profile\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserved\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"arbitrationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"charge\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deactivate\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freeBond\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"store_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrationFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registrationBond_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"arbitrationFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onArbiteredBet\",\"inputs\":[{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"yesAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registrationBond\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reserveTaker\",\"inputs\":[{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setArbitrationFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFactory\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRevenueWallet\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signerOf\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAgentStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sweepFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePolicy\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSigner\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminTransferred\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentDeactivated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentPolicyUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRegistered\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentSignerUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ArbitrationFeeUpdated\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondCharged\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondRefunded\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondReleased\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondReserved\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondSlashed\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FactoryUpdated\",\"inputs\":[{\"name\":\"factory\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesSwept\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevenueWalletUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BondReservedOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBond\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEscrow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFactory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// AgentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentRegistryMetaData.ABI instead.
var AgentRegistryABI = AgentRegistryMetaData.ABI

// AgentRegistry is an auto generated Go binding around an Ethereum contract.
type AgentRegistry struct {
	AgentRegistryCaller     // Read-only binding to the contract
	AgentRegistryTransactor // Write-only binding to the contract
	AgentRegistryFilterer   // Log filterer for contract events
}

// AgentRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentRegistrySession struct {
	Contract     *AgentRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentRegistryCallerSession struct {
	Contract *AgentRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AgentRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentRegistryTransactorSession struct {
	Contract     *AgentRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgentRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRegistryRaw struct {
	Contract *AgentRegistry // Generic contract binding to access the raw methods on
}

// AgentRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentRegistryCallerRaw struct {
	Contract *AgentRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AgentRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentRegistryTransactorRaw struct {
	Contract *AgentRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentRegistry creates a new instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistry(address common.Address, backend bind.ContractBackend) (*AgentRegistry, error) {
	contract, err := bindAgentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentRegistry{AgentRegistryCaller: AgentRegistryCaller{contract: contract}, AgentRegistryTransactor: AgentRegistryTransactor{contract: contract}, AgentRegistryFilterer: AgentRegistryFilterer{contract: contract}}, nil
}

// NewAgentRegistryCaller creates a new read-only instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryCaller(address common.Address, caller bind.ContractCaller) (*AgentRegistryCaller, error) {
	contract, err := bindAgentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryCaller{contract: contract}, nil
}

// NewAgentRegistryTransactor creates a new write-only instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentRegistryTransactor, error) {
	contract, err := bindAgentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryTransactor{contract: contract}, nil
}

// NewAgentRegistryFilterer creates a new log filterer instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentRegistryFilterer, error) {
	contract, err := bindAgentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryFilterer{contract: contract}, nil
}

// bindAgentRegistry binds a generic wrapper to an already deployed contract.
func bindAgentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistry *AgentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistry.Contract.AgentRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistry *AgentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistry.Contract.AgentRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistry *AgentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistry.Contract.AgentRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistry *AgentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistry *AgentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistry *AgentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AgentRegistry *AgentRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AgentRegistry *AgentRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AgentRegistry.Contract.UPGRADEINTERFACEVERSION(&_AgentRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AgentRegistry *AgentRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AgentRegistry.Contract.UPGRADEINTERFACEVERSION(&_AgentRegistry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AgentRegistry *AgentRegistryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AgentRegistry *AgentRegistrySession) Admin() (common.Address, error) {
	return _AgentRegistry.Contract.Admin(&_AgentRegistry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) Admin() (common.Address, error) {
	return _AgentRegistry.Contract.Admin(&_AgentRegistry.CallOpts)
}

// Agent is a free data retrieval call binding the contract method 0x92e423b5.
//
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistryCaller) Agent(opts *bind.CallOpts, wallet common.Address) (AgentStorageProfile, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "agent", wallet)

	if err != nil {
		return *new(AgentStorageProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStorageProfile)).(*AgentStorageProfile)

	return out0, err

}

// Agent is a free data retrieval call binding the contract method 0x92e423b5.
//
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistrySession) Agent(wallet common.Address) (AgentStorageProfile, error) {
	return _AgentRegistry.Contract.Agent(&_AgentRegistry.CallOpts, wallet)
}

// Agent is a free data retrieval call binding the contract method 0x92e423b5.
//
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistryCallerSession) Agent(wallet common.Address) (AgentStorageProfile, error) {
	return _AgentRegistry.Contract.Agent(&_AgentRegistry.CallOpts, wallet)
}

// ArbitrationFee is a free data retrieval call binding the contract method 0x1cfa7ba8.
//
// Solidity: function arbitrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) ArbitrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "arbitrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitrationFee is a free data retrieval call binding the contract method 0x1cfa7ba8.
//
// Solidity: function arbitrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) ArbitrationFee() (*big.Int, error) {
	return _AgentRegistry.Contract.ArbitrationFee(&_AgentRegistry.CallOpts)
}

// ArbitrationFee is a free data retrieval call binding the contract method 0x1cfa7ba8.
//
// Solidity: function arbitrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) ArbitrationFee() (*big.Int, error) {
	return _AgentRegistry.Contract.ArbitrationFee(&_AgentRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AgentRegistry *AgentRegistryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AgentRegistry *AgentRegistrySession) Factory() (common.Address, error) {
	return _AgentRegistry.Contract.Factory(&_AgentRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) Factory() (common.Address, error) {
	return _AgentRegistry.Contract.Factory(&_AgentRegistry.CallOpts)
}

// FreeBond is a free data retrieval call binding the contract method 0xe132b025.
//
// Solidity: function freeBond(address wallet) view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) FreeBond(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "freeBond", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeBond is a free data retrieval call binding the contract method 0xe132b025.
//
// Solidity: function freeBond(address wallet) view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) FreeBond(wallet common.Address) (*big.Int, error) {
	return _AgentRegistry.Contract.FreeBond(&_AgentRegistry.CallOpts, wallet)
}

// FreeBond is a free data retrieval call binding the contract method 0xe132b025.
//
// Solidity: function freeBond(address wallet) view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) FreeBond(wallet common.Address) (*big.Int, error) {
	return _AgentRegistry.Contract.FreeBond(&_AgentRegistry.CallOpts, wallet)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address wallet) view returns(bool)
func (_AgentRegistry *AgentRegistryCaller) IsActive(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "isActive", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address wallet) view returns(bool)
func (_AgentRegistry *AgentRegistrySession) IsActive(wallet common.Address) (bool, error) {
	return _AgentRegistry.Contract.IsActive(&_AgentRegistry.CallOpts, wallet)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address wallet) view returns(bool)
func (_AgentRegistry *AgentRegistryCallerSession) IsActive(wallet common.Address) (bool, error) {
	return _AgentRegistry.Contract.IsActive(&_AgentRegistry.CallOpts, wallet)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AgentRegistry *AgentRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AgentRegistry *AgentRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _AgentRegistry.Contract.ProxiableUUID(&_AgentRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AgentRegistry *AgentRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AgentRegistry.Contract.ProxiableUUID(&_AgentRegistry.CallOpts)
}

// RegistrationBond is a free data retrieval call binding the contract method 0x373b9dbd.
//
// Solidity: function registrationBond() view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) RegistrationBond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "registrationBond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationBond is a free data retrieval call binding the contract method 0x373b9dbd.
//
// Solidity: function registrationBond() view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) RegistrationBond() (*big.Int, error) {
	return _AgentRegistry.Contract.RegistrationBond(&_AgentRegistry.CallOpts)
}

// RegistrationBond is a free data retrieval call binding the contract method 0x373b9dbd.
//
// Solidity: function registrationBond() view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) RegistrationBond() (*big.Int, error) {
	return _AgentRegistry.Contract.RegistrationBond(&_AgentRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "registrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) RegistrationFee() (*big.Int, error) {
	return _AgentRegistry.Contract.RegistrationFee(&_AgentRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _AgentRegistry.Contract.RegistrationFee(&_AgentRegistry.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_AgentRegistry *AgentRegistryCaller) RevenueWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "revenueWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_AgentRegistry *AgentRegistrySession) RevenueWallet() (common.Address, error) {
	return _AgentRegistry.Contract.RevenueWallet(&_AgentRegistry.CallOpts)
}

// RevenueWallet is a free data retrieval call binding the contract method 0x44478425.
//
// Solidity: function revenueWallet() view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) RevenueWallet() (common.Address, error) {
	return _AgentRegistry.Contract.RevenueWallet(&_AgentRegistry.CallOpts)
}

// SignerOf is a free data retrieval call binding the contract method 0x5b5175d5.
//
// Solidity: function signerOf(address wallet) view returns(address)
func (_AgentRegistry *AgentRegistryCaller) SignerOf(opts *bind.CallOpts, wallet common.Address) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "signerOf", wallet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerOf is a free data retrieval call binding the contract method 0x5b5175d5.
//
// Solidity: function signerOf(address wallet) view returns(address)
func (_AgentRegistry *AgentRegistrySession) SignerOf(wallet common.Address) (common.Address, error) {
	return _AgentRegistry.Contract.SignerOf(&_AgentRegistry.CallOpts, wallet)
}

// SignerOf is a free data retrieval call binding the contract method 0x5b5175d5.
//
// Solidity: function signerOf(address wallet) view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) SignerOf(wallet common.Address) (common.Address, error) {
	return _AgentRegistry.Contract.SignerOf(&_AgentRegistry.CallOpts, wallet)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_AgentRegistry *AgentRegistryCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "store")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_AgentRegistry *AgentRegistrySession) Store() (common.Address, error) {
	return _AgentRegistry.Contract.Store(&_AgentRegistry.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) Store() (common.Address, error) {
	return _AgentRegistry.Contract.Store(&_AgentRegistry.CallOpts)
}

// Charge is a paid mutator transaction binding the contract method 0xc709a4f1.
//
// Solidity: function charge(address wallet, address to, uint256 amount) returns()
func (_AgentRegistry *AgentRegistryTransactor) Charge(opts *bind.TransactOpts, wallet common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "charge", wallet, to, amount)
}

// Charge is a paid mutator transaction binding the contract method 0xc709a4f1.
//
// Solidity: function charge(address wallet, address to, uint256 amount) returns()
func (_AgentRegistry *AgentRegistrySession) Charge(wallet common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Charge(&_AgentRegistry.TransactOpts, wallet, to, amount)
}

// Charge is a paid mutator transaction binding the contract method 0xc709a4f1.
//
// Solidity: function charge(address wallet, address to, uint256 amount) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Charge(wallet common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Charge(&_AgentRegistry.TransactOpts, wallet, to, amount)
}

// Deactivate is a paid mutator transaction binding the contract method 0x3ea053eb.
//
// Solidity: function deactivate(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactor) Deactivate(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "deactivate", wallet)
}

// Deactivate is a paid mutator transaction binding the contract method 0x3ea053eb.
//
// Solidity: function deactivate(address wallet) returns()
func (_AgentRegistry *AgentRegistrySession) Deactivate(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Deactivate(&_AgentRegistry.TransactOpts, wallet)
}

// Deactivate is a paid mutator transaction binding the contract method 0x3ea053eb.
//
// Solidity: function deactivate(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Deactivate(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Deactivate(&_AgentRegistry.TransactOpts, wallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x49bc779b.
//
// Solidity: function initialize(address store_, uint256 registrationFee_, uint256 registrationBond_, uint256 arbitrationFee_, address admin_, address revenueWallet_, address factory_) returns()
func (_AgentRegistry *AgentRegistryTransactor) Initialize(opts *bind.TransactOpts, store_ common.Address, registrationFee_ *big.Int, registrationBond_ *big.Int, arbitrationFee_ *big.Int, admin_ common.Address, revenueWallet_ common.Address, factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "initialize", store_, registrationFee_, registrationBond_, arbitrationFee_, admin_, revenueWallet_, factory_)
}

// Initialize is a paid mutator transaction binding the contract method 0x49bc779b.
//
// Solidity: function initialize(address store_, uint256 registrationFee_, uint256 registrationBond_, uint256 arbitrationFee_, address admin_, address revenueWallet_, address factory_) returns()
func (_AgentRegistry *AgentRegistrySession) Initialize(store_ common.Address, registrationFee_ *big.Int, registrationBond_ *big.Int, arbitrationFee_ *big.Int, admin_ common.Address, revenueWallet_ common.Address, factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Initialize(&_AgentRegistry.TransactOpts, store_, registrationFee_, registrationBond_, arbitrationFee_, admin_, revenueWallet_, factory_)
}

// Initialize is a paid mutator transaction binding the contract method 0x49bc779b.
//
// Solidity: function initialize(address store_, uint256 registrationFee_, uint256 registrationBond_, uint256 arbitrationFee_, address admin_, address revenueWallet_, address factory_) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Initialize(store_ common.Address, registrationFee_ *big.Int, registrationBond_ *big.Int, arbitrationFee_ *big.Int, admin_ common.Address, revenueWallet_ common.Address, factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Initialize(&_AgentRegistry.TransactOpts, store_, registrationFee_, registrationBond_, arbitrationFee_, admin_, revenueWallet_, factory_)
}

// OnArbiteredBet is a paid mutator transaction binding the contract method 0xdf1d0b3a.
//
// Solidity: function onArbiteredBet(address escrow, address yesAgent, address noAgent) returns()
func (_AgentRegistry *AgentRegistryTransactor) OnArbiteredBet(opts *bind.TransactOpts, escrow common.Address, yesAgent common.Address, noAgent common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "onArbiteredBet", escrow, yesAgent, noAgent)
}

// OnArbiteredBet is a paid mutator transaction binding the contract method 0xdf1d0b3a.
//
// Solidity: function onArbiteredBet(address escrow, address yesAgent, address noAgent) returns()
func (_AgentRegistry *AgentRegistrySession) OnArbiteredBet(escrow common.Address, yesAgent common.Address, noAgent common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.OnArbiteredBet(&_AgentRegistry.TransactOpts, escrow, yesAgent, noAgent)
}

// OnArbiteredBet is a paid mutator transaction binding the contract method 0xdf1d0b3a.
//
// Solidity: function onArbiteredBet(address escrow, address yesAgent, address noAgent) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) OnArbiteredBet(escrow common.Address, yesAgent common.Address, noAgent common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.OnArbiteredBet(&_AgentRegistry.TransactOpts, escrow, yesAgent, noAgent)
}

// Register is a paid mutator transaction binding the contract method 0xe8527e38.
//
// Solidity: function register(address wallet, address signer, bytes32 policyHash, bytes32 metadataHash) returns()
func (_AgentRegistry *AgentRegistryTransactor) Register(opts *bind.TransactOpts, wallet common.Address, signer common.Address, policyHash [32]byte, metadataHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "register", wallet, signer, policyHash, metadataHash)
}

// Register is a paid mutator transaction binding the contract method 0xe8527e38.
//
// Solidity: function register(address wallet, address signer, bytes32 policyHash, bytes32 metadataHash) returns()
func (_AgentRegistry *AgentRegistrySession) Register(wallet common.Address, signer common.Address, policyHash [32]byte, metadataHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Register(&_AgentRegistry.TransactOpts, wallet, signer, policyHash, metadataHash)
}

// Register is a paid mutator transaction binding the contract method 0xe8527e38.
//
// Solidity: function register(address wallet, address signer, bytes32 policyHash, bytes32 metadataHash) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Register(wallet common.Address, signer common.Address, policyHash [32]byte, metadataHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Register(&_AgentRegistry.TransactOpts, wallet, signer, policyHash, metadataHash)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactor) Release(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "release", wallet)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address wallet) returns()
func (_AgentRegistry *AgentRegistrySession) Release(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Release(&_AgentRegistry.TransactOpts, wallet)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Release(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Release(&_AgentRegistry.TransactOpts, wallet)
}

// ReserveTaker is a paid mutator transaction binding the contract method 0x82f81702.
//
// Solidity: function reserveTaker(address taker) returns()
func (_AgentRegistry *AgentRegistryTransactor) ReserveTaker(opts *bind.TransactOpts, taker common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "reserveTaker", taker)
}

// ReserveTaker is a paid mutator transaction binding the contract method 0x82f81702.
//
// Solidity: function reserveTaker(address taker) returns()
func (_AgentRegistry *AgentRegistrySession) ReserveTaker(taker common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.ReserveTaker(&_AgentRegistry.TransactOpts, taker)
}

// ReserveTaker is a paid mutator transaction binding the contract method 0x82f81702.
//
// Solidity: function reserveTaker(address taker) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) ReserveTaker(taker common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.ReserveTaker(&_AgentRegistry.TransactOpts, taker)
}

// SetArbitrationFee is a paid mutator transaction binding the contract method 0x5ea7b4fc.
//
// Solidity: function setArbitrationFee(uint256 fee) returns()
func (_AgentRegistry *AgentRegistryTransactor) SetArbitrationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "setArbitrationFee", fee)
}

// SetArbitrationFee is a paid mutator transaction binding the contract method 0x5ea7b4fc.
//
// Solidity: function setArbitrationFee(uint256 fee) returns()
func (_AgentRegistry *AgentRegistrySession) SetArbitrationFee(fee *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetArbitrationFee(&_AgentRegistry.TransactOpts, fee)
}

// SetArbitrationFee is a paid mutator transaction binding the contract method 0x5ea7b4fc.
//
// Solidity: function setArbitrationFee(uint256 fee) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) SetArbitrationFee(fee *big.Int) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetArbitrationFee(&_AgentRegistry.TransactOpts, fee)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_AgentRegistry *AgentRegistryTransactor) SetFactory(opts *bind.TransactOpts, factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "setFactory", factory_)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_AgentRegistry *AgentRegistrySession) SetFactory(factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetFactory(&_AgentRegistry.TransactOpts, factory_)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address factory_) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) SetFactory(factory_ common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetFactory(&_AgentRegistry.TransactOpts, factory_)
}

// SetRevenueWallet is a paid mutator transaction binding the contract method 0xfb235f34.
//
// Solidity: function setRevenueWallet(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactor) SetRevenueWallet(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "setRevenueWallet", wallet)
}

// SetRevenueWallet is a paid mutator transaction binding the contract method 0xfb235f34.
//
// Solidity: function setRevenueWallet(address wallet) returns()
func (_AgentRegistry *AgentRegistrySession) SetRevenueWallet(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetRevenueWallet(&_AgentRegistry.TransactOpts, wallet)
}

// SetRevenueWallet is a paid mutator transaction binding the contract method 0xfb235f34.
//
// Solidity: function setRevenueWallet(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) SetRevenueWallet(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.SetRevenueWallet(&_AgentRegistry.TransactOpts, wallet)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactor) Slash(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "slash", wallet)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address wallet) returns()
func (_AgentRegistry *AgentRegistrySession) Slash(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Slash(&_AgentRegistry.TransactOpts, wallet)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address wallet) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) Slash(wallet common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.Slash(&_AgentRegistry.TransactOpts, wallet)
}

// SweepFees is a paid mutator transaction binding the contract method 0xd113b95c.
//
// Solidity: function sweepFees() returns()
func (_AgentRegistry *AgentRegistryTransactor) SweepFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "sweepFees")
}

// SweepFees is a paid mutator transaction binding the contract method 0xd113b95c.
//
// Solidity: function sweepFees() returns()
func (_AgentRegistry *AgentRegistrySession) SweepFees() (*types.Transaction, error) {
	return _AgentRegistry.Contract.SweepFees(&_AgentRegistry.TransactOpts)
}

// SweepFees is a paid mutator transaction binding the contract method 0xd113b95c.
//
// Solidity: function sweepFees() returns()
func (_AgentRegistry *AgentRegistryTransactorSession) SweepFees() (*types.Transaction, error) {
	return _AgentRegistry.Contract.SweepFees(&_AgentRegistry.TransactOpts)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AgentRegistry *AgentRegistryTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AgentRegistry *AgentRegistrySession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.TransferAdmin(&_AgentRegistry.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.TransferAdmin(&_AgentRegistry.TransactOpts, newAdmin)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x97cc56fa.
//
// Solidity: function updatePolicy(address wallet, bytes32 policyHash) returns()
func (_AgentRegistry *AgentRegistryTransactor) UpdatePolicy(opts *bind.TransactOpts, wallet common.Address, policyHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "updatePolicy", wallet, policyHash)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x97cc56fa.
//
// Solidity: function updatePolicy(address wallet, bytes32 policyHash) returns()
func (_AgentRegistry *AgentRegistrySession) UpdatePolicy(wallet common.Address, policyHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpdatePolicy(&_AgentRegistry.TransactOpts, wallet, policyHash)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x97cc56fa.
//
// Solidity: function updatePolicy(address wallet, bytes32 policyHash) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) UpdatePolicy(wallet common.Address, policyHash [32]byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpdatePolicy(&_AgentRegistry.TransactOpts, wallet, policyHash)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8fd016e8.
//
// Solidity: function updateSigner(address wallet, address signer) returns()
func (_AgentRegistry *AgentRegistryTransactor) UpdateSigner(opts *bind.TransactOpts, wallet common.Address, signer common.Address) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "updateSigner", wallet, signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8fd016e8.
//
// Solidity: function updateSigner(address wallet, address signer) returns()
func (_AgentRegistry *AgentRegistrySession) UpdateSigner(wallet common.Address, signer common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpdateSigner(&_AgentRegistry.TransactOpts, wallet, signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8fd016e8.
//
// Solidity: function updateSigner(address wallet, address signer) returns()
func (_AgentRegistry *AgentRegistryTransactorSession) UpdateSigner(wallet common.Address, signer common.Address) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpdateSigner(&_AgentRegistry.TransactOpts, wallet, signer)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AgentRegistry *AgentRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AgentRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AgentRegistry *AgentRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpgradeToAndCall(&_AgentRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AgentRegistry *AgentRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AgentRegistry.Contract.UpgradeToAndCall(&_AgentRegistry.TransactOpts, newImplementation, data)
}

// AgentRegistryAdminTransferredIterator is returned from FilterAdminTransferred and is used to iterate over the raw logs and unpacked data for AdminTransferred events raised by the AgentRegistry contract.
type AgentRegistryAdminTransferredIterator struct {
	Event *AgentRegistryAdminTransferred // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAdminTransferred)
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
		it.Event = new(AgentRegistryAdminTransferred)
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
func (it *AgentRegistryAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAdminTransferred represents a AdminTransferred event raised by the AgentRegistry contract.
type AgentRegistryAdminTransferred struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminTransferred is a free log retrieval operation binding the contract event 0xe22b4f506b7da9a528a95d7063a6cde8d2b6268773b1f7f65c00057531704a3d.
//
// Solidity: event AdminTransferred(address indexed admin)
func (_AgentRegistry *AgentRegistryFilterer) FilterAdminTransferred(opts *bind.FilterOpts, admin []common.Address) (*AgentRegistryAdminTransferredIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AdminTransferred", adminRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAdminTransferredIterator{contract: _AgentRegistry.contract, event: "AdminTransferred", logs: logs, sub: sub}, nil
}

// WatchAdminTransferred is a free log subscription operation binding the contract event 0xe22b4f506b7da9a528a95d7063a6cde8d2b6268773b1f7f65c00057531704a3d.
//
// Solidity: event AdminTransferred(address indexed admin)
func (_AgentRegistry *AgentRegistryFilterer) WatchAdminTransferred(opts *bind.WatchOpts, sink chan<- *AgentRegistryAdminTransferred, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AdminTransferred", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAdminTransferred)
				if err := _AgentRegistry.contract.UnpackLog(event, "AdminTransferred", log); err != nil {
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

// ParseAdminTransferred is a log parse operation binding the contract event 0xe22b4f506b7da9a528a95d7063a6cde8d2b6268773b1f7f65c00057531704a3d.
//
// Solidity: event AdminTransferred(address indexed admin)
func (_AgentRegistry *AgentRegistryFilterer) ParseAdminTransferred(log types.Log) (*AgentRegistryAdminTransferred, error) {
	event := new(AgentRegistryAdminTransferred)
	if err := _AgentRegistry.contract.UnpackLog(event, "AdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentDeactivatedIterator is returned from FilterAgentDeactivated and is used to iterate over the raw logs and unpacked data for AgentDeactivated events raised by the AgentRegistry contract.
type AgentRegistryAgentDeactivatedIterator struct {
	Event *AgentRegistryAgentDeactivated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentDeactivated)
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
		it.Event = new(AgentRegistryAgentDeactivated)
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
func (it *AgentRegistryAgentDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentDeactivated represents a AgentDeactivated event raised by the AgentRegistry contract.
type AgentRegistryAgentDeactivated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentDeactivated is a free log retrieval operation binding the contract event 0xa58c733780f5c1527d6106ff1e731cbf2e9ffc0eb11ad399ca33ea55f057ddf4.
//
// Solidity: event AgentDeactivated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentDeactivated(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryAgentDeactivatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentDeactivated", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentDeactivatedIterator{contract: _AgentRegistry.contract, event: "AgentDeactivated", logs: logs, sub: sub}, nil
}

// WatchAgentDeactivated is a free log subscription operation binding the contract event 0xa58c733780f5c1527d6106ff1e731cbf2e9ffc0eb11ad399ca33ea55f057ddf4.
//
// Solidity: event AgentDeactivated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentDeactivated(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentDeactivated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentDeactivated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentDeactivated)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentDeactivated", log); err != nil {
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

// ParseAgentDeactivated is a log parse operation binding the contract event 0xa58c733780f5c1527d6106ff1e731cbf2e9ffc0eb11ad399ca33ea55f057ddf4.
//
// Solidity: event AgentDeactivated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentDeactivated(log types.Log) (*AgentRegistryAgentDeactivated, error) {
	event := new(AgentRegistryAgentDeactivated)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentPolicyUpdatedIterator is returned from FilterAgentPolicyUpdated and is used to iterate over the raw logs and unpacked data for AgentPolicyUpdated events raised by the AgentRegistry contract.
type AgentRegistryAgentPolicyUpdatedIterator struct {
	Event *AgentRegistryAgentPolicyUpdated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentPolicyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentPolicyUpdated)
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
		it.Event = new(AgentRegistryAgentPolicyUpdated)
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
func (it *AgentRegistryAgentPolicyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentPolicyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentPolicyUpdated represents a AgentPolicyUpdated event raised by the AgentRegistry contract.
type AgentRegistryAgentPolicyUpdated struct {
	Wallet     common.Address
	PolicyHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentPolicyUpdated is a free log retrieval operation binding the contract event 0xcc8946cc5e90c32269cfbfe3ca38dea5bf8e331038589b2f3ede2c35c42fffa3.
//
// Solidity: event AgentPolicyUpdated(address indexed wallet, bytes32 policyHash)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentPolicyUpdated(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryAgentPolicyUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentPolicyUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentPolicyUpdatedIterator{contract: _AgentRegistry.contract, event: "AgentPolicyUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentPolicyUpdated is a free log subscription operation binding the contract event 0xcc8946cc5e90c32269cfbfe3ca38dea5bf8e331038589b2f3ede2c35c42fffa3.
//
// Solidity: event AgentPolicyUpdated(address indexed wallet, bytes32 policyHash)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentPolicyUpdated(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentPolicyUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentPolicyUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentPolicyUpdated)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentPolicyUpdated", log); err != nil {
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

// ParseAgentPolicyUpdated is a log parse operation binding the contract event 0xcc8946cc5e90c32269cfbfe3ca38dea5bf8e331038589b2f3ede2c35c42fffa3.
//
// Solidity: event AgentPolicyUpdated(address indexed wallet, bytes32 policyHash)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentPolicyUpdated(log types.Log) (*AgentRegistryAgentPolicyUpdated, error) {
	event := new(AgentRegistryAgentPolicyUpdated)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentPolicyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentRegisteredIterator is returned from FilterAgentRegistered and is used to iterate over the raw logs and unpacked data for AgentRegistered events raised by the AgentRegistry contract.
type AgentRegistryAgentRegisteredIterator struct {
	Event *AgentRegistryAgentRegistered // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentRegistered)
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
		it.Event = new(AgentRegistryAgentRegistered)
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
func (it *AgentRegistryAgentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentRegistered represents a AgentRegistered event raised by the AgentRegistry contract.
type AgentRegistryAgentRegistered struct {
	Wallet common.Address
	Owner  common.Address
	Signer common.Address
	Bond   *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentRegistered is a free log retrieval operation binding the contract event 0xcc059d6320dd91f6d72ca6e83caaf137242df66d024a6e14412a0472557dba63.
//
// Solidity: event AgentRegistered(address indexed wallet, address indexed owner, address signer, uint256 bond, uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentRegistered(opts *bind.FilterOpts, wallet []common.Address, owner []common.Address) (*AgentRegistryAgentRegisteredIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentRegistered", walletRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentRegisteredIterator{contract: _AgentRegistry.contract, event: "AgentRegistered", logs: logs, sub: sub}, nil
}

// WatchAgentRegistered is a free log subscription operation binding the contract event 0xcc059d6320dd91f6d72ca6e83caaf137242df66d024a6e14412a0472557dba63.
//
// Solidity: event AgentRegistered(address indexed wallet, address indexed owner, address signer, uint256 bond, uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentRegistered(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentRegistered, wallet []common.Address, owner []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentRegistered", walletRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentRegistered)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
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

// ParseAgentRegistered is a log parse operation binding the contract event 0xcc059d6320dd91f6d72ca6e83caaf137242df66d024a6e14412a0472557dba63.
//
// Solidity: event AgentRegistered(address indexed wallet, address indexed owner, address signer, uint256 bond, uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentRegistered(log types.Log) (*AgentRegistryAgentRegistered, error) {
	event := new(AgentRegistryAgentRegistered)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentSignerUpdatedIterator is returned from FilterAgentSignerUpdated and is used to iterate over the raw logs and unpacked data for AgentSignerUpdated events raised by the AgentRegistry contract.
type AgentRegistryAgentSignerUpdatedIterator struct {
	Event *AgentRegistryAgentSignerUpdated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentSignerUpdated)
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
		it.Event = new(AgentRegistryAgentSignerUpdated)
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
func (it *AgentRegistryAgentSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentSignerUpdated represents a AgentSignerUpdated event raised by the AgentRegistry contract.
type AgentRegistryAgentSignerUpdated struct {
	Wallet common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSignerUpdated is a free log retrieval operation binding the contract event 0x15c539f3c72a7122b9be60c5a7d9929c4882896eca5158969a1877c067192d9b.
//
// Solidity: event AgentSignerUpdated(address indexed wallet, address signer)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentSignerUpdated(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryAgentSignerUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentSignerUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentSignerUpdatedIterator{contract: _AgentRegistry.contract, event: "AgentSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentSignerUpdated is a free log subscription operation binding the contract event 0x15c539f3c72a7122b9be60c5a7d9929c4882896eca5158969a1877c067192d9b.
//
// Solidity: event AgentSignerUpdated(address indexed wallet, address signer)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentSignerUpdated(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentSignerUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentSignerUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentSignerUpdated)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentSignerUpdated", log); err != nil {
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

// ParseAgentSignerUpdated is a log parse operation binding the contract event 0x15c539f3c72a7122b9be60c5a7d9929c4882896eca5158969a1877c067192d9b.
//
// Solidity: event AgentSignerUpdated(address indexed wallet, address signer)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentSignerUpdated(log types.Log) (*AgentRegistryAgentSignerUpdated, error) {
	event := new(AgentRegistryAgentSignerUpdated)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryArbitrationFeeUpdatedIterator is returned from FilterArbitrationFeeUpdated and is used to iterate over the raw logs and unpacked data for ArbitrationFeeUpdated events raised by the AgentRegistry contract.
type AgentRegistryArbitrationFeeUpdatedIterator struct {
	Event *AgentRegistryArbitrationFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryArbitrationFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryArbitrationFeeUpdated)
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
		it.Event = new(AgentRegistryArbitrationFeeUpdated)
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
func (it *AgentRegistryArbitrationFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryArbitrationFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryArbitrationFeeUpdated represents a ArbitrationFeeUpdated event raised by the AgentRegistry contract.
type AgentRegistryArbitrationFeeUpdated struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterArbitrationFeeUpdated is a free log retrieval operation binding the contract event 0xb1484c2bf00d94a00783b6081ebc5f5d02be4675f6eb8cf4c0c95bfe5a3f06ed.
//
// Solidity: event ArbitrationFeeUpdated(uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) FilterArbitrationFeeUpdated(opts *bind.FilterOpts) (*AgentRegistryArbitrationFeeUpdatedIterator, error) {

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "ArbitrationFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentRegistryArbitrationFeeUpdatedIterator{contract: _AgentRegistry.contract, event: "ArbitrationFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchArbitrationFeeUpdated is a free log subscription operation binding the contract event 0xb1484c2bf00d94a00783b6081ebc5f5d02be4675f6eb8cf4c0c95bfe5a3f06ed.
//
// Solidity: event ArbitrationFeeUpdated(uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) WatchArbitrationFeeUpdated(opts *bind.WatchOpts, sink chan<- *AgentRegistryArbitrationFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "ArbitrationFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryArbitrationFeeUpdated)
				if err := _AgentRegistry.contract.UnpackLog(event, "ArbitrationFeeUpdated", log); err != nil {
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

// ParseArbitrationFeeUpdated is a log parse operation binding the contract event 0xb1484c2bf00d94a00783b6081ebc5f5d02be4675f6eb8cf4c0c95bfe5a3f06ed.
//
// Solidity: event ArbitrationFeeUpdated(uint256 fee)
func (_AgentRegistry *AgentRegistryFilterer) ParseArbitrationFeeUpdated(log types.Log) (*AgentRegistryArbitrationFeeUpdated, error) {
	event := new(AgentRegistryArbitrationFeeUpdated)
	if err := _AgentRegistry.contract.UnpackLog(event, "ArbitrationFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryBondChargedIterator is returned from FilterBondCharged and is used to iterate over the raw logs and unpacked data for BondCharged events raised by the AgentRegistry contract.
type AgentRegistryBondChargedIterator struct {
	Event *AgentRegistryBondCharged // Event containing the contract specifics and raw log

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
func (it *AgentRegistryBondChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryBondCharged)
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
		it.Event = new(AgentRegistryBondCharged)
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
func (it *AgentRegistryBondChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryBondChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryBondCharged represents a BondCharged event raised by the AgentRegistry contract.
type AgentRegistryBondCharged struct {
	Wallet common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondCharged is a free log retrieval operation binding the contract event 0x422abfeee0012328b4dd6b71257f61b3dc984217260bf85f4c610558be90ddd9.
//
// Solidity: event BondCharged(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterBondCharged(opts *bind.FilterOpts, wallet []common.Address, to []common.Address) (*AgentRegistryBondChargedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "BondCharged", walletRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryBondChargedIterator{contract: _AgentRegistry.contract, event: "BondCharged", logs: logs, sub: sub}, nil
}

// WatchBondCharged is a free log subscription operation binding the contract event 0x422abfeee0012328b4dd6b71257f61b3dc984217260bf85f4c610558be90ddd9.
//
// Solidity: event BondCharged(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchBondCharged(opts *bind.WatchOpts, sink chan<- *AgentRegistryBondCharged, wallet []common.Address, to []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "BondCharged", walletRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryBondCharged)
				if err := _AgentRegistry.contract.UnpackLog(event, "BondCharged", log); err != nil {
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

// ParseBondCharged is a log parse operation binding the contract event 0x422abfeee0012328b4dd6b71257f61b3dc984217260bf85f4c610558be90ddd9.
//
// Solidity: event BondCharged(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseBondCharged(log types.Log) (*AgentRegistryBondCharged, error) {
	event := new(AgentRegistryBondCharged)
	if err := _AgentRegistry.contract.UnpackLog(event, "BondCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryBondRefundedIterator is returned from FilterBondRefunded and is used to iterate over the raw logs and unpacked data for BondRefunded events raised by the AgentRegistry contract.
type AgentRegistryBondRefundedIterator struct {
	Event *AgentRegistryBondRefunded // Event containing the contract specifics and raw log

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
func (it *AgentRegistryBondRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryBondRefunded)
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
		it.Event = new(AgentRegistryBondRefunded)
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
func (it *AgentRegistryBondRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryBondRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryBondRefunded represents a BondRefunded event raised by the AgentRegistry contract.
type AgentRegistryBondRefunded struct {
	Wallet common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondRefunded is a free log retrieval operation binding the contract event 0x4fac74d6b4ed45cb641c8e730951b01c4dbcae0565cf330e909230492bd07d78.
//
// Solidity: event BondRefunded(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterBondRefunded(opts *bind.FilterOpts, wallet []common.Address, to []common.Address) (*AgentRegistryBondRefundedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "BondRefunded", walletRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryBondRefundedIterator{contract: _AgentRegistry.contract, event: "BondRefunded", logs: logs, sub: sub}, nil
}

// WatchBondRefunded is a free log subscription operation binding the contract event 0x4fac74d6b4ed45cb641c8e730951b01c4dbcae0565cf330e909230492bd07d78.
//
// Solidity: event BondRefunded(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchBondRefunded(opts *bind.WatchOpts, sink chan<- *AgentRegistryBondRefunded, wallet []common.Address, to []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "BondRefunded", walletRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryBondRefunded)
				if err := _AgentRegistry.contract.UnpackLog(event, "BondRefunded", log); err != nil {
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

// ParseBondRefunded is a log parse operation binding the contract event 0x4fac74d6b4ed45cb641c8e730951b01c4dbcae0565cf330e909230492bd07d78.
//
// Solidity: event BondRefunded(address indexed wallet, address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseBondRefunded(log types.Log) (*AgentRegistryBondRefunded, error) {
	event := new(AgentRegistryBondRefunded)
	if err := _AgentRegistry.contract.UnpackLog(event, "BondRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryBondReleasedIterator is returned from FilterBondReleased and is used to iterate over the raw logs and unpacked data for BondReleased events raised by the AgentRegistry contract.
type AgentRegistryBondReleasedIterator struct {
	Event *AgentRegistryBondReleased // Event containing the contract specifics and raw log

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
func (it *AgentRegistryBondReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryBondReleased)
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
		it.Event = new(AgentRegistryBondReleased)
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
func (it *AgentRegistryBondReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryBondReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryBondReleased represents a BondReleased event raised by the AgentRegistry contract.
type AgentRegistryBondReleased struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondReleased is a free log retrieval operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterBondReleased(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryBondReleasedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "BondReleased", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryBondReleasedIterator{contract: _AgentRegistry.contract, event: "BondReleased", logs: logs, sub: sub}, nil
}

// WatchBondReleased is a free log subscription operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchBondReleased(opts *bind.WatchOpts, sink chan<- *AgentRegistryBondReleased, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "BondReleased", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryBondReleased)
				if err := _AgentRegistry.contract.UnpackLog(event, "BondReleased", log); err != nil {
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

// ParseBondReleased is a log parse operation binding the contract event 0x60b8ef4216791426b3d7acfb0b6d11a400872350afd70a3ce5ebf62bea7cb0d4.
//
// Solidity: event BondReleased(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseBondReleased(log types.Log) (*AgentRegistryBondReleased, error) {
	event := new(AgentRegistryBondReleased)
	if err := _AgentRegistry.contract.UnpackLog(event, "BondReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryBondReservedIterator is returned from FilterBondReserved and is used to iterate over the raw logs and unpacked data for BondReserved events raised by the AgentRegistry contract.
type AgentRegistryBondReservedIterator struct {
	Event *AgentRegistryBondReserved // Event containing the contract specifics and raw log

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
func (it *AgentRegistryBondReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryBondReserved)
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
		it.Event = new(AgentRegistryBondReserved)
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
func (it *AgentRegistryBondReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryBondReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryBondReserved represents a BondReserved event raised by the AgentRegistry contract.
type AgentRegistryBondReserved struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondReserved is a free log retrieval operation binding the contract event 0xba521e6ccb16052ac552526db8eeca34c80e9bc3364012d209fd8a16387e6396.
//
// Solidity: event BondReserved(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterBondReserved(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryBondReservedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "BondReserved", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryBondReservedIterator{contract: _AgentRegistry.contract, event: "BondReserved", logs: logs, sub: sub}, nil
}

// WatchBondReserved is a free log subscription operation binding the contract event 0xba521e6ccb16052ac552526db8eeca34c80e9bc3364012d209fd8a16387e6396.
//
// Solidity: event BondReserved(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchBondReserved(opts *bind.WatchOpts, sink chan<- *AgentRegistryBondReserved, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "BondReserved", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryBondReserved)
				if err := _AgentRegistry.contract.UnpackLog(event, "BondReserved", log); err != nil {
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

// ParseBondReserved is a log parse operation binding the contract event 0xba521e6ccb16052ac552526db8eeca34c80e9bc3364012d209fd8a16387e6396.
//
// Solidity: event BondReserved(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseBondReserved(log types.Log) (*AgentRegistryBondReserved, error) {
	event := new(AgentRegistryBondReserved)
	if err := _AgentRegistry.contract.UnpackLog(event, "BondReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryBondSlashedIterator is returned from FilterBondSlashed and is used to iterate over the raw logs and unpacked data for BondSlashed events raised by the AgentRegistry contract.
type AgentRegistryBondSlashedIterator struct {
	Event *AgentRegistryBondSlashed // Event containing the contract specifics and raw log

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
func (it *AgentRegistryBondSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryBondSlashed)
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
		it.Event = new(AgentRegistryBondSlashed)
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
func (it *AgentRegistryBondSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryBondSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryBondSlashed represents a BondSlashed event raised by the AgentRegistry contract.
type AgentRegistryBondSlashed struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondSlashed is a free log retrieval operation binding the contract event 0x9c78aa0eb382c004d853205ff9929a2caed00de0cdd490849c8b84d7f461a39d.
//
// Solidity: event BondSlashed(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterBondSlashed(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryBondSlashedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "BondSlashed", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryBondSlashedIterator{contract: _AgentRegistry.contract, event: "BondSlashed", logs: logs, sub: sub}, nil
}

// WatchBondSlashed is a free log subscription operation binding the contract event 0x9c78aa0eb382c004d853205ff9929a2caed00de0cdd490849c8b84d7f461a39d.
//
// Solidity: event BondSlashed(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchBondSlashed(opts *bind.WatchOpts, sink chan<- *AgentRegistryBondSlashed, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "BondSlashed", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryBondSlashed)
				if err := _AgentRegistry.contract.UnpackLog(event, "BondSlashed", log); err != nil {
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

// ParseBondSlashed is a log parse operation binding the contract event 0x9c78aa0eb382c004d853205ff9929a2caed00de0cdd490849c8b84d7f461a39d.
//
// Solidity: event BondSlashed(address indexed wallet, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseBondSlashed(log types.Log) (*AgentRegistryBondSlashed, error) {
	event := new(AgentRegistryBondSlashed)
	if err := _AgentRegistry.contract.UnpackLog(event, "BondSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryFactoryUpdatedIterator is returned from FilterFactoryUpdated and is used to iterate over the raw logs and unpacked data for FactoryUpdated events raised by the AgentRegistry contract.
type AgentRegistryFactoryUpdatedIterator struct {
	Event *AgentRegistryFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryFactoryUpdated)
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
		it.Event = new(AgentRegistryFactoryUpdated)
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
func (it *AgentRegistryFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryFactoryUpdated represents a FactoryUpdated event raised by the AgentRegistry contract.
type AgentRegistryFactoryUpdated struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFactoryUpdated is a free log retrieval operation binding the contract event 0x24cd1310c8883cbeaf5b805ab13586ce018b79c022827158ff3e8df14d344936.
//
// Solidity: event FactoryUpdated(address indexed factory)
func (_AgentRegistry *AgentRegistryFilterer) FilterFactoryUpdated(opts *bind.FilterOpts, factory []common.Address) (*AgentRegistryFactoryUpdatedIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "FactoryUpdated", factoryRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryFactoryUpdatedIterator{contract: _AgentRegistry.contract, event: "FactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchFactoryUpdated is a free log subscription operation binding the contract event 0x24cd1310c8883cbeaf5b805ab13586ce018b79c022827158ff3e8df14d344936.
//
// Solidity: event FactoryUpdated(address indexed factory)
func (_AgentRegistry *AgentRegistryFilterer) WatchFactoryUpdated(opts *bind.WatchOpts, sink chan<- *AgentRegistryFactoryUpdated, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "FactoryUpdated", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryFactoryUpdated)
				if err := _AgentRegistry.contract.UnpackLog(event, "FactoryUpdated", log); err != nil {
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

// ParseFactoryUpdated is a log parse operation binding the contract event 0x24cd1310c8883cbeaf5b805ab13586ce018b79c022827158ff3e8df14d344936.
//
// Solidity: event FactoryUpdated(address indexed factory)
func (_AgentRegistry *AgentRegistryFilterer) ParseFactoryUpdated(log types.Log) (*AgentRegistryFactoryUpdated, error) {
	event := new(AgentRegistryFactoryUpdated)
	if err := _AgentRegistry.contract.UnpackLog(event, "FactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryFeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the AgentRegistry contract.
type AgentRegistryFeesSweptIterator struct {
	Event *AgentRegistryFeesSwept // Event containing the contract specifics and raw log

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
func (it *AgentRegistryFeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryFeesSwept)
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
		it.Event = new(AgentRegistryFeesSwept)
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
func (it *AgentRegistryFeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryFeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryFeesSwept represents a FeesSwept event raised by the AgentRegistry contract.
type AgentRegistryFeesSwept struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0xb6c81b526e1bd55a3845d8caa6fe0c1ce87b3c9e534eec43f43b3c1849d4e2e8.
//
// Solidity: event FeesSwept(address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) FilterFeesSwept(opts *bind.FilterOpts, to []common.Address) (*AgentRegistryFeesSweptIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "FeesSwept", toRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryFeesSweptIterator{contract: _AgentRegistry.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0xb6c81b526e1bd55a3845d8caa6fe0c1ce87b3c9e534eec43f43b3c1849d4e2e8.
//
// Solidity: event FeesSwept(address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *AgentRegistryFeesSwept, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "FeesSwept", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryFeesSwept)
				if err := _AgentRegistry.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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

// ParseFeesSwept is a log parse operation binding the contract event 0xb6c81b526e1bd55a3845d8caa6fe0c1ce87b3c9e534eec43f43b3c1849d4e2e8.
//
// Solidity: event FeesSwept(address indexed to, uint256 amount)
func (_AgentRegistry *AgentRegistryFilterer) ParseFeesSwept(log types.Log) (*AgentRegistryFeesSwept, error) {
	event := new(AgentRegistryFeesSwept)
	if err := _AgentRegistry.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentRegistry contract.
type AgentRegistryInitializedIterator struct {
	Event *AgentRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *AgentRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryInitialized)
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
		it.Event = new(AgentRegistryInitialized)
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
func (it *AgentRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryInitialized represents a Initialized event raised by the AgentRegistry contract.
type AgentRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentRegistry *AgentRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentRegistryInitializedIterator, error) {

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentRegistryInitializedIterator{contract: _AgentRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentRegistry *AgentRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryInitialized)
				if err := _AgentRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AgentRegistry *AgentRegistryFilterer) ParseInitialized(log types.Log) (*AgentRegistryInitialized, error) {
	event := new(AgentRegistryInitialized)
	if err := _AgentRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryRevenueWalletUpdatedIterator is returned from FilterRevenueWalletUpdated and is used to iterate over the raw logs and unpacked data for RevenueWalletUpdated events raised by the AgentRegistry contract.
type AgentRegistryRevenueWalletUpdatedIterator struct {
	Event *AgentRegistryRevenueWalletUpdated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryRevenueWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryRevenueWalletUpdated)
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
		it.Event = new(AgentRegistryRevenueWalletUpdated)
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
func (it *AgentRegistryRevenueWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryRevenueWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryRevenueWalletUpdated represents a RevenueWalletUpdated event raised by the AgentRegistry contract.
type AgentRegistryRevenueWalletUpdated struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevenueWalletUpdated is a free log retrieval operation binding the contract event 0xacfd38a8c0e371d5f091a49bf6aafad89ed17a551f8b2f152420408a0d4b3e61.
//
// Solidity: event RevenueWalletUpdated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) FilterRevenueWalletUpdated(opts *bind.FilterOpts, wallet []common.Address) (*AgentRegistryRevenueWalletUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "RevenueWalletUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryRevenueWalletUpdatedIterator{contract: _AgentRegistry.contract, event: "RevenueWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchRevenueWalletUpdated is a free log subscription operation binding the contract event 0xacfd38a8c0e371d5f091a49bf6aafad89ed17a551f8b2f152420408a0d4b3e61.
//
// Solidity: event RevenueWalletUpdated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) WatchRevenueWalletUpdated(opts *bind.WatchOpts, sink chan<- *AgentRegistryRevenueWalletUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "RevenueWalletUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryRevenueWalletUpdated)
				if err := _AgentRegistry.contract.UnpackLog(event, "RevenueWalletUpdated", log); err != nil {
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

// ParseRevenueWalletUpdated is a log parse operation binding the contract event 0xacfd38a8c0e371d5f091a49bf6aafad89ed17a551f8b2f152420408a0d4b3e61.
//
// Solidity: event RevenueWalletUpdated(address indexed wallet)
func (_AgentRegistry *AgentRegistryFilterer) ParseRevenueWalletUpdated(log types.Log) (*AgentRegistryRevenueWalletUpdated, error) {
	event := new(AgentRegistryRevenueWalletUpdated)
	if err := _AgentRegistry.contract.UnpackLog(event, "RevenueWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AgentRegistry contract.
type AgentRegistryUpgradedIterator struct {
	Event *AgentRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *AgentRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryUpgraded)
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
		it.Event = new(AgentRegistryUpgraded)
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
func (it *AgentRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryUpgraded represents a Upgraded event raised by the AgentRegistry contract.
type AgentRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AgentRegistry *AgentRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AgentRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryUpgradedIterator{contract: _AgentRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AgentRegistry *AgentRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AgentRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryUpgraded)
				if err := _AgentRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AgentRegistry *AgentRegistryFilterer) ParseUpgraded(log types.Log) (*AgentRegistryUpgraded, error) {
	event := new(AgentRegistryUpgraded)
	if err := _AgentRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
