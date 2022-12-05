import {ENDPOINT} from "./App";
import axios from "axios";

export function getFromBackend(url: string) {
    return axios.get(`${ENDPOINT}/${url}`).then(r => r.data)
}

export function getFromBackendToken(url: string) {
    let access_token = getToken()
    return axios.get(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
}

export function deleteFromBackendToken(url: string) {
    let access_token = getToken()
    return axios.delete(`${ENDPOINT}/${url}`, {withCredentials: true, headers: {
            "Authorization": `Bearer ${access_token}`
        }}).then(r => r.data)
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

export function getToken() {
    let tokens = document.cookie.split(' ')
    let access_token = ''
    for (var i = 0; i < tokens.length; i++) {
        if (tokens[i].startsWith("access_token=")) {
            access_token = tokens[i].replace("access_token=", "")
        }
    }
    return access_token.replace(";", "")
}