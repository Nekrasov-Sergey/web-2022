import {ENDPOINT} from "./App";
import axios from "axios";

export function getJsonStores (url: string)  {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getJsonPromo(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data.promo)
}






