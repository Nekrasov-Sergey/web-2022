import React from "react"
import {Link} from "react-router-dom";

export function Payment() {
    return (
        <div className="text-center">
            <p className=" font-normal text-6xl text-indigo-700">
                Страница оплаты
            </p>

            <p className="mt-8 font-normal text-4xl text-indigo-700">
                Ваш промокод:
            </p>

            <p className="mt-8">
                <Link to="/"
                      className="border-2 border-blue-500 text-blue-700 hover:bg-blue-500 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>

        </div>
    )
}