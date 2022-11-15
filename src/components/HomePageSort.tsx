import {Store} from "./Store"
import {GetStores} from "../requests/GetStores";
import {IStore} from "../models";
import {Link, useLocation} from "react-router-dom";
import React from "react";

export function HomePageSort() {

    return (
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                Freebie shop
            </p>

            <p className="text-center text-5xl font-bold text-pink-500">
                Доступные промокоды на Ноябрь 2022
            </p>

            <p className="text-center text-2xl font-normal text-black-500">
                Сортировать по:{" "}

                <Link to="/store/name" className="mr-2" state={{state: "name"}}>
                    названию
                </Link>
                <Link to="/store/price" className="mr-2" state={{state: "price"}}>
                    цене
                </Link>
                <Link to="/store/quantity" className="mr-2" state={{state: "quantity"}}>
                    количеству
                </Link>
            </p>

            {GetStores(useLocation().state.state).map((store: IStore) => {
                return <Store store={store}/>
            })}
        </div>
    )
}