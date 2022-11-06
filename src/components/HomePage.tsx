import {Promo} from "./Promo"
import {getJson} from "../modules";
import {useEffect, useState} from "react";

export function HomePage() {
    const [Promos, setPromos] = useState([])

    useEffect(() => {
        async function getAllPromos() {
            setPromos(await getJson("promos"))
        }
        getAllPromos()
    }, [])

    return (
        <div className="bg-yellow-50">
            <p className="ml-4 text-2xl font-normal text-black">
                Freebie shop
            </p>

            <p className="text-center text-6xl font-bold text-pink-500">
                Доступные промокоды на Ноябрь 2022
            </p>

            {Promos.map((promo) => {
                return <Promo promo={promo}/>
            })}
        </div>
    )
}