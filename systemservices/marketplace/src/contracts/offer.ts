import BigNumber from "bignumber.js"
import { Marketplace } from "./Marketplace"
import { Offer } from "../types/offer";
import { fromUnit, parseTimestamp, stringToHex } from "./utils";

const getServiceOffers = async (contract: Marketplace, sid: string): Promise<Offer[]> => {
  const sidHex = stringToHex(sid)
  if (!await contract.methods.isServiceExist(sidHex).call()) {
    throw new Error(`service ${sid} does not exist`)
  }
  const offersLength = new BigNumber(await contract.methods.serviceOffersLength(sidHex).call())
  const offersPromise: Promise<Offer>[] = []
  for (let j = new BigNumber(0); offersLength.isGreaterThan(j); j = j.plus(1)) {
    offersPromise.push(getServiceOffer(contract, sid, j))
  }
  return Promise.all(offersPromise)
}

const getServiceOffer = async (contract: Marketplace, sid: string, offerIndex: BigNumber): Promise<Offer> => {
  const sidHex = stringToHex(sid)
  if (!await contract.methods.isServiceOfferExist(sidHex, offerIndex.toString()).call()) {
    throw new Error(`offer for service '${sid}' with offer index '${offerIndex.toString()}' does not exist`)
  }
  const offer = await contract.methods.serviceOffer(sidHex, offerIndex.toString()).call()
  return {
    offerIndex: offerIndex,
    price: fromUnit(offer.price),
    duration: new BigNumber(offer.duration),
    active: offer.active,
    createTime: parseTimestamp(offer.createTime),
  }
}

export { getServiceOffers, getServiceOffer }
