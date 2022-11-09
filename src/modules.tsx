import {ENDPOINT} from "./App";
import {useEffect, useState} from "react";
import {useLocation} from "react-router-dom";

export const getJsonStores = async (url: string) => {
    return await fetch(`${ENDPOINT}/${url}`).then(r => r.json())
}

export const getJsonPromo = async (url: string) => {
    return await fetch(`${ENDPOINT}/${url}`).then(r => r.json()).then(data => data.promo)
}

export function AllStores() {
    const [Stores, setStore] = useState([])

    useEffect(() => {
        async function getAllStores() {
            setStore(await getJsonStores("store"))
        }

        getAllStores()
    }, [])
    return (
        Stores
    )
}


export function Promocode() {
    const [Promo, setPromo] = useState([])

    const uuid = `store/promo/${useLocation().state.UUID}`

    useEffect(() => {
        async function getPromo() {
            setPromo(await getJsonPromo(uuid))
        }
        getPromo()
    }, [uuid])

    return (
        <p>
            {Promo}
        </p>
    )
}
