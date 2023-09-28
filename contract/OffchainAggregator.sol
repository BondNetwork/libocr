// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
pragma abicoder v2;

import {LibOcrFunctions} from  "./libraries/ocr/LibOcrFunctions.sol";
import {LibOcrTypes} from  "./libraries/ocr/LibOcrTypes.sol";
import "./AccessControllerInterface.sol";
import "./AggregatorV2V3Interface.sol";
import "./AggregatorValidatorInterface.sol";
import "./LinkTokenInterface.sol";
import "./Owned.sol";
import "./OffchainAggregatorBilling.sol";
import "./TypeAndVersionInterface.sol";

/**
  * @notice Onchain verification of reports from the offchain reporting protocol

  * @dev For details on its operation, see the offchain reporting protocol design
  * @dev doc, which refers to this contract as simply the "contract".
*/

contract OffchainAggregator is Owned, OffchainAggregatorBilling, AggregatorV2V3Interface, TypeAndVersionInterface {
  
  using LibOcrFunctions for LibOcrFunctions.SelfData;

  /*struct SelfData{
    mapping (address => LibOcrTypes.Oracle)  s_oracles;
    uint16[maxNumOracles]  s_oracleObservationsCounts;
    uint256[maxNumOracles]  s_gasReimbursementsLinkWei;
    LibOcrTypes.HotVars  s_hotVars;
    mapping (address => address )  s_payees;
    LibOcrTypes.Billing  s_billing;
    LinkTokenInterface  s_linkToken;
    uint32  s_latestConfigBlockNumber;
    uint32  s_configCount;
    address[]  s_signers;
    address[]  s_transmitters;
  }*/

  
  LibOcrTypes.HotVars internal s_hotVars;
  mapping(uint32 /* aggregator round ID */ => LibOcrTypes.Transmission) internal s_transmissions;

  // incremented each time a new config is posted. This count is incorporated
  // into the config digest, to prevent replay attacks.
  uint32 internal s_configCount;
  uint32 internal s_latestConfigBlockNumber; // makes it easier for offchain systems
                                             // to extract config from logs.

  // Lowest answer the system is allowed to report in response to transmissions
  int192 immutable public minAnswer;
  // Highest answer the system is allowed to report in response to transmissions
  int192 immutable public maxAnswer;

  bool internal s_lock;

  constructor(
    LibOcrTypes.InitOCR memory data
  )
    OffchainAggregatorBilling(data._maximumGasPrice, data._reasonableGasPrice, data._microLinkPerEth,
      data._linkGweiPerObservation, data._linkGweiPerTransmission,data._link,
      data._billingAccessController
    )
  {
    decimals = data._decimals;
    s_description = data._description;
    setRequesterAccessController(data._requesterAccessController);
    setValidatorConfig(AggregatorValidatorInterface(0x0), 0);
    minAnswer = data._minAnswer;
    maxAnswer = data._maxAnswer;
  }

  /*
   * Versioning
   */
  function typeAndVersion()
    external
    override
    pure
    virtual
    returns (string memory)
  {
    return "OffchainAggregator 5.0.0";
  }

  // Reverts transaction if config args are invalid
  modifier checkConfigValid (
    uint256 _numSigners, uint256 _numTransmitters, uint256 _threshold
  ) {
    require(_numSigners <= maxNumOracles, "1");
    require(_threshold > 0, "2");
    require(
      _numSigners == _numTransmitters,
      "3"
    );
    require(_numSigners > 3*_threshold, "3");
    _;
  }

  /**
   * @notice sets offchain reporting protocol configuration incl. participating oracles
   * @param _signers addresses with which oracles sign the reports
   * @param _transmitters addresses oracles use to transmit the reports
   * @param _threshold number of faulty oracles the system can tolerate
   * @param _encodedConfigVersion version number for offchainEncoding schema
   * @param _encoded encoded off-chain oracle configuration
   */
  function setConfig(
    address[] calldata _signers,
    address[] calldata _transmitters,
    uint8 _threshold,
    uint64 _encodedConfigVersion,
    bytes calldata _encoded
  )
    external
    checkConfigValid(_signers.length, _transmitters.length, _threshold)
    onlyOwner()
  {

    /*while (s_signers.length != 0) { // remove any old signer/transmitter addresses
      uint lastIdx = s_signers.length - 1;
      address signer = s_signers[lastIdx];
      address transmitter = s_transmitters[lastIdx];
      payOracle(transmitter);
      delete s_oracles[signer];
      delete s_oracles[transmitter];
      s_signers.pop();
      s_transmitters.pop();
    }

    for (uint i = 0; i < _signers.length; i++) { // add new signer/transmitter addresses
      require(
        s_oracles[_signers[i]].role == Role.Unset,
        "4"
      );
      s_oracles[_signers[i]] = Oracle(uint8(i), Role.Signer);
      require(s_payees[_transmitters[i]] != address(0), "23");
      require(
        s_oracles[_transmitters[i]].role == Role.Unset,
        "5"
      );
      s_oracles[_transmitters[i]] = Oracle(uint8(i), Role.Transmitter);
      s_signers.push(_signers[i]);
      s_transmitters.push(_transmitters[i]);
    }
    s_hotVars.threshold = _threshold;
    uint32 previousConfigBlockNumber = s_latestConfigBlockNumber;
    s_latestConfigBlockNumber = uint32(block.number);
    s_configCount += 1;
    uint64 configCount = s_configCount;
    {
      s_hotVars.latestConfigDigest = configDigestFromConfigData(
        address(this),
        configCount,
        _signers,
        _transmitters,
        _threshold,
        _encodedConfigVersion,
        _encoded
      );
      s_hotVars.latestEpochAndRound = 0;
    }
    emit LibOcrTypes.ConfigSet(
      previousConfigBlockNumber,
      configCount,
      _signers,
      _transmitters,
      _threshold,
      _encodedConfigVersion,
      _encoded
    );*/
    /*struct SelfData{
    mapping (address => LibOcrTypes.Oracle)  s_oracles;
    uint16[maxNumOracles]  s_oracleObservationsCounts;
    uint256[maxNumOracles]  s_gasReimbursementsLinkWei;
    LibOcrTypes.HotVars  s_hotVars;
    mapping (address => address )  s_payees;
    LibOcrTypes.Billing  s_billing;
    LinkTokenInterface  s_linkToken;
    uint32  s_latestConfigBlockNumber;
    uint32  s_configCount;
    address[]  s_signers;
    address[]  s_transmitters;
  }*/
    LibOcrFunctions.SelfData storage data;
    data.s_oracleObservationsCounts = s_oracleObservationsCounts;
    data.s_gasReimbursementsLinkWei = s_gasReimbursementsLinkWei;
    data.s_hotVars = s_hotVars;
    data.s_billing = s_billing;
    data.s_linkToken = s_linkToken;
    data.s_latestConfigBlockNumber = s_latestConfigBlockNumber;
    data.s_configCount = s_configCount;
    data.s_signers = s_signers;
    data.s_transmitters = s_transmitters;
    LibOcrFunctions.ConfigData memory configData = LibOcrFunctions.ConfigData(_signers, _transmitters, _threshold, _encodedConfigVersion, _encoded);
    data.setConfig(s_oracles, s_payees, configData);
  }

  function configDigestFromConfigData(
    address _contractAddress,
    uint64 _configCount,
    address[] calldata _signers,
    address[] calldata _transmitters,
    uint8 _threshold,
    uint64 _encodedConfigVersion,
    bytes calldata _encodedConfig
  ) internal pure returns (bytes16) {
    return LibOcrFunctions.configDigestFromConfigData(_contractAddress, _configCount,
     _signers, _transmitters, _threshold, _encodedConfigVersion, _encodedConfig);
  }
  /**
   * @notice information about current offchain reporting protocol configuration

   * @return configCount ordinal number of current config, out of all configs applied to this contract so far
   * @return blockNumber block at which this config was set
   * @return configDigest domain-separation tag for current config (see configDigestFromConfigData)
   */
  function latestConfigDetails()
    external
    view
    returns (
      uint32 configCount,
      uint32 blockNumber,
      bytes16 configDigest
    )
  {
    return (s_configCount, s_latestConfigBlockNumber, s_hotVars.latestConfigDigest);
  }

  /**
   * @return list of addresses permitted to transmit reports to this contract

   * @dev The list will match the order used to specify the transmitter during setConfig
   */
  function transmitters()
    external
    view
    returns(address[] memory)
  {
      return s_transmitters;
  }

  LibOcrTypes.ValidatorConfig private s_validatorConfig;

  /**
   * @notice validator configuration
   * @return validator validator contract
   * @return gasLimit gas limit for validate calls
   */
  function validatorConfig()
    external
    view
    returns (AggregatorValidatorInterface validator, uint32 gasLimit)
  {
    LibOcrTypes.ValidatorConfig memory vc = s_validatorConfig;
    return (vc.validator, vc.gasLimit);
  }

  /**
   * @notice sets validator configuration
   * @dev set _newValidator to 0x0 to disable validate calls
   * @param _newValidator address of the new validator contract
   * @param _newGasLimit new gas limit for validate calls
   */
  function setValidatorConfig(AggregatorValidatorInterface _newValidator, uint32 _newGasLimit)
    public
    onlyOwner()
  {
    LibOcrTypes.ValidatorConfig memory previous = s_validatorConfig;

    if (previous.validator != _newValidator || previous.gasLimit != _newGasLimit) {
      s_validatorConfig = LibOcrTypes.ValidatorConfig({
        validator: _newValidator,
        gasLimit: _newGasLimit
      });

      emit LibOcrTypes.ValidatorConfigSet(previous.validator, previous.gasLimit, _newValidator, _newGasLimit);
    }
  }

  uint256 private constant CALL_WITH_EXACT_GAS_CUSHION = 5_000;

  /*
   * requestNewRound logic
   */

  AccessControllerInterface internal s_requesterAccessController;


  /**
   * @notice address of the requester access controller contract
   * @return requester access controller address
   */
  function requesterAccessController()
    external
    view
    returns (AccessControllerInterface)
  {
    return s_requesterAccessController;
  }

  /**
   * @notice sets the requester access controller
   * @param _requesterAccessController designates the address of the new requester access controller
   */
  function setRequesterAccessController(AccessControllerInterface _requesterAccessController)
    public
    onlyOwner()
  {
    AccessControllerInterface oldController = s_requesterAccessController;
    if (_requesterAccessController != oldController) {
      s_requesterAccessController = AccessControllerInterface(_requesterAccessController);
      emit LibOcrTypes.RequesterAccessControllerSet(oldController, _requesterAccessController);
    }
  }

  /**
   * @notice immediately requests a new round
   * @return the aggregatorRoundId of the next round. Note: The report for this round may have been
   * transmitted (but not yet mined) *before* requestNewRound() was even called. There is *no*
   * guarantee of causality between the request and the report at aggregatorRoundId.
   */
  function requestNewRound() external returns (uint80) {
    require(msg.sender == owner || s_requesterAccessController.hasAccess(msg.sender, msg.data),
      "6");

    LibOcrTypes.HotVars memory hotVars = s_hotVars;

    emit LibOcrTypes.RoundRequested(
      msg.sender,
      hotVars.latestConfigDigest,
      uint32(s_hotVars.latestEpochAndRound >> 8),
      uint8(s_hotVars.latestEpochAndRound)
    );
    return hotVars.latestAggregatorRoundId + 1;
  }

  /**
   * @notice transmit is called to post a new report to the contract
   * @param _report serialized report, which the signatures are signing. See parsing code below for format. The ith element of the observers component must be the index in s_signers of the address for the ith signature
   * @param _rs ith element is the R components of the ith signature on report. Must have at most maxNumOracles entries
   * @param _ss ith element is the S components of the ith signature on report. Must have at most maxNumOracles entries
   * @param _rawVs ith element is the the V component of the ith signature
   */
  function transmit(
    // NOTE: If these parameters are changed, expectedMsgDataLength and/or
    // TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT need to be changed accordingly
    bytes calldata _report,
    bytes32[] calldata _rs, bytes32[] calldata _ss, bytes32 _rawVs // signatures
  )
    external
  {
    uint256 initialGas = gasleft(); // This line must come first
    // Make sure the transmit message-length matches the inputs. Otherwise, the
    // transmitter could append an arbitrarily long (up to gas-block limit)
    // string of 0 bytes, which we would reimburse at a rate of 16 gas/byte, but
    // which would only cost the transmitter 4 gas/byte. (Appendix G of the
    // yellow paper, p. 25, for G_txdatazero and EIP 2028 for G_txdatanonzero.)
    // This could amount to reimbursement profit of 36 million gas, given a 3MB
    // zero tail.
    require(msg.data.length == LibOcrFunctions.expectedMsgDataLength(_report, _rs, _ss),
      "9");
    LibOcrTypes.ReportData memory r; // Relieves stack pressure
    {
      r.hotVars = s_hotVars; // cache read from storage

      bytes32 rawObservers;
      (r.rawReportContext, rawObservers, r.observationsRoot) = abi.decode(
        _report, (bytes32, bytes32, LibOcrTypes.ProjectTaskData[])
      );

      // rawReportContext consists of:
      // 11-byte zero padding
      // 16-byte configDigest
      // 4-byte epoch
      // 1-byte round

      bytes16 configDigest = bytes16(r.rawReportContext << 88);
      require(
        r.hotVars.latestConfigDigest == configDigest,
        "10"
      );

      uint40 epochAndRound = uint40(uint256(r.rawReportContext));

      // direct numerical comparison works here, because
      //
      //   ((e,r) <= (e',r')) implies (epochAndRound <= epochAndRound')
      //
      // because alphabetic ordering implies e <= e', and if e = e', then r<=r',
      // so e*256+r <= e'*256+r', because r, r' < 256
      require(r.hotVars.latestEpochAndRound < epochAndRound, "11");

      require(_rs.length > r.hotVars.threshold, "12");
      require(_rs.length <= maxNumOracles, "13");
      require(_ss.length == _rs.length, "14");
      require(r.observationsRoot.length <= maxNumOracles,
              "15");
      require(r.observationsRoot.length > 2 * r.hotVars.threshold,
              "16");

      // Copy signature parities in bytes32 _rawVs to bytes r.v
      r.vs = new bytes(_rs.length);
      for (uint8 i = 0; i < _rs.length; i++) {
        r.vs[i] = _rawVs[i];
      }

      // Copy observer identities in bytes32 rawObservers to bytes r.observers
      r.observers = new bytes(r.observationsRoot.length);
      bool[maxNumOracles] memory seen;
      for (uint8 i = 0; i < r.observationsRoot.length; i++) {
        uint8 observerIdx = uint8(rawObservers[i]);
        require(!seen[observerIdx], "17");
        seen[observerIdx] = true;
        r.observers[i] = rawObservers[i];
      }

      LibOcrTypes.Oracle memory transmitter = s_oracles[msg.sender];
      require( // Check that sender is authorized to report
        transmitter.role == LibOcrTypes.Role.Transmitter &&
        msg.sender == s_transmitters[transmitter.index],
        "18"
      );
      // record epochAndRound here, so that we don't have to carry the local
      // variable in transmit. The change is reverted if something fails later.
      r.hotVars.latestEpochAndRound = epochAndRound;
    }

    { // Verify signatures attached to report
      bytes32 h = keccak256(_report);
      bool[maxNumOracles] memory signed;

      LibOcrTypes.Oracle memory o;
      for (uint i = 0; i < _rs.length; i++) {
        address signer = ecrecover(h, uint8(r.vs[i])+27, _rs[i], _ss[i]);
        o = s_oracles[signer];
        require(o.role == LibOcrTypes.Role.Signer, "19");
        require(!signed[o.index], "20");
        signed[o.index] = true;
      }
    }

    { 
      LibOcrTypes.ProjectTaskData  memory  root = r.observationsRoot[r.observationsRoot.length/2];
      //require(minAnswer <= median && median <= maxAnswer, "median is out of min-max range");
      r.hotVars.latestAggregatorRoundId++;
      LibOcrTypes.Transmission storage item = s_transmissions[r.hotVars.latestAggregatorRoundId];
      for (uint32 i = 0; i < uint32(root.taskItems.length); i++) {
        item.merkleRoot.taskItems.push(root.taskItems[i]);
      }
      item.merkleRoot.batchId = root.batchId;
      item.merkleRoot.projectId = root.projectId;
      item.merkleRoot.taskCount = root.taskCount;      
      item.timestamp = uint64(block.timestamp);

      //s_transmissions[r.hotVars.latestAggregatorRoundId] =
      //  Transmission(root, uint64(block.timestamp));  
      emit LibOcrTypes.NewTransmission(
        r.hotVars.latestAggregatorRoundId,
        item.merkleRoot.projectId,
        msg.sender,
        r.observers,
        r.rawReportContext
      );
      // Emit these for backwards compatability with offchain consumers
      // that only support legacy events
      emit NewRound(
        r.hotVars.latestAggregatorRoundId,
        address(0x0), // use zero address since we don't have anybody "starting" the round here
        block.timestamp
      );
      emit AnswerUpdated(
        root.batchId,
        r.hotVars.latestAggregatorRoundId,
        block.timestamp
      );
      //validateAnswer(r.hotVars.latestAggregatorRoundId, root);
    }
    s_hotVars = r.hotVars;
    assert(initialGas < LibOcrTypes.maxUint32);
    reimburseAndRewardOracles(uint32(initialGas), r.observers);
  }

  /*
   * v2 Aggregator interface
   */

  /**
   * @notice median from the most recent report
   */
  function latestAnswer()
    public
    override
    view
    virtual
    returns (int256)
  {
    return 1;
  }

  /**
   * @notice timestamp of block in which last report was transmitted
   */
  function latestTimestamp()
    public
    override
    view
    virtual
    returns (uint256)
  {
    return s_transmissions[s_hotVars.latestAggregatorRoundId].timestamp;
  }

  /**
   * @notice Aggregator round (NOT OCR round) in which last report was transmitted
   */
  function latestRound()
    public
    override
    view
    virtual
    returns (uint256)
  {
    return s_hotVars.latestAggregatorRoundId;
  }

  /**
   * @notice median of report from given aggregator round (NOT OCR round)
   * @param _roundId the aggregator round of the target report
   */
  function getAnswer(uint256 _roundId)
    public
    override
    view
    virtual
    returns (int256)
  {
    if (_roundId > 0xFFFFFFFF) { return 0; }
    return 1;
  }

  /**
   * @notice timestamp of block in which report from given aggregator round was transmitted
   * @param _roundId aggregator round (NOT OCR round) of target report
   */
  function getTimestamp(uint256 _roundId)
    public
    override
    view
    virtual
    returns (uint256)
  {
    if (_roundId > 0xFFFFFFFF) { return 0; }
    return s_transmissions[uint32(_roundId)].timestamp;
  }

  /*
   * v3 Aggregator interface
   */

  string constant private V3_NO_DATA_ERROR = "No data present";

  /**
   * @return answers are stored in fixed-point format, with this many digits of precision
   */
  uint8 immutable public override decimals;

  /**
   * @notice aggregator contract version
   */
  uint256 constant public override version = 4;

  string internal s_description;

  /**
   * @notice human-readable description of observable this contract is reporting on
   */
  function description()
    public
    override
    view
    virtual
    returns (string memory)
  {
    return s_description;
  }

  /**
   * @notice details for the given aggregator round
   * @param _roundId target aggregator round (NOT OCR round). Must fit in uint32
   * @return roundId _roundId
   * @return answer median of report from given _roundId
   * @return startedAt timestamp of block in which report from given _roundId was transmitted
   * @return updatedAt timestamp of block in which report from given _roundId was transmitted
   * @return answeredInRound _roundId
   */
  function getRoundData(uint80 _roundId)
    public
    override
    view
    virtual
    returns (
      uint80 roundId,
      int256 answer,
      uint256 startedAt,
      uint256 updatedAt,
      uint80 answeredInRound
    )
  {
    require(_roundId <= 0xFFFFFFFF, LibOcrTypes.V3_NO_DATA_ERROR);
    LibOcrTypes.Transmission memory transmission = s_transmissions[uint32(_roundId)];
    return (
      _roundId,
      1,
      transmission.timestamp,
      transmission.timestamp,
      _roundId
    );
  }

  /**
   * @notice aggregator details for the most recently transmitted report
   * @return roundId aggregator round of latest report (NOT OCR round)
   * @return answer median of latest report
   * @return startedAt timestamp of block containing latest report
   * @return updatedAt timestamp of block containing latest report
   * @return answeredInRound aggregator round of latest report
   */
  function latestRoundData()
    public
    override
    view
    virtual
    returns (
      uint80 roundId,
      int256 answer,
      uint256 startedAt,
      uint256 updatedAt,
      uint80 answeredInRound
    )
  {
    roundId = s_hotVars.latestAggregatorRoundId;

    // Skipped for compatability with existing FluxAggregator in which latestRoundData never reverts.
    // require(roundId != 0, V3_NO_DATA_ERROR);

    LibOcrTypes.Transmission memory transmission = s_transmissions[uint32(roundId)];
    return (
      roundId,
      1,
      transmission.timestamp,
      transmission.timestamp,
      roundId
    );
  }

  function isLocked()
  external
  override
  view
  virtual
  returns (bool)
  {
    return s_lock;
  }

  function setLock(bool isLock) external override virtual
  {
    s_lock = isLock;
  }

  function latestMerkleRoundData(string calldata taskId)
  external
  override
  view
  virtual
  returns (
    uint80  roundId,
    uint256 batchId,
    bytes32 merkelAnswer,
    uint256 startedAt,
    uint256 updatedAt
  )
  {
    roundId = s_hotVars.latestAggregatorRoundId;
    LibOcrTypes.Transmission memory transmission = s_transmissions[uint32(roundId)];
    for(uint32 nIndex = 0; nIndex < uint32(transmission.merkleRoot.taskItems.length); nIndex++){
      if(LibOcrFunctions.isEqual(transmission.merkleRoot.taskItems[nIndex].tId, taskId)){
        return (
            roundId,
            transmission.merkleRoot.batchId,
            transmission.merkleRoot.taskItems[nIndex].tMerkleRoot,
            transmission.timestamp,
            transmission.timestamp
          );
      }
    }
  }
}
