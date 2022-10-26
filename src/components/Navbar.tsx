import React from "react"
import {Link} from "react-router-dom";


export function Navbar() {
    return (
        <nav className="relative flex flex-wrap items-center px-2 py-3 bg-gray-600 mb-3">
            <p className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 uppercase text-white">
                <Link to="/">Home Page</Link>
            </p>

            <p className="text-sm font-bold leading-relaxed inline-block mr-4 py-2 uppercase text-white">
                <Link to="/info">About Us</Link>
            </p>
        </nav>
    );
}