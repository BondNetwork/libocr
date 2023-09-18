// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
interface AggregatorFactoryInterface {
  function createAggregator(address deployer, string memory pId) external returns (address);
  event AggregatorCreated(address deployer, string  pId, address aggregator, uint256 timestamp);
}
