import {createContext, useEffect, useReducer} from "react";
import {getFromBackend} from "../modules";
import {CartContext} from "../context";

export const MyContext1 = createContext(CartContext);
const initialState = {cart: []}
const success = "Success"

function reducer(state: any, action: { type: any; payload: any; }) {
    switch (action.type) {
        case success:
            return {
                cart: action.payload
            }
        default:
            return state
    }
}

export function GetCart() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `cart`

    useEffect(() => {
        getFromBackend(url).then((result) => {
            dispatch({type: success, payload: result})
        })
    }, [url])

    return state.cart
}