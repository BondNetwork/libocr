// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;
//pragma abicoder v2;
import "./AccessControlledOffchainAggregator.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract AggregatorFactory is Initializable{
    event AggregatorCreated(string  pId, address aggregator, uint256 timestamp);
    event AggregatorUpdate(string  pId, address aggregator, uint256 timestamp);
    struct AggregatorItem {
     bool isInit;
     address aggregator;
    }
    mapping(string => AggregatorItem) private projectIds;
    
    function initialize (
    ) external initializer{

    }

    function createAggregator(string memory pId)
    external
    {
        require(projectIds[pId].isInit == false, "project was created");
        address aggregatorObj = deployAggregator();
        projectIds[pId] = AggregatorItem(true, aggregatorObj);
        emit AggregatorCreated(pId, aggregatorObj, block.timestamp);
    }
    
    function updateAggregator(string memory pId)
    external
    returns (address) {
        require(projectIds[pId].isInit == true, "pid is not exist");
        address aggregatorObj = deployAggregator();
        projectIds[pId].aggregator = aggregatorObj;
        emit AggregatorUpdate(pId, aggregatorObj, block.timestamp);
        return address(aggregatorObj);
    }

    function deployAggregator()
    internal
    returns (address) {
        AccessControlledOffchainAggregator aggregatorObj = new AccessControlledOffchainAggregator(
        2000, 1000, 102829, 6000, 3000,
        BondTokenInterface(0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e),
        1,
        95780971304118053647396689196894323976171195136475135,
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
        AccessControllerInterface(0x0000000000000000000000000000000000000000),
         18, "BondOffchainAggregator");
        return address(aggregatorObj);
    }

    function aggregatorFactoryVersion()
    external
    pure
    returns (string memory)
    {
        return "1.0.3";
    }

    function getAggregatorAddress(string memory pId)
    external
    view
    returns (address)
    {
        return projectIds[pId].aggregator;
    }
}
