import React from "react"
import {Link} from "react-router-dom";


export function Navbar() {
    return (
        <nav className="py-4 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
            <p className="pl-40 font-bold inline-block uppercase text-green-300">
                <Link to="/">ГЛАВНАЯ</Link>
            </p>

            <p className="pl-8 font-bold inline-block uppercase text-green-300">
                <Link to="/info">О НАС</Link>
            </p>
        </nav>
    );
}