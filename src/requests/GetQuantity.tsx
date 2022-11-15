import {useEffect, useReducer} from "react";
import {getJsonQuantity} from "../modules";

const initialState = {quantity: "4"}
const success = "Success"

function reducer(state: any, action: { type: any; payload: any; }) {
    switch (action.type) {
        case success:
            return {
                quantity: action.payload
            }
        default:
            return state
    }
}

export function GetQuantity(store: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `cart/${store}`

    useEffect(() => {
        getJsonQuantity(url).then((result) => {
            dispatch({type: success, payload: result})
        })
    }, [url])

    return state.quantity
}