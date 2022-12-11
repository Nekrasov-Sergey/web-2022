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
                className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-base font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
                Добавить
            </button>
        </>
    );

}