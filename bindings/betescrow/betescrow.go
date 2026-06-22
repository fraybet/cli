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
	ClaimDeadline   uint64
	ChallengeWindow uint64
	TermsHash       [32]byte
	Visibility      uint8
}

// BetEscrowMetaData contains all meta data concerning the BetEscrow contract.
var BetEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"t\",\"type\":\"tuple\",\"internalType\":\"structBetEscrow.Terms\",\"components\":[{\"name\":\"yesAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbiter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"yesStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimDeadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"challengeWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"termsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"visibility\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"baseFeeBps_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accept\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agreeOutcome\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"arbiter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengeDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimedOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executionFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"noAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revoke\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"status\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"termsHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"visibility\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voidUnclaimed\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yesAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Accepted\",\"inputs\":[{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Challenged\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"challengeDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Funded\",\"inputs\":[{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutcomeAgreed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Revoked\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settled\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WentLive\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadTerms\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcome\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotContested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFunding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRefundable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooEarly\",\"inputs\":[]}]",
	Bin: "0x6101e060405234801562000011575f80fd5b506040516200249f3803806200249f8339810160408190526200003491620002c6565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055835160208501516001600160a01b0391821615911615818015620000795750805b8062000090575060608601516001600160a01b0316155b806200009e57506080860151155b80620000ac575060a0860151155b80620000c3575060e08601516001600160401b0316155b80620000fc575081158015620000d7575080155b8015620000fc575085602001516001600160a01b0316865f01516001600160a01b0316145b156200011b57604051638dc5992160e01b815260040160405180910390fd5b60408601516001600160a01b0316158015906200014157505f8511806200014157505f84115b80156200015557506001600160a01b038316155b156200017457604051638dc5992160e01b815260040160405180910390fd5b6127108511156200019857604051638dc5992160e01b815260040160405180910390fd5b505083515f80546001600160a01b0319166001600160a01b039283161790556020850151600180546040880151841660809081526060890151851660a09081529089015160c09081529089015160e0908152908901516001600160401b03908116610100908152918a01511661012090815290890151610140529097015160ff1661016052610180959095526101a0939093529081166101c0526001600160a81b03199093169216919091179055620003c1565b60405161014081016001600160401b03811182821017156200027c57634e487b7160e01b5f52604160045260245ffd5b60405290565b80516001600160a01b038116811462000299575f80fd5b919050565b80516001600160401b038116811462000299575f80fd5b805160ff8116811462000299575f80fd5b5f805f808486036101a0811215620002dc575f80fd5b61014080821215620002ec575f80fd5b620002f66200024c565b9150620003038762000282565b8252620003136020880162000282565b6020830152620003266040880162000282565b6040830152620003396060880162000282565b60608301526080870151608083015260a087015160a08301526200036060c088016200029e565b60c08301526200037360e088016200029e565b60e0830152610100878101519083015261012062000393818901620002b5565b9083015286015161016087015191955093509150620003b6610180860162000282565b905092959194509250565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c051611f536200054c5f395f818161037c01526119d801525f8181610355015281816107bb01528181610ca801528181610fa001528181611784015281816117f601528181611a1601528181611a9d01528181611af101528181611b4a0152611b9e01525f81816104b5015261197201525f6102c801525f61047e01525f8181610445015261161b01525f818161030101526105e001525f81816102430152818161089f01528181610d9e01528181610fe90152818161181c015261188f01525f81816102a10152818161081701528181610d270152818161106a015281816117aa01526118b001525f81816105250152818161084501528181611017015281816119b601528181611a5a01528181611ace01528181611b270152611b7b01525f818161054c01528181610787015281816109bd01528181610ae101528181610c7401528181610f6c01528181611472015281816117070152611a7c0152611f535ff3fe608060405234801561000f575f80fd5b50600436106101e7575f3560e01c806362ccd3b811610109578063b6549f751161009e578063d2ef73981161006e578063d2ef739814610505578063fb4b66eb1461050d578063fc0c546a14610520578063fe25e00a14610547575f80fd5b8063b6549f75146104a8578063bf5a5940146104b0578063c94e8e1d146104d7578063ca593dad146104ea575f80fd5b8063861a1412116100d9578063861a1412146104405780638f77583914610467578063b311d9fd14610479578063b60d4288146104a0575f80fd5b806362ccd3b8146103f657806367cce5e11461040a57806369523a201461041e5780636dc87dac1461042b575f80fd5b80633ba86c441161017f57806346ef47881161014f57806346ef4788146103b65780634bb278f3146103c95780635c3c10d7146103d157806362897e20146103e4575f80fd5b80633ba86c44146102fc578063404002a61461033c57806340e9903b146103505780634447842514610377575f80fd5b8063200d2ed2116101ba578063200d2ed2146102735780632852b71c1461029457806328ea6dda1461029c57806329adec14146102c3575f80fd5b80630695c46c146101eb5780630f383d1114610213578063138887741461023457806314bc116f1461023e575b5f80fd5b6002546101fe9062010000900460ff1681565b60405190151581526020015b60405180910390f35b60015461022790600160a81b900460ff1681565b60405161020a9190611db0565b61023c61056e565b005b6102657f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161020a565b60015461028790600160a01b900460ff1681565b60405161020a9190611dc4565b61023c610678565b6102657f000000000000000000000000000000000000000000000000000000000000000081565b6102ea7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff909116815260200161020a565b6103237f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff909116815260200161020a565b60015461022790600160b01b900460ff1681565b6102657f000000000000000000000000000000000000000000000000000000000000000081565b61039e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161020a565b61023c6103c4366004611df1565b61096c565b61023c610b66565b60015461039e906001600160a01b031681565b5f5461039e906001600160a01b031681565b600254610227906301000000900460ff1681565b6001546101fe90600160f81b900460ff1681565b6002546101fe9060ff1681565b60025461022790640100000000900460ff1681565b6103237f000000000000000000000000000000000000000000000000000000000000000081565b6002546101fe90610100900460ff1681565b6102657f000000000000000000000000000000000000000000000000000000000000000081565b61023c610c2d565b61023c610e78565b6102657f000000000000000000000000000000000000000000000000000000000000000081565b61023c6104e5366004611e19565b611108565b60015461032390600160b81b900467ffffffffffffffff1681565b61023c6113bc565b61023c61051b366004611df1565b611507565b61039e7f000000000000000000000000000000000000000000000000000000000000000081565b61039e7f000000000000000000000000000000000000000000000000000000000000000081565b6105766116b3565b5f600154600160a01b900460ff16600581111561059557610595611d88565b141580156105c0575060018054600160a01b900460ff1660058111156105bd576105bd611d88565b14155b156105de57604051631ba168fb60e11b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff164210156106295760405163085de62560e01b815260040160405180910390fd5b60018054600360b01b60ff60b01b198216811783556005929162ff00ff60a01b191617600160a01b835b02179055506106606116ce565b61067660015f80516020611efe83398151915255565b565b6106806116b3565b5f600154600160a01b900460ff16600581111561069f5761069f611d88565b146106bc576040516282354d60e81b815260040160405180910390fd5b5f546001546001600160a01b03918216159116158115826106db575080155b156106f957604051631bb5f5b360e31b815260040160405180910390fd5b5f8261070f575f546001600160a01b031661071c565b6001546001600160a01b03165b90506001600160a01b03811633036107475760405163721c7c6760e11b815260040160405180910390fd5b8261075f57600154600160f81b900460ff1615610767565b60025460ff16155b15610784576040516282354d60e81b815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107b9575f6107db565b7f00000000000000000000000000000000000000000000000000000000000000005b90508315610872575f80546001600160a01b03191633908117909155600180546001600160f81b0316600160f81b17905561086d903061083b847f0000000000000000000000000000000000000000000000000000000000000000611e4d565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016929190611c0f565b6108c3565b600180546001600160a01b0319163390811782556002805460ff19169092179091556108c3903061083b847f0000000000000000000000000000000000000000000000000000000000000000611e4d565b60405133907f2e7bcc1a1a1c1bea0ff9a677c29b43662d6560afbebd99c213389fc15a6279d2905f90a260405133907fb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c905f90a26001805460ff60a01b1916600160a01b1790556040517f12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb905f90a15050505061067660015f80516020611efe83398151915255565b6109746116b3565b6003600154600160a01b900460ff16600581111561099457610994611d88565b146109b257604051631ba026e160e11b815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109fb5760405163665b32d360e11b815260040160405180910390fd5b6001826003811115610a0f57610a0f611d88565b14158015610a2f57506002826003811115610a2c57610a2c611d88565b14155b8015610a4d57506003826003811115610a4a57610a4a611d88565b14155b15610a6b5760405163c74a206d60e01b815260040160405180910390fd5b6001805483919060ff60b01b1916600160b01b836003811115610a9057610a90611d88565b02179055506003826003811115610aa957610aa9611d88565b14610ab5576004610ab8565b60055b6001805460ff60a01b1916600160a01b836005811115610ada57610ada611d88565b02179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167fa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd4411318383835f604051610b3c93929190611e60565b60405180910390a2610b4c6116ce565b610b6260015f80516020611efe83398151915255565b5050565b610b6e6116b3565b6002600154600160a01b900460ff166005811115610b8e57610b8e611d88565b14610bac57604051635b95129160e11b815260040160405180910390fd5b600154600160b81b900467ffffffffffffffff16421015610be05760405163fa7bc54760e01b815260040160405180910390fd5b6001805460ff600160a81b820416919060ff60b01b1916600160b01b836003811115610c0e57610c0e611d88565b0217905550600180546004919060ff60a01b1916600160a01b83610653565b610c356116b3565b5f600154600160a01b900460ff166005811115610c5457610c54611d88565b14610c71576040516282354d60e81b815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610ca6575f610cc8565b7f00000000000000000000000000000000000000000000000000000000000000005b5f549091506001600160a01b03163303610d5057600154600160f81b900460ff1615610d0757604051635adf638760e01b815260040160405180910390fd5b600180546001600160f81b0316600160f81b179055610d4b333061083b847f0000000000000000000000000000000000000000000000000000000000000000611e4d565b610ddb565b6001546001600160a01b03163303610dc25760025460ff1615610d8657604051635adf638760e01b815260040160405180910390fd5b6002805460ff19166001179055610d4b333061083b847f0000000000000000000000000000000000000000000000000000000000000000611e4d565b60405163721c7c6760e11b815260040160405180910390fd5b60405133907fb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c905f90a2600154600160f81b900460ff168015610e20575060025460ff165b15610e61576001805460ff60a01b1916600160a01b1790556040517f12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb905f90a15b5061067660015f80516020611efe83398151915255565b610e806116b3565b5f600154600160a01b900460ff166005811115610e9f57610e9f611d88565b14610ebc576040516282354d60e81b815260040160405180910390fd5b5f546001546001600160a01b0391821615911615811582610edb575080155b15610ef957604051631bb5f5b360e31b815260040160405180910390fd5b5f82610f0f575f546001600160a01b0316610f1c565b6001546001600160a01b03165b9050336001600160a01b03821614610f475760405163721c7c6760e11b815260040160405180910390fd5b6001805460ff60a01b1916600560a01b1790556002805461ff0019166101001790555f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610f9e575f610fc0565b7f00000000000000000000000000000000000000000000000000000000000000005b905083156110435760025460ff161561103e5760015461103e906001600160a01b031661100d837f0000000000000000000000000000000000000000000000000000000000000000611e4d565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190611c50565b61108e565b600154600160f81b900460ff161561108e575f5461108e906001600160a01b031661100d837f0000000000000000000000000000000000000000000000000000000000000000611e4d565b6040517f44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6905f90a17f56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f52160036040516110e69190611db0565b60405180910390a15050505061067660015f80516020611efe83398151915255565b6111106116b3565b60018054600160a01b900460ff16600581111561112f5761112f611d88565b1415801561115b57506002600154600160a01b900460ff16600581111561115857611158611d88565b14155b156111795760405163baf13b3f60e01b815260040160405180910390fd5b600181600381111561118d5761118d611d88565b141580156111ad575060028160038111156111aa576111aa611d88565b14155b80156111cb575060038160038111156111c8576111c8611d88565b14155b156111e95760405163c74a206d60e01b815260040160405180910390fd5b5f546001600160a01b03163303611229576002805482919063ff0000001916630100000083600381111561121f5761121f611d88565b0217905550611262565b6001546001600160a01b03163303610dc2576002805482919064ff00000000191664010000000083600381111561121f5761121f611d88565b336001600160a01b03167f79a012f7d2faaa3daa2d4a00d83acc1a9334b76556f42bd053d2d45bf5a5b6a88260405161129b9190611db0565b60405180910390a25f6002546301000000900460ff1660038111156112c2576112c2611d88565b1415801561130a5750600254640100000000900460ff1660038111156112ea576112ea611d88565b6002546301000000900460ff16600381111561130857611308611d88565b145b156113a35760025460018054630100000090920460ff169160ff60b01b1916600160b01b83600381111561134057611340611d88565b02179055506003600154600160b01b900460ff16600381111561136557611365611d88565b14611371576004611374565b60055b6001805460ff60a01b1916600160a01b83600581111561139657611396611d88565b02179055506113a36116ce565b6113b960015f80516020611efe83398151915255565b50565b6002600154600160a01b900460ff1660058111156113dc576113dc611d88565b146113fa57604051635b95129160e11b815260040160405180910390fd5b600154600160b81b900467ffffffffffffffff16421061142d57604051631588dee160e11b815260040160405180910390fd5b5f546001600160a01b0316331480159061145257506001546001600160a01b03163314155b156114705760405163721c7c6760e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166114b7576040516304e2778b60e31b815260040160405180910390fd5b6001805460ff60a01b1916600360a01b1790556002805462ff000019166201000017905560405133907f98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c905f90a2565b60018054600160a01b900460ff16600581111561152657611526611d88565b146115445760405163baf13b3f60e01b815260040160405180910390fd5b600182600381111561155857611558611d88565b141580156115785750600282600381111561157557611575611d88565b14155b156115965760405163c74a206d60e01b815260040160405180910390fd5b5f546001600160a01b031633148015906115bb57506001546001600160a01b03163314155b156115d95760405163721c7c6760e11b815260040160405180910390fd5b6001805483919060ff60a81b1916600160a81b8360038111156115fe576115fe611d88565b02179055506001805460ff60a01b1916600160a11b1790556116407f000000000000000000000000000000000000000000000000000000000000000042611e8c565b6001805467ffffffffffffffff60b81b1916600160b81b67ffffffffffffffff9384168102919091179182905560405133937fa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd44113183936116a79388938893919092041690611e60565b60405180910390a25050565b6116bb611c8a565b60025f80516020611efe83398151915255565b600254610100900460ff16156116f65760405162560ff960e81b815260040160405180910390fd5b6002805461ff0019166101001790557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615156003600154600160b01b900460ff16600381111561175157611751611d88565b0361188957600154600160f81b900460ff16156117ce575f546117ce906001600160a01b031682611782575f6117a4565b7f00000000000000000000000000000000000000000000000000000000000000005b61100d907f0000000000000000000000000000000000000000000000000000000000000000611e4d565b60025460ff161561184057600154611840906001600160a01b0316826117f4575f611816565b7f00000000000000000000000000000000000000000000000000000000000000005b61100d907f0000000000000000000000000000000000000000000000000000000000000000611e4d565b7f56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f521600160169054906101000a900460ff1660405161187e9190611db0565b60405180910390a150565b5f6118d47f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611e4d565b90505f60018054600160b01b900460ff1660038111156118f6576118f6611d88565b1461190c576001546001600160a01b0316611918565b5f546001600160a01b03165b90505f60018054600160b01b900460ff16600381111561193a5761193a611d88565b1461194f575f546001600160a01b031661195c565b6001546001600160a01b03165b90505f8461196a575f6119a1565b6127106119977f000000000000000000000000000000000000000000000000000000000000000086611eb4565b6119a19190611ecb565b905080156119fd576119fd6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000083611c50565b611a0b8361100d8387611eea565b848015611a3757505f7f0000000000000000000000000000000000000000000000000000000000000000115b15611bc25760025462010000900460ff1615611b1a57611ac16001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611c50565b611b156001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016847f0000000000000000000000000000000000000000000000000000000000000000611c50565b611bc2565b611b6e6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016847f0000000000000000000000000000000000000000000000000000000000000000611c50565b611bc26001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016837f0000000000000000000000000000000000000000000000000000000000000000611c50565b7f56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f521600160169054906101000a900460ff16604051611c009190611db0565b60405180910390a15050505050565b611c1d848484846001611cb9565b611c4a57604051635274afe760e01b81526001600160a01b03851660048201526024015b60405180910390fd5b50505050565b611c5d8383836001611d26565b611c8557604051635274afe760e01b81526001600160a01b0384166004820152602401611c41565b505050565b5f80516020611efe8339815191525460020361067657604051633ee5aeb560e01b815260040160405180910390fd5b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f51148316611d15578383151615611d09573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316611d7c578383151615611d70573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b634e487b7160e01b5f52602160045260245ffd5b60048110611dac57611dac611d88565b9052565b60208101611dbe8284611d9c565b92915050565b6020810160068310611dd857611dd8611d88565b91905290565b803560048110611dec575f80fd5b919050565b5f8060408385031215611e02575f80fd5b611e0b83611dde565b946020939093013593505050565b5f60208284031215611e29575f80fd5b611e3282611dde565b9392505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115611dbe57611dbe611e39565b60608101611e6e8286611d9c565b83602083015267ffffffffffffffff83166040830152949350505050565b67ffffffffffffffff818116838216019080821115611ead57611ead611e39565b5092915050565b8082028115828204841417611dbe57611dbe611e39565b5f82611ee557634e487b7160e01b5f52601260045260245ffd5b500490565b81810381811115611dbe57611dbe611e3956fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220ab75888a607874031cfd4505e232b34620751d4a68fe62b73fabd38ce6fa1d2164736f6c63430008180033",
}

// BetEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use BetEscrowMetaData.ABI instead.
var BetEscrowABI = BetEscrowMetaData.ABI

// BetEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BetEscrowMetaData.Bin instead.
var BetEscrowBin = BetEscrowMetaData.Bin

// DeployBetEscrow deploys a new Ethereum contract, binding an instance of BetEscrow to it.
func DeployBetEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, t BetEscrowTerms, baseFeeBps_ *big.Int, executionFee_ *big.Int, revenueWallet_ common.Address) (common.Address, *types.Transaction, *BetEscrow, error) {
	parsed, err := BetEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BetEscrowBin), backend, t, baseFeeBps_, executionFee_, revenueWallet_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BetEscrow{BetEscrowCaller: BetEscrowCaller{contract: contract}, BetEscrowTransactor: BetEscrowTransactor{contract: contract}, BetEscrowFilterer: BetEscrowFilterer{contract: contract}}, nil
}

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

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrow *BetEscrowCaller) ExecutionFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BetEscrow.contract.Call(opts, &out, "executionFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrow *BetEscrowSession) ExecutionFee() (*big.Int, error) {
	return _BetEscrow.Contract.ExecutionFee(&_BetEscrow.CallOpts)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x40e9903b.
//
// Solidity: function executionFee() view returns(uint256)
func (_BetEscrow *BetEscrowCallerSession) ExecutionFee() (*big.Int, error) {
	return _BetEscrow.Contract.ExecutionFee(&_BetEscrow.CallOpts)
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

// Resolve is a paid mutator transaction binding the contract method 0x46ef4788.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowTransactor) Resolve(opts *bind.TransactOpts, outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.contract.Transact(opts, "resolve", outcome, evidenceHash)
}

// Resolve is a paid mutator transaction binding the contract method 0x46ef4788.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowSession) Resolve(outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.Contract.Resolve(&_BetEscrow.TransactOpts, outcome, evidenceHash)
}

// Resolve is a paid mutator transaction binding the contract method 0x46ef4788.
//
// Solidity: function resolve(uint8 outcome, bytes32 evidenceHash) returns()
func (_BetEscrow *BetEscrowTransactorSession) Resolve(outcome uint8, evidenceHash [32]byte) (*types.Transaction, error) {
	return _BetEscrow.Contract.Resolve(&_BetEscrow.TransactOpts, outcome, evidenceHash)
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
