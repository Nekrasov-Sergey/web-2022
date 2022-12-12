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
                className="border-4 border-red-500 bg-white text-red-500 hover:bg-red-500 hover:text-white py-1 px-2 place-self-center rounded-full text-2xl font-bold"
            >
                Добавить
            </button>
        </>
    );

}