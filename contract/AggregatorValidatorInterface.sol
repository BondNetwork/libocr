// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface AggregatorValidatorInterface {
  function validate(
    uint256 previousRoundId,
    bytes32 previousAnswer,
    uint256  previousBatchId,
    uint256 currentRoundId,
    bytes32 currentAnswer,
    uint256 currentBatchId
  ) external returns (bool);
}