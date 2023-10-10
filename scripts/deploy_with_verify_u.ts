import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";
import {deployProxy} from "./hardhatUpdate/utils"
import {verifyContract} from "./hardhatUpdate/utils"
import hre from "hardhat";

async function main() {
    //await verifyContract("Params", "0xA85eAb3fEFC21cAFC0851CB493064bbaFb153eF3");
    const accounts = await ethers.getSigners();
    const options = {
      signer: accounts[0]
    };
    /*const aggregatorFactory = await ethers.getContractFactory("AggregatorFactory", accounts[0]);
    const aggregatorFactoryC = await deployWithVerify(await aggregatorFactory.deploy(), [], 'contract/AggregatorFactory.sol:AggregatorFactory');
    console.log('AggregatorFactory address', aggregatorFactoryC.address);*/

    const contractData = await deployProxy("AggregatorFactory",
      {
          options: options,
          args: [
          ]
      },
      true
  );
  console.log('Params address', contractData.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
