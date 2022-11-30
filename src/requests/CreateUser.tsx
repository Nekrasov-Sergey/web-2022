import {createUser} from "../modules";
import React from "react";
import {Link} from "react-router-dom";

export function CreateUser(name: string, pass: string) {
    const url = `sign_up`

    function Create() {
        createUser(url, name, pass)
    }


    return (
        <>
            <Link to="/login"
                  className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-indigo-700 rounded-md hover:bg-indigo-600 focus:outline-none focus:bg-indigo-600"
                  onClick={() => Create()}
            >
                Зарегестрироваться
            </Link>
        </>
    );

}