import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";
require('@openzeppelin/hardhat-upgrades');

async function main() {
    const [ownerSigner] = await ethers.getSigners();
    const aggregatorFactory = await ethers.getContractFactory("AggregatorFactory");
    const aggregatorFactorySigner =  aggregatorFactory.connect(ownerSigner);
    const aggregatorFactoryC = await aggregatorFactorySigner.deploy();
    console.log('aggregatorFactoryC address', aggregatorFactoryC.address);

    const transparentUpgradeableProxy = await ethers.getContractFactory("TransparentUpgradeableProxy", ownerSigner);
    const merkleDistributorFProxyC = await transparentUpgradeableProxy.deploy(
        aggregatorFactoryC.address,
        "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8",
        [],
    );
    console.log('merkleDistributorFProxyC address', merkleDistributorFProxyC.address);
    const aggregatorFactoryObj = await ethers.getContractAt('AggregatorFactory', "0xc0bC14B4463e552E9384C1458B6CC35E3AECC3ee", ownerSigner);
    const version = await aggregatorFactoryObj.aggregatorFactoryVersion({gasPrice:1000000000, gasLimit:2100000});
    console.log('AggregatorFactory version', version);
    const tx = await aggregatorFactoryObj.createAggregator("abc123333", {gasPrice:1000000000, gasLimit:2100000});
    const receipt = await tx.wait()
    console.log('AggregatorFactory create receipt:', receipt);
    const aaddress = await aggregatorFactoryObj.getAggregatorAddress("abc123333", {gasPrice:1000000000, gasLimit:210000});
    console.log('AggregatorFactory aaddress', aaddress);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
