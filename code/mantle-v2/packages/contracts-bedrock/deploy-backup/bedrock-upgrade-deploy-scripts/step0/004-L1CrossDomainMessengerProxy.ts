import { DeployFunction } from 'hardhat-deploy/dist/types'

import {deploy, deploySleepTime, getDeploymentAddress} from '../src/deploy-utils'
import {sleep} from "@eth-optimism/core-utils";

const deployFn: DeployFunction = async (hre) => {
  const addressManager = await getDeploymentAddress(hre, 'Lib_AddressManager')

  await sleep(deploySleepTime)
  await deploy({
    hre,
    name: 'Proxy__BVM_L1CrossDomainMessenger',
    contract: 'ResolvedDelegateProxy',
    args: [addressManager, 'BVM_L1CrossDomainMessenger'],
  })
}

deployFn.tags = ['L1CrossDomainMessengerProxy', 'setup', 'l1']

export default deployFn
