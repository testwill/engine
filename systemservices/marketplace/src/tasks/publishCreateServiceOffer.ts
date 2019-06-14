import { TaskInputs } from "mesg-js/lib/service"
import Web3 from "web3"
import { extractEventFromLogs } from "../contracts/utils";
import { Marketplace } from "../contracts/Marketplace";
import { serviceOfferCreated } from "../contracts/parseEvents";


export default (
  web3: Web3,
  marketplace: Marketplace,
) => async (inputs: TaskInputs): Promise<object> => {
  const receipt = await web3.eth.sendSignedTransaction(inputs.signedTransaction)
  if (receipt.logs === undefined) throw new Error('receipt does not contain logs')
  const decodedLog = extractEventFromLogs(web3, marketplace, 'ServiceOfferCreated', receipt.logs)
  return serviceOfferCreated(decodedLog)
}