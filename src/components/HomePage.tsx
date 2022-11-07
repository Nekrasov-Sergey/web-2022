import {Store} from "./Store"
import {getJson} from "../modules";
import {useEffect, useState} from "react";

export function HomePage() {
    const [Stores, setStore] = useState([])

    useEffect(() => {
        async function getAllStores() {
            setStore(await getJson("store"))
        }
        getAllStores()
    }, [])

    return (
        <div className="bg-yellow-50">
            <p className="ml-4 text-2xl font-normal text-black">
                Freebie shop
            </p>

            <p className="text-center text-6xl font-bold text-pink-500">
                Доступные промокоды на Ноябрь 2022
            </p>

            {Stores.map((store) => {
                return <Store store={store}/>
            })}
        </div>
    )
}