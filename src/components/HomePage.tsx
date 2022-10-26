import {Promos} from "../repository/Promos";
import {Promo} from "./Promo";
import React from "react";

export function HomePage() {
    return (
        <div className="container mx-auto max-w-5xl pt-5 flex justify-between">
            {Promos.map((promo, key) => {
                return <Promo promo={promo} key={key}/>
            })}
        </div>
    )
}