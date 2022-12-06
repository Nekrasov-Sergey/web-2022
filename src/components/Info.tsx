import {Link} from "react-router-dom";
import {Navbar} from "./Navbar";
import React from "react";

export function Info() {
    return (
        <>
            <Navbar/>
            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/store" className="mr-2">
                        Freebie shop
                    </Link>
                    / info
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Freebie shop
                </p>

                <p className="text-center sm:mt-4 mx-8 font-medium mob:font-normal text-3xl mob:text-2xl text-indigo-700">
                    Это магазин промокодов, где вы можете купить промокоды для многих популярных магазинов.
                    Экономьте деньги вместе с нами!
                </p>

                <img src="https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Discount_hgs7rl.png"
                     width="29%" className="mx-auto" alt="Discount"/>
            </div>
        </>
    )
}

