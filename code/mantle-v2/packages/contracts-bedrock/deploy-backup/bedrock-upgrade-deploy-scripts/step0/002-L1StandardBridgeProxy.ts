import { DeployFunction } from 'hardhat-deploy/dist/types'

import {assertContractVariable, deploy, deploySleepTime} from '../src/deploy-utils'
import {sleep} from "@eth-optimism/core-utils";

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await sleep(deploySleepTime)
  await deploy({
    hre,
    name: 'Proxy__BVM_L1StandardBridge',
    contract: 'L1ChugSplashProxy',
    args: [deployer],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'getOwner', deployer)
    },
  })
}

deployFn.tags = ['L1StandardBridgeProxy', 'setup', 'l1']

export default deployFn
