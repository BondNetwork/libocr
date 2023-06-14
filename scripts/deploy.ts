import { ethers } from "hardhat";
import { BigNumber } from "ethers";
import { BigNumberish } from "ethers";

async function main() {
    //01 部署聚合
    const OffchainAggregator = await ethers.getContractFactory("AccessControlledOffchainAggregator");
    const tx = await OffchainAggregator.deploy('2000','1000',
    '102829','6000','3000',
    "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853", 1,
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