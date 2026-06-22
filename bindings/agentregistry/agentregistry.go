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

// AgentRegistryAgentProfile is an auto generated low-level Go binding around an user-defined struct.
type AgentRegistryAgentProfile struct {
	Owner        common.Address
	Wallet       common.Address
	Signer       common.Address
	PolicyHash   [32]byte
	MetadataHash [32]byte
	Bond         *big.Int
	CreatedAt    uint64
	Active       bool
}

// AgentRegistryMetaData contains all meta data concerning the AgentRegistry contract.
var AgentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"bondToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrationFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registrationBond_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accruedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"agent\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAgentRegistry.AgentProfile\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deactivate\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metadataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registrationBond\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRevenueWallet\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signerOf\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweepFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePolicy\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSigner\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminTransferred\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentDeactivated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentPolicyUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"policyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRegistered\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentSignerUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondRefunded\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondSlashed\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesSwept\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevenueWalletUpdated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60e060405234801562000010575f80fd5b506040516200125238038062001252833981016040819052620000339162000105565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00556001600160a01b03851615806200007557506001600160a01b038216155b806200008857506001600160a01b038116155b15620000a75760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0394851660805260a09390935260c0919091525f80549184166001600160a01b0319928316179055600180549290931691161790556200015f565b80516001600160a01b038116811462000100575f80fd5b919050565b5f805f805f60a086880312156200011a575f80fd5b6200012586620000e9565b945060208601519350604086015192506200014360608701620000e9565b91506200015360808701620000e9565b90509295509295909350565b60805160a05160c051611089620001c95f395f818161014901528181610a0101528181610b020152610c8901525f818161010f01528181610a2201528181610a850152610caa01525f81816102ed0152818161043e015281816108fb0152610a5b01526110895ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c806392e423b51161009e578063c96be4cb1161006e578063c96be4cb1461030f578063d113b95c14610322578063e8527e381461032a578063f851a4401461033d578063fb235f341461034f575f80fd5b806392e423b51461020857806397cc56fa146102905780639f8a13d7146102a3578063c28f4392146102e8575f80fd5b80635b5175d5116100d95780635b5175d5146101ab578063682c2058146101d957806375829def146101e25780638fd016e8146101f5575f80fd5b806314c44e091461010a578063373b9dbd146101445780633ea053eb1461016b5780634447842514610180575b5f80fd5b6101317f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b6101317f000000000000000000000000000000000000000000000000000000000000000081565b61017e610179366004610f5c565b610362565b005b600154610193906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b6101936101b9366004610f5c565b6001600160a01b039081165f908152600360205260409020600201541690565b61013160025481565b61017e6101f0366004610f5c565b6104c9565b61017e610203366004610f7c565b610561565b61021b610216366004610f5c565b610658565b6040805182516001600160a01b039081168252602080850151821690830152838301511691810191909152606080830151908201526080808301519082015260a0808301519082015260c08083015167ffffffffffffffff169082015260e0918201511515918101919091526101000161013b565b61017e61029e366004610fad565b610748565b6102d86102b1366004610f5c565b6001600160a01b03165f90815260036020526040902060060154600160401b900460ff1690565b604051901515815260200161013b565b6101937f000000000000000000000000000000000000000000000000000000000000000081565b61017e61031d366004610f5c565b6107fb565b61017e6108d4565b61017e610338366004610fd5565b610980565b5f54610193906001600160a01b031681565b61017e61035d366004610f5c565b610d19565b61036a610db3565b6001600160a01b038082165f90815260036020526040902080548392166103a45760405163aba4733960e01b815260040160405180910390fd5b80546001600160a01b031633146103ce576040516330cd747160e01b815260040160405180910390fd5b6001600160a01b0383165f8181526003602052604080822060068101805460ff60401b19169055600581018054908490559151909391927fa58c733780f5c1527d6106ff1e731cbf2e9ffc0eb11ad399ca33ea55f057ddf491a280156104ac578154610467906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116911683610dce565b81546040518281526001600160a01b03918216918716907f4fac74d6b4ed45cb641c8e730951b01c4dbcae0565cf330e909230492bd07d789060200160405180910390a35b505050506104c660015f8051602061103483398151915255565b50565b5f546001600160a01b031633146104f357604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b03811661051a5760405163d92e233d60e01b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b038316908117825560405190917fe22b4f506b7da9a528a95d7063a6cde8d2b6268773b1f7f65c00057531704a3d91a250565b6001600160a01b038083165f908152600360205260409020805484921661059b5760405163aba4733960e01b815260040160405180910390fd5b80546001600160a01b031633146105c5576040516330cd747160e01b815260040160405180910390fd5b6001600160a01b0383166105ec5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038481165f8181526003602090815260409182902060020180546001600160a01b0319169488169485179055905192835290917f15c539f3c72a7122b9be60c5a7d9929c4882896eca5158969a1877c067192d9b91015b60405180910390a250505050565b60408051610100810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091526001600160a01b038281165f9081526003602081815260409283902083516101008101855281548616808252600183015487169382019390935260028201549095169385019390935290820154606084015260048201546080840152600582015460a084015260069091015467ffffffffffffffff811660c0840152600160401b900460ff16151560e08301526107425760405163aba4733960e01b815260040160405180910390fd5b92915050565b6001600160a01b038083165f90815260036020526040902080548492166107825760405163aba4733960e01b815260040160405180910390fd5b80546001600160a01b031633146107ac576040516330cd747160e01b815260040160405180910390fd5b6001600160a01b0384165f8181526003602081905260409182902001859055517fcc8946cc5e90c32269cfbfe3ca38dea5bf8e331038589b2f3ede2c35c42fffa39061064a9086815260200190565b5f546001600160a01b0316331461082557604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b038082165f908152600360205260409020805490911661085f5760405163aba4733960e01b815260040160405180910390fd5b6005810180545f9182905560068301805460ff60401b19169055600280549192839261088c908490611014565b90915550506040518181526001600160a01b038416907f9c78aa0eb382c004d853205ff9929a2caed00de0cdd490849c8b84d7f461a39d9060200160405180910390a2505050565b6108dc610db3565b600280545f909155801561096757600154610924906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116911683610dce565b6001546040518281526001600160a01b03909116907fb6c81b526e1bd55a3845d8caa6fe0c1ce87b3c9e534eec43f43b3c1849d4e2e89060200160405180910390a25b5061097e60015f8051602061103483398151915255565b565b610988610db3565b6001600160a01b03841615806109a557506001600160a01b038316155b156109c35760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038481165f9081526003602052604090205416156109fb57604051630ea075bf60e21b815260040160405180910390fd5b5f610a467f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611014565b90508015610a8357610a836001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084610e0d565b7f000000000000000000000000000000000000000000000000000000000000000060025f828254610ab49190611014565b92505081905550604051806101000160405280336001600160a01b03168152602001866001600160a01b03168152602001856001600160a01b031681526020018481526020018381526020017f000000000000000000000000000000000000000000000000000000000000000081526020014267ffffffffffffffff1681526020016001151581525060035f876001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506040820151816002015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160060160086101000a81548160ff021916908315150217905550905050336001600160a01b0316856001600160a01b03167fcc059d6320dd91f6d72ca6e83caaf137242df66d024a6e14412a0472557dba63867f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051610cf4939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a350610d1360015f8051602061103483398151915255565b50505050565b5f546001600160a01b03163314610d4357604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b038116610d6a5760405163d92e233d60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040517facfd38a8c0e371d5f091a49bf6aafad89ed17a551f8b2f152420408a0d4b3e61905f90a250565b610dbb610e43565b60025f8051602061103483398151915255565b610ddb8383836001610e72565b610e0857604051635274afe760e01b81526001600160a01b03841660048201526024015b60405180910390fd5b505050565b610e1b848484846001610ed4565b610d1357604051635274afe760e01b81526001600160a01b0385166004820152602401610dff565b5f805160206110348339815191525460020361097e57604051633ee5aeb560e01b815260040160405180910390fd5b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316610ec8578383151615610ebc573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f51148316610f30578383151615610f24573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b80356001600160a01b0381168114610f57575f80fd5b919050565b5f60208284031215610f6c575f80fd5b610f7582610f41565b9392505050565b5f8060408385031215610f8d575f80fd5b610f9683610f41565b9150610fa460208401610f41565b90509250929050565b5f8060408385031215610fbe575f80fd5b610fc783610f41565b946020939093013593505050565b5f805f8060808587031215610fe8575f80fd5b610ff185610f41565b9350610fff60208601610f41565b93969395505050506040820135916060013590565b8082018082111561074257634e487b7160e01b5f52601160045260245ffdfe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122074cf9454198a8707acc79b3272c85d7be3bf7d0b59ec86f5a454184a7474ade964736f6c63430008180033",
}

// AgentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentRegistryMetaData.ABI instead.
var AgentRegistryABI = AgentRegistryMetaData.ABI

// AgentRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentRegistryMetaData.Bin instead.
var AgentRegistryBin = AgentRegistryMetaData.Bin

// DeployAgentRegistry deploys a new Ethereum contract, binding an instance of AgentRegistry to it.
func DeployAgentRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, bondToken_ common.Address, registrationFee_ *big.Int, registrationBond_ *big.Int, admin_ common.Address, revenueWallet_ common.Address) (common.Address, *types.Transaction, *AgentRegistry, error) {
	parsed, err := AgentRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentRegistryBin), backend, bondToken_, registrationFee_, registrationBond_, admin_, revenueWallet_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentRegistry{AgentRegistryCaller: AgentRegistryCaller{contract: contract}, AgentRegistryTransactor: AgentRegistryTransactor{contract: contract}, AgentRegistryFilterer: AgentRegistryFilterer{contract: contract}}, nil
}

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

// AccruedFees is a free data retrieval call binding the contract method 0x682c2058.
//
// Solidity: function accruedFees() view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) AccruedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "accruedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFees is a free data retrieval call binding the contract method 0x682c2058.
//
// Solidity: function accruedFees() view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) AccruedFees() (*big.Int, error) {
	return _AgentRegistry.Contract.AccruedFees(&_AgentRegistry.CallOpts)
}

