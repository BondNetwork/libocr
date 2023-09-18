// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
import "./AggregatorFactoryInterface.sol";
import "./AccessControlledOffchainAggregator.sol";

contract AggregatorFactory is Owned, AggregatorFactoryInterface{
    function createAggregator(address deployer, string memory pId)
    external
    override
    virtual returns (address) 
    {
        require(deployer == msg.sender);
        AccessControlledOffchainAggregator aggregatorObj = new AccessControlledOffchainAggregator(
        2000, 1000, 102829, 6000, 3000,
        LinkTokenInterface(0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e),
        1,
        95780971304118053647396689196894323976171195136475135,
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
         18, "BondOffchainAggregator");
        emit AggregatorCreated(deployer, pId, address(aggregatorObj), block.timestamp);
        return address(aggregatorObj);
    }
}
