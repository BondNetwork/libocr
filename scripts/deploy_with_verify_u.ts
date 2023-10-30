import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";
import {deployProxy, upgradeProxy} from "./hardhatUpdate/utils"
import {verifyContract} from "./hardhatUpdate/utils"
import hre from "hardhat";

async function main() {
    const accounts = await ethers.getSigners();
    const options = {
      signer: accounts[0]
    };

    /*const contractData = await deployProxy("AggregatorFactory",
      {
          options: options,
          args: [
          ]
      },
      true
    );
    console.log('upgradeContractData address', contractData.address);*/

    const upgradeContractData = await upgradeProxy(
        "0xDc1716E795821b1b39ac912Bc180c8fd5A0404Ab",
        "AggregatorFactory",
        {
            options: options,
            args: [
            ]
        },
        true
    );
  console.log('upgradeContractData address', upgradeContractData.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
