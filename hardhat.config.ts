import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-etherscan'
import '@nomiclabs/hardhat-waffle'
import '@openzeppelin/hardhat-upgrades'
import '@typechain/hardhat'
import 'hardhat-abi-exporter'
import 'hardhat-contract-sizer'
import 'solidity-coverage'

const COMPILER_SETTINGS = {
  optimizer: {
    enabled: true,
    runs: 100,
  },
  metadata: {
    bytecodeHash: 'none',
  },
}

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
export default {
  abiExporter: {
    path: './abi',
  },
  paths: {
    artifacts: './artifacts',
    cache: './cache',
    sources: './contract',
    tests: './test',
  },
  typechain: {
    outDir: './typechain',
    target: 'ethers-v5',
  },
  networks: {
    hardhat: {
    },
    localhost: {
      url: "http://127.0.0.1:8545",
      chainId: 31337
    }
  },
  solidity: {
    compilers: [
      {
        version: '0.4.24',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.5.0',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.6.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.7.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.15',
        settings: COMPILER_SETTINGS,
      },
    ],
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: false,
    disambiguatePaths: false,
  },
  mocha: {
    timeout: 100000,
    forbidOnly: Boolean(process.env.CI),
  },
}