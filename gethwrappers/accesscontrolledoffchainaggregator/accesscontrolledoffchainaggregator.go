// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accesscontrolledoffchainaggregator

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

// AccessControlledOffchainAggregatorMetaData contains all meta data concerning the AccessControlledOffchainAggregator contract.
var AccessControlledOffchainAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerTransmission\",\"type\":\"uint32\"},{\"internalType\":\"contractBondTokenInterface\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"int192\",\"name\":\"_minAnswer\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"_maxAnswer\",\"type\":\"int192\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"},{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAccessController\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentBatchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPrice\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPrice\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"microLinkPerEth\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"linkGweiPerObservation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"linkGweiPerTransmission\",\"type\":\"uint32\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"encodedConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encoded\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"_oldLinkToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"_newLinkToken\",\"type\":\"address\"}],\"name\":\"LinkTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"observers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rawReportContext\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractBondTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"RequesterAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"name\":\"RoundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"previousValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousGasLimit\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"currentValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"currentGasLimit\",\"type\":\"uint32\"}],\"name\":\"ValidatorConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"billingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkGweiPerTransmission\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinkToken\",\"outputs\":[{\"internalType\":\"contractBondTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"}],\"name\":\"latestMerkleRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkelAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTransmissionDetails\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"configDigest\",\"type\":\"bytes16\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"},{\"internalType\":\"int192\",\"name\":\"latestAnswer\",\"type\":\"int192\"},{\"internalType\":\"uint64\",\"name\":\"latestTimestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerOrTransmitter\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requesterAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maximumGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_reasonableGasPrice\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_microLinkPerEth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerObservation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_linkGweiPerTransmission\",\"type\":\"uint32\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_encodedConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_encoded\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBondTokenInterface\",\"name\":\"_linkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"setLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_requesterAccessController\",\"type\":\"address\"}],\"name\":\"setRequesterAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"_newValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"}],\"name\":\"setValidatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorConfig\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e0604052346200086a5762004fb280380390816200001e81620008af565b9182396101808183810103126200086a576200003a81620008d5565b6200004860208301620008d5565b6200005660408401620008d5565b6200006460608501620008d5565b926200007360808601620008d5565b60a086015194906001600160a01b03861686036200086a576200009960c08801620008e7565b97620000a860e08901620008e7565b98620000b86101008a01620008f6565b98620000c86101208201620008f6565b6101408201519160ff831683036200086a57610160810151906001600160401b0382116200086a57858101601f8383010112156200086a5780820151916001600160401b0383116200079e5762000129601f8401601f1916602001620008af565b9683885282016020848385010101116200086a575f5b838110620008525750505f91860160200191909152506040519b6001600160401b036101808e01908111908e11176200079e578c60c060ff95610180830160405263ffffffff8b16835263ffffffff8d16602084015263ffffffff8c16604084015263ffffffff8a16606084015263ffffffff891660808401528d60018060a01b031660a084015260170b91015260170b60e08d015260018060a01b038b166101008d015260018060a01b03166101208c0152166101408a01526101608901523360018060a01b03195f5416175f556040519460a086019286841060018060401b038511176200079e5763ffffffff60a096818096947fd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b69a8296604052828a16815282821660208201528284166040820152828616606082015260808389169101526002548260801b8860801b1690838b16906001808e1b0319161767ffffffff000000008360201b16176bffffffff00000000000000008560401b16176fffffffff0000000000000000000000008760601b1617176002558160405199168952166020880152166040860152166060840152166080820152a1600380546001600160a01b0319166001600160a01b03929092169182179055604051905f7f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a8180a36004546001600160a01b038082169391908216849003620007fb575b505050506200036b6200088e565b6103e03682376200037b6200088e565b906103e03683375f5b601f60ff82161015620003df57600180620003a360ff8416856200090b565b52620003b360ff8316856200090b565b5260ff80821614620003cb5760ff1660010162000384565b634e487b7160e01b5f52601160045260245ffd5b505f5b600181101562000430575f805b601081106200040757506005820155600101620003e2565b835190939160019160209161ffff600488901b81811b199092169216901b1792019301620003ef565b505f90815b600f8110620007c85750506006555f5b601f8110620007b25761014083015160ff1660c05261016083015180519084906001600160401b0383116200079e57603054600181811c9116801562000793575b60208210146200077f57601f81116200071b575b50602090601f8411600114620006995773f4888aed11bfa9adcfa25b42e11cb6e064a354b893915f91836200068d575b50508160011b915f199060031b1c1916176030555b6101208101515f549160e0916200052a906001600160a01b03808616331491166200050a8262000931565b602f54906001600160a01b03821680820362000635575b50505062000931565b620005346200086e565b602e546001600160a01b03811680835260a082901c63ffffffff1660208401818152929115801591906200062a575b50620005ba575b50505060c0810151601790810b608052910151900b60a0526001600160a01b031916175f5560405161463390816200097f8239608051816132ac015260a05181612a8d015260c0518161320b0152f35b604063ffffffff5f947fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca154193866020620005f26200086e565b82815201526001600160c01b031916602e55519351825191168152602081018590526001600160a01b0390931692a38480806200056a565b905015158862000563565b6001600160a01b0319929092168117602f55604080516001600160a01b0393841681529290911660208301527f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae63491a186808062000521565b015190508480620004ca565b9060305f5260205f20915f5b601f198616811062000702575091849160019373f4888aed11bfa9adcfa25b42e11cb6e064a354b896601f19811610620006e9575b505050811b01603055620004df565b01515f1960f88460031b161c19169055848080620006da565b91926020600181928685015181550194019201620006a5565b60305f527f6ff97a59c90d62cc7236ba3a37cd85351bf564556780cf8c1157a220f31f0cbb601f850160051c81016020861062000777575b601f830160051c820181106200076b5750506200049a565b5f815560010162000753565b508062000753565b634e487b7160e01b5f52602260045260245ffd5b90607f169062000486565b634e487b7160e01b5f52604160045260245ffd5b6001906020835193019281600901550162000445565b90916020620007f160019261ffff8651169085851b60031b9161ffff809116831b921b19161790565b9301910162000435565b6001600160a01b0319166001600160a01b039182169081176004559216815260208101919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d4891290604090a15f8080806200035d565b806020809284860101015182828b010152016200013f565b5f80fd5b60408051919082016001600160401b038111838210176200079e57604052565b604051906103e082016001600160401b038111838210176200079e57604052565b6040519190601f01601f191682016001600160401b038111838210176200079e57604052565b519063ffffffff821682036200086a57565b51908160170b82036200086a57565b51906001600160a01b03821682036200086a57565b90601f8110156200091d5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b156200093957565b60405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606490fdfe6101206040526004361015610012575f80fd5b5f3560e01c80630eafb25b1461331a578063181f5a77146132d257806322adbc7814613295578063299372681461322f578063313ce567146131f25780634fb17470146131ae57806350d25bcd1461319357806354fd4d5014613178578063585aa7de14612b1e578063619d519414612ad9578063668a0f0214612ab357806370da2f6714612a7657806370efdf2d14612a505780637284e4161461297157806379ba5097146128c7578063814118341461286457806381ff7048146128205780638205bf6a146127e25780638ac28d5a146127765780638da5cb5b146127515780638e0566de1461271557806398e5b12a146125dc578063996e8298146125b65780639a6fc8f5146124115780639c849b30146122935780639e3ceeab146121ea578063a4e2d634146121c5578063b121e14714612109578063b5ab58dc146120eb578063b633620c146120c5578063bd82470614611ec3578063c107532914611ca2578063c980753914610841578063d09dc3391461077f578063d4c520d114610718578063e4902f82146106e9578063e5fe457714610647578063e76d516814610621578063eb45716314610532578063eb5dcd6c14610442578063f2fde38b146103d7578063fbffd2c11461032d5763feaf968c146101f3575f80fd5b34610329575f3660031901126103295763ffffffff80602b5460b01c1690815f52602090602c825260405f20906040519161022d83613417565b60405161023981613432565b61024282613947565b8152600180830154936001600160401b03948581168885015260401c166040830152600290818401918254906102778261377d565b936102856040519586613484565b8285525f9081528981208a8087015b8584106102f6578d828d8d60038e8e8e6060820152845201541691829101526102f26040519282849382859293608092959460a085019669ffffffffffffffffffff8094168652600160208701526040860152606085015216910152565b0390f35b8691859160405161030681613417565b61030f86613947565b815284860154838201528152019201920191908b90610294565b5f80fd5b34610329576020366003190112610329576001600160a01b0360043581811680820361032957610361835f541633146134d3565b6004549283169081810361037157005b73ffffffffffffffffffffffffffffffffffffffff1990931692909217600455604080516001600160a01b0393841681529190921660208201527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d4891291819081015b0390a1005b34610329576020366003190112610329576103f061333d565b6001600160a01b03805f5416916104088333146134d3565b1690816001600160a01b031960015416176001557fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12785f80a3005b346103295760403660031901126103295761045b61333d565b610463613353565b6001600160a01b0380921690815f5260076020528260405f2054163303610508578216918233146104de578290825f52600860205260405f20805490836001600160a01b0319831617905516036104b657005b33907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e383675f80a4005b60405162461bcd60e51b81526020600482015260026024820152610d4d60f21b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261353360f01b6044820152606490fd5b34610329576040366003190112610329576004356001600160a01b0380821680920361032957610560613404565b9061056f815f541633146134d3565b6105776136ec565b9183828451161480159061060a575b61058c57005b6020604091817fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541948451926105c084613417565b88845263ffffffff93848216958691015288602e549163ffffffff60a01b9060a01b16916001600160401b0360c01b161717602e55865116950151169082519182526020820152a3005b5063ffffffff806020850151169082161415610586565b34610329575f3660031901126103295760206001600160a01b0360035416604051908152f35b34610329575f366003190112610329573233036106bf5760a0602b5463ffffffff90818160b01c165f52602c60205260ff6001600160401b03600360405f2001541691604051936001600160801b03198260801b1685528160881c16602085015260801c166040830152600160608301526080820152f35b60405162461bcd60e51b8152602060048201526002602482015261303960f01b6044820152606490fd5b3461032957602036600319011261032957602061070c61070761333d565b614597565b61ffff60405191168152f35b34610329576020366003190112610329576004356001600160401b0381116103295761075361074d60a09236906004016133d7565b906139e6565b9269ffffffffffffffffffff604093929351951685526020850152604084015260608301526080820152f35b34610329575f36600319011261032957602460206001600160a01b0360035416604051928380926370a0823160e01b82523060048301525afa908115610836575f91610805575b506107cf61444a565b905f82820392128183128116918313901516176107f157602090604051908152f35b634e487b7160e01b5f52601160045260245ffd5b90506020813d821161082e575b8161081f60209383613484565b810103126103295751816107c6565b3d9150610812565b6040513d5f823e3d90fd5b34610329576080366003190112610329576004356001600160401b038111610329576108719036906004016133d7565b908160a0526080526024356001600160401b038111610329576108989036906004016133a7565b6044929192356001600160401b038111610329576108ba9036906004016133a7565b90915a9360e4018060e4116107f1578160051b828104602014831517156107f1576108e4916138f8565b8260051b838104602014841517156107f1576108ff916138f8565b3603611c7857604051926109128461344d565b60405161091e81613432565b5f81525f60208201525f60408201525f6060820152845260606020850152606060408501526060808501525f6080850152610957613730565b84526080519460a051860196606087890312610329576040870135976001600160401b03891161032957601f898901011215610329576109998888013561377d565b6109a66040519182613484565b888801358152602081019860a051906080519182016020828c013560051b60408501358501010111610329578989898960208f9897966040810135010160e0525b60805160208686013560051b60408301358301010160e0511015610db9576001600160401b0360e051351161032957608090601f199060e0513560408201358201019060a051010301126103295760405199610a428b613432565b6080516001600160401b03602060e0513560408401358401010135116103295760a05160e051610a889291820191604081013501903501602081810135909101016137ca565b8b5260805160e051604080830135830191359190910101356001600160401b03811690036103295760e0805160408381013590930190350182013560208e01526080519051909182013582019035016060013563ffffffff811690036103295760609060e051359060408101350101013560408c0152608051610100526001600160401b03608060e0513560406101005101356101005101010135116103295760a0516101005160e0519181016040820135909101913591909101608081013501603f0112156103295760e05161010051610b779160408201359091019035016080810135016020013561377d565b99610b856040519b8c613484565b8a60c0526020608060e051356040610100510135610100510101013560e0513560406101005101356101005101010101358b5260208b0160c05260a05160805190810160406020608060e0513583610100510135610100510101013560e051358361010051013561010051010101013560051b608060e0513583610100510135610100510101013560e05135838601358601010101011161032957604090608060e051358361010051013561010051010101359060e051359083810135010101015b6080805160e0516101005191926020604080850135909401923592830191820135918201013560051b83850135850192909201010101821015610d97576001600160401b0382351161032957604090603f19908335608060e0513585610100510135610100510101013560e051358584013584010101019060a051010301126103295760405190610cd782613417565b608051916001600160401b0360408335608060e0513583610100510135610100510101013560e051358388013588010101010135116103295760a05160e05161010051602095610d50938101926040808301359092019035908101838301359093010160800135919091018535018082013501016137ca565b81526060608051833590608060e05135604061010051013561010051010101359060e051359060408101350101010101358382015260c051528160c0510160c05201610c47565b505060608c019a909a5299885260e080516020908101909152909701966109e7565b508989888860408701526080513560808701526001600160801b0319806080513560581b16908751511603611c4e5764ffffffffff608087015116928364ffffffffff602089510151161015611c245760ff60408851015116851115611bfa57601f8511611bd057848603611ba657601f60408801515111611b7c5760408701515160408851015160011b906101fe60fe8316921682036107f1571115611b5257610e63856137e8565b60608801525f5b8560ff82161015610ebd57602060ff82161015610ea957610ea49060643560ff82161a610e9e60ff831660608c015161382b565b5361381a565b610e6a565b634e487b7160e01b5f52603260045260245ffd5b5086939495610ed06040860151516137e8565b602086015260405196610ee288613468565b6103e03689375f5b60408701515160ff82161015610f8857602060ff82161015610ea957610f1b6020608051013560ff83161a8a61383c565b51610f5e57610f59906001610f3b6020608051013560ff84161a8c61383c565b526020608051013560ff82161a610e9e60ff831660208b015161382b565b610eea565b60405162461bcd60e51b8152602060048201526002602482015261062760f31b6044820152606490fd5b5086335f52602860205260405f20610fbb60ff60405192610fa884613417565b54818116845260081c16602083016135c8565b60208101516003811015611780576002149081611b2b575b5015611b0157602087510152610fee3660a051608051613794565b60208151910120936040519461100386613468565b6103e03687375f602060405161101881613417565b82815201525f5b8381106119d257888863ffffffff606083510151165f52602c60205260405f20916110536040820151805160011c9061384d565b51606082510163ffffffff61106a81835116613861565b1690525f5b606082015163ffffffff81511663ffffffff831690811015611211576110949161384d565b516002860154680100000000000000008110156111ea578060016110c19201600289015560028801613876565b9190916111fe5780518051906001600160401b0382116111ea576110e5845461388f565b601f81116111af575b50602090601f831160011461113f57928260019360209361112f9897965f92611134575b50505f19600383901b1c191690841b1784555b0151910155613861565b61106f565b015190508c80611112565b90845f5260205f20915f5b601f198516811061119757508360209361112f98979693600196938794601f1981161061117f575b505050811b018455611125565b01515f1960f88460031b161c191690558c8080611172565b9192602060018192868501518155019401920161114a565b6111da90855f5260205f20601f850160051c810191602086106111e0575b601f0160051c01906138c7565b896110ee565b90915081906111cd565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f525f60045260245ffd5b505050602081015160018501805467ffffffffffffffff19166001600160401b03928316179055815180519182116111ea5761124d865461388f565b601f81116119a2575b50602090601f83116001146119395760409392915f918361192e575b50508160011b915f199060031b1c19161785555b01516bffffffff000000000000000060018501549160401b166bffffffff0000000000000000198216176001850155600384016001600160401b0342166001600160401b031982541617905563ffffffff6060835101511660208301519460808401519060405196608088525f918054906113008261388f565b918260808c01526001811690815f1461190757506001146118cb575b5050967f62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db19261135c82936020999a338b8601528482036040860152613369565b9060608301520390a25f63ffffffff606084510151167f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac6027185604051428152a363ffffffff60608351015116907f5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093846001600160401b03604051934285521692a38051805160801c90602b549164ffffffffff60801b8583015160801b1690606060ff60a81b604085015160a81b1693015163ffffffff60b01b9060b01b169365ffffffffffff60d01b1617171717602b5561143c63ffffffff84106138dd565b0151335f52602860205260405f209061147060ff6040519361145d85613417565b54818116855260081c16602084016135c8565b611478613ce8565b91604051938460055f5b601f600f820110611806575061ffff6101c091548181168452818160101c166020850152818160201c166040850152818160301c166060850152818160401c166080850152818160501c1660a0850152818160601c1660c0850152818160701c1660e0850152818160801c16610100850152818160901c16610120850152818160a01c16610140850152818160b01c16610160850152818160c01c16610180850152818160d01c166101a085015260e01c1691015261154085613468565b6103e060405161154f81613468565b3690375f5b83518110156115ae57611567818561382b565b5160f81c9061ffff611579838961383c565b5116916001830183116107f1576115a361ffff61159b60016115a99601614607565b16918961383c565b52613596565b611554565b5083855f5b600181106117c557505f905f5b600f8110611794575050600655602083015160038110156117805760020361175657633b9aca0091611613833a0463ffffffff60208501511663ffffffff855116918091808210611737575b505061461b565b907f0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff361636036107f15763ffffffff5a911681811061170d5761165d611662923660041b9261351f565b6138f8565b9061179382018092116107f157611678916134c0565b828102908082048414901517156107f1576116b6620f4240916001600160801b036116a48183106138dd565b63ffffffff60408601511691166134c0565b0460ff84511690601f821015610ea9576116db60809163ffffffff93600901546138f8565b920151168281029281840414901517156107f15760ff916116fb916138f8565b915116601f811015610ea95760090155005b60405162461bcd60e51b81526020600482015260026024820152611a1b60f11b6044820152606490fd5b61174f9250816117469161351f565b60011c906138f8565b878061160c565b60405162461bcd60e51b8152602060048201526002602482015261035360f41b6044820152606490fd5b634e487b7160e01b5f52602160045260245ffd5b909160206117bc60019261ffff8651169085851b60031b9161ffff809116831b921b19161790565b930191016115c0565b5f805b601081106117de575060058201556001016115b3565b835190939160019160209161ffff600488901b81811b199092169216901b17920193016117c8565b90916001610200601092855461ffff80821683528082871c166020840152808260201c166040840152808260301c166060840152808260401c166080840152808260501c1660a0840152808260c0828260601c1681870152828260701c1660e0870152828260801c16610100870152828260901c16610120870152828260a01c16610140870152828260b01c166101608701521c16610180840152808260d01c166101a08401528160e01c166101c083015260f01c6101e08201520193019101611482565b5f90815260208120929350915b8183106118f0575050870160a001908261135c61131c565b8054838b0160a001526020909201916001016118d8565b60ff191660a0808d019190915292151560051b8b01909201935084915061135c905061131c565b015190508780611272565b90865f5260205f20915f5b601f198516811061198a575091839160019360409695601f19811610611972575b505050811b018555611286565b01515f1960f88460031b161c19169055878080611965565b91926020600181928685015181550194019201611944565b6119cc90875f5260205f20601f850160051c810191602086106111e057601f0160051c01906138c7565b86611256565b601b6119e28260608c015161382b565b5160f81c0160ff81116107f1575f6080602092611a008589896135a4565b35611a0c868b8d6135a4565b359060ff6040519389855216868401526040830152606082015282805260015afa15610836576001600160a01b035f51165f52602860205260405f209060405191611a5683613417565b5460ff81168352611a7160ff602085019260081c16826135c8565b51600381101561178057600103611ad757611a9060ff8351168961383c565b51611aad5760016115a360ff611aa89451168a61383c565b61101f565b60405162461bcd60e51b8152602060048201526002602482015261191960f11b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261323160f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261313960f01b6044820152606490fd5b6001600160a01b03915060ff611b4291511661352c565b90549060031b1c16331488610fd3565b60405162461bcd60e51b8152602060048201526002602482015261313760f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261189b60f11b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261313560f01b6044820152606490fd5b60405162461bcd60e51b81526020600482015260026024820152610c4d60f21b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261313360f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261189960f11b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261313160f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261031360f41b6044820152606490fd5b3461032957604036600319011261032957611cbb61333d565b6001600160a01b0390815f541633148015611e5a575b15611e3057611cde61444a565b91600354166040516370a0823160e01b81523060048201526020938482602481865afa918215610836575f92611e01575b50808210611dd75791611d33611d679492611d2e87956024359261351f565b61461b565b60405163a9059cbb60e01b81526001600160a01b0390931660048401526024830152909283919082905f9082906044820190565b03925af1908115610836575f91611daa575b5015611d8157005b6064906040519062461bcd60e51b825260048201526002602482015261343960f01b6044820152fd5b611dca9150823d8411611dd0575b611dc28183613484565b810190613718565b82611d79565b503d611db8565b60405162461bcd60e51b8152600481018690526002602482015261068760f31b6044820152606490fd5b9091508481813d8311611e29575b611e198183613484565b8101031261032957519085611d0f565b503d611e0f565b60405162461bcd60e51b8152602060048201526002602482015261343760f01b6044820152606490fd5b50816004541660206040518092630d629b5f60e31b8252336004830152604060248301528180611e8d6044820136613636565b03915afa908115610836575f91611ea5575b50611cd1565b611ebd915060203d8111611dd057611dc28183613484565b83611e9f565b346103295760a03660031901126103295763ffffffff6004358181169081810361032957611eef613404565b9160443584811690818103610329576064359086821693848303610329576084359388851690818603610329576001600160a01b038060045416905f5416331490811561205f575b5015612035577fd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6996103d297608092611f6e6141ad565b60405192611f7b8461344d565b8684528c16602084015260408301526060820152015260025473ffffffff000000000000000000000000000000008560801b16916fffffffff0000000000000000000000008560601b16916001600160a01b0319161767ffffffff000000008960201b16176bffffffff00000000000000008460401b16171717600255604051958695869391608093929695919660a086019763ffffffff94858094818094168a5216602089015216604087015216606085015216910152565b60405162461bcd60e51b81526020600482015260026024820152611a1960f11b6044820152606490fd5b905060206040518092630d629b5f60e31b825233600483015260406024830152818061208e6044820136613636565b03915afa908115610836575f916120a7575b508b611f37565b6120bf915060203d8111611dd057611dc28183613484565b8b6120a0565b346103295760203660031901126103295760206120e3600435613918565b604051908152f35b346103295760203660031901126103295760206120e3600435613905565b34610329576020366003190112610329576001600160a01b038061212b61333d565b16805f5260086020528160405f205416330361219b57805f52600760205260405f20918254926001600160a01b03199033828616179055600860205260405f20908154169055339216907f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b35f80a4005b60405162461bcd60e51b8152602060048201526002602482015261353560f01b6044820152606490fd5b34610329575f36600319011261032957602060ff602d5460401c166040519015158152f35b34610329576020366003190112610329576001600160a01b036004358181168082036103295761221e835f541633146134d3565b602f549283169081810361222e57005b73ffffffffffffffffffffffffffffffffffffffff1990931692909217602f55604080516001600160a01b0393841681529190921660208201527f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae63491819081016103d2565b34610329576040366003190112610329576001600160401b03600435818111610329576122c49036906004016133a7565b6024928335908111610329576122de9036906004016133a7565b90916001600160a01b03916122f7835f541633146134d3565b8082036123e8575f5b82811061230957005b61231c6123178285896135a4565b6135b4565b8461232b61231784868a6135a4565b9116805f526007906020918083528760405f20541692831580156123dd575b156123b557908291612384969594935f52528760405f20931692836001600160a01b0319825416179055828203612389575b505050613596565b612300565b7f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b35f80a488808061237c565b60649060028d6040519262461bcd60e51b84526004840152820152611a9960f11b6044820152fd5b50888516841461234a565b60405162461bcd60e51b81526020600482015260028188015261353160f01b6044820152606490fd5b3461032957602080600319360112610329576004359069ffffffffffffffffffff82168083036103295760405161244781613417565b600f81527f4e6f20646174612070726573656e7400000000000000000000000000000000008382015263ffffffff80921161259057508083165f52602c825260405f20906040519161249883613417565b6040516124a481613432565b6124ad82613947565b8152600180830154936001600160401b03948581168885015260401c166040830152600290818401918254906124e28261377d565b936124f06040519586613484565b8285525f9081528981208a8087015b85841061255d578d828d8d60038e8e8e6060820152845201541691829101526102f26040519282849382859293608092959460a085019669ffffffffffffffffffff8094168652600160208701526040860152606085015216910152565b8691859160405161256d81613417565b61257686613947565b815284860154838201528152019201920191908b906124ff565b826125b260405192839262461bcd60e51b845260048401526024830190613369565b0390fd5b34610329575f3660031901126103295760206001600160a01b0360045416604051908152f35b34610329575f366003190112610329576001600160a01b03805f541633149081156126ab575b501561268157602060ff612678612617613730565b60606001600160801b031982511691602b549260405190815263ffffffff958487809660881c168984015260801c1660408201527f3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037833392a20151166135d4565b16604051908152f35b60405162461bcd60e51b8152602060048201526002602482015261060760f31b6044820152606490fd5b9050602f541660206040518092630d629b5f60e31b82523360048301526040602483015281806126de6044820136613636565b03915afa908115610836575f916126f7575b5081612602565b61270f915060203d8111611dd057611dc28183613484565b816126f0565b34610329575f36600319011261032957604061272f6136ec565b63ffffffff60206001600160a01b038351169201511682519182526020820152f35b34610329575f3660031901126103295760206001600160a01b035f5416604051908152f35b346103295760203660031901126103295761278f61333d565b6001600160a01b038082165f52600760205260405f20541633036127b8576127b690613e68565b005b60405162461bcd60e51b8152602060048201526002602482015261343360f01b6044820152606490fd5b34610329575f3660031901126103295763ffffffff602b5460b01c165f52602c60205260206001600160401b03600360405f20015416604051908152f35b34610329575f366003190112610329576060602d54602b5460801b6040519163ffffffff90818116845260201c1660208301526001600160801b0319166040820152f35b34610329575f3660031901126103295761287c613676565b60405180916020808301818452825180915281604085019301915f5b8281106128a757505050500390f35b83516001600160a01b031685528695509381019392810192600101612898565b34610329575f366003190112610329576001546001600160a01b03808216330361292c575f54916001600160a01b03199033828516175f55166001553391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b60405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606490fd5b34610329575f366003190112610329576040515f6030546129918161388f565b80845290600190818116908115612a2957506001146129cf575b6102f2846129bb81860382613484565b604051918291602083526020830190613369565b60305f90815292507f6ff97a59c90d62cc7236ba3a37cd85351bf564556780cf8c1157a220f31f0cbb5b828410612a115750505081016020016129bb826129ab565b805460208587018101919091529093019281016129f9565b60ff191660208087019190915292151560051b850190920192506129bb91508390506129ab565b34610329575f3660031901126103295760206001600160a01b03602f5416604051908152f35b34610329575f3660031901126103295760206040517f000000000000000000000000000000000000000000000000000000000000000060170b8152f35b34610329575f36600319011261032957602063ffffffff602b5460b01c16604051908152f35b34610329576020366003190112610329576004358015158091036103295768ff0000000000000000602d549160401b169068ff0000000000000000191617602d555f80f35b346103295760a0366003190112610329576001600160401b0360043581811161032957612b4f9036906004016133a7565b60243583811161032957612b679036906004016133a7565b9360ff6044351660443503610329576064359080821682036103295760843581811161032957612b9b9036906004016133d7565b919093601f861161314e5760ff6044351615613124578786036130fa57600360443560ff1681810291820490036107f1578611156130d057612be86001600160a01b035f541633146134d3565b6029548015612ccf57805f198101116107f157805f1981011015610ea95760295f526001600160a01b03612c44817fcb7c14ce178f56e2e8d86ab33ebc0ae081ba8556a00cd122038841867181caab84015416925f190161352c565b90549060031b1c1690612c5682613e68565b5f5260286020525f60408120555f525f60408120556029548015612cbb575f1901612c98612c8382613561565b6001600160a01b0382549160031b1b19169055565b602955602a548015612cbb575f1901612cb3612c838261352c565b602a55612be8565b634e487b7160e01b5f52603160045260245ffd5b50875f5b878110612e5a5750602b5497602d549763ffffffff8916612cf3906135d4565b63ffffffff16998a4360201b67ffffffff00000000166001600160401b03198c161717602d5560405190602082013081528c60408401526060830160e090526101008301612d429085876135e9565b838c8c8b8b601f19958c87878303016080880152612d5f926135e9565b9160443560ff1660a08601521660c08401528c848483030160e0850152612d8592613656565b039081018452612d959084613484565b64ffffffffff60801b199251902060801c9060ff60a81b60443560a81b169075ff0000000000ffffffffffffffffffffffffffffffff1916171716602b55604051998a9960201c63ffffffff168a5260208a01526040890160e0905260e0890190612dff926135e9565b908782036060890152612e11926135e9565b9260443560ff1660808701521660a085015283820360c0850152612e3492613656565b037f25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb991a1005b6001600160a01b03612e70612317838b8d6135a4565b165f52602860205260ff60405f205460081c166003811015611780576130a657604051612e9c81613417565b60ff82168152896001600160a01b03612ec1612317858d6020870195600187526135a4565b165f52602860205260ff60405f209251168254915160038110156117805761ff009060081b169161ffff1916171790556001600160a01b03612f076123178385876135a4565b165f5260076020526001600160a01b0360405f2054161561307c576001600160a01b03612f386123178385876135a4565b165f52602860205260ff60405f205460081c1660038110156117805761305257604051612f6481613417565b60ff8216815260208101600281526001600160a01b03612f886123178587896135a4565b165f52602860205260ff60405f209251168254915160038110156117805761ff009060081b169161ffff191617179055612fc6612317828a8c6135a4565b60295490680100000000000000008210156111ea57612fee82600161300e9401602955613561565b90919082549060031b916001600160a01b03809116831b921b1916179055565b61301c6123178284866135a4565b90602a54680100000000000000008110156111ea5761304d92612fee8260016130489401602a5561352c565b613596565b612cd3565b60405162461bcd60e51b8152602060048201526002602482015261303760f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261181b60f11b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261303560f01b6044820152606490fd5b60405162461bcd60e51b81526020600482015260026024820152610c0d60f21b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261303360f01b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261181960f11b6044820152606490fd5b60405162461bcd60e51b8152602060048201526002602482015261303160f01b6044820152606490fd5b34610329575f36600319011261032957602060405160048152f35b34610329575f36600319011261032957602060405160018152f35b34610329576040366003190112610329576004356001600160a01b038082168203610329576127b6916131ed6131e2613353565b925f541633146134d3565b613b33565b34610329575f36600319011261032957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610329575f36600319011261032957613247613ce8565b8051602080830151604080850151606080870151608097880151845163ffffffff9889168152958816968601969096529186169284019290925284169082015291169181019190915260a090f35b34610329575f3660031901126103295760206040517f000000000000000000000000000000000000000000000000000000000000000060170b8152f35b34610329575f366003190112610329576102f26040516132f181613417565b6008815267061675f352e302e360c41b6020820152604051918291602083526020830190613369565b346103295760203660031901126103295760206120e361333861333d565b613d5d565b600435906001600160a01b038216820361032957565b602435906001600160a01b038216820361032957565b91908251928382525f5b848110613393575050825f602080949584010152601f8019910116010190565b602081830181015184830182015201613373565b9181601f84011215610329578235916001600160401b038311610329576020808501948460051b01011161032957565b9181601f84011215610329578235916001600160401b038311610329576020838186019501011161032957565b6024359063ffffffff8216820361032957565b604081019081106001600160401b038211176111ea57604052565b608081019081106001600160401b038211176111ea57604052565b60a081019081106001600160401b038211176111ea57604052565b6103e081019081106001600160401b038211176111ea57604052565b90601f801991011681019081106001600160401b038211176111ea57604052565b6001600160401b0381116111ea57601f01601f191660200190565b818102929181159184041417156107f157565b156134da57565b60405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606490fd5b919082039182116107f157565b602a54811015610ea957602a5f527fbeced09521047d05b8960b7e7bcc1d1292cf3e4b2a6b63f48335cbde5f7545d201905f90565b602954811015610ea95760295f527fcb7c14ce178f56e2e8d86ab33ebc0ae081ba8556a00cd122038841867181caac01905f90565b5f1981146107f15760010190565b9190811015610ea95760051b0190565b356001600160a01b03811681036103295790565b60038210156117805752565b90600163ffffffff809316019182116107f157565b9190808252602080920192915f805b83821061360757505050505090565b909192939485356001600160a01b0381168091036136325781528301948301939291600101906135f8565b8280fd5b90602091808252805f848401375f828201840152601f01601f1916010190565b908060209392818452848401375f828201840152601f01601f1916010190565b60405190602a548083528260209182820190602a5f527fbeced09521047d05b8960b7e7bcc1d1292cf3e4b2a6b63f48335cbde5f7545d2935f905b8282106136c9575050506136c792500383613484565b565b85546001600160a01b0316845260019586019588955093810193909101906136b1565b604051906136f982613417565b602e546001600160a01b038116835260a01c63ffffffff166020830152565b90816020910312610329575180151581036103295790565b6040519061373d82613432565b81606063ffffffff602b546001600160801b03198160801b16845264ffffffffff8160801c16602085015260ff8160a81c16604085015260b01c16910152565b6001600160401b0381116111ea5760051b60200190565b9291926137a0826134a5565b916137ae6040519384613484565b829481845281830111610329578281602093845f960137010152565b9080601f83011215610329578160206137e593359101613794565b90565b906137f2826134a5565b6137ff6040519182613484565b8281528092613810601f19916134a5565b0190602036910137565b60ff1660ff81146107f15760010190565b908151811015610ea9570160200190565b90601f811015610ea95760051b0190565b8051821015610ea95760209160051b010190565b63ffffffff8091169081146107f15760010190565b8054821015610ea9575f5260205f209060011b01905f90565b90600182811c921680156138bd575b60208310146138a957565b634e487b7160e01b5f52602260045260245ffd5b91607f169161389e565b8181106138d2575050565b5f81556001016138c7565b156138e457565b634e487b7160e01b5f52600160045260245ffd5b919082018092116107f157565b63ffffffff1061391457600190565b5f90565b63ffffffff9081811161394157165f52602c6020526001600160401b03600360405f2001541690565b50505f90565b9060405191825f82549261395a8461388f565b9081845260019485811690815f146139c35750600114613983575b50506136c792500383613484565b909391505f52602090815f20935f915b8183106139ab5750506136c793508201015f80613975565b85548884018501529485019487945091830191613993565b9150506136c794506020925060ff191682840152151560051b8201015f80613975565b915f925f928384808193829563ffffffff80602b5460b01c16975b888652602c60209080825260029060409082828b2001868154168787161015613aae57613a3a613a3487613a4b93613876565b50613947565b613a45368b8b613794565b90613abe565b613a615750505050613a5c90613861565b613a01565b965097965098509a50915050613a9a995060039450868352818952838320998a916001600160401b039788600180950154169c01613876565b500154978683525220015416909392918190565b5050505050505050509350509392565b908151815103613941575f90815b8351811015613b2a577fff0000000000000000000000000000000000000000000000000000000000000080613b01838761382b565b511690613b0e838561382b565b511603613b2357613b1e90613596565b613acc565b5050905090565b50505050600190565b6001600160a01b038060035416911691818314613ce357604080516370a0823160e01b8082523060048301526020939184816024818a5afa8015613cd957908591613cb0575b5050613b836141ad565b82519081523060048201528381602481885afa908115613ca6579084915f91613c73575b50835163a9059cbb60e01b81526001600160a01b03909316600484015260248301528180604481015b03815f885af1908115613c69575f91613c4c575b5015613c24575050816001600160a01b031960035416176003557f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a5f80a3565b60649250519062461bcd60e51b825260048201526002602482015261343160f01b6044820152fd5b613c639150833d8511611dd057611dc28183613484565b5f613be4565b82513d5f823e3d90fd5b9182813d8311613c9f575b613c888183613484565b81010312613c9c5750518390613bd0613ba7565b80fd5b503d613c7e565b83513d5f823e3d90fd5b813d8311613cd2575b613cc38183613484565b8101031261032957835f613b79565b503d613cb9565b84513d5f823e3d90fd5b505050565b60405190613cf58261344d565b81608060025463ffffffff908181168452818160201c166020850152818160401c166040850152818160601c166060850152821c16910152565b90601f821015610ea957601e8260041c6005019260011b1690565b61ffff9081165f1901919082116107f157565b906001600160a01b035f9216825260286020526040822060405190613d8182613417565b5460ff81168252613d9c60ff602084019260081c16826135c8565b516003811015613e545715613e5157613dea613db6613ce8565b63ffffffff6060613ddf613dcd60ff875116613d2f565b905461ffff929160031b1c8216613d4a565b1692015116906134c0565b90633b9aca0091828102928184041490151715613e29575160ff16601f811015613e3d57600901545f19810191908211613e29576137e59293506138f8565b634e487b7160e01b84526011600452602484fd5b634e487b7160e01b84526032600452602484fd5b50565b634e487b7160e01b84526021600452602484fd5b6001600160a01b03908181165f81815260209060288252604094613eac8683209561333860ff895198613e9a8a613417565b548181168a5260081c168689016135c8565b9182613ebc575b50505050505050565b8481526007845286812054600354885163a9059cbb60e01b81529184166001600160a01b0381166004840152602483018690529791908690829086168186816044810103925af1908115613fe3578391613fc6575b5015613f9d5780613f2660ff80935116613d2f565b81549060031b9061ffff6001831b921b1916179055511690601f821015613f8957509060017fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c94939260090155600354169551908152a45f808080808080613eb3565b634e487b7160e01b81526032600452602490fd5b875162461bcd60e51b81526004810186905260026024820152610d0d60f21b6044820152606490fd5b613fdd9150863d8811611dd057611dc28183613484565b5f613f11565b89513d85823e3d90fd5b604090815191829060055f5b601f600f8201106140b85750906101c09154809161ffff9283918282168752828260101c166020880152828260201c1681880152828260301c1660608801521c166080850152818160501c1660a0850152818160601c1660c0850152818160701c1660e0850152818160801c16610100850152818160901c16610120850152818160a01c16610140850152818160b01c16610160850152818160c01c16610180850152818160d01c166101a085015260e01c169101526136c782613468565b90926001610200601092865461ffff80821683528082602082828a1c16818701521c1688840152818882826060828260301c168189015282826080951c1684890152828260a0828260501c16818c0152828260c0951c16848c0152828260e0988d8a848460701c169101521c166101008c0152828260901c166101208c01521c16610140890152828260b01c166101608901521c16610180860152828260d01c166101a08601521c166101c083015260f01c6101e08201520194019101613ff9565b6040519060095f835b601f8210614197575050506136c782613468565b6001602081928554815201930191019091614183565b6141b5613ce8565b906001600160a01b03918260035416906141cd613fed565b906141d661417a565b946141df613676565b916060015f5b8351811015614381576141f8818961383c565b515f1981019081116107f15761422e61ffff61421f81614218868b61383c565b5116613d4a565b1663ffffffff855116906134c0565b90633b9aca00918281029281840414901517156107f15761424e916138f8565b80614263575b5061425e90613596565b6141e5565b8361426e838761384d565b51165f5260206007815260409085825f20541692825163a9059cbb60e01b815260049083818d815f816142ba898d8a8401602090939291936001600160a01b0360408201951681520152565b03925af1908115614377575f9161435a575b50156143325750917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8594928b94868f61425e99614318918f614312600195869261383c565b5261383c565b5289614324888d61384d565b51169351908152a490614254565b835162461bcd60e51b81529081018390526002602482015261343560f01b6044820152606490fd5b6143719150843d8611611dd057611dc28183613484565b5f6142cc565b85513d5f823e3d90fd5b5050505090505f5b6001808210156143e0575f9081905b601082106143af5750506005820155600101614389565b60206143d6869461ffff84959851169088851b60031b9161ffff809116831b921b19161790565b9301940190614398565b50509190915f90815b600f81106144195750506006555f5b601f8110614404575050565b600190602083519301928160090155016143f8565b9091602061444160019261ffff8651169085851b60031b9161ffff809116831b921b19161790565b930191016143e9565b5f9081614455613fed565b83905b601f821061456a57505063ffffffff6060614471613ce8565b01511690633b9aca0091828102928184041490151715613e295790614495916134c0565b604051918283602a54808352602080930190602a88527fbeced09521047d05b8960b7e7bcc1d1292cf3e4b2a6b63f48335cbde5f7545d29388905b828210614547575050506144e692500384613484565b6144ee61417a565b9084925b845184101561453f57614505848461383c565b515f1981019190821161452b576145259161451f916138f8565b93613596565b926144f2565b634e487b7160e01b87526011600452602487fd5b945092505050565b85546001600160a01b0316845260019586019589955093810193909101906144d0565b909161458b6145919161ffff61458481614218888861383c565b16906138f8565b92613596565b90614458565b6001600160a01b03165f52602860205260405f20604051906145b882613417565b5460ff811682526145d360ff602084019260081c16826135c8565b51600381101561178057156146025761ffff6145f560ff6137e5935116613d2f565b90549060031b1c16613d4a565b505f90565b61ffff90818110614616575090565b905090565b81811061461657509056fea164736f6c6343000814000a",
}

// AccessControlledOffchainAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlledOffchainAggregatorMetaData.ABI instead.
var AccessControlledOffchainAggregatorABI = AccessControlledOffchainAggregatorMetaData.ABI

// AccessControlledOffchainAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessControlledOffchainAggregatorMetaData.Bin instead.
var AccessControlledOffchainAggregatorBin = AccessControlledOffchainAggregatorMetaData.Bin

// DeployAccessControlledOffchainAggregator deploys a new Ethereum contract, binding an instance of AccessControlledOffchainAggregator to it.
func DeployAccessControlledOffchainAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32, _link common.Address, _minAnswer *big.Int, _maxAnswer *big.Int, _billingAccessController common.Address, _requesterAccessController common.Address, _decimals uint8, desc string) (common.Address, *types.Transaction, *AccessControlledOffchainAggregator, error) {
	parsed, err := AccessControlledOffchainAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessControlledOffchainAggregatorBin), backend, _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission, _link, _minAnswer, _maxAnswer, _billingAccessController, _requesterAccessController, _decimals, desc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlledOffchainAggregator{AccessControlledOffchainAggregatorCaller: AccessControlledOffchainAggregatorCaller{contract: contract}, AccessControlledOffchainAggregatorTransactor: AccessControlledOffchainAggregatorTransactor{contract: contract}, AccessControlledOffchainAggregatorFilterer: AccessControlledOffchainAggregatorFilterer{contract: contract}}, nil
}

