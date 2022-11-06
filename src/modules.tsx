import {ENDPOINT} from "./App";

export const getJson = async (url: string) => {
    return await fetch(`${ENDPOINT}/${url}`).then(r => r.json())
}
export const deleteJson = async (url: string) => {
    return await fetch(`${ENDPOINT}/${url}`).then(r => r.json()).then(data => data.promo)
}