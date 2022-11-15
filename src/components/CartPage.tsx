import {Link} from "react-router-dom";
import {ICart} from "../models";
import {Cart} from "./Cart";
import {GetCart, MyContext1} from "../requests/GetCart";
import React from "react";

export function CartPage() {
    return (
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/store" className="mr-2">
                    Freebie shop
                </Link>
                / cart
            </p>

            <p className="text-center font-bold text-5xl text-pink-500">
                Корзина
            </p>

            {GetCart().map((cart: ICart) => {
                return(
                    <MyContext1.Provider value={cart}>
                        <Cart/>
                    </MyContext1.Provider>
                )
            })}

            <p className="my-8 text-center">
                <Link to="/store"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>
        </div>
    )
}

