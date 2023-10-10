// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;

interface AggregatorMerkleInterface {
  function isLocked() external view returns (bool);
  function setLock(bool isLock) external;
  function latestMerkleRoundData(string calldata taskId)
    external
    view
    returns (
      uint80 roundId,
      uint256 batchId,
      bytes32 merkelAnswer,
      uint256 startedAt,
      uint256 updatedAt
    );
}