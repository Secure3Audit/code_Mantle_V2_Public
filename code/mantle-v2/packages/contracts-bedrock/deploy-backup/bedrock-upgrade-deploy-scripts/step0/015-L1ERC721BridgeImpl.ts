import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../src'
import {
  assertContractVariable,
  deploy, deploySleepTime,
  getContractFromArtifact,
} from '../src/deploy-utils'
import {sleep} from "@eth-optimism/core-utils";

const deployFn: DeployFunction = async (hre) => {
  const L1CrossDomainMessengerProxy = await getContractFromArtifact(
    hre,
    'Proxy__BVM_L1CrossDomainMessenger'
  )

  await sleep(deploySleepTime)
  await deploy({
    hre,
    name: 'L1ERC721Bridge',
    args: [L1CrossDomainMessengerProxy.address, predeploys.L2ERC721Bridge],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxy.address
      )
    },
  })
}

deployFn.tags = ['L1ERC721BridgeImpl', 'setup', 'l1']

export default deployFn
