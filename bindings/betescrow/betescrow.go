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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"t\",\"type\":\"tuple\",\"internalType\":\"structBetEscrow.Terms\",\"components\":[{\"name\":\"yesAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"noAgent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbiter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"yesStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimDeadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"challengeWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"termsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"visibility\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"baseFeeBps_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revenueWallet_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"agreeOutcome\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"arbiter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengeDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimedOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executionFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalOutcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"noAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revenueWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"settled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"status\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"termsHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"visibility\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voidUnclaimed\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yesAgent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesAgentAgreed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBetEscrow.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesFunded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Challenged\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"challengeDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Funded\",\"inputs\":[{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutcomeAgreed\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settled\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumBetEscrow.Outcome\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WentLive\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadTerms\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeWindowOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcome\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotArbiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotContested\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFunding\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRefundable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TooEarly\",\"inputs\":[]}]",
	Bin: "0x61022060405234801562000011575f80fd5b506040516200207738038062002077833981016040819052620000349162000298565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005583516001600160a01b031615806200007b575060208401516001600160a01b0316155b8062000092575060608401516001600160a01b0316155b80620000b6575083602001516001600160a01b0316845f01516001600160a01b0316145b80620000c457506080840151155b80620000d2575060a0840151155b80620000e9575060e08401516001600160401b0316155b156200010857604051638dc5992160e01b815260040160405180910390fd5b60408401516001600160a01b0316158015906200012e57505f8311806200012e57505f82115b80156200014257506001600160a01b038116155b156200016157604051638dc5992160e01b815260040160405180910390fd5b6127108311156200018557604051638dc5992160e01b815260040160405180910390fd5b83516001600160a01b0390811660809081526020860151821660a09081526040870151831660c09081526060880151841660e09081529288015161010090815291880151610120908152908801516001600160401b03908116610140529288015190921661016052860151610180529094015160ff166101a0526101c0929092526101e05216610200525f805460ff1916905562000393565b60405161014081016001600160401b03811182821017156200024e57634e487b7160e01b5f52604160045260245ffd5b60405290565b80516001600160a01b03811681146200026b575f80fd5b919050565b80516001600160401b03811681146200026b575f80fd5b805160ff811681146200026b575f80fd5b5f805f808486036101a0811215620002ae575f80fd5b61014080821215620002be575f80fd5b620002c86200021e565b9150620002d58762000254565b8252620002e56020880162000254565b6020830152620002f86040880162000254565b60408301526200030b6060880162000254565b60608301526080870151608083015260a087015160a08301526200033260c0880162000270565b60c08301526200034560e0880162000270565b60e083015261010087810151908301526101206200036581890162000287565b908301528601516101608701519195509350915062000388610180860162000254565b905092959194509250565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516101e05161020051611b1b6200055c5f395f818161035101526115a101525f818161032a01528181610965015281816112af0152818161136c015281816115df01528181611666015281816116ba01528181611713015261176701525f81816104ae015261153b01525f61029f01525f61047f01525f8181610445015261113101525f81816102d801526105c901525f818161022a01528181610ad001528181611392015261140401525f8181610278015281816109fe015281816112d5015261142501525f818161051d01528181610a2c015281816113030152818161157f0152818161162301528181611697015281816116f0015261174401525f8181610544015281816106a1015281816107bb0152818161093101528181610f5c01528181611220015261164501525f81816103ab01528181610a6301528181610ce101528181610f19015281816110ba0152818161134001528181611471015261150501525f81816103d20152818161099101528181610c8301528181610ee501528181611086015281816112830152818161149701526114df0152611b1b5ff3fe608060405234801561000f575f80fd5b50600436106101d1575f3560e01c806362ccd3b8116100fe578063b60d42881161009e578063d2ef73981161006e578063d2ef7398146104fd578063fb4b66eb14610505578063fc0c546a14610518578063fe25e00a1461053f575f80fd5b8063b60d4288146104a1578063bf5a5940146104a9578063c94e8e1d146104d0578063ca593dad146104e3575f80fd5b80636dc87dac116100d95780636dc87dac1461042d578063861a1412146104405780638f77583914610467578063b311d9fd1461047a575f80fd5b806362ccd3b8146103f457806367cce5e11461040757806369523a201461041a575f80fd5b80633ba86c441161017457806346ef47881161014457806346ef47881461038b5780634bb278f31461039e5780635c3c10d7146103a657806362897e20146103cd575f80fd5b80633ba86c44146102d3578063404002a61461031357806340e9903b14610325578063444784251461034c575f80fd5b806314bc116f116101af57806314bc116f14610225578063200d2ed21461025a57806328ea6dda1461027357806329adec141461029a575f80fd5b80630695c46c146101d55780630f383d11146101fd578063138887741461021b575b5f80fd5b5f546101e890600160701b900460ff1681565b60405190151581526020015b60405180910390f35b5f5461020e90610100900460ff1681565b6040516101f49190611978565b610223610566565b005b61024c7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101f4565b5f546102669060ff1681565b6040516101f4919061198c565b61024c7f000000000000000000000000000000000000000000000000000000000000000081565b6102c17f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101f4565b6102fa7f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020016101f4565b5f5461020e9062010000900460ff1681565b61024c7f000000000000000000000000000000000000000000000000000000000000000081565b6103737f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f4565b6102236103993660046119b9565b610658565b610223610840565b6103737f000000000000000000000000000000000000000000000000000000000000000081565b6103737f000000000000000000000000000000000000000000000000000000000000000081565b5f5461020e90600160781b900460ff1681565b5f546101e890600160581b900460ff1681565b5f546101e890600160601b900460ff1681565b5f5461020e90600160801b900460ff1681565b6102fa7f000000000000000000000000000000000000000000000000000000000000000081565b5f546101e890600160681b900460ff1681565b61024c7f000000000000000000000000000000000000000000000000000000000000000081565b6102236108f2565b61024c7f000000000000000000000000000000000000000000000000000000000000000081565b6102236104de3660046119e1565b610ba7565b5f546102fa906301000000900467ffffffffffffffff1681565b610223610e72565b6102236105133660046119b9565b610ff3565b6103737f000000000000000000000000000000000000000000000000000000000000000081565b6103737f000000000000000000000000000000000000000000000000000000000000000081565b61056e6111c8565b5f805460ff16600581111561058557610585611950565b141580156105a9575060015f5460ff1660058111156105a6576105a6611950565b14155b156105c757604051631ba168fb60e11b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff164210156106125760405163085de62560e01b815260040160405180910390fd5b5f80546203000062ff0000198216811783556005929162ff00ff1916176001835b02179055506106406111e3565b61065660015f80516020611ac683398151915255565b565b6106606111c8565b60035f5460ff16600581111561067857610678611950565b1461069657604051631ba026e160e11b815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106df5760405163665b32d360e11b815260040160405180910390fd5b60018260038111156106f3576106f3611950565b141580156107135750600282600381111561071057610710611950565b14155b80156107315750600382600381111561072e5761072e611950565b14155b1561074f5760405163c74a206d60e01b815260040160405180910390fd5b5f805483919062ff000019166201000083600381111561077157610771611950565b0217905550600382600381111561078a5761078a611950565b14610796576004610799565b60055b5f805460ff191660018360058111156107b4576107b4611950565b02179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167fa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd4411318383835f60405161081693929190611a01565b60405180910390a26108266111e3565b61083c60015f80516020611ac683398151915255565b5050565b6108486111c8565b60025f5460ff16600581111561086057610860611950565b1461087e57604051635b95129160e11b815260040160405180910390fd5b5f546301000000900467ffffffffffffffff164210156108b15760405163fa7bc54760e01b815260040160405180910390fd5b5f805460ff610100820416919062ff00001916620100008360038111156108da576108da611950565b02179055505f80546004919060ff1916600183610633565b6108fa6111c8565b5f805460ff16600581111561091157610911611950565b1461092e576040516282354d60e81b815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610963575f610985565b7f00000000000000000000000000000000000000000000000000000000000000005b90506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163303610a59575f54600160581b900460ff16156109e157604051635adf638760e01b815260040160405180910390fd5b5f805460ff60581b1916600160581b179055610a543330610a22847f0000000000000000000000000000000000000000000000000000000000000000611a41565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169291906117d7565b610b0d565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163303610af4575f54600160601b900460ff1615610ab357604051635adf638760e01b815260040160405180910390fd5b5f805460ff60601b1916600160601b179055610a543330610a22847f0000000000000000000000000000000000000000000000000000000000000000611a41565b60405163721c7c6760e11b815260040160405180910390fd5b60405133907fb436c2bf863ccd7b8f63171201efd4792066b4ce8e543dde9c3e9e9ab98e216c905f90a25f54600160581b900460ff168015610b5757505f54600160601b900460ff165b15610b90575f805460ff191660011781556040517f12a54b87db8a631dcf48764fce631da6bc92767389cdec321c4dd8af987726bb9190a15b5061065660015f80516020611ac683398151915255565b610baf6111c8565b60015f5460ff166005811115610bc757610bc7611950565b14158015610beb575060025f5460ff166005811115610be857610be8611950565b14155b15610c095760405163baf13b3f60e01b815260040160405180910390fd5b6001816003811115610c1d57610c1d611950565b14158015610c3d57506002816003811115610c3a57610c3a611950565b14155b8015610c5b57506003816003811115610c5857610c58611950565b14155b15610c795760405163c74a206d60e01b815260040160405180910390fd5b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163303610cd7575f805482919060ff60781b1916600160781b836003811115610ccd57610ccd611950565b0217905550610d2b565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163303610af4575f805482919060ff60801b1916600160801b836003811115610ccd57610ccd611950565b336001600160a01b03167f79a012f7d2faaa3daa2d4a00d83acc1a9334b76556f42bd053d2d45bf5a5b6a882604051610d649190611978565b60405180910390a25f8054600160781b900460ff166003811115610d8a57610d8a611950565b14158015610dcf57505f54600160801b900460ff166003811115610db057610db0611950565b5f54600160781b900460ff166003811115610dcd57610dcd611950565b145b15610e59575f805460ff600160781b820416919062ff0000191662010000836003811115610dff57610dff611950565b021790555060035f5462010000900460ff166003811115610e2257610e22611950565b14610e2e576004610e31565b60055b5f805460ff19166001836005811115610e4c57610e4c611950565b0217905550610e596111e3565b610e6f60015f80516020611ac683398151915255565b50565b60025f5460ff166005811115610e8a57610e8a611950565b14610ea857604051635b95129160e11b815260040160405180910390fd5b5f546301000000900467ffffffffffffffff164210610eda57604051631588dee160e11b815260040160405180910390fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801590610f3c5750336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614155b15610f5a5760405163721c7c6760e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610fa1576040516304e2778b60e31b815260040160405180910390fd5b5f80546eff00000000000000000000000000ff19166e01000000000000000000000000000317815560405133917f98027b38153f995c4b802a5c7e6365bee3addb25af6b29818c0c304684d8052c91a2565b60015f5460ff16600581111561100b5761100b611950565b146110295760405163baf13b3f60e01b815260040160405180910390fd5b600182600381111561103d5761103d611950565b1415801561105d5750600282600381111561105a5761105a611950565b14155b1561107b5760405163c74a206d60e01b815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015906110dd5750336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614155b156110fb5760405163721c7c6760e11b815260040160405180910390fd5b5f805483919061ff00191661010083600381111561111b5761111b611950565b02179055505f805460ff191660021790556111567f000000000000000000000000000000000000000000000000000000000000000042611a54565b5f80546affffffffffffffff0000001916630100000067ffffffffffffffff9384168102919091179182905560405133937fa5c1e4e52593bfa9aacfee333f44dba84d7df350744f82ada7a542bd44113183936111bc9388938893919092041690611a01565b60405180910390a25050565b6111d0611818565b60025f80516020611ac683398151915255565b5f54600160681b900460ff161561120c5760405162560ff960e81b815260040160405180910390fd5b5f805460ff60681b1916600160681b1790557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316151560035f5462010000900460ff16600381111561126857611268611950565b036113fe575f54600160581b900460ff161561132a5761132a7f0000000000000000000000000000000000000000000000000000000000000000826112ad575f6112cf565b7f00000000000000000000000000000000000000000000000000000000000000005b6112f9907f0000000000000000000000000000000000000000000000000000000000000000611a41565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190611847565b5f54600160601b900460ff16156113b6576113b67f00000000000000000000000000000000000000000000000000000000000000008261136a575f61138c565b7f00000000000000000000000000000000000000000000000000000000000000005b6112f9907f0000000000000000000000000000000000000000000000000000000000000000611a41565b7f56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f5215f60029054906101000a900460ff166040516113f39190611978565b60405180910390a150565b5f6114497f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611a41565b90505f60015f5462010000900460ff16600381111561146a5761146a611950565b14611495577f00000000000000000000000000000000000000000000000000000000000000006114b7565b7f00000000000000000000000000000000000000000000000000000000000000005b90505f60015f5462010000900460ff1660038111156114d8576114d8611950565b14611503577f0000000000000000000000000000000000000000000000000000000000000000611525565b7f00000000000000000000000000000000000000000000000000000000000000005b90505f84611533575f61156a565b6127106115607f000000000000000000000000000000000000000000000000000000000000000086611a7c565b61156a9190611a93565b905080156115c6576115c66001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000083611847565b6115d4836112f98387611ab2565b84801561160057505f7f0000000000000000000000000000000000000000000000000000000000000000115b1561178b575f54600160701b900460ff16156116e35761168a6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611847565b6116de6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016847f0000000000000000000000000000000000000000000000000000000000000000611847565b61178b565b6117376001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016847f0000000000000000000000000000000000000000000000000000000000000000611847565b61178b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016837f0000000000000000000000000000000000000000000000000000000000000000611847565b7f56d734158477dd7d8c47ddb978b91e6ffb34d55ac626b4cd8095129d0829f5215f60029054906101000a900460ff166040516117c89190611978565b60405180910390a15050505050565b6117e5848484846001611881565b61181257604051635274afe760e01b81526001600160a01b03851660048201526024015b60405180910390fd5b50505050565b5f80516020611ac68339815191525460020361065657604051633ee5aeb560e01b815260040160405180910390fd5b61185483838360016118ee565b61187c57604051635274afe760e01b81526001600160a01b0384166004820152602401611809565b505050565b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f511483166118dd5783831516156118d1573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316611944578383151615611938573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b634e487b7160e01b5f52602160045260245ffd5b6004811061197457611974611950565b9052565b602081016119868284611964565b92915050565b60208101600683106119a0576119a0611950565b91905290565b8035600481106119b4575f80fd5b919050565b5f80604083850312156119ca575f80fd5b6119d3836119a6565b946020939093013593505050565b5f602082840312156119f1575f80fd5b6119fa826119a6565b9392505050565b60608101611a0f8286611964565b83602083015267ffffffffffffffff83166040830152949350505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561198657611986611a2d565b67ffffffffffffffff818116838216019080821115611a7557611a75611a2d565b5092915050565b808202811582820484141761198657611986611a2d565b5f82611aad57634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561198657611986611a2d56fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212205d2ebf1302a490c44845b5694b47a41069717f97cd29ea4771ec5d73dcee152b64736f6c63430008180033",
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
