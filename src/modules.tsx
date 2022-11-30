import {ENDPOINT} from "./App";
import axios from "axios";

export function getFromBackend(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function deleteFromBackend(url: string) {
    return axios.delete(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function createUser(url: string, name: string, pass: string) {
    const body = {name: name, pass: pass}
    return axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true}).then(function (response) {
        console.log(response)
        window.location.replace("/login")
    }).catch(function () {
        window.location.replace("/registration")
    })
}

export function loginUser(url: string, name: string, pass: string) {
    const body = {login: name, password: pass}
    return axios.post(`${ENDPOINT}/${url}`, body, {withCredentials: true}).then(function (response) {
        console.log(response)
        window.location.replace("/store")
    }).catch(function () {
        window.location.replace("/login")
    })
}

export function logoutUser(url: string) {
    let access_token = document.cookie.replace("access_token=", "")
    console.log(access_token)
    return axios.get(`${ENDPOINT}/${url}`, {
        withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }
    }).then(function (r) {
        console.log(r.data)
        window.location.replace("/login")
    })
}

