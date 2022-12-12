import {IOrder} from "../models";
import React, {createContext} from "react";
import {OrdersContext} from "../context";
import {Navbar} from "./Navbar";
import {GetOrders} from "../requests/GetOrders";
import {Order} from "./Order";
import {Link} from "react-router-dom";


export const MyContext = createContext(OrdersContext);

export function OrderPage() {
    let orders = GetOrders()

    return (
        <>
            <Navbar/>

            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/store" className="mr-2">
                        Freebie shop
                    </Link>
                    / orders
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Заказы
                </p>

                <div className="px-2 sm:px-0 flex flex-col gap-4 mx-auto container">
                    {orders.map((order: IOrder, key: any) => {
                        return (
                            <MyContext.Provider value={order} key={key}>
                                <Order/>
                            </MyContext.Provider>
                        )
                    })}
                </div>

                <p className="my-8 text-center">
                    <Link to="/store"
                          className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                    >
                        Обратно на главную
                    </Link>
                </p>
            </div>
        </>
    )
}