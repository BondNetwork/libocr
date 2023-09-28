// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
import {LibOcrTypes} from  "./LibOcrTypes.sol";
import "../../LinkTokenInterface.sol";
pragma abicoder v2;

library LibOcrFunctions{
  uint256 constant  maxNumOracles = 31;

  uint256 constant  maxUint16 = (1 << 16) - 1;

  uint256 constant  maxUint128 = (1 << 128) - 1;

  struct SelfData{
    uint16[maxNumOracles]  s_oracleObservationsCounts;
    uint256[maxNumOracles]  s_gasReimbursementsLinkWei;
    LibOcrTypes.HotVars  s_hotVars;
    LibOcrTypes.Billing  s_billing;
    LinkTokenInterface  s_linkToken;
    uint32  s_latestConfigBlockNumber;
    uint32  s_configCount;
    address[]  s_signers;
    address[]  s_transmitters;
  }

  struct ConfigData{
    address[]  _signers;
    address[]  _transmitters;
    uint8 _threshold;
    uint64 _encodedConfigVersion;
    bytes  _encoded;
  }

  function isEqual(string memory a, string memory b) external pure returns (bool) {
      bytes memory aa = bytes(a);
      bytes memory bb = bytes(b);
      if (aa.length != bb.length) return false;
      for(uint i = 0; i < aa.length; i ++) {
          if(aa[i] != bb[i]) return false;
      }
      return true;
    }
  function configDigestFromConfigData(
    address _contractAddress,
    uint64 _configCount,
    address[] calldata _signers,
    address[] calldata _transmitters,
    uint8 _threshold,
    uint64 _encodedConfigVersion,
    bytes calldata _encodedConfig
  )internal pure returns (bytes16) {
    return bytes16(keccak256(abi.encode(_contractAddress, _configCount,
      _signers, _transmitters, _threshold, _encodedConfigVersion, _encodedConfig
    )));
  }

    function expectedMsgDataLength(
    bytes calldata _report, bytes32[] calldata _rs, bytes32[] calldata _ss
  ) external pure returns (uint256 length)
  {
    // calldata will never be big enough to make this overflow
    return uint256(LibOcrTypes.TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT) +
      _report.length + // one byte pure entry in _report
      _rs.length * 32 + // 32 bytes per entry in _rs
      _ss.length * 32 + // 32 bytes per entry in _ss
      0; // placeholder
  }

  function min(uint256 a, uint256 b)
    public
    pure
    returns (uint256)
  {
    if (a < b) { return a; }
    return b;
  }

  function saturatingAddUint16(uint16 _x, uint16 _y)
    public
    pure
    returns (uint16)
  {
    return uint16(min(uint256(_x)+uint256(_y), maxUint16));
  }

function oracleRewards(
    bytes memory observers,
    uint16[maxNumOracles] memory observations
  )
    external
    pure
    returns (uint16[maxNumOracles] memory)
  {
    // reward each observer-participant with the observer reward
    for (uint obsIdx = 0; obsIdx < observers.length; obsIdx++) {
      uint8 observer = uint8(observers[obsIdx]);
      observations[observer] = saturatingAddUint16(observations[observer], 1);
    }
    return observations;
  }
  
  // payOracle pays out _transmitter's balance to the corresponding payee, and zeros it out
  function payOracle(
  mapping (address /* signer OR transmitter address */ => LibOcrTypes.Oracle) storage s_oracles,
  mapping (address /* transmitter */ => address /* payment address */) storage s_payees,
  LinkTokenInterface  s_linkToken,
  uint16[maxNumOracles]  storage s_oracleObservationsCounts,
  uint256[maxNumOracles]  storage s_gasReimbursementsLinkWei,
  LibOcrTypes.Billing  storage s_billing,
  address _transmitter)
    public
  {
    LibOcrTypes.Oracle memory oracle = s_oracles[_transmitter];
    uint256 linkWeiAmount = owedPayment(s_billing, s_oracleObservationsCounts, s_gasReimbursementsLinkWei, 
    s_oracles, _transmitter);
    if (linkWeiAmount > 0) {
      address payee = s_payees[_transmitter];
      // Poses no re-entrancy issues, because LINK.transfer does not yield
      // control flow.
      require(s_linkToken.transfer(payee, linkWeiAmount), "insufficient funds");
      s_oracleObservationsCounts[oracle.index] = 1; // "zero" the counts. see var's docstring
      s_gasReimbursementsLinkWei[oracle.index] = 1; // "zero" the counts. see var's docstring
      emit LibOcrTypes.OraclePaid(_transmitter, payee, linkWeiAmount, s_linkToken);
    }
  }

  function setConfig(
    SelfData storage self, 
    mapping (address => LibOcrTypes.Oracle) storage s_oracles,
    mapping (address => address) storage s_payees,
    ConfigData calldata configData
  ) 
    external
    returns (uint32, uint32)
  {
    while (self.s_signers.length != 0) { // remove any old signer/transmitter addresses
      uint lastIdx = self.s_signers.length - 1;
      address signer = self.s_signers[lastIdx];
      address transmitter = self.s_transmitters[lastIdx];
      payOracle(s_oracles, s_payees, self.s_linkToken, self.s_oracleObservationsCounts,
       self.s_gasReimbursementsLinkWei, self.s_billing, transmitter);
      delete s_oracles[signer];
      delete s_oracles[transmitter];
      self.s_signers.pop();
      self.s_transmitters.pop();
    }

    for (uint i = 0; i < configData._signers.length; i++) { // add new signer/transmitter addresses
      require(
        s_oracles[configData._signers[i]].role == LibOcrTypes.Role.Unset,
        "4"
      );
      s_oracles[configData._signers[i]] = LibOcrTypes.Oracle(uint8(i), LibOcrTypes.Role.Signer);
      require(s_payees[configData._transmitters[i]] != address(0), "23");
      require(
        s_oracles[configData._transmitters[i]].role == LibOcrTypes.Role.Unset,
        "5"
      );
      s_oracles[configData._transmitters[i]] = LibOcrTypes.Oracle(uint8(i), LibOcrTypes.Role.Transmitter);
      self.s_signers.push(configData._signers[i]);
      self.s_transmitters.push(configData._transmitters[i]);
    }
    self.s_hotVars.threshold = configData._threshold;
    uint32 previousConfigBlockNumber = self.s_latestConfigBlockNumber;
    self.s_latestConfigBlockNumber = uint32(block.number);
    self.s_configCount += 1;
    uint64 configCount = self.s_configCount;
    {
      self.s_hotVars.latestConfigDigest = configDigestFromConfigData(
        address(this),
        configCount,
        configData._signers,
        configData._transmitters,
        configData._threshold,
        configData._encodedConfigVersion,
        configData._encoded
      );
      self.s_hotVars.latestEpochAndRound = 0;
    }
    emit LibOcrTypes.ConfigSet(
      previousConfigBlockNumber,
      configCount,
      configData._signers,
      configData._transmitters,
      configData._threshold,
      configData._encodedConfigVersion,
      configData._encoded
    );
    return (self.s_latestConfigBlockNumber, self.s_configCount);
  }

  /**
   * @notice query an oracle's payment amount
   * @param _transmitter the transmitter address of the oracle
   */
  function owedPayment(
  LibOcrTypes.Billing  storage s_billing,
  uint16[maxNumOracles] storage s_oracleObservationsCounts,
  uint256[maxNumOracles] storage s_gasReimbursementsLinkWei,
  mapping (address /* signer OR transmitter address */ => LibOcrTypes.Oracle) storage s_oracles,
  address _transmitter)
    public
    view
    returns (uint256)
  {
    LibOcrTypes.Oracle memory oracle = s_oracles[_transmitter];
    if (oracle.role == LibOcrTypes.Role.Unset) { return 0; }
    LibOcrTypes.Billing memory billing = s_billing;
    uint256 linkWeiAmount =
      uint256(s_oracleObservationsCounts[oracle.index] - 1) *
      uint256(billing.linkGweiPerObservation) *
      (1 gwei);
    linkWeiAmount += s_gasReimbursementsLinkWei[oracle.index] - 1;
    return linkWeiAmount;
  }

  function setPayees(
    mapping (address => address) storage s_payees,
    address[] calldata _transmitters,
    address[] calldata _payees
  )
    external
  {
    require(_transmitters.length == _payees.length, "31");

    for (uint i = 0; i < _transmitters.length; i++) {
      address transmitter = _transmitters[i];
      address payee = _payees[i];
      address currentPayee = s_payees[transmitter];
      bool zeroedOut = currentPayee == address(0);
      require(zeroedOut || currentPayee == payee, "32");
      s_payees[transmitter] = payee;

      if (currentPayee != payee) {
        emit LibOcrTypes.PayeeshipTransferred(transmitter, currentPayee, payee);
      }
    }
  }

  function payOracles(
  mapping (address /* transmitter */ => address /* payment address */) storage s_payees,
  LinkTokenInterface  s_linkToken,
  uint16[maxNumOracles]  storage s_oracleObservationsCounts,
  uint256[maxNumOracles]  storage s_gasReimbursementsLinkWei,
  address[] storage s_transmitters,
  LibOcrTypes.Billing  storage s_billing)
    external
  {
    LibOcrTypes.Billing memory billing = s_billing;
    LinkTokenInterface linkToken = s_linkToken;
    uint16[maxNumOracles] memory observationsCounts = s_oracleObservationsCounts;
    uint256[maxNumOracles] memory gasReimbursementsLinkWei =
      s_gasReimbursementsLinkWei;
    address[] memory transmitters = s_transmitters;
    for (uint transmitteridx = 0; transmitteridx < transmitters.length; transmitteridx++) {
      uint256 reimbursementAmountLinkWei = gasReimbursementsLinkWei[transmitteridx] - 1;
      uint256 obsCount = observationsCounts[transmitteridx] - 1;
      uint256 linkWeiAmount =
        obsCount * uint256(billing.linkGweiPerObservation) * (1 gwei) + reimbursementAmountLinkWei;
      if (linkWeiAmount > 0) {
          address payee = s_payees[transmitters[transmitteridx]];
          // Poses no re-entrancy issues, because LINK.transfer does not yield
          // control flow.
          require(linkToken.transfer(payee, linkWeiAmount), "44");
          observationsCounts[transmitteridx] = 1;       // "zero" the counts.
          gasReimbursementsLinkWei[transmitteridx] = 1; // "zero" the counts.
          emit LibOcrTypes.OraclePaid(transmitters[transmitteridx], payee, linkWeiAmount, linkToken);
        }
    }
    // "Zero" the accounting storage variables
    for (uint transmitteridx = 0; transmitteridx < observationsCounts.length; transmitteridx++) {
          s_oracleObservationsCounts[transmitteridx] = observationsCounts[transmitteridx];
    }
    for (uint transmitteridx = 0; transmitteridx < gasReimbursementsLinkWei.length; transmitteridx++) {
          s_gasReimbursementsLinkWei[transmitteridx] = gasReimbursementsLinkWei[transmitteridx];
    }
  }
}