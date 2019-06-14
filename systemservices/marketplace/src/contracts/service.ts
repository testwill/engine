import BigNumber from "bignumber.js"
import { Marketplace } from "./Marketplace"
import { Service } from "../types/service";
import { hexToString, parseTimestamp, stringToHex } from "./utils";

const getAllServices = async (contract: Marketplace): Promise<Service[]> => {
  const servicesLength = new BigNumber(await contract.methods.servicesLength().call())
  const servicesPromise: Promise<Service>[] = []
  for (let i = new BigNumber(0); servicesLength.isGreaterThan(i); i = i.plus(1)) {
    servicesPromise.push(getServiceWithIndex(contract, i))
  }
  return Promise.all(servicesPromise)
}

const getServiceWithIndex = async (contract: Marketplace, serviceIndex: BigNumber): Promise<Service> => {
  const sidHashed = await contract.methods.servicesList(serviceIndex.toString()).call()
  const service = await contract.methods.services(sidHashed).call()
  return getService(contract, hexToString(service.sid))
}

const getService = async (contract: Marketplace, sid: string): Promise<Service> => {
  await requireServiceExist(contract, sid)
  const service = await contract.methods.service(stringToHex(sid)).call()
  return {
    owner: service.owner,
    sid: sid,
    createTime: parseTimestamp(service.createTime),
  }
}

const isServiceExist = async (contract: Marketplace, sid: string): Promise<boolean> => {
  return contract.methods.isServiceExist(stringToHex(sid)).call()
}

const requireServiceExist = async (contract: Marketplace, sid: string): Promise<any> => {
  if (!await isServiceExist(contract, sid)) {
    throw new Error(`service '${sid}' does not exist`)
  }
}

export {
  getAllServices,
  getService,
  isServiceExist,
  requireServiceExist
}