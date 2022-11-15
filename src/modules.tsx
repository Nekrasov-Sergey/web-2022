import {ENDPOINT} from "./App";
import axios from "axios";

export function getJsonStores (url: string)  {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getJsonStore (url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getJsonPromo(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getJsonCart (url: string)  {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getJsonQuantity (url: string)  {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function increaseQuantity(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data.quantity)
}

export function decreaseQuantity(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data.quantity)
}