// AccessControlledOffchainAggregator is an auto generated Go binding around an Ethereum contract.
type AccessControlledOffchainAggregator struct {
	AccessControlledOffchainAggregatorCaller     // Read-only binding to the contract
	AccessControlledOffchainAggregatorTransactor // Write-only binding to the contract
	AccessControlledOffchainAggregatorFilterer   // Log filterer for contract events
}

// AccessControlledOffchainAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlledOffchainAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlledOffchainAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlledOffchainAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlledOffchainAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlledOffchainAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlledOffchainAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlledOffchainAggregatorSession struct {
	Contract     *AccessControlledOffchainAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AccessControlledOffchainAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlledOffchainAggregatorCallerSession struct {
	Contract *AccessControlledOffchainAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// AccessControlledOffchainAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlledOffchainAggregatorTransactorSession struct {
	Contract     *AccessControlledOffchainAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// AccessControlledOffchainAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlledOffchainAggregatorRaw struct {
	Contract *AccessControlledOffchainAggregator // Generic contract binding to access the raw methods on
}

// AccessControlledOffchainAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlledOffchainAggregatorCallerRaw struct {
	Contract *AccessControlledOffchainAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlledOffchainAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlledOffchainAggregatorTransactorRaw struct {
	Contract *AccessControlledOffchainAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlledOffchainAggregator creates a new instance of AccessControlledOffchainAggregator, bound to a specific deployed contract.
func NewAccessControlledOffchainAggregator(address common.Address, backend bind.ContractBackend) (*AccessControlledOffchainAggregator, error) {
	contract, err := bindAccessControlledOffchainAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregator{AccessControlledOffchainAggregatorCaller: AccessControlledOffchainAggregatorCaller{contract: contract}, AccessControlledOffchainAggregatorTransactor: AccessControlledOffchainAggregatorTransactor{contract: contract}, AccessControlledOffchainAggregatorFilterer: AccessControlledOffchainAggregatorFilterer{contract: contract}}, nil
}

// NewAccessControlledOffchainAggregatorCaller creates a new read-only instance of AccessControlledOffchainAggregator, bound to a specific deployed contract.
func NewAccessControlledOffchainAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AccessControlledOffchainAggregatorCaller, error) {
	contract, err := bindAccessControlledOffchainAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorCaller{contract: contract}, nil
}

