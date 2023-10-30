import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";
import {verifyContract} from "./hardhatUpdate/utils";

async function main() {
    const accounts = await ethers.getSigners();
    const options = {
        signer: accounts[0]
    };
    //const aggregatorFactory = await ethers.getContractFactory("AggregatorFactory", accounts[0]);
    //const aggregatorFactoryC = await deployWithVerify(await aggregatorFactory.deploy(), [], 'contract/AggregatorFactory.sol:AggregatorFactory');
    //console.log('AggregatorFactory address', aggregatorFactoryC.address);

    await verifyContract("AccessControlledOffchainAggregator", "0x1be52A9F4d3115edE8c1772deFaAa463fFf433cA",
        [2000, 1000, 102829, 6000, 3000,
        "0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e",
        1,
        95780971304118053647396689196894323976171195136475135,
        "0x0000000000000000000000000000000000000000",
        "0x0000000000000000000000000000000000000000",
        18, "BondOffchainAggregator"])
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
