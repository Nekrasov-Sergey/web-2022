import {addStore} from "../modules";
import React from "react";


export function AddingStore(name: string, discount: number, price: number, quantity: number, promo: string[], image: string) {

    const url = `store`

    function Add() {
        addStore(url, name, discount, price, quantity, promo, image)
    }


    return (
        <>
            <button
                onClick={() => Add()}
                className="border-4 border-red-500 bg-white text-red-500 hover:bg-red-500 hover:text-white py-1 px-2 place-self-center rounded-lg text-2xl font-bold"

                // className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
                Добавить
            </button>
        </>
    );

}