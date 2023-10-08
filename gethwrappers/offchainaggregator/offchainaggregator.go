// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offchainaggregator

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

// OffchainAggregatorInitOCR is an auto generated low-level Go binding around an user-defined struct.
type OffchainAggregatorInitOCR struct {
	MaximumGasPrice           uint32
	ReasonableGasPrice        uint32
	MicroLinkPerEth           uint32
	LinkGweiPerObservation    uint32
	LinkGweiPerTransmission   uint32
	Link                      common.Address
	MinAnswer                 *big.Int
	MaxAnswer                 *big.Int
	BillingAccessController   common.Address
	RequesterAccessController common.Address
	Decimals                  uint8
	Description               string
}

// OffchainAggregatorMetaData contains all meta data concerning the OffchainAggregator contract.
var OffchainAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"_maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerTransmission\",\"type\":\"uint32\"},{\"internalType\":\"contractBondTokenInterface\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"int192\",\"name\":\"_minAnswer\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"_maxAnswer\",\"type\":\"int192\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAccessController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"internalType\":\"structOffchainAggregator.InitOCR\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentBatchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPrice\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPrice\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"microLinkPerEth\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"linkGweiPerObservation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"linkGweiPerTransmission\",\"type\":\"uint32\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"encodedConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"_oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"_newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rawReportContext\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"billingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkGweiPerTransmission\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractBondTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"}],\"name\":\"latestMerkleRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkelAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerOrTransmitter\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerTransmission\",\"type\":\"uint32\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_encodedConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_encoded\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBondTokenInterface\",\"name\":\"_linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"setLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"_newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620056a7380380620056a78339810160408190526200003491620007dd565b8051602082015160408301516060840151608085015160a0860151610100870151600080546001600160a01b03191633179055620000768787878787620001ea565b600380546001600160a01b0319166001600160a01b0384169081179091556040516000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a3620000ca81620002dc565b620000d462000577565b620000de62000577565b60005b601f8160ff1610156200012e576001838260ff16601f81106200010057fe5b61ffff909216602092909202015260018260ff8316601f81106200012057fe5b6020020152600101620000e1565b506200013e600583601f62000596565b506200014e600982601f62000633565b5050505061014087015160f81b7fff000000000000000000000000000000000000000000000000000000000000001660c0525050506101608401518051620001a29450603093506020909101915062000664565b50610120810151620001b49062000355565b620001c160008062000428565b60c0810151601790810b810b604090811b60805260e090920151810b900b901b60a05262000980565b6040805160a0808201835263ffffffff88811680845288821660208086018290528984168688018190528985166060808901829052958a1660809889018190526002805463ffffffff1916871763ffffffff60201b191664010000000087021763ffffffff60401b19166801000000000000000085021763ffffffff60601b19166c0100000000000000000000000084021763ffffffff60801b1916600160801b830217905589519586529285019390935283880152928201529283015291517fd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6929181900390910190a15050505050565b6004546001600160a01b0390811690821681146200035157600480546001600160a01b0319166001600160a01b03848116918217909255604080519284168352602083019190915280517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d489129281900390910190a15b5050565b6000546001600160a01b03163314620003b5576040805162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b602f546001600160a01b0390811690821681146200035157602f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634906200041c90839085906200092b565b60405180910390a15050565b6000546001600160a01b0316331462000488576040805162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015290519081900360640190fd5b60408051808201909152602e546001600160a01b03808216808452600160a01b90920463ffffffff1660208401528416141580620004d657508163ffffffff16816020015163ffffffff1614155b1562000572576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052602e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191620005699190879062000945565b60405180910390a35b505050565b604051806103e00160405280601f906020820280368337509192915050565b600283019183908215620006215791602002820160005b83821115620005ef57835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302620005ad565b80156200061f5782816101000a81549061ffff0219169055600201602081600101049283019260010302620005ef565b505b506200062f929150620006e6565b5090565b82601f810192821562000621579160200282015b828111156200062157825182559160200191906001019062000647565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200069c576000855562000621565b82601f10620006b757805160ff191683800117855562000621565b828001600101855582156200062157918201828111156200062157825182559160200191906001019062000647565b5b808211156200062f5760008155600101620006e7565b80516001600160a01b03811681146200071557600080fd5b919050565b8051601781900b81146200071557600080fd5b600082601f8301126200073e578081fd5b81516001600160401b038111156200075257fe5b602062000768601f8301601f191682016200095c565b82815285828487010111156200077c578384fd5b835b838110156200079b5785810183015182820184015282016200077e565b83811115620007ac57848385840101525b5095945050505050565b805163ffffffff811681146200071557600080fd5b805160ff811681146200071557600080fd5b600060208284031215620007ef578081fd5b81516001600160401b038082111562000806578283fd5b81840191506101808083870312156200081d578384fd5b62000828816200095c565b90506200083583620007b6565b81526200084560208401620007b6565b60208201526200085860408401620007b6565b60408201526200086b60608401620007b6565b60608201526200087e60808401620007b6565b60808201526200089160a08401620006fd565b60a0820152620008a460c084016200071a565b60c0820152620008b760e084016200071a565b60e0820152610100620008cc818501620006fd565b90820152610120620008e0848201620006fd565b90820152610140620008f4848201620007cb565b9082015261016083810151838111156200090c578586fd5b6200091a888287016200072d565b918301919091525095945050505050565b6001600160a01b0392831681529116602082015260400190565b63ffffffff92831681529116602082015260400190565b6040518181016001600160401b03811182821017156200097857fe5b604052919050565b60805160401c60a05160401c60c05160f81c614cf1620009b66000398061077052508061103f5250806106eb5250614cf16000f3fe608060405234801561001057600080fd5b50600436106102695760003560e01c806398e5b12a11610151578063c9807539116100c3578063e76d516811610087578063e76d516814610527578063eb4571631461052f578063eb5dcd6c14610542578063f2fde38b14610555578063fbffd2c114610568578063feaf968c1461057b57610269565b8063c9807539146104af578063d09dc339146104c2578063d4c520d1146104ca578063e4902f82146104ee578063e5fe45771461050e57610269565b8063a4e2d63411610115578063a4e2d6341461043b578063b121e14714610450578063b5ab58dc14610463578063b633620c14610476578063bd82470614610489578063c10753291461049c57610269565b806398e5b12a146103d4578063996e8298146103e95780639a6fc8f5146103f15780639c849b30146104155780639e3ceeab1461042857610269565b8063668a0f02116101ea57806381411834116101ae578063814118341461036f57806381ff7048146103845780638205bf6a1461039b5780638ac28d5a146103a35780638da5cb5b146103b65780638e0566de146103be57610269565b8063668a0f021461033a57806370da2f671461034257806370efdf2d1461034a5780637284e4161461035f57806379ba50971461036757610269565b80634fb17470116102315780634fb17470146102ef57806350d25bcd1461030457806354fd4d501461030c578063585aa7de14610314578063619d51941461032757610269565b80630eafb25b1461026e578063181f5a771461029757806322adbc78146102ac57806329937268146102c1578063313ce567146102da575b600080fd5b61028161027c36600461407a565b610583565b60405161028e919061480b565b60405180910390f35b61029f6106b2565b60405161028e9190614814565b6102b46106e9565b60405161028e91906147fd565b6102c961070d565b60405161028e959493929190614b76565b6102e261076e565b60405161028e9190614c47565b6103026102fd36600461445b565b610792565b005b6102816109ff565b610281610a04565b610302610322366004614161565b610a09565b610302610335366004614221565b611007565b61028161102a565b6102b461103d565b610352611061565b60405161028e9190614629565b61029f611070565b610302611106565b6103776111b5565b60405161028e91906146df565b61038c611216565b60405161028e93929190614b4c565b610281611236565b6103026103b136600461407a565b611264565b6103526112c3565b6103c66112d2565b60405161028e9291906147de565b6103dc611305565b60405161028e9190614bd5565b610352611469565b6104046103ff366004614528565b611478565b60405161028e959493929190614be9565b6103026104233660046140f9565b6116df565b61030261043636600461407a565b6118b8565b610443611975565b60405161028e919061472c565b61030261045e36600461407a565b611985565b6102816104713660046144ac565b611a4b565b6102816104843660046144ac565b611a69565b6103026104973660046144c4565b611aa4565b6103026104aa3660046140ce565b611bb8565b6103026104bd366004614389565b611e3e565b610281612716565b6104dd6104d836600461446d565b6127a7565b60405161028e959493929190614c19565b6105016104fc36600461407a565b61299f565b60405161028e9190614b26565b610516612a4c565b60405161028e959493929190614762565b610352612ac9565b61030261053d366004614427565b612ad8565b610302610550366004614096565b612c10565b61030261056336600461407a565b612d1d565b61030261057636600461407a565b612dbb565b610404612e11565b6001600160a01b03811660009081526028602090815260408083208151808301909252805460ff8082168452859484019161010090041660028111156105c557fe5b60028111156105d057fe5b90525090506000816020015160028111156105e757fe5b14156105f75760009150506106ad565b6040805160a08101825260025463ffffffff8082168352600160201b820481166020840152600160401b8204811693830193909352600160601b8104831660608301819052600160801b90910490921660808201528251909160009160019060059060ff16601f811061066657fe5b601091828204019190066002029054906101000a900461ffff160361ffff1602633b9aca0002905060016009846000015160ff16601f81106106a457fe5b01540301925050505b919050565b60408051808201909152601881527f4f6666636861696e41676772656761746f7220352e302e300000000000000000602082015290565b7f000000000000000000000000000000000000000000000000000000000000000081565b6040805160a08101825260025463ffffffff808216808452600160201b8304821660208501819052600160401b84048316958501869052600160601b8404831660608601819052600160801b90940490921660809094018490529490939290565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031633146107df576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b6003546001600160a01b039081169083168114156107fd57506109fb565b604080516370a0823160e01b815230600482015290516001600160a01b038516916370a08231916024808301926020929190829003018186803b15801561084357600080fd5b505afa158015610857573d6000803e3d6000fd5b505050506040513d602081101561086d57600080fd5b506108789050613028565b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156108c757600080fd5b505afa1580156108db573d6000803e3d6000fd5b505050506040513d60208110156108f157600080fd5b50516040805163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905291519293509084169163a9059cbb916044808201926020929091908290030181600087803b15801561094957600080fd5b505af115801561095d573d6000803e3d6000fd5b505050506040513d602081101561097357600080fd5b50516109ab576040805162461bcd60e51b8152602060048201526002602482015261343160f01b604482015290519081900360640190fd5b600380546001600160a01b0319166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b600190565b600481565b868560ff8616601f831115610a395760405162461bcd60e51b8152600401610a309061499e565b60405180910390fd5b60008111610a595760405162461bcd60e51b8152600401610a3090614ad2565b818314610a785760405162461bcd60e51b8152600401610a3090614a9a565b806003028311610a9a5760405162461bcd60e51b8152600401610a3090614966565b6000546001600160a01b03163314610ae7576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b60295415610be4576029805460001981019160009183908110610b0657fe5b6000918252602082200154602a80546001600160a01b0390921693509084908110610b2d57fe5b6000918252602090912001546001600160a01b03169050610b4d816133a1565b6001600160a01b03808316600090815260286020526040808220805461ffff1990811690915592841682529020805490911690556029805480610b8c57fe5b600082815260209020810160001990810180546001600160a01b0319169055019055602a805480610bb957fe5b600082815260209020810160001990810180546001600160a01b031916905501905550610ae7915050565b60005b8a811015610ef9576000602860008e8e85818110610c0157fe5b9050602002016020810190610c16919061407a565b6001600160a01b03168152602081019190915260400160002054610100900460ff166002811115610c4357fe5b14610c605760405162461bcd60e51b8152600401610a3090614ab6565b6040805180820190915260ff8216815260016020820152602860008e8e85818110610c8757fe5b9050602002016020810190610c9c919061407a565b6001600160a01b031681526020808201929092526040016000208251815460ff191660ff90911617808255918301519091829061ff001916610100836002811115610ce357fe5b02179055506000915060079050818c8c85818110610cfd57fe5b9050602002016020810190610d12919061407a565b6001600160a01b039081168252602082019290925260400160002054161415610d4d5760405162461bcd60e51b8152600401610a30906148f6565b6000602860008c8c85818110610d5f57fe5b9050602002016020810190610d74919061407a565b6001600160a01b03168152602081019190915260400160002054610100900460ff166002811115610da157fe5b14610dbe5760405162461bcd60e51b8152600401610a30906148da565b6040805180820190915260ff8216815260026020820152602860008c8c85818110610de557fe5b9050602002016020810190610dfa919061407a565b6001600160a01b031681526020808201929092526040016000208251815460ff191660ff90911617808255918301519091829061ff001916610100836002811115610e4157fe5b021790555090505060298c8c83818110610e5757fe5b9050602002016020810190610e6c919061407a565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055602a8a8a83818110610eac57fe5b9050602002016020810190610ec1919061407a565b815460018082018455600093845260209093200180546001600160a01b0319166001600160a01b039290921691909117905501610be7565b50602b805460ff8916600160a81b0260ff60a81b19909116179055602d80544363ffffffff908116600160201b90810267ffffffff0000000019841617808316600101831663ffffffff19909116179384905590910481169116610f6530828f8f8f8f8f8f8f8f613579565b602b60000160006101000a8154816001600160801b03021916908360801c02179055506000602b60000160106101000a81548164ffffffffff021916908364ffffffffff1602179055507f25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb982828f8f8f8f8f8f8f8f604051610ff09a99989796959493929190614ba5565b60405180910390a150505050505050505050505050565b602d8054911515600160401b0268ff000000000000000019909216919091179055565b602b54600160b01b900463ffffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000081565b602f546001600160a01b031690565b60308054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156110fc5780601f106110d1576101008083540402835291602001916110fc565b820191906000526020600020905b8154815290600101906020018083116110df57829003601f168201915b5050505050905090565b6001546001600160a01b0316331461115e576040805162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b604482015290519081900360640190fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060602a8054806020026020016040519081016040528092919081815260200182805480156110fc57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111ef575050505050905090565b602d54602b5463ffffffff80831692600160201b9004169060801b909192565b602b54600160b01b900463ffffffff166000908152602c60205260409020600301546001600160401b031690565b6001600160a01b038181166000908152600760205260409020541633146112b7576040805162461bcd60e51b8152602060048201526002602482015261343360f01b604482015290519081900360640190fd5b6112c0816133a1565b50565b6000546001600160a01b031681565b60408051808201909152602e546001600160a01b038116808352600160a01b90910463ffffffff16602090920182905291565b600080546001600160a01b031633148061139f5750602f54604051630d629b5f60e31b81526001600160a01b0390911690636b14daf89061134f903390600090369060040161463d565b60206040518083038186803b15801561136757600080fd5b505afa15801561137b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139f919061423d565b6113bb5760405162461bcd60e51b8152600401610a3090614982565b6040805160808082018352602b549081901b6001600160801b031916808352600160801b820464ffffffffff811660208501819052600160a81b840460ff1685870152600160b01b90930463ffffffff90811660608601529451939433947f3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b520379461144d949360081c9092169190614737565b60405180910390a2806060015160010163ffffffff1691505090565b6004546001600160a01b031690565b600080600080600063ffffffff866001600160501b031611156040518060400160405280600f81526020016e139bc819185d18481c1c995cd95b9d608a1b815250906114d75760405162461bcd60e51b8152600401610a309190614814565b5063ffffffff86166000908152602c602090815260408083208151815460026001821615610100026000190190911604601f8101859004909402810160e0908101845260c08201858152919492938593908401928592849284918701828280156115825780601f1061155757610100808354040283529160200191611582565b820191906000526020600020905b81548152906001019060200180831161156557829003601f168201915b505050918352505060018201546001600160401b038116602080840191909152600160401b90910463ffffffff16604080840191909152600284018054825181850281018501909352808352606090940193919290919060009084015b828210156116a55760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f8101859004909402810160609081018352918101848152909384928491908401828280156116835780601f1061165857610100808354040283529160200191611683565b820191906000526020600020905b81548152906001019060200180831161166657829003601f168201915b50505050508152602001600182015481525050815260200190600101906115df565b505050915250508152600391909101546001600160401b039081166020928301529101519798600198909116965086955088945092505050565b6000546001600160a01b0316331461172c576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b828114611765576040805162461bcd60e51b8152602060048201526002602482015261353160f01b604482015290519081900360640190fd5b60005b838110156118b157600085858381811061177e57fe5b905060200201356001600160a01b03169050600084848481811061179e57fe5b6001600160a01b0385811660009081526007602090815260409091205492029390930135831693509091169050801580806117ea5750826001600160a01b0316826001600160a01b0316145b611820576040805162461bcd60e51b81526020600482015260026024820152611a9960f11b604482015290519081900360640190fd5b6001600160a01b03848116600090815260076020526040902080546001600160a01b031916858316908117909155908316146118a157826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b5050600190920191506117689050565b5050505050565b6000546001600160a01b03163314611905576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b602f546001600160a01b0390811690821681146109fb57602f80546001600160a01b0319166001600160a01b0384161790556040517f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae6349061196990839085906147c4565b60405180910390a15050565b602d54600160401b900460ff1690565b6001600160a01b038181166000908152600860205260409020541633146119d8576040805162461bcd60e51b8152602060048201526002602482015261353560f01b604482015290519081900360640190fd5b6001600160a01b0381811660008181526007602090815260408083208054336001600160a01b031980831682179093556008909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b600063ffffffff821115611a61575060006106ad565b506001919050565b600063ffffffff821115611a7f575060006106ad565b5063ffffffff166000908152602c60205260409020600301546001600160401b031690565b6004546000546001600160a01b039182169116331480611b65575060408051630d629b5f60e31b815233600482018181526024830193845236604484018190526001600160a01b03861694636b14daf8946000939190606401848480828437600083820152604051601f909101601f1916909201965060209550909350505081840390508186803b158015611b3857600080fd5b505afa158015611b4c573d6000803e3d6000fd5b505050506040513d6020811015611b6257600080fd5b50515b611b9b576040805162461bcd60e51b81526020600482015260026024820152611a1960f11b604482015290519081900360640190fd5b611ba3613028565b611bb086868686866135c4565b505050505050565b6000546001600160a01b0316331480611c7957506004805460408051630d629b5f60e31b8152339381018481526024820192835236604483018190526001600160a01b0390941694636b14daf8949093600093919291606401848480828437600083820152604051601f909101601f1916909201965060209550909350505081840390508186803b158015611c4c57600080fd5b505afa158015611c60573d6000803e3d6000fd5b505050506040513d6020811015611c7657600080fd5b50515b611caf576040805162461bcd60e51b8152602060048201526002602482015261343760f01b604482015290519081900360640190fd5b6000611cb96136a8565b600354604080516370a0823160e01b815230600482015290519293506000926001600160a01b03909216916370a0823191602480820192602092909190829003018186803b158015611d0a57600080fd5b505afa158015611d1e573d6000803e3d6000fd5b505050506040513d6020811015611d3457600080fd5b5051905081811015611d72576040805162461bcd60e51b8152602060048201526002602482015261068760f31b604482015290519081900360640190fd5b6003546001600160a01b031663a9059cbb85611d908585038761385c565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611dd657600080fd5b505af1158015611dea573d6000803e3d6000fd5b505050506040513d6020811015611e0057600080fd5b5051611e38576040805162461bcd60e51b8152602060048201526002602482015261343960f01b604482015290519081900360640190fd5b50505050565b60005a9050611e51888888888888613876565b3614611e6f5760405162461bcd60e51b8152600401610a3090614912565b611e77613cbc565b6040805160808082018352602b549081901b6001600160801b0319168252600160801b810464ffffffffff166020830152600160a81b810460ff1692820192909252600160b01b90910463ffffffff16606082015281526000611edc8a8a018b614259565b60408501526080840182905283515190925060589190911b906001600160801b0319808316911614611f205760405162461bcd60e51b8152600401610a30906149f2565b608083015183516020015164ffffffffff808316911610611f535760405162461bcd60e51b8152600401610a3090614a2a565b83516040015160ff168911611f7a5760405162461bcd60e51b8152600401610a3090614a0e565b601f891115611f9b5760405162461bcd60e51b8152600401610a30906149ba565b868914611fba5760405162461bcd60e51b8152600401610a309061492e565b601f8460400151511115611fe05760405162461bcd60e51b8152600401610a309061494a565b83600001516040015160020260ff16846040015151116120125760405162461bcd60e51b8152600401610a3090614a46565b886001600160401b038111801561202857600080fd5b506040519080825280601f01601f191660200182016040528015612053576020820181803683370190505b50606085015260005b60ff81168a11156120ac57868160ff166020811061207657fe5b1a60f81b85606001518260ff168151811061208d57fe5b60200101906001600160f81b031916908160001a90535060010161205c565b508360400151516001600160401b03811180156120c857600080fd5b506040519080825280601f01601f1916602001820160405280156120f3576020820181803683370190505b506020850152612101613cf0565b60005b8560400151518160ff1610156121ba576000858260ff166020811061212557fe5b1a90508281601f811061213457fe5b6020020151156121565760405162461bcd60e51b8152600401610a3090614a62565b6001838260ff16601f811061216757fe5b91151560209283029190910152869060ff841690811061218357fe5b1a60f81b87602001518360ff168151811061219a57fe5b60200101906001600160f81b031916908160001a90535050600101612104565b503360009081526028602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156121f857fe5b600281111561220357fe5b905250905060028160200151600281111561221a57fe5b14801561224e5750602a816000015160ff168154811061223657fe5b6000918252602090912001546001600160a01b031633145b61226a5760405162461bcd60e51b8152600401610a3090614a7e565b5050835164ffffffffff9091166020909101525050604051600090612292908b908b90614619565b604051809103902090506122a4613cf0565b6122ac613d0f565b60005b89811015612426576000600185876060015184815181106122cc57fe5b60209101015160f81c601b018e8e868181106122e457fe5b905060200201358d8d878181106122f757fe5b905060200201356040516000815260200160405260405161231b94939291906147a6565b6020604051602081039080840390855afa15801561233d573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526028602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561238c57fe5b600281111561239757fe5b90525092506001836020015160028111156123ae57fe5b146123cb5760405162461bcd60e51b8152600401610a3090614b0a565b8251849060ff16601f81106123dc57fe5b6020020151156123fe5760405162461bcd60e51b8152600401610a3090614aee565b600184846000015160ff16601f811061241357fe5b91151560209092020152506001016122af565b505082516060015163ffffffff166000908152602c60205260408082209085015180519194509192506002810490811061245c57fe5b60209081029190910101518351606001805160010163ffffffff169052905060005b81606001515163ffffffff168163ffffffff161015612501578260000160020182606001518263ffffffff16815181106124b457fe5b6020908102919091018101518254600181018455600093845292829020815180519294600202909101926124ed92849290910190613d26565b50602091909101516001918201550161247e565b5060208181015160018401805467ffffffffffffffff19166001600160401b039092169190911790558151805161253b9285920190613d26565b5060409081015160018301805463ffffffff60401b1916600160401b63ffffffff9384160217905560038301805467ffffffffffffffff1916426001600160401b031617905583516060015160208501516080860151935191909216927f62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db1926125c8928692339291614827565b60405180910390a281516060015160405160009163ffffffff16907f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac602719061261090429061480b565b60405180910390a3815160600151600182015460405163ffffffff909216916001600160401b03909116907f5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa0939061266890429061480b565b60405180910390a35080518051602b8054602084015160408501516060909501516001600160801b031990921660809490941c9390931764ffffffffff60801b1916600160801b64ffffffffff909416939093029290921760ff60a81b1916600160a81b60ff909416939093029290921763ffffffff60b01b1916600160b01b63ffffffff928316021790915582106126fd57fe5b61270b82826020015161388e565b505050505050505050565b600354604080516370a0823160e01b8152306004820152905160009283926001600160a01b03909116916370a0823191602480820192602092909190829003018186803b15801561276657600080fd5b505afa15801561277a573d6000803e3d6000fd5b505050506040513d602081101561279057600080fd5b50519050600061279e6136a8565b90910391505090565b600080808080805b602b5463ffffffff600160b01b90910481166000908152602c60205260409020600201548116908216101561299357602b5463ffffffff600160b01b90910481166000908152602c6020526040902060020180546128e692841690811061281257fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f8101859004850282018501909352828152929091908301828280156128a55780601f1061287a576101008083540402835291602001916128a5565b820191906000526020600020905b81548152906001019060200180831161288857829003601f168201915b505050505089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613a9b92505050565b1561298b57602b5463ffffffff600160b01b90910481166000818152602c6020526040902060018101546002909101805492936001600160401b03909216929091851690811061293257fe5b6000918252602080832060016002909302019190910154602b5463ffffffff600160b01b90910481168452602c909252604090922060030154931698506001600160401b03918216975095501692508291506129959050565b6001016127af565b505b9295509295909350565b6001600160a01b03811660009081526028602090815260408083208151808301909252805460ff8082168452859484019161010090041660028111156129e157fe5b60028111156129ec57fe5b9052509050600081602001516002811115612a0357fe5b1415612a135760009150506106ad565b60016005826000015160ff16601f8110612a2957fe5b601091828204019190066002029054906101000a900461ffff1603915050919050565b600080808080333214612a715760405162461bcd60e51b8152600401610a30906149d6565b5050602b5463ffffffff600160b01b820481166000908152602c6020526040902060030154608083901b96600160801b909304600881901c909216955064ffffffffff9091169350600192506001600160401b031690565b6003546001600160a01b031690565b6000546001600160a01b03163314612b25576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b60408051808201909152602e546001600160a01b03808216808452600160a01b90920463ffffffff1660208401528416141580612b7257508163ffffffff16816020015163ffffffff1614155b15612c0b576040805180820182526001600160a01b0385811680835263ffffffff86166020938401819052602e80546001600160a01b031916831763ffffffff60a01b1916600160a01b9092029190911790558451928501519351909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154191612c0291908790614b35565b60405180910390a35b505050565b6001600160a01b03828116600090815260076020526040902054163314612c63576040805162461bcd60e51b8152602060048201526002602482015261353360f01b604482015290519081900360640190fd5b336001600160a01b0382161415612ca6576040805162461bcd60e51b81526020600482015260026024820152610d4d60f21b604482015290519081900360640190fd5b6001600160a01b03808316600090815260086020526040902080548383166001600160a01b031982168117909255909116908114612c0b576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b6000546001600160a01b03163314612d6a576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000546001600160a01b03163314612e08576040805162461bcd60e51b81526020600482015260166024820152600080516020614cc5833981519152604482015290519081900360640190fd5b6112c081613b23565b602b54600160b01b900463ffffffff166000818152602c602090815260408083208151815460026101006001831615026000190190911604601f8101859004909402810160e0908101845260c08201858152869586958695869594909385939284019285928492918491870182828015612ecc5780601f10612ea157610100808354040283529160200191612ecc565b820191906000526020600020905b815481529060010190602001808311612eaf57829003601f168201915b505050918352505060018201546001600160401b038116602080840191909152600160401b90910463ffffffff16604080840191909152600284018054825181850281018501909352808352606090940193919290919060009084015b82821015612fef5760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f810185900490940281016060908101835291810184815290938492849190840182828015612fcd5780601f10612fa257610100808354040283529160200191612fcd565b820191906000526020600020905b815481529060010190602001808311612fb057829003601f168201915b5050505050815260200160018201548152505081526020019060010190612f29565b505050915250508152600391909101546001600160401b0390811660209283015291015160019650169350839250859150509091929394565b6040805160a08101825260025463ffffffff8082168352600160201b820481166020840152600160401b8204811683850152600160601b820481166060840152600160801b90910416608082015260035482516103e081019384905291926001600160a01b0390911691600091600590601f908285855b82829054906101000a900461ffff1661ffff168152602001906002019060208260010104928301926001038202915080841161309f575050604080516103e0810191829052959650600095945060099350601f9250905082845b8154815260200190600101908083116130f957505050505090506000602a80548060200260200160405190810160405280929190818152602001828054801561316b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161314d575b5050505050905060005b815181101561338557600060018483601f811061318e57fe5b6020020151039050600060018684601f81106131a657fe5b60200201510361ffff169050600082896060015163ffffffff168302633b9aca0002019050600081111561337a576000600760008787815181106131e657fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050886001600160a01b031663a9059cbb82846040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561327b57600080fd5b505af115801561328f573d6000803e3d6000fd5b505050506040513d60208110156132a557600080fd5b50516132dd576040805162461bcd60e51b8152602060048201526002602482015261343560f01b604482015290519081900360640190fd5b60018886601f81106132eb57fe5b61ffff909216602092909202015260018786601f811061330757fe5b602002018181525050886001600160a01b0316816001600160a01b031687878151811061333057fe5b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c856040518082815260200191505060405180910390a4505b505050600101613175565b50613393600584601f613db2565b50611bb0600983601f613e43565b6001600160a01b03811660009081526028602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156133e757fe5b60028111156133f257fe5b9052509050600061340283610583565b90508015612c0b576001600160a01b03808416600090815260076020908152604080832054600354825163a9059cbb60e01b8152918616600483018190526024830188905292519295169363a9059cbb9360448084019491939192918390030190829087803b15801561347457600080fd5b505af1158015613488573d6000803e3d6000fd5b505050506040513d602081101561349e57600080fd5b50516134d6576040805162461bcd60e51b81526020600482015260026024820152610d0d60f21b604482015290519081900360640190fd5b60016005846000015160ff16601f81106134ec57fe5b601091828204019190066002026101000a81548161ffff021916908361ffff16021790555060016009846000015160ff16601f811061352757fe5b01556003546040805184815290516001600160a01b039283169284811692908816917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c9181900360200190a450505050565b60008a8a8a8a8a8a8a8a8a8a60405160200161359e9a99989796959493929190614662565b6040516020818303038152906040528051906020012090509a9950505050505050505050565b6040805160a0808201835263ffffffff88811680845288821660208086018290528984168688018190528985166060808901829052958a1660809889018190526002805463ffffffff1916871767ffffffff000000001916600160201b87021763ffffffff60401b1916600160401b85021763ffffffff60601b1916600160601b84021763ffffffff60801b1916600160801b830217905589519586529285019390935283880152928201529283015291517fd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6929181900390910190a15050505050565b604080516103e0810191829052600091829190600590601f908285855b82829054906101000a900461ffff1661ffff16815260200190600201906020826001010492830192600103820291508084116136c55790505050505050905060005b601f8110156137355760018282601f811061371e57fe5b60200201510361ffff169290920191600101613707565b506040805160a08101825260025463ffffffff8082168352600160201b82048116602080850191909152600160401b8304821684860152600160601b8304821660608501819052600160801b9093049091166080840152602a805485518184028101840190965280865296909202633b9aca00029592936000939092918301828280156137eb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116137cd575b5050604080516103e0810191829052949550600094935060099250601f915082845b81548152602001906001019080831161380d575050505050905060005b82518110156138545760018282601f811061384157fe5b602002015103959095019460010161382a565b505050505090565b60008183101561386d575081613870565b50805b92915050565b602083810286019082020160e4019695505050505050565b3360009081526028602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156138cb57fe5b60028111156138d657fe5b9052506040805160a08101825260025463ffffffff8082168352600160201b820481166020840152600160401b8204811683850152600160601b820481166060840152600160801b90910416608082015281516103e08101928390529293509161398a91859190600590601f90826000855b82829054906101000a900461ffff1661ffff16815260200190600201906020826001010492830192600103820291508084116139485790505050505050613b9a565b61399890600590601f613db2565b506002826020015160028111156139ab57fe5b146139e2576040805162461bcd60e51b8152602060048201526002602482015261035360f41b604482015290519081900360640190fd5b6000613a09633b9aca003a04836020015163ffffffff16846000015163ffffffff16613c0f565b90506010360260005a90506000613a288863ffffffff16858585613c35565b6001600160801b031690506000620f4240866040015163ffffffff16830281613a4d57fe5b049050856080015163ffffffff16633b9aca0002816009896000015160ff16601f8110613a7657fe5b015401016009886000015160ff16601f8110613a8e57fe5b0155505050505050505050565b805182516000918491849114613ab657600092505050613870565b60005b8251811015613b1757818181518110613ace57fe5b602001015160f81c60f81b6001600160f81b031916838281518110613aef57fe5b01602001516001600160f81b03191614613b0f5760009350505050613870565b600101613ab9565b50600195945050505050565b6004546001600160a01b0390811690821681146109fb57600480546001600160a01b0319166001600160a01b03848116918217909255604080519284168352602083019190915280517f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d489129281900390910190a15050565b613ba2613cf0565b60005b8351811015613c07576000848281518110613bbc57fe5b016020015160f81c9050613be18482601f8110613bd557fe5b60200201516001613c9d565b848260ff16601f8110613bf057fe5b61ffff909216602092909202015250600101613ba5565b509092915050565b60008383811015613c2257600285850304015b613c2c818461385c565b95945050505050565b600081851015613c71576040805162461bcd60e51b81526020600482015260026024820152611a1b60f11b604482015290519081900360640190fd5b818503830161179301633b9aca00858202026001600160801b038110613c9357fe5b9695505050505050565b6000613cb58261ffff168461ffff160161ffff61385c565b9392505050565b6040518060a00160405280613ccf613e70565b81526060602082018190526040820181905280820152600060809091015290565b604051806103e00160405280601f906020820280368337509192915050565b604080518082019091526000808252602082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282613d5c5760008555613da2565b82601f10613d7557805160ff1916838001178555613da2565b82800160010185558215613da2579182015b82811115613da2578251825591602001919060010190613d87565b50613dae929150613e97565b5090565b600283019183908215613da25791602002820160005b83821115613e0857835183826101000a81548161ffff021916908361ffff1602179055509260200192600201602081600101049283019260010302613dc8565b8015613e365782816101000a81549061ffff0219169055600201602081600101049283019260010302613e08565b5050613dae929150613e97565b82601f8101928215613da25791602002820182811115613da2578251825591602001919060010190613d87565b60408051608081018252600080825260208201819052918101829052606081019190915290565b5b80821115613dae5760008155600101613e98565b60008083601f840112613ebd578182fd5b5081356001600160401b03811115613ed3578182fd5b6020830191508360208083028501011115613eed57600080fd5b9250929050565b600082601f830112613f04578081fd5b81356020613f19613f1483614c78565b614c55565b82815281810190858301855b85811015613f9f5781358801604080601f19838d03011215613f45578889fd5b80518181016001600160401b038282108183111715613f6057fe5b908352838901359080821115613f74578b8cfd5b50613f838d8a83870101613feb565b8252509101358682015284529284019290840190600101613f25565b5090979650505050505050565b60008083601f840112613fbd578182fd5b5081356001600160401b03811115613fd3578182fd5b602083019150836020828501011115613eed57600080fd5b600082601f830112613ffb578081fd5b81356001600160401b0381111561400e57fe5b614021601f8201601f1916602001614c55565b818152846020838601011115614035578283fd5b816020850160208301379081016020019190915292915050565b803563ffffffff811681146106ad57600080fd5b80356001600160401b03811681146106ad57600080fd5b60006020828403121561408b578081fd5b8135613cb581614ca1565b600080604083850312156140a8578081fd5b82356140b381614ca1565b915060208301356140c381614ca1565b809150509250929050565b600080604083850312156140e0578182fd5b82356140eb81614ca1565b946020939093013593505050565b6000806000806040858703121561410e578182fd5b84356001600160401b0380821115614124578384fd5b61413088838901613eac565b90965094506020870135915080821115614148578384fd5b5061415587828801613eac565b95989497509550505050565b60008060008060008060008060a0898b03121561417c578586fd5b88356001600160401b0380821115614192578788fd5b61419e8c838d01613eac565b909a50985060208b01359150808211156141b6578788fd5b6141c28c838d01613eac565b909850965060408b0135915060ff821682146141dc578586fd5b8195506141eb60608c01614063565b945060808b0135915080821115614200578384fd5b5061420d8b828c01613fac565b999c989b5096995094979396929594505050565b600060208284031215614232578081fd5b8135613cb581614cb6565b60006020828403121561424e578081fd5b8151613cb581614cb6565b60008060006060848603121561426d578081fd5b833592506020840135915060408401356001600160401b0380821115614291578283fd5b818601915086601f8301126142a4578283fd5b81356142b2613f1482614c78565b818152602080820191908501865b8481101561437857813587016080818e03601f190112156142df578889fd5b6040516080810181811089821117156142f457fe5b604052602082013588811115614308578a8bfd5b6143178f602083860101613feb565b82525061432660408301614063565b60208201526143376060830161404f565b604082015260808201358881111561434d578a8bfd5b61435c8f602083860101613ef4565b60608301525085525060209384019391909101906001016142c0565b505080955050505050509250925092565b60008060008060008060006080888a0312156143a3578081fd5b87356001600160401b03808211156143b9578283fd5b6143c58b838c01613fac565b909950975060208a01359150808211156143dd578283fd5b6143e98b838c01613eac565b909750955060408a0135915080821115614401578283fd5b5061440e8a828b01613eac565b989b979a50959894979596606090950135949350505050565b60008060408385031215614439578182fd5b823561444481614ca1565b91506144526020840161404f565b90509250929050565b600080604083850312156140a8578182fd5b6000806020838503121561447f578182fd5b82356001600160401b03811115614494578283fd5b6144a085828601613fac565b90969095509350505050565b6000602082840312156144bd578081fd5b5035919050565b600080600080600060a086880312156144db578283fd5b6144e48661404f565b94506144f26020870161404f565b93506145006040870161404f565b925061450e6060870161404f565b915061451c6080870161404f565b90509295509295909350565b600060208284031215614539578081fd5b81356001600160501b0381168114613cb5578182fd5b6001600160a01b03169052565b60008284526020808501945082825b8581101561459957813561457e81614ca1565b6001600160a01b03168752958201959082019060010161456b565b509495945050505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452815b818110156145f3576020818501810151868301820152016145d7565b818111156146045782602083870101525b50601f01601f19169290920160200192915050565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6001600160a01b0384168152604060208201819052600090613c2c90830184866145a4565b600060018060a01b038c1682526001600160401b03808c16602084015260e0604084015261469460e084018b8d61455c565b83810360608501526146a7818a8c61455c565b905060ff8816608085015281871660a085015283810360c08501526146cd8186886145a4565b9e9d5050505050505050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156147205783516001600160a01b0316835292840192918401916001016146fb565b50909695505050505050565b901515815260200190565b6001600160801b031993909316835263ffffffff91909116602083015260ff16604082015260600190565b6001600160801b031995909516855263ffffffff93909316602085015260ff91909116604084015260170b60608301526001600160401b0316608082015260a00190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0392909216825263ffffffff16602082015260400190565b60179190910b815260200190565b90815260200190565b600060208252613cb560208301846145ce565b6000608082016080835281875460018082166000811461484e576001811461486c576148a5565b60028304607f16855260ff19831660a088015260c0870193506148a5565b6002830480865261487c8c614c95565b875b8281101561489b5781548a820160a001529084019060200161487e565b890160a001955050505b5050506148b5602085018861454f565b83810360408501526148c781876145ce565b9250505082606083015295945050505050565b602080825260029082015261303760f01b604082015260600190565b602080825260029082015261181b60f11b604082015260600190565b602080825260029082015261031360f41b604082015260600190565b602080825260029082015261313560f01b604082015260600190565b602080825260029082015261189b60f11b604082015260600190565b6020808252600290820152610c0d60f21b604082015260600190565b602080825260029082015261060760f31b604082015260600190565b602080825260029082015261303160f01b604082015260600190565b6020808252600290820152610c4d60f21b604082015260600190565b602080825260029082015261303960f01b604082015260600190565b602080825260029082015261313160f01b604082015260600190565b602080825260029082015261313360f01b604082015260600190565b602080825260029082015261189960f11b604082015260600190565b602080825260029082015261313760f01b604082015260600190565b602080825260029082015261062760f31b604082015260600190565b602080825260029082015261313960f01b604082015260600190565b602080825260029082015261303360f01b604082015260600190565b602080825260029082015261303560f01b604082015260600190565b602080825260029082015261181960f11b604082015260600190565b602080825260029082015261191960f11b604082015260600190565b602080825260029082015261323160f01b604082015260600190565b61ffff91909116815260200190565b63ffffffff92831681529116602082015260400190565b63ffffffff93841681529190921660208201526001600160801b0319909116604082015260600190565b63ffffffff95861681529385166020850152918416604084015283166060830152909116608082015260a00190565b600063ffffffff8c1682526001600160401b03808c16602084015260e0604084015261469460e084018b8d61455c565b6001600160501b0391909116815260200190565b6001600160501b039586168152602081019490945260408401929092526060830152909116608082015260a00190565b6001600160501b03959095168552602085019390935260408401919091526060830152608082015260a00190565b60ff91909116815260200190565b6040518181016001600160401b0381118282101715614c7057fe5b604052919050565b60006001600160401b03821115614c8b57fe5b5060209081020190565b60009081526020902090565b6001600160a01b03811681146112c057600080fd5b80151581146112c057600080fdfe4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000a164736f6c6343000706000a",
}

// OffchainAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OffchainAggregatorMetaData.ABI instead.
var OffchainAggregatorABI = OffchainAggregatorMetaData.ABI

// OffchainAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OffchainAggregatorMetaData.Bin instead.
var OffchainAggregatorBin = OffchainAggregatorMetaData.Bin

// DeployOffchainAggregator deploys a new Ethereum contract, binding an instance of OffchainAggregator to it.
func DeployOffchainAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, data OffchainAggregatorInitOCR) (common.Address, *types.Transaction, *OffchainAggregator, error) {
	parsed, err := OffchainAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffchainAggregatorBin), backend, data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffchainAggregator{OffchainAggregatorCaller: OffchainAggregatorCaller{contract: contract}, OffchainAggregatorTransactor: OffchainAggregatorTransactor{contract: contract}, OffchainAggregatorFilterer: OffchainAggregatorFilterer{contract: contract}}, nil
}

// OffchainAggregator is an auto generated Go binding around an Ethereum contract.
type OffchainAggregator struct {
	OffchainAggregatorCaller     // Read-only binding to the contract
	OffchainAggregatorTransactor // Write-only binding to the contract
	OffchainAggregatorFilterer   // Log filterer for contract events
}

// OffchainAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OffchainAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OffchainAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OffchainAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OffchainAggregatorSession struct {
	Contract     *OffchainAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OffchainAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OffchainAggregatorCallerSession struct {
	Contract *OffchainAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OffchainAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OffchainAggregatorTransactorSession struct {
	Contract     *OffchainAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OffchainAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OffchainAggregatorRaw struct {
	Contract *OffchainAggregator // Generic contract binding to access the raw methods on
}

// OffchainAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OffchainAggregatorCallerRaw struct {
	Contract *OffchainAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// OffchainAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OffchainAggregatorTransactorRaw struct {
	Contract *OffchainAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOffchainAggregator creates a new instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregator(address common.Address, backend bind.ContractBackend) (*OffchainAggregator, error) {
	contract, err := bindOffchainAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregator{OffchainAggregatorCaller: OffchainAggregatorCaller{contract: contract}, OffchainAggregatorTransactor: OffchainAggregatorTransactor{contract: contract}, OffchainAggregatorFilterer: OffchainAggregatorFilterer{contract: contract}}, nil
}

// NewOffchainAggregatorCaller creates a new read-only instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorCaller(address common.Address, caller bind.ContractCaller) (*OffchainAggregatorCaller, error) {
	contract, err := bindOffchainAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorCaller{contract: contract}, nil
}

// NewOffchainAggregatorTransactor creates a new write-only instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OffchainAggregatorTransactor, error) {
	contract, err := bindOffchainAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorTransactor{contract: contract}, nil
}

