// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
import "../../LinkTokenInterface.sol";
import "../../AccessControllerInterface.sol";
import "../../AggregatorValidatorInterface.sol";
pragma abicoder v2;

library LibOcrTypes {
    /*
   * @param _maximumGasPrice highest gas price for which transmitter will be compensated
   * @param _reasonableGasPrice transmitter will receive reward for gas prices under this value
   * @param _microLinkPerEth reimbursement per ETH of gas cost, in 1e-6LINK units
   * @param _linkGweiPerObservation reward to oracle for contributing an observation to a successfully transmitted report, in 1e-9LINK units
   * @param _linkGweiPerTransmission reward to transmitter of a successful report, in 1e-9LINK units
   * @param _link address of the LINK contract
   * @param _minAnswer lowest answer the median of a report is allowed to be
   * @param _maxAnswer highest answer the median of a report is allowed to be
   * @param _billingAccessController access controller for billing admin functions
   * @param _requesterAccessController access controller for requesting new rounds
   * @param _decimals answers are stored in fixed-point format, with this many digits of precision
   * @param _description short human-readable description of observable this contract's answers pertain to
   */
  struct InitOCR{
    uint32 _maximumGasPrice;
    uint32 _reasonableGasPrice;
    uint32 _microLinkPerEth;
    uint32 _linkGweiPerObservation;
    uint32 _linkGweiPerTransmission;
    LinkTokenInterface _link;
    int192 _minAnswer;
    int192 _maxAnswer;
    AccessControllerInterface _billingAccessController;
    AccessControllerInterface _requesterAccessController;
    uint8 _decimals;
    string _description;
  }

  struct TaskItem {
    string   tId;
    bytes32  tMerkleRoot;
  }
  struct ProjectTaskData {
    string  projectId;
    uint64  batchId;
    uint32  taskCount;
    TaskItem []taskItems;
  }

    // Transmission records the median answer from the transmit transaction at
  // time timestamp
  struct Transmission {
    ProjectTaskData merkleRoot;
    uint64 timestamp;
  }

  uint256 constant  maxUint32 = (1 << 32) - 1;

  uint256 constant CALL_WITH_EXACT_GAS_CUSHION = 5_000;

  string  constant  V3_NO_DATA_ERROR = "No data present";

  // Storing these fields used on the hot path in a HotVars variable reduces the
  // retrieval of all of them to a single SLOAD. If any further fields are
  // added, make sure that storage of the struct still takes at most 32 bytes.
  struct HotVars {
    // Provides 128 bits of security against 2nd pre-image attacks, but only
    // 64 bits against collisions. This is acceptable, since a malicious owner has
    // easier way of messing up the protocol than to find hash collisions.
    bytes16 latestConfigDigest;
    uint40 latestEpochAndRound; // 32 most sig bits for epoch, 8 least sig bits for round
    // Current bound assumed on number of faulty/dishonest oracles participating
    // in the protocol, this value is referred to as f in the design
    uint8 threshold;
    // Chainlink Aggregators expose a roundId to consumers. The offchain reporting
    // protocol does not use this id anywhere. We increment it whenever a new
    // transmission is made to provide callers with contiguous ids for successive
    // reports.
    uint32 latestAggregatorRoundId;
  }

    /*
   * Config logic
   */

  /**
   * @notice triggers a new run of the offchain reporting protocol
   * @param previousConfigBlockNumber block in which the previous config was set, to simplify historic analysis
   * @param configCount ordinal number of this config setting among all config settings over the life of this contract
   * @param signers ith element is address ith oracle uses to sign a report
   * @param transmitters ith element is address ith oracle uses to transmit a report via the transmit method
   * @param threshold maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly
   * @param encodedConfigVersion version of the serialization format used for "encoded" parameter
   * @param encoded serialized data used by oracles to configure their offchain operation
   */
  event ConfigSet(
    uint32 previousConfigBlockNumber,
    uint64 configCount,
    address[] signers,
    address[] transmitters,
    uint8 threshold,
    uint64 encodedConfigVersion,
    bytes encoded
  );


    /*
   * On-chain validation logc
   */

  // Configuration for validator
  struct ValidatorConfig {
    AggregatorValidatorInterface validator;
    uint32 gasLimit;
  }

    /**
     * @notice indicates that the validator configuration has been set
     * @param previousValidator previous validator contract
     * @param previousGasLimit previous gas limit for validate calls
     * @param currentValidator current validator contract
     * @param currentGasLimit current gas limit for validate calls
     */
  event ValidatorConfigSet(
    AggregatorValidatorInterface indexed previousValidator,
    uint32 previousGasLimit,
    AggregatorValidatorInterface indexed currentValidator,
    uint32 currentGasLimit
  );

      /**
   * @notice emitted when a new requester access controller contract is set
   * @param old the address prior to the current setting
   * @param current the address of the new access controller contract
   */
  event RequesterAccessControllerSet(AccessControllerInterface old, AccessControllerInterface current);

  /**
   * @notice emitted to immediately request a new round
   * @param requester the address of the requester
   * @param configDigest the latest transmission's configDigest
   * @param epoch the latest transmission's epoch
   * @param round the latest transmission's round
   */
  event RoundRequested(address indexed requester, bytes16 configDigest, uint32 epoch, uint8 round);


    /**
   * @notice indicates that a new report was transmitted
   * @param aggregatorRoundId the round to which this report was assigned
   * @param projectId projectId of the observations attached this report
   * @param transmitter address from which the report was transmitted
   * @param observers observers transmitted with this report
   * @param rawReportContext signature-replay-prevention domain-separation tag
   */
  event NewTransmission(
    uint32 indexed aggregatorRoundId,
    string  projectId,
    address transmitter,
    bytes observers,
    bytes32 rawReportContext
  );

  // Used to relieve stack pressure in transmit
  struct ReportData {
    HotVars hotVars; // Only read from storage once
    bytes observers; // ith element is the index of the ith observer
    ProjectTaskData[] observationsRoot; // ith element is the ith observation
    bytes vs; // jth element is the v component of the jth signature
    bytes32 rawReportContext;
  }

    // The constant-length components of the msg.data sent to transmit.
  // See the "If we wanted to call sam" example on for example reasoning
  // https://solidity.readthedocs.io/en/v0.7.2/abi-spec.html
  uint16 constant TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT =
    4 + // function selector
    32 + // word containing start location of abiencoded _report value
    32 + // word containing location start of abiencoded  _rs value
    32 + // word containing start location of abiencoded _ss value
    32 + // _rawVs value
    32 + // word containing length of _report
    32 + // word containing length _rs
    32 + // word containing length of _ss
    0; // placeholder

  struct Oracle {
    uint8 index; // Index of oracle in s_signers/s_transmitters
    Role role;   // Role of the address which mapped to this struct
  }

  enum Role {
    // No oracle role has been set for address a
    Unset,
    // Signing address for the s_oracles[a].index'th oracle. I.e., report
    // signatures from this oracle should ecrecover back to address a.
    Signer,
    // Transmission address for the s_oracles[a].index'th oracle. I.e., if a
    // report is received by OffchainAggregator.transmit in which msg.sender is
    // a, it is attributed to the s_oracles[a].index'th oracle.
    Transmitter
  }

  // Parameters for oracle payments
  struct Billing {

    // Highest compensated gas price, in ETH-gwei uints
    uint32 maximumGasPrice;

    // If gas price is less (in ETH-gwei units), transmitter gets half the savings
    uint32 reasonableGasPrice;

    // Pay transmitter back this much LINK per unit eth spent on gas
    // (1e-6LINK/ETH units)
    uint32 microLinkPerEth;

    // Fixed LINK reward for each observer, in LINK-gwei units
    uint32 linkGweiPerObservation;

    // Fixed reward for transmitter, in linkGweiPerObservation units
    uint32 linkGweiPerTransmission;
  }

  event OraclePaid(
    address indexed transmitter,
    address indexed payee,
    uint256 amount,
    LinkTokenInterface indexed linkToken
  );

    /**
   * @notice emitted when a transfer of an oracle's payee address has been completed
   * @param transmitter address from which the oracle sends reports to the transmit method
   * @param current the payeee address for the oracle, prior to this setting
   */
  event PayeeshipTransferred(
    address indexed transmitter,
    address indexed previous,
    address indexed current
  );
}