// NewAccessControlledOffchainAggregatorTransactor creates a new write-only instance of AccessControlledOffchainAggregator, bound to a specific deployed contract.
func NewAccessControlledOffchainAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlledOffchainAggregatorTransactor, error) {
	contract, err := bindAccessControlledOffchainAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorTransactor{contract: contract}, nil
}

// NewAccessControlledOffchainAggregatorFilterer creates a new log filterer instance of AccessControlledOffchainAggregator, bound to a specific deployed contract.
func NewAccessControlledOffchainAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlledOffchainAggregatorFilterer, error) {
	contract, err := bindAccessControlledOffchainAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorFilterer{contract: contract}, nil
}

// bindAccessControlledOffchainAggregator binds a generic wrapper to an already deployed contract.
func bindAccessControlledOffchainAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlledOffchainAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlledOffchainAggregator.Contract.AccessControlledOffchainAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AccessControlledOffchainAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AccessControlledOffchainAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlledOffchainAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.contract.Transact(opts, method, params...)
}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) BillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "billingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) BillingAccessController() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.BillingAccessController(&_AccessControlledOffchainAggregator.CallOpts)
}

// BillingAccessController is a free data retrieval call binding the contract method 0x996e8298.
//
// Solidity: function billingAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) BillingAccessController() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.BillingAccessController(&_AccessControlledOffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Decimals() (uint8, error) {
	return _AccessControlledOffchainAggregator.Contract.Decimals(&_AccessControlledOffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) Decimals() (uint8, error) {
	return _AccessControlledOffchainAggregator.Contract.Decimals(&_AccessControlledOffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Description() (string, error) {
	return _AccessControlledOffchainAggregator.Contract.Description(&_AccessControlledOffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) Description() (string, error) {
	return _AccessControlledOffchainAggregator.Contract.Description(&_AccessControlledOffchainAggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.GetAnswer(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.GetAnswer(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "getBilling")

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) GetBilling() (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	return _AccessControlledOffchainAggregator.Contract.GetBilling(&_AccessControlledOffchainAggregator.CallOpts)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) GetBilling() (struct {
	MaximumGasPrice         uint32
	ReasonableGasPrice      uint32
	MicroLinkPerEth         uint32
	LinkGweiPerObservation  uint32
	LinkGweiPerTransmission uint32
}, error) {
	return _AccessControlledOffchainAggregator.Contract.GetBilling(&_AccessControlledOffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) GetLinkToken() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.GetLinkToken(&_AccessControlledOffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.GetLinkToken(&_AccessControlledOffchainAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "getRoundData", _roundId)

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.GetRoundData(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.GetRoundData(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.GetTimestamp(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.GetTimestamp(&_AccessControlledOffchainAggregator.CallOpts, _roundId)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) IsLocked() (bool, error) {
	return _AccessControlledOffchainAggregator.Contract.IsLocked(&_AccessControlledOffchainAggregator.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) IsLocked() (bool, error) {
	return _AccessControlledOffchainAggregator.Contract.IsLocked(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes16 configDigest)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestConfigDetails")

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestConfigDetails(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes16 configDigest)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [16]byte
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestConfigDetails(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xd4c520d1.
//
// Solidity: function latestMerkleRoundData(string taskId) view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestMerkleRoundData(opts *bind.CallOpts, taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestMerkleRoundData", taskId)

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestMerkleRoundData(taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestMerkleRoundData(&_AccessControlledOffchainAggregator.CallOpts, taskId)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xd4c520d1.
//
// Solidity: function latestMerkleRoundData(string taskId) view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestMerkleRoundData(taskId string) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestMerkleRoundData(&_AccessControlledOffchainAggregator.CallOpts, taskId)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestRound() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestRound(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestRound(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestRoundData")

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestRoundData(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestRoundData(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestTimestamp(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestTimestamp(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes16 configDigest, uint32 epoch, uint8 round, int192 latestAnswer, uint64 latestTimestamp)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "latestTransmissionDetails")

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestTransmissionDetails(&_AccessControlledOffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes16 configDigest, uint32 epoch, uint8 round, int192 latestAnswer, uint64 latestTimestamp)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [16]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _AccessControlledOffchainAggregator.Contract.LatestTransmissionDetails(&_AccessControlledOffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LinkAvailableForPayment(&_AccessControlledOffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.LinkAvailableForPayment(&_AccessControlledOffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) MaxAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.MaxAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.MaxAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) MinAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.MinAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.MinAnswer(&_AccessControlledOffchainAggregator.CallOpts)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) OracleObservationCount(opts *bind.CallOpts, _signerOrTransmitter common.Address) (uint16, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "oracleObservationCount", _signerOrTransmitter)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) OracleObservationCount(_signerOrTransmitter common.Address) (uint16, error) {
	return _AccessControlledOffchainAggregator.Contract.OracleObservationCount(&_AccessControlledOffchainAggregator.CallOpts, _signerOrTransmitter)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address _signerOrTransmitter) view returns(uint16)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) OracleObservationCount(_signerOrTransmitter common.Address) (uint16, error) {
	return _AccessControlledOffchainAggregator.Contract.OracleObservationCount(&_AccessControlledOffchainAggregator.CallOpts, _signerOrTransmitter)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) OwedPayment(opts *bind.CallOpts, _transmitter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "owedPayment", _transmitter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) OwedPayment(_transmitter common.Address) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.OwedPayment(&_AccessControlledOffchainAggregator.CallOpts, _transmitter)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address _transmitter) view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) OwedPayment(_transmitter common.Address) (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.OwedPayment(&_AccessControlledOffchainAggregator.CallOpts, _transmitter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Owner() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.Owner(&_AccessControlledOffchainAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) Owner() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.Owner(&_AccessControlledOffchainAggregator.CallOpts)
}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) RequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "requesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) RequesterAccessController() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.RequesterAccessController(&_AccessControlledOffchainAggregator.CallOpts)
}

// RequesterAccessController is a free data retrieval call binding the contract method 0x70efdf2d.
//
// Solidity: function requesterAccessController() view returns(address)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) RequesterAccessController() (common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.RequesterAccessController(&_AccessControlledOffchainAggregator.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Transmitters() ([]common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.Transmitters(&_AccessControlledOffchainAggregator.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) Transmitters() ([]common.Address, error) {
	return _AccessControlledOffchainAggregator.Contract.Transmitters(&_AccessControlledOffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) TypeAndVersion() (string, error) {
	return _AccessControlledOffchainAggregator.Contract.TypeAndVersion(&_AccessControlledOffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) TypeAndVersion() (string, error) {
	return _AccessControlledOffchainAggregator.Contract.TypeAndVersion(&_AccessControlledOffchainAggregator.CallOpts)
}

// ValidatorConfig is a free data retrieval call binding the contract method 0x8e0566de.
//
// Solidity: function validatorConfig() view returns(address validator, uint32 gasLimit)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) ValidatorConfig(opts *bind.CallOpts) (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "validatorConfig")

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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) ValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _AccessControlledOffchainAggregator.Contract.ValidatorConfig(&_AccessControlledOffchainAggregator.CallOpts)
}

// ValidatorConfig is a free data retrieval call binding the contract method 0x8e0566de.
//
// Solidity: function validatorConfig() view returns(address validator, uint32 gasLimit)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) ValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _AccessControlledOffchainAggregator.Contract.ValidatorConfig(&_AccessControlledOffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlledOffchainAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Version() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.Version(&_AccessControlledOffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorCallerSession) Version() (*big.Int, error) {
	return _AccessControlledOffchainAggregator.Contract.Version(&_AccessControlledOffchainAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AcceptOwnership(&_AccessControlledOffchainAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AcceptOwnership(&_AccessControlledOffchainAggregator.TransactOpts)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, _transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "acceptPayeeship", _transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) AcceptPayeeship(_transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AcceptPayeeship(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) AcceptPayeeship(_transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.AcceptPayeeship(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.RequestNewRound(&_AccessControlledOffchainAggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.RequestNewRound(&_AccessControlledOffchainAggregator.TransactOpts)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetBilling(opts *bind.TransactOpts, _maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setBilling", _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetBilling(_maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetBilling(&_AccessControlledOffchainAggregator.TransactOpts, _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBilling is a paid mutator transaction binding the contract method 0xbd824706.
//
// Solidity: function setBilling(uint32 _maximumGasPrice, uint32 _reasonableGasPrice, uint32 _microLinkPerEth, uint32 _linkGweiPerObservation, uint32 _linkGweiPerTransmission) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetBilling(_maximumGasPrice uint32, _reasonableGasPrice uint32, _microLinkPerEth uint32, _linkGweiPerObservation uint32, _linkGweiPerTransmission uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetBilling(&_AccessControlledOffchainAggregator.TransactOpts, _maximumGasPrice, _reasonableGasPrice, _microLinkPerEth, _linkGweiPerObservation, _linkGweiPerTransmission)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetBillingAccessController(&_AccessControlledOffchainAggregator.TransactOpts, _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetBillingAccessController(&_AccessControlledOffchainAggregator.TransactOpts, _billingAccessController)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setConfig", _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetConfig(&_AccessControlledOffchainAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetConfig(&_AccessControlledOffchainAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, _linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setLinkToken", _linkToken, _recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetLinkToken(_linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetLinkToken(&_AccessControlledOffchainAggregator.TransactOpts, _linkToken, _recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address _linkToken, address _recipient) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetLinkToken(_linkToken common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetLinkToken(&_AccessControlledOffchainAggregator.TransactOpts, _linkToken, _recipient)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetLock(opts *bind.TransactOpts, isLock bool) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setLock", isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetLock(&_AccessControlledOffchainAggregator.TransactOpts, isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetLock(&_AccessControlledOffchainAggregator.TransactOpts, isLock)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetPayees(opts *bind.TransactOpts, _transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setPayees", _transmitters, _payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetPayees(_transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetPayees(&_AccessControlledOffchainAggregator.TransactOpts, _transmitters, _payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] _transmitters, address[] _payees) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetPayees(_transmitters []common.Address, _payees []common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetPayees(&_AccessControlledOffchainAggregator.TransactOpts, _transmitters, _payees)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, _requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setRequesterAccessController", _requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetRequesterAccessController(_requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetRequesterAccessController(&_AccessControlledOffchainAggregator.TransactOpts, _requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address _requesterAccessController) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetRequesterAccessController(_requesterAccessController common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetRequesterAccessController(&_AccessControlledOffchainAggregator.TransactOpts, _requesterAccessController)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, _newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "setValidatorConfig", _newValidator, _newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) SetValidatorConfig(_newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetValidatorConfig(&_AccessControlledOffchainAggregator.TransactOpts, _newValidator, _newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address _newValidator, uint32 _newGasLimit) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) SetValidatorConfig(_newValidator common.Address, _newGasLimit uint32) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.SetValidatorConfig(&_AccessControlledOffchainAggregator.TransactOpts, _newValidator, _newGasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.TransferOwnership(&_AccessControlledOffchainAggregator.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.TransferOwnership(&_AccessControlledOffchainAggregator.TransactOpts, _to)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, _transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "transferPayeeship", _transmitter, _proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) TransferPayeeship(_transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.TransferPayeeship(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter, _proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address _transmitter, address _proposed) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) TransferPayeeship(_transmitter common.Address, _proposed common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.TransferPayeeship(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter, _proposed)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) Transmit(opts *bind.TransactOpts, _report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "transmit", _report, _rs, _ss, _rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) Transmit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.Transmit(&_AccessControlledOffchainAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xc9807539.
//
// Solidity: function transmit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) Transmit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.Transmit(&_AccessControlledOffchainAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.WithdrawFunds(&_AccessControlledOffchainAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.WithdrawFunds(&_AccessControlledOffchainAggregator.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, _transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.contract.Transact(opts, "withdrawPayment", _transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorSession) WithdrawPayment(_transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.WithdrawPayment(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address _transmitter) returns()
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorTransactorSession) WithdrawPayment(_transmitter common.Address) (*types.Transaction, error) {
	return _AccessControlledOffchainAggregator.Contract.WithdrawPayment(&_AccessControlledOffchainAggregator.TransactOpts, _transmitter)
}

// AccessControlledOffchainAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorAnswerUpdatedIterator struct {
	Event *AccessControlledOffchainAggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorAnswerUpdated)
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
		it.Event = new(AccessControlledOffchainAggregatorAnswerUpdated)
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
func (it *AccessControlledOffchainAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorAnswerUpdated represents a AnswerUpdated event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorAnswerUpdated struct {
	CurrentBatchId *big.Int
	RoundId        *big.Int
	UpdatedAt      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093.
//
// Solidity: event AnswerUpdated(uint256 indexed currentBatchId, uint256 indexed roundId, uint256 updatedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, currentBatchId []*big.Int, roundId []*big.Int) (*AccessControlledOffchainAggregatorAnswerUpdatedIterator, error) {

	var currentBatchIdRule []interface{}
	for _, currentBatchIdItem := range currentBatchId {
		currentBatchIdRule = append(currentBatchIdRule, currentBatchIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentBatchIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorAnswerUpdatedIterator{contract: _AccessControlledOffchainAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x5a898c50e6c03bbff0088c38aacacb2500461f91ab987eff3b054f233a9aa093.
//
// Solidity: event AnswerUpdated(uint256 indexed currentBatchId, uint256 indexed roundId, uint256 updatedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorAnswerUpdated, currentBatchId []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentBatchIdRule []interface{}
	for _, currentBatchIdItem := range currentBatchId {
		currentBatchIdRule = append(currentBatchIdRule, currentBatchIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentBatchIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorAnswerUpdated)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*AccessControlledOffchainAggregatorAnswerUpdated, error) {
	event := new(AccessControlledOffchainAggregatorAnswerUpdated)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorBillingAccessControllerSetIterator is returned from FilterBillingAccessControllerSet and is used to iterate over the raw logs and unpacked data for BillingAccessControllerSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorBillingAccessControllerSetIterator struct {
	Event *AccessControlledOffchainAggregatorBillingAccessControllerSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorBillingAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorBillingAccessControllerSet)
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
		it.Event = new(AccessControlledOffchainAggregatorBillingAccessControllerSet)
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
func (it *AccessControlledOffchainAggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorBillingAccessControllerSet represents a BillingAccessControllerSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBillingAccessControllerSet is a free log retrieval operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOffchainAggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorBillingAccessControllerSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchBillingAccessControllerSet is a free log subscription operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorBillingAccessControllerSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*AccessControlledOffchainAggregatorBillingAccessControllerSet, error) {
	event := new(AccessControlledOffchainAggregatorBillingAccessControllerSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorBillingSetIterator is returned from FilterBillingSet and is used to iterate over the raw logs and unpacked data for BillingSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorBillingSetIterator struct {
	Event *AccessControlledOffchainAggregatorBillingSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorBillingSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorBillingSet)
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
		it.Event = new(AccessControlledOffchainAggregatorBillingSet)
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
func (it *AccessControlledOffchainAggregatorBillingSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorBillingSet represents a BillingSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorBillingSet struct {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*AccessControlledOffchainAggregatorBillingSetIterator, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorBillingSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

// WatchBillingSet is a free log subscription operation binding the contract event 0xd0d9486a2c673e2a4b57fc82e4c8a556b3e2b82dd5db07e2c04a920ca0f469b6.
//
// Solidity: event BillingSet(uint32 maximumGasPrice, uint32 reasonableGasPrice, uint32 microLinkPerEth, uint32 linkGweiPerObservation, uint32 linkGweiPerTransmission)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorBillingSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseBillingSet(log types.Log) (*AccessControlledOffchainAggregatorBillingSet, error) {
	event := new(AccessControlledOffchainAggregatorBillingSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorConfigSetIterator struct {
	Event *AccessControlledOffchainAggregatorConfigSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorConfigSet)
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
		it.Event = new(AccessControlledOffchainAggregatorConfigSet)
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
func (it *AccessControlledOffchainAggregatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorConfigSet represents a ConfigSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorConfigSet struct {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*AccessControlledOffchainAggregatorConfigSetIterator, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorConfigSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x25d719d88a4512dd76c7442b910a83360845505894eb444ef299409e180f8fb9.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, uint64 configCount, address[] signers, address[] transmitters, uint8 threshold, uint64 encodedConfigVersion, bytes encoded)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorConfigSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseConfigSet(log types.Log) (*AccessControlledOffchainAggregatorConfigSet, error) {
	event := new(AccessControlledOffchainAggregatorConfigSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorLinkTokenSetIterator is returned from FilterLinkTokenSet and is used to iterate over the raw logs and unpacked data for LinkTokenSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorLinkTokenSetIterator struct {
	Event *AccessControlledOffchainAggregatorLinkTokenSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorLinkTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorLinkTokenSet)
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
		it.Event = new(AccessControlledOffchainAggregatorLinkTokenSet)
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
func (it *AccessControlledOffchainAggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorLinkTokenSet represents a LinkTokenSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLinkTokenSet is a free log retrieval operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed _oldLinkToken, address indexed _newLinkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, _oldLinkToken []common.Address, _newLinkToken []common.Address) (*AccessControlledOffchainAggregatorLinkTokenSetIterator, error) {

	var _oldLinkTokenRule []interface{}
	for _, _oldLinkTokenItem := range _oldLinkToken {
		_oldLinkTokenRule = append(_oldLinkTokenRule, _oldLinkTokenItem)
	}
	var _newLinkTokenRule []interface{}
	for _, _newLinkTokenItem := range _newLinkToken {
		_newLinkTokenRule = append(_newLinkTokenRule, _newLinkTokenItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "LinkTokenSet", _oldLinkTokenRule, _newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorLinkTokenSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

// WatchLinkTokenSet is a free log subscription operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed _oldLinkToken, address indexed _newLinkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorLinkTokenSet, _oldLinkToken []common.Address, _newLinkToken []common.Address) (event.Subscription, error) {

	var _oldLinkTokenRule []interface{}
	for _, _oldLinkTokenItem := range _oldLinkToken {
		_oldLinkTokenRule = append(_oldLinkTokenRule, _oldLinkTokenItem)
	}
	var _newLinkTokenRule []interface{}
	for _, _newLinkTokenItem := range _newLinkToken {
		_newLinkTokenRule = append(_newLinkTokenRule, _newLinkTokenItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "LinkTokenSet", _oldLinkTokenRule, _newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorLinkTokenSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseLinkTokenSet(log types.Log) (*AccessControlledOffchainAggregatorLinkTokenSet, error) {
	event := new(AccessControlledOffchainAggregatorLinkTokenSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorNewRoundIterator struct {
	Event *AccessControlledOffchainAggregatorNewRound // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorNewRound)
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
		it.Event = new(AccessControlledOffchainAggregatorNewRound)
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
func (it *AccessControlledOffchainAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorNewRound represents a NewRound event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AccessControlledOffchainAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorNewRoundIterator{contract: _AccessControlledOffchainAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorNewRound)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseNewRound(log types.Log) (*AccessControlledOffchainAggregatorNewRound, error) {
	event := new(AccessControlledOffchainAggregatorNewRound)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorNewTransmissionIterator is returned from FilterNewTransmission and is used to iterate over the raw logs and unpacked data for NewTransmission events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorNewTransmissionIterator struct {
	Event *AccessControlledOffchainAggregatorNewTransmission // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorNewTransmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorNewTransmission)
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
		it.Event = new(AccessControlledOffchainAggregatorNewTransmission)
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
func (it *AccessControlledOffchainAggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorNewTransmission represents a NewTransmission event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorNewTransmission struct {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*AccessControlledOffchainAggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorNewTransmissionIterator{contract: _AccessControlledOffchainAggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

// WatchNewTransmission is a free log subscription operation binding the contract event 0x62dd611419e1478005128e47b35bab9b72d945c034c7c0f3851c84fc189d9db1.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, string projectId, address transmitter, bytes observers, bytes32 rawReportContext)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorNewTransmission)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseNewTransmission(log types.Log) (*AccessControlledOffchainAggregatorNewTransmission, error) {
	event := new(AccessControlledOffchainAggregatorNewTransmission)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorOraclePaidIterator is returned from FilterOraclePaid and is used to iterate over the raw logs and unpacked data for OraclePaid events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOraclePaidIterator struct {
	Event *AccessControlledOffchainAggregatorOraclePaid // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorOraclePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorOraclePaid)
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
		it.Event = new(AccessControlledOffchainAggregatorOraclePaid)
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
func (it *AccessControlledOffchainAggregatorOraclePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorOraclePaid represents a OraclePaid event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePaid is a free log retrieval operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*AccessControlledOffchainAggregatorOraclePaidIterator, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorOraclePaidIterator{contract: _AccessControlledOffchainAggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

// WatchOraclePaid is a free log subscription operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorOraclePaid)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseOraclePaid(log types.Log) (*AccessControlledOffchainAggregatorOraclePaid, error) {
	event := new(AccessControlledOffchainAggregatorOraclePaid)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator struct {
	Event *AccessControlledOffchainAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorOwnershipTransferRequested)
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
		it.Event = new(AccessControlledOffchainAggregatorOwnershipTransferRequested)
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
func (it *AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorOwnershipTransferRequestedIterator{contract: _AccessControlledOffchainAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorOwnershipTransferRequested)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*AccessControlledOffchainAggregatorOwnershipTransferRequested, error) {
	event := new(AccessControlledOffchainAggregatorOwnershipTransferRequested)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOwnershipTransferredIterator struct {
	Event *AccessControlledOffchainAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorOwnershipTransferred)
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
		it.Event = new(AccessControlledOffchainAggregatorOwnershipTransferred)
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
func (it *AccessControlledOffchainAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AccessControlledOffchainAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorOwnershipTransferredIterator{contract: _AccessControlledOffchainAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorOwnershipTransferred)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*AccessControlledOffchainAggregatorOwnershipTransferred, error) {
	event := new(AccessControlledOffchainAggregatorOwnershipTransferred)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator is returned from FilterPayeeshipTransferRequested and is used to iterate over the raw logs and unpacked data for PayeeshipTransferRequested events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator struct {
	Event *AccessControlledOffchainAggregatorPayeeshipTransferRequested // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorPayeeshipTransferRequested)
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
		it.Event = new(AccessControlledOffchainAggregatorPayeeshipTransferRequested)
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
func (it *AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorPayeeshipTransferRequested represents a PayeeshipTransferRequested event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferRequested is a free log retrieval operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorPayeeshipTransferRequestedIterator{contract: _AccessControlledOffchainAggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferRequested is a free log subscription operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorPayeeshipTransferRequested)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*AccessControlledOffchainAggregatorPayeeshipTransferRequested, error) {
	event := new(AccessControlledOffchainAggregatorPayeeshipTransferRequested)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorPayeeshipTransferredIterator is returned from FilterPayeeshipTransferred and is used to iterate over the raw logs and unpacked data for PayeeshipTransferred events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorPayeeshipTransferredIterator struct {
	Event *AccessControlledOffchainAggregatorPayeeshipTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorPayeeshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorPayeeshipTransferred)
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
		it.Event = new(AccessControlledOffchainAggregatorPayeeshipTransferred)
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
func (it *AccessControlledOffchainAggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorPayeeshipTransferred represents a PayeeshipTransferred event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferred is a free log retrieval operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*AccessControlledOffchainAggregatorPayeeshipTransferredIterator, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorPayeeshipTransferredIterator{contract: _AccessControlledOffchainAggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferred is a free log subscription operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorPayeeshipTransferred)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*AccessControlledOffchainAggregatorPayeeshipTransferred, error) {
	event := new(AccessControlledOffchainAggregatorPayeeshipTransferred)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator is returned from FilterRequesterAccessControllerSet and is used to iterate over the raw logs and unpacked data for RequesterAccessControllerSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator struct {
	Event *AccessControlledOffchainAggregatorRequesterAccessControllerSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorRequesterAccessControllerSet)
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
		it.Event = new(AccessControlledOffchainAggregatorRequesterAccessControllerSet)
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
func (it *AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorRequesterAccessControllerSet represents a RequesterAccessControllerSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequesterAccessControllerSet is a free log retrieval operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorRequesterAccessControllerSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchRequesterAccessControllerSet is a free log subscription operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorRequesterAccessControllerSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*AccessControlledOffchainAggregatorRequesterAccessControllerSet, error) {
	event := new(AccessControlledOffchainAggregatorRequesterAccessControllerSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorRoundRequestedIterator is returned from FilterRoundRequested and is used to iterate over the raw logs and unpacked data for RoundRequested events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorRoundRequestedIterator struct {
	Event *AccessControlledOffchainAggregatorRoundRequested // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorRoundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorRoundRequested)
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
		it.Event = new(AccessControlledOffchainAggregatorRoundRequested)
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
func (it *AccessControlledOffchainAggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorRoundRequested represents a RoundRequested event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [16]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRoundRequested is a free log retrieval operation binding the contract event 0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037.
//
// Solidity: event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*AccessControlledOffchainAggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorRoundRequestedIterator{contract: _AccessControlledOffchainAggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

// WatchRoundRequested is a free log subscription operation binding the contract event 0x3ea16a923ff4b1df6526e854c9e3a995c43385d70e73359e10623c74f0b52037.
//
// Solidity: event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorRoundRequested)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseRoundRequested(log types.Log) (*AccessControlledOffchainAggregatorRoundRequested, error) {
	event := new(AccessControlledOffchainAggregatorRoundRequested)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlledOffchainAggregatorValidatorConfigSetIterator is returned from FilterValidatorConfigSet and is used to iterate over the raw logs and unpacked data for ValidatorConfigSet events raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorValidatorConfigSetIterator struct {
	Event *AccessControlledOffchainAggregatorValidatorConfigSet // Event containing the contract specifics and raw log

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
func (it *AccessControlledOffchainAggregatorValidatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlledOffchainAggregatorValidatorConfigSet)
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
		it.Event = new(AccessControlledOffchainAggregatorValidatorConfigSet)
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
func (it *AccessControlledOffchainAggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlledOffchainAggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlledOffchainAggregatorValidatorConfigSet represents a ValidatorConfigSet event raised by the AccessControlledOffchainAggregator contract.
type AccessControlledOffchainAggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorConfigSet is a free log retrieval operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*AccessControlledOffchainAggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlledOffchainAggregatorValidatorConfigSetIterator{contract: _AccessControlledOffchainAggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

// WatchValidatorConfigSet is a free log subscription operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *AccessControlledOffchainAggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _AccessControlledOffchainAggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlledOffchainAggregatorValidatorConfigSet)
				if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
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
func (_AccessControlledOffchainAggregator *AccessControlledOffchainAggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*AccessControlledOffchainAggregatorValidatorConfigSet, error) {
	event := new(AccessControlledOffchainAggregatorValidatorConfigSet)
	if err := _AccessControlledOffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
