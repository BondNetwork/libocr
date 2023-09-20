// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
interface AggregatorFactoryInterface {
  function createAggregator(string memory pId) external returns (address);
  event AggregatorCreated(string  pId, address aggregator, uint256 timestamp);
}
