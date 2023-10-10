import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";

async function main() {
    const accounts = await ethers.getSigners();
    const aggregatorFactory = await ethers.getContractFactory("AggregatorFactory", accounts[0]);
    const aggregatorFactoryC = await deployWithVerify(await aggregatorFactory.deploy(), [], 'contract/AggregatorFactory.sol:AggregatorFactory');
    console.log('AggregatorFactory address', aggregatorFactoryC.address);
    
    /*const transparentUpgradeableProxy = await ethers.getContractFactory("TransparentUpgradeableProxy", accounts[0]);
    const merkleDistributorFProxy = await deployWithVerify(
      await transparentUpgradeableProxy.deploy(
        aggregatorFactoryC.address,
        "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8",
        [],
      ),
      [
        aggregatorFactoryC.address,
        "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8",
        [],
      ],
      'contract/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy'
  );
  console.log('merkleDistribmerkleDistributorFProxyutorF address', merkleDistributorFProxy.address);
  */
  /*const aggregatorFactoryObj = await ethers.getContractAt('AggregatorFactory', merkleDistributorFProxy.address, accounts[0]);

  const tx = await aggregatorFactoryObj.createAggregator("zxc123456",{gasPrice:1000000000, gasLimit:210000});
  const replicate = await tx.wait();
  console.log('createagg tx', replicate);

  const aaddress = await aggregatorFactoryObj.getAggregatorAddress("zxc123456");
  console.log('createagg aaddress', aaddress);*/
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