// AccruedFees is a free data retrieval call binding the contract method 0x682c2058.
//
// Solidity: function accruedFees() view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) AccruedFees() (*big.Int, error) {
	return _AgentRegistry.Contract.AccruedFees(&_AgentRegistry.CallOpts)
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
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistryCaller) Agent(opts *bind.CallOpts, wallet common.Address) (AgentRegistryAgentProfile, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "agent", wallet)

	if err != nil {
		return *new(AgentRegistryAgentProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentRegistryAgentProfile)).(*AgentRegistryAgentProfile)

	return out0, err

}

// Agent is a free data retrieval call binding the contract method 0x92e423b5.
//
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistrySession) Agent(wallet common.Address) (AgentRegistryAgentProfile, error) {
	return _AgentRegistry.Contract.Agent(&_AgentRegistry.CallOpts, wallet)
}

// Agent is a free data retrieval call binding the contract method 0x92e423b5.
//
// Solidity: function agent(address wallet) view returns((address,address,address,bytes32,bytes32,uint256,uint64,bool))
func (_AgentRegistry *AgentRegistryCallerSession) Agent(wallet common.Address) (AgentRegistryAgentProfile, error) {
	return _AgentRegistry.Contract.Agent(&_AgentRegistry.CallOpts, wallet)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AgentRegistry *AgentRegistryCaller) BondToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "bondToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AgentRegistry *AgentRegistrySession) BondToken() (common.Address, error) {
	return _AgentRegistry.Contract.BondToken(&_AgentRegistry.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) BondToken() (common.Address, error) {
	return _AgentRegistry.Contract.BondToken(&_AgentRegistry.CallOpts)
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
