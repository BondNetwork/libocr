import { ethers } from "hardhat";
import { BigNumber } from "ethers";
import { BigNumberish } from "ethers";

async function main() {
    //01 部署聚合
    const OffchainAggregator = await ethers.getContractFactory("AccessControlledOffchainAggregator");
    const tx = await OffchainAggregator.deploy('2000','1000',
    '102829','6000','3000',
    "0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e", 1,
        '95780971304118053647396689196894323976171195136475135',
        "0x0000000000000000000000000000000000000000",
        "0x0000000000000000000000000000000000000000",
        18,"test_OffchainAggregator");
    //console.log(tx);
    console.log("OffchainAggregator deployed to " + tx.address);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});