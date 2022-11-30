import {Link} from "react-router-dom"
import {useLocation} from "react-router-dom"
import React from 'react';
import {GetPromo} from "../requests/GetPromo";
import {Navbar} from "./Navbar";

export function Payment() {
    return (
        <>
            <Navbar/>
            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 text-2xl font-normal text-black">
                    <Link to="/store" className="mr-2">
                        Freebie shop
                    </Link>
                    <Link to="/cart" className="mr-2">
                        / cart
                    </Link>
                    / {useLocation().state.Name}
                </p>

                <p className=" font-bold text-5xl text-pink-500 text-center">
                    Страница оплаты
                </p>

                <p className="mt-8 font-medium text-4xl text-green-500 text-center">
                    Ваши промокоды для магазина "{useLocation().state.Name}":
                    <p className="font-bold italic text-4xl text-red-700">
                        {GetPromo()}
                    </p>
                </p>

                <p className="my-8 text-center">
                    <Link to="/cart"
                          className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                    >
                        Обратно в корзину
                    </Link>
                </p>

                <img src="https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Payment_iw36pk.png"
                     width="23%" className="mx-auto" alt="Payment"/>
            </div>
        </>
    )
}