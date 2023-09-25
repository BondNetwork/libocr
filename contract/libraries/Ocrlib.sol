// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;
library Ocrlib {
  // Reverts transaction if config args are invalid
  modifier checkConfigValid ( uint256  maxNumOracles, 
    uint256 _numSigners, uint256 _numTransmitters, uint256 _threshold
  ) {
    require(_numSigners <= maxNumOracles, "too many signers");
    require(_threshold > 0, "threshold must be positive");
    require(
      _numSigners == _numTransmitters,
      "oracle addresses out of registration"
    );
    require(_numSigners > 3*_threshold, "faulty-oracle threshold too high");
    _;
  }

    /**
   * @dev Reverts if called by anyone other than the contract owner.
   */
  modifier onlyOwner(address owner) {
    require(msg.sender == owner, "Only callable by owner");
    _;
  }

  function isEqual(string memory a, string memory b) public pure returns (bool) {
            bytes memory aa = bytes(a);
            bytes memory bb = bytes(b);
            if (aa.length != bb.length) return false;
            for(uint i = 0; i < aa.length; i ++) {
                if(aa[i] != bb[i]) return false;
            }
            return true;
    }
}