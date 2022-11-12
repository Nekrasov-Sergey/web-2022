import React, {useReducer} from "react";
import {useLocation} from "react-router-dom";
import {getJsonPromo} from "../modules";

const initialState = {
    promo: "",
};

const reducer = (state: any, action: { type: any; payload: any; }) => {
    switch (action.type) {
        case "FETCH_SUCCESS":
            return {
                promo: action.payload
            }
        default:
            return {
                state
            }
    }
}

export function GetPromo() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `store/promo/${useLocation().state.UUID}`

    function Promo() {
        getJsonPromo(url).then((result) => {
            dispatch({type: "FETCH_SUCCESS", payload: result})
        })
    }

    return (
        <p>
            <button onClick={() => Promo()}>
                Показать:
            </button>
            <>
                {" "}{state.promo}
            </>
        </p>
    )
}
