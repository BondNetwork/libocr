// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
pragma abicoder v2;

import {LibOcrFunctions} from  "./libraries/ocr/LibOcrFunctions.sol";
import {LibOcrTypes} from  "./libraries/ocr/LibOcrTypes.sol";

import "./AggregatorFactoryInterface.sol";
import "./AccessControlledOffchainAggregator.sol";

contract AggregatorFactory is Owned {
    event AggregatorCreated(string  pId, address aggregator, uint256 timestamp);
    //mapping(string => bool)  projectIds;
    function createAggregator(string memory pId)
    external
    returns (address) 
    {
        //require(projectIds[pId] == false, "project was created");
        AccessControlledOffchainAggregator aggregatorObj = new AccessControlledOffchainAggregator(
        2000, 1000, 102829, 6000, 3000,
        LinkTokenInterface(0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e),
        1,
        95780971304118053647396689196894323976171195136475135,
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
         18, "BondOffchainAggregator");
        //projectIds[pId] = true;
        emit AggregatorCreated(pId, address(aggregatorObj), block.timestamp);
        return address(aggregatorObj);
    }
}