// NewOffchainAggregatorFilterer creates a new log filterer instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OffchainAggregatorFilterer, error) {
	contract, err := bindOffchainAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorFilterer{contract: contract}, nil
}

// bindOffchainAggregator binds a generic wrapper to an already deployed contract.
func bindOffchainAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OffchainAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainAggregator *OffchainAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainAggregator.Contract.OffchainAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainAggregator *OffchainAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.OffchainAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainAggregator *OffchainAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.OffchainAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainAggregator *OffchainAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainAggregator *OffchainAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainAggregator *OffchainAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.contract.Transact(opts, method, params...)
}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) BillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "billingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) BillingAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.BillingAccessController(&_OffchainAggregator.CallOpts)
}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) BillingAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.BillingAccessController(&_OffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorSession) Decimals() (uint8, error) {
	return _OffchainAggregator.Contract.Decimals(&_OffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Decimals() (uint8, error) {
	return _OffchainAggregator.Contract.Decimals(&_OffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorSession) Description() (string, error) {
	return _OffchainAggregator.Contract.Description(&_OffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Description() (string, error) {
	return _OffchainAggregator.Contract.Description(&_OffchainAggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetAnswer(&_OffchainAggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetAnswer(&_OffchainAggregator.CallOpts, _roundId)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPrice         uint32
		ReasonableGasPrice      uint32
		MicroLinkPerEth         uint32
		LinkGweiPerObservation  uint32
		LinkGweiPerTransmission uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPrice = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPrice = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.MicroLinkPerEth = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.LinkGweiPerObservation = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.LinkGweiPerTransmission = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorSession) GetBilling() (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	return _OffchainAggregator.Contract.GetBilling(&_OffchainAggregator.CallOpts)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetBilling() (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	return _OffchainAggregator.Contract.GetBilling(&_OffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorSession) GetLinkToken() (common.Address, error) {
	return _OffchainAggregator.Contract.GetLinkToken(&_OffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _OffchainAggregator.Contract.GetLinkToken(&_OffchainAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetRoundData(&_OffchainAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetRoundData(&_OffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetTimestamp(&_OffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetTimestamp(&_OffchainAggregator.CallOpts, _roundId)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorSession) IsLocked() (bool, error) {
	return _OffchainAggregator.Contract.IsLocked(&_OffchainAggregator.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCallerSession) IsLocked() (bool, error) {
	return _OffchainAggregator.Contract.IsLocked(&_OffchainAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestAnswer(&_OffchainAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestAnswer(&_OffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes16 configDigest)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [16]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([16]byte)).(*[16]byte)

	return *outstruct, err

}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes16 configDigest)
func (_OffchainAggregator *OffchainAggregatorSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDetails(&_OffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes16 configDigest)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDetails(&_OffchainAggregator.CallOpts)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xd4c520d1.
//
// Solidity: function latestMerkleRoundData(string taskId) view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestMerkleRoundData(opts *bind.CallOpts, taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestMerkleRoundData", taskId)

	outstruct := new(struct {
		RoundId      *big.Int
		BatchId      *big.Int
		MerkelAnswer [32]byte
		StartedAt    *big.Int
		UpdatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MerkelAnswer = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StartedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xd4c520d1.
//
// Solidity: function latestMerkleRoundData(string taskId) view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorSession) LatestMerkleRoundData(taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestMerkleRoundData(&_OffchainAggregator.CallOpts, taskId)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xd4c520d1.
//
// Solidity: function latestMerkleRoundData(string taskId) view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestMerkleRoundData(taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestMerkleRoundData(&_OffchainAggregator.CallOpts, taskId)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestRound() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestRound(&_OffchainAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestRound(&_OffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestRoundData(&_OffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestRoundData(&_OffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestTimestamp(&_OffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestTimestamp(&_OffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes16 configDigest, uint32 epoch, uint8 round, int192 latestAnswer, uint64 latestTimestamp)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestTransmissionDetails")

	outstruct := new(struct {
		ConfigDigest    [16]byte
		Epoch           uint32
		Round           uint8
		LatestAnswer    *big.Int
		LatestTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigDigest = *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)
	outstruct.Epoch = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Round = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.LatestAnswer = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LatestTimestamp = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes16 configDigest, uint32 epoch, uint8 round, int192 latestAnswer, uint64 latestTimestamp)
func (_OffchainAggregator *OffchainAggregatorSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _OffchainAggregator.Contract.LatestTransmissionDetails(&_OffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes16 configDigest, uint32 epoch, uint8 round, int192 latestAnswer, uint64 latestTimestamp)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _OffchainAggregator.Contract.LatestTransmissionDetails(&_OffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OffchainAggregator.Contract.LinkAvailableForPayment(&_OffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OffchainAggregator.Contract.LinkAvailableForPayment(&_OffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorSession) MaxAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MaxAnswer(&_OffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MaxAnswer(&_OffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorSession) MinAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MinAnswer(&_OffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MinAnswer(&_OffchainAggregator.CallOpts)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_OffchainAggregator *OffchainAggregatorCaller) OracleObservationCount(opts *bind.CallOpts, _signerOrTransmitter common.Address) (uint16, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "oracleObservationCount", _signerOrTransmitter)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_OffchainAggregator *OffchainAggregatorSession) OracleObservationCount(_signerOrTransmitter common.Address) (uint16, error) {
	return _OffchainAggregator.Contract.OracleObservationCount(&_OffchainAggregator.CallOpts, _signerOrTransmitter)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_OffchainAggregator *OffchainAggregatorCallerSession) OracleObservationCount(_signerOrTransmitter common.Address) (uint16, error) {
	return _OffchainAggregator.Contract.OracleObservationCount(&_OffchainAggregator.CallOpts, _signerOrTransmitter)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) OwedPayment(opts *bind.CallOpts, _transmitter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "owedPayment", _transmitter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) OwedPayment(_transmitter common.Address) (*big.Int, error) {
	return _OffchainAggregator.Contract.OwedPayment(&_OffchainAggregator.CallOpts, _transmitter)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) OwedPayment(_transmitter common.Address) (*big.Int, error) {
	return _OffchainAggregator.Contract.OwedPayment(&_OffchainAggregator.CallOpts, _transmitter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) Owner() (common.Address, error) {
	return _OffchainAggregator.Contract.Owner(&_OffchainAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Owner() (common.Address, error) {
	return _OffchainAggregator.Contract.Owner(&_OffchainAggregator.CallOpts)
}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) RequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "requesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) RequesterAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.RequesterAccessController(&_OffchainAggregator.CallOpts)
}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) RequesterAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.RequesterAccessController(&_OffchainAggregator.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorSession) Transmitters() ([]common.Address, error) {
	return _OffchainAggregator.Contract.Transmitters(&_OffchainAggregator.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorCallerSession) Transmitters() ([]common.Address, error) {
	return _OffchainAggregator.Contract.Transmitters(&_OffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorSession) TypeAndVersion() (string, error) {
	return _OffchainAggregator.Contract.TypeAndVersion(&_OffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorCallerSession) TypeAndVersion() (string, error) {
	return _OffchainAggregator.Contract.TypeAndVersion(&_OffchainAggregator.CallOpts)
}

// ValidatorConfig is a free data retrieval call binding the contract method 0x8e0566de.
//
// Solidity: function validatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorCaller) ValidatorConfig(opts *bind.CallOpts) (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "validatorConfig")

	outstruct := new(struct {
		Validator common.Address
		GasLimit  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// ValidatorConfig is a free data retrieval call binding the contract method 0x8e0566de.
//
// Solidity: function validatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorSession) ValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _OffchainAggregator.Contract.ValidatorConfig(&_OffchainAggregator.CallOpts)
}

// ValidatorConfig is a free data retrieval call binding the contract method 0x8e0566de.
//
// Solidity: function validatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorCallerSession) ValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _OffchainAggregator.Contract.ValidatorConfig(&_OffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) Version() (*big.Int, error) {
	return _OffchainAggregator.Contract.Version(&_OffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Version() (*big.Int, error) {
	return _OffchainAggregator.Contract.Version(&_OffchainAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptOwnership(&_OffchainAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptOwnership(&_OffchainAggregator.TransactOpts)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, _transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "acceptPayeeship", _transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorSession) AcceptPayeeship(_transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptPayeeship(&_OffchainAggregator.TransactOpts, _transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) AcceptPayeeship(_transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptPayeeship(&_OffchainAggregator.TransactOpts, _transmitter)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RequestNewRound(&_OffchainAggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RequestNewRound(&_OffchainAggregator.TransactOpts)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetBilling(opts *bind.TransactOpts, _maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setBilling", _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetBilling(_maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBilling(&_OffchainAggregator.TransactOpts, _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetBilling(_maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBilling(&_OffchainAggregator.TransactOpts, _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBillingAccessController(&_OffchainAggregator.TransactOpts, _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBillingAccessController(&_OffchainAggregator.TransactOpts, _billingAccessController)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setConfig", _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetConfig(&_OffchainAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetConfig(&_OffchainAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, _linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setLinkToken", _linkToken, _recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetLinkToken(_linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLinkToken(&_OffchainAggregator.TransactOpts, _linkToken, _recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetLinkToken(_linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLinkToken(&_OffchainAggregator.TransactOpts, _linkToken, _recipient)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetLock(opts *bind.TransactOpts, isLock bool) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setLock", isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLock(&_OffchainAggregator.TransactOpts, isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLock(&_OffchainAggregator.TransactOpts, isLock)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetPayees(opts *bind.TransactOpts, _transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setPayees", _transmitters, _payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetPayees(_transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetPayees(&_OffchainAggregator.TransactOpts, _transmitters, _payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetPayees(_transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetPayees(&_OffchainAggregator.TransactOpts, _transmitters, _payees)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, _requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setRequesterAccessController", _requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetRequesterAccessController(_requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetRequesterAccessController(&_OffchainAggregator.TransactOpts, _requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetRequesterAccessController(_requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetRequesterAccessController(&_OffchainAggregator.TransactOpts, _requesterAccessController)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, _newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setValidatorConfig", _newValidator, _newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetValidatorConfig(_newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetValidatorConfig(&_OffchainAggregator.TransactOpts, _newValidator, _newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetValidatorConfig(_newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetValidatorConfig(&_OffchainAggregator.TransactOpts, _newValidator, _newGasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_OffchainAggregator *OffchainAggregatorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferOwnership(&_OffchainAggregator.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferOwnership(&_OffchainAggregator.TransactOpts, _to)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, _transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transferPayeeship", _transmitter, _proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_OffchainAggregator *OffchainAggregatorSession) TransferPayeeship(_transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferPayeeship(&_OffchainAggregator.TransactOpts, _transmitter, _proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) TransferPayeeship(_transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferPayeeship(&_OffchainAggregator.TransactOpts, _transmitter, _proposed)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) Transmit(opts *bind.TransactOpts, _report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transmit", _report, _rs, _ss, _rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorSession) Transmit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.Transmit(&_OffchainAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) Transmit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.Transmit(&_OffchainAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_OffchainAggregator *OffchainAggregatorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawFunds(&_OffchainAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawFunds(&_OffchainAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, _transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "withdrawPayment", _transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorSession) WithdrawPayment(_transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawPayment(&_OffchainAggregator.TransactOpts, _transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) WithdrawPayment(_transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawPayment(&_OffchainAggregator.TransactOpts, _transmitter)
}

// OffchainAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the OffchainAggregator contract.
type OffchainAggregatorAnswerUpdatedIterator struct {
	Event *OffchainAggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorAnswerUpdated)
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
		it.Event = new(OffchainAggregatorAnswerUpdated)
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
func (it *OffchainAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorAnswerUpdated represents a AnswerUpdated event raised by the OffchainAggregator contract.
type OffchainAggregatorAnswerUpdated struct {
	CurrentBatchId *big.Int
	RoundId        *big.Int
	UpdatedAt      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093.
//
// Solidity: event AnswerUpdated(uint256 indexed currentBatchId, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, currentBatchId []*big.Int, roundId []*big.Int) (*OffchainAggregatorAnswerUpdatedIterator, error) {

	var currentBatchIdRule []interface{}
	for _, currentBatchIdItem := range currentBatchId {
		currentBatchIdRule = append(currentBatchIdRule, currentBatchIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentBatchIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorAnswerUpdatedIterator{contract: _OffchainAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093.
//
// Solidity: event AnswerUpdated(uint256 indexed currentBatchId, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorAnswerUpdated, currentBatchId []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentBatchIdRule []interface{}
	for _, currentBatchIdItem := range currentBatchId {
		currentBatchIdRule = append(currentBatchIdRule, currentBatchIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentBatchIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorAnswerUpdated)
				if err := _OffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093.
//
// Solidity: event AnswerUpdated(uint256 indexed currentBatchId, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*OffchainAggregatorAnswerUpdated, error) {
	event := new(OffchainAggregatorAnswerUpdated)
	if err := _OffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorBillingAccessControllerSetIterator is returned from FilterBillingAccessControllerSet and is used to iterate over the raw logs and unpacked data for BillingAccessControllerSet events raised by the OffchainAggregator contract.
type OffchainAggregatorBillingAccessControllerSetIterator struct {
	Event *OffchainAggregatorBillingAccessControllerSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorBillingAccessControllerSet)
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
		it.Event = new(OffchainAggregatorBillingAccessControllerSet)
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
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorBillingAccessControllerSet represents a BillingAccessControllerSet event raised by the OffchainAggregator contract.
type OffchainAggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBillingAccessControllerSet is a free log retrieval operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*OffchainAggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorBillingAccessControllerSetIterator{contract: _OffchainAggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchBillingAccessControllerSet is a free log subscription operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorBillingAccessControllerSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

// ParseBillingAccessControllerSet is a log parse operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*OffchainAggregatorBillingAccessControllerSet, error) {
	event := new(OffchainAggregatorBillingAccessControllerSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorBillingSetIterator is returned from FilterBillingSet and is used to iterate over the raw logs and unpacked data for BillingSet events raised by the OffchainAggregator contract.
type OffchainAggregatorBillingSetIterator struct {
	Event *OffchainAggregatorBillingSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorBillingSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorBillingSet)
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
		it.Event = new(OffchainAggregatorBillingSet)
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
func (it *OffchainAggregatorBillingSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorBillingSet represents a BillingSet event raised by the OffchainAggregator contract.
type OffchainAggregatorBillingSet struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBillingSet is a free log retrieval operation binding the contract event 0xd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6.
//
// Solidity: event BillingSet(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*OffchainAggregatorBillingSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorBillingSetIterator{contract: _OffchainAggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

// WatchBillingSet is a free log subscription operation binding the contract event 0xd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6.
//
// Solidity: event BillingSet(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorBillingSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

// ParseBillingSet is a log parse operation binding the contract event 0xd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6.
//
// Solidity: event BillingSet(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseBillingSet(log types.Log) (*OffchainAggregatorBillingSet, error) {
	event := new(OffchainAggregatorBillingSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the OffchainAggregator contract.
type OffchainAggregatorConfigSetIterator struct {
	Event *OffchainAggregatorConfigSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorConfigSet)
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
		it.Event = new(OffchainAggregatorConfigSet)
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
func (it *OffchainAggregatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorConfigSet represents a ConfigSet event raised by the OffchainAggregator contract.
type OffchainAggregatorConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	Threshold                 uint8
	EncodedConfigVersion      uint64
	Encoded                   []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb9.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, uint64 configCount, address[] signers, address[] transmitters, uint8 threshold, uint64 encodedConfigVersion, bytes encoded)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffchainAggregatorConfigSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorConfigSetIterator{contract: _OffchainAggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb9.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, uint64 configCount, address[] signers, address[] transmitters, uint8 threshold, uint64 encodedConfigVersion, bytes encoded)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorConfigSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

// ParseConfigSet is a log parse operation binding the contract event 0x25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb9.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, uint64 configCount, address[] signers, address[] transmitters, uint8 threshold, uint64 encodedConfigVersion, bytes encoded)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseConfigSet(log types.Log) (*OffchainAggregatorConfigSet, error) {
	event := new(OffchainAggregatorConfigSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorLinkTokenSetIterator is returned from FilterLinkTokenSet and is used to iterate over the raw logs and unpacked data for LinkTokenSet events raised by the OffchainAggregator contract.
type OffchainAggregatorLinkTokenSetIterator struct {
	Event *OffchainAggregatorLinkTokenSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorLinkTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorLinkTokenSet)
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
		it.Event = new(OffchainAggregatorLinkTokenSet)
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
func (it *OffchainAggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorLinkTokenSet represents a LinkTokenSet event raised by the OffchainAggregator contract.
type OffchainAggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLinkTokenSet is a free log retrieval operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed _oldLinkToken, address indexed _newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, _oldLinkToken []common.Address, _newLinkToken []common.Address) (*OffchainAggregatorLinkTokenSetIterator, error) {

	var _oldLinkTokenRule []interface{}
	for _, _oldLinkTokenItem := range _oldLinkToken {
		_oldLinkTokenRule = append(_oldLinkTokenRule, _oldLinkTokenItem)
	}
	var _newLinkTokenRule []interface{}
	for _, _newLinkTokenItem := range _newLinkToken {
		_newLinkTokenRule = append(_newLinkTokenRule, _newLinkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "LinkTokenSet", _oldLinkTokenRule, _newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorLinkTokenSetIterator{contract: _OffchainAggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

// WatchLinkTokenSet is a free log subscription operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed _oldLinkToken, address indexed _newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorLinkTokenSet, _oldLinkToken []common.Address, _newLinkToken []common.Address) (event.Subscription, error) {

	var _oldLinkTokenRule []interface{}
	for _, _oldLinkTokenItem := range _oldLinkToken {
		_oldLinkTokenRule = append(_oldLinkTokenRule, _oldLinkTokenItem)
	}
	var _newLinkTokenRule []interface{}
	for _, _newLinkTokenItem := range _newLinkToken {
		_newLinkTokenRule = append(_newLinkTokenRule, _newLinkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "LinkTokenSet", _oldLinkTokenRule, _newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorLinkTokenSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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

// ParseLinkTokenSet is a log parse operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed _oldLinkToken, address indexed _newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseLinkTokenSet(log types.Log) (*OffchainAggregatorLinkTokenSet, error) {
	event := new(OffchainAggregatorLinkTokenSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the OffchainAggregator contract.
type OffchainAggregatorNewRoundIterator struct {
	Event *OffchainAggregatorNewRound // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorNewRound)
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
		it.Event = new(OffchainAggregatorNewRound)
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
func (it *OffchainAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorNewRound represents a NewRound event raised by the OffchainAggregator contract.
type OffchainAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OffchainAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorNewRoundIterator{contract: _OffchainAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorNewRound)
				if err := _OffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseNewRound(log types.Log) (*OffchainAggregatorNewRound, error) {
	event := new(OffchainAggregatorNewRound)
	if err := _OffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorNewTransmissionIterator is returned from FilterNewTransmission and is used to iterate over the raw logs and unpacked data for NewTransmission events raised by the OffchainAggregator contract.
type OffchainAggregatorNewTransmissionIterator struct {
	Event *OffchainAggregatorNewTransmission // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorNewTransmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorNewTransmission)
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
		it.Event = new(OffchainAggregatorNewTransmission)
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
func (it *OffchainAggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorNewTransmission represents a NewTransmission event raised by the OffchainAggregator contract.
type OffchainAggregatorNewTransmission struct {
	AggregatorRoundId uint32
	ProjectId         string
	Transmitter       common.Address
	Observers         []byte
	RawReportContext  [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewTransmission is a free log retrieval operation binding the contract event 0x62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db1.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, string projectId, address transmitter, bytes observers, bytes32 rawReportContext)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*OffchainAggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorNewTransmissionIterator{contract: _OffchainAggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

// WatchNewTransmission is a free log subscription operation binding the contract event 0x62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db1.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, string projectId, address transmitter, bytes observers, bytes32 rawReportContext)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorNewTransmission)
				if err := _OffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

// ParseNewTransmission is a log parse operation binding the contract event 0x62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db1.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, string projectId, address transmitter, bytes observers, bytes32 rawReportContext)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseNewTransmission(log types.Log) (*OffchainAggregatorNewTransmission, error) {
	event := new(OffchainAggregatorNewTransmission)
	if err := _OffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOraclePaidIterator is returned from FilterOraclePaid and is used to iterate over the raw logs and unpacked data for OraclePaid events raised by the OffchainAggregator contract.
type OffchainAggregatorOraclePaidIterator struct {
	Event *OffchainAggregatorOraclePaid // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorOraclePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOraclePaid)
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
		it.Event = new(OffchainAggregatorOraclePaid)
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
func (it *OffchainAggregatorOraclePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOraclePaid represents a OraclePaid event raised by the OffchainAggregator contract.
type OffchainAggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePaid is a free log retrieval operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*OffchainAggregatorOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOraclePaidIterator{contract: _OffchainAggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

// WatchOraclePaid is a free log subscription operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOraclePaid)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

// ParseOraclePaid is a log parse operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOraclePaid(log types.Log) (*OffchainAggregatorOraclePaid, error) {
	event := new(OffchainAggregatorOraclePaid)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferRequestedIterator struct {
	Event *OffchainAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOwnershipTransferRequested)
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
		it.Event = new(OffchainAggregatorOwnershipTransferRequested)
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
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffchainAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOwnershipTransferRequestedIterator{contract: _OffchainAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOwnershipTransferRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffchainAggregatorOwnershipTransferRequested, error) {
	event := new(OffchainAggregatorOwnershipTransferRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferredIterator struct {
	Event *OffchainAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOwnershipTransferred)
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
		it.Event = new(OffchainAggregatorOwnershipTransferred)
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
func (it *OffchainAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffchainAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOwnershipTransferredIterator{contract: _OffchainAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOwnershipTransferred)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*OffchainAggregatorOwnershipTransferred, error) {
	event := new(OffchainAggregatorOwnershipTransferred)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorPayeeshipTransferRequestedIterator is returned from FilterPayeeshipTransferRequested and is used to iterate over the raw logs and unpacked data for PayeeshipTransferRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferRequestedIterator struct {
	Event *OffchainAggregatorPayeeshipTransferRequested // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorPayeeshipTransferRequested)
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
		it.Event = new(OffchainAggregatorPayeeshipTransferRequested)
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
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorPayeeshipTransferRequested represents a PayeeshipTransferRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferRequested is a free log retrieval operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*OffchainAggregatorPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorPayeeshipTransferRequestedIterator{contract: _OffchainAggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferRequested is a free log subscription operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorPayeeshipTransferRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

// ParsePayeeshipTransferRequested is a log parse operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*OffchainAggregatorPayeeshipTransferRequested, error) {
	event := new(OffchainAggregatorPayeeshipTransferRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorPayeeshipTransferredIterator is returned from FilterPayeeshipTransferred and is used to iterate over the raw logs and unpacked data for PayeeshipTransferred events raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferredIterator struct {
	Event *OffchainAggregatorPayeeshipTransferred // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorPayeeshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorPayeeshipTransferred)
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
		it.Event = new(OffchainAggregatorPayeeshipTransferred)
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
func (it *OffchainAggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorPayeeshipTransferred represents a PayeeshipTransferred event raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferred is a free log retrieval operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*OffchainAggregatorPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorPayeeshipTransferredIterator{contract: _OffchainAggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferred is a free log subscription operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorPayeeshipTransferred)
				if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

// ParsePayeeshipTransferred is a log parse operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*OffchainAggregatorPayeeshipTransferred, error) {
	event := new(OffchainAggregatorPayeeshipTransferred)
	if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorRequesterAccessControllerSetIterator is returned from FilterRequesterAccessControllerSet and is used to iterate over the raw logs and unpacked data for RequesterAccessControllerSet events raised by the OffchainAggregator contract.
type OffchainAggregatorRequesterAccessControllerSetIterator struct {
	Event *OffchainAggregatorRequesterAccessControllerSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorRequesterAccessControllerSet)
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
		it.Event = new(OffchainAggregatorRequesterAccessControllerSet)
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
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorRequesterAccessControllerSet represents a RequesterAccessControllerSet event raised by the OffchainAggregator contract.
type OffchainAggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequesterAccessControllerSet is a free log retrieval operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*OffchainAggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorRequesterAccessControllerSetIterator{contract: _OffchainAggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchRequesterAccessControllerSet is a free log subscription operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorRequesterAccessControllerSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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

// ParseRequesterAccessControllerSet is a log parse operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*OffchainAggregatorRequesterAccessControllerSet, error) {
	event := new(OffchainAggregatorRequesterAccessControllerSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorRoundRequestedIterator is returned from FilterRoundRequested and is used to iterate over the raw logs and unpacked data for RoundRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorRoundRequestedIterator struct {
	Event *OffchainAggregatorRoundRequested // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorRoundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorRoundRequested)
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
		it.Event = new(OffchainAggregatorRoundRequested)
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
func (it *OffchainAggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorRoundRequested represents a RoundRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [16]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRoundRequested is a free log retrieval operation binding the contract event 0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037.
//
// Solidity: event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*OffchainAggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorRoundRequestedIterator{contract: _OffchainAggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

// WatchRoundRequested is a free log subscription operation binding the contract event 0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037.
//
// Solidity: event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorRoundRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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

// ParseRoundRequested is a log parse operation binding the contract event 0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037.
//
// Solidity: event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseRoundRequested(log types.Log) (*OffchainAggregatorRoundRequested, error) {
	event := new(OffchainAggregatorRoundRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorValidatorConfigSetIterator is returned from FilterValidatorConfigSet and is used to iterate over the raw logs and unpacked data for ValidatorConfigSet events raised by the OffchainAggregator contract.
type OffchainAggregatorValidatorConfigSetIterator struct {
	Event *OffchainAggregatorValidatorConfigSet // Event containing the contract specifics and raw log

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
func (it *OffchainAggregatorValidatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorValidatorConfigSet)
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
		it.Event = new(OffchainAggregatorValidatorConfigSet)
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
func (it *OffchainAggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorValidatorConfigSet represents a ValidatorConfigSet event raised by the OffchainAggregator contract.
type OffchainAggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorConfigSet is a free log retrieval operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*OffchainAggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorValidatorConfigSetIterator{contract: _OffchainAggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

// WatchValidatorConfigSet is a free log subscription operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorValidatorConfigSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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

// ParseValidatorConfigSet is a log parse operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*OffchainAggregatorValidatorConfigSet, error) {
	event := new(OffchainAggregatorValidatorConfigSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
