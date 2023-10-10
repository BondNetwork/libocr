// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import {BondUpgradeable} from "../upgradeability/BondUpgradeable.sol";

contract Params is BondUpgradeable {
    mapping(string => uint256) private uint256Params;

    event Uint256ParamSetted(string indexed _key,uint256 _value);
    uint32 internal constant REVISION = 10;

    function initialize(
    ) external initializer {
        __BondUpgradeable_init();
    }

    function SetUint256Param(string memory _key,uint256 _value) external{
        uint256Params[_key] = _value;
        emit Uint256ParamSetted(_key,_value);
    }

    function GetUint256Param(string memory _key)public view returns(uint256){
        return uint256Params[_key];
    }
    
    function getVersion() external pure returns (uint32) {
        return REVISION;
    }

}