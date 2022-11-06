import {ENDPOINT} from "./App";
import {IPromo} from "./models";


export const getJson = async (url: string) => {
    return await fetch(`${ENDPOINT}/${url}`).then((r) => r.json() as Promise<IPromo[]>)
}