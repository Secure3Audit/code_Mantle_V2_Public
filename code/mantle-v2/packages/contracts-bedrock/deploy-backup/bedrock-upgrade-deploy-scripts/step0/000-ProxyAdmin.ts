import { DeployFunction } from 'hardhat-deploy/dist/types'

import {assertContractVariable, deploy, deploySleepTime} from '../src/deploy-utils'
import {sleep} from "@eth-optimism/core-utils";

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await sleep(deploySleepTime)
  await deploy({
    hre,
    name: 'ProxyAdmin',
    args: [deployer],
    postDeployAction: async (contract) => {
      // Owner is temporarily set to the deployer. We transfer ownership of the ProxyAdmin to the
      // SystemDictator before we trigger the dictator steps.
      await assertContractVariable(contract, 'owner', deployer)
    },
  })
}

deployFn.tags = ['ProxyAdmin', 'setup', 'l1']

export default deployFn
