import {useEffect, useReducer} from "react";
import {getFromBackendToken} from "../modules";

const initialState = {order: []}
const success = "Success"

function reducer(state: any, action: { type: any; payload: any; }) {
    switch (action.type) {
        case success:
            return {
                order: action.payload
            }
        default:
            return state
    }
}

export function GetOrders() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `orders`

    useEffect(() => {
        getFromBackendToken(url).then((result) => {
            dispatch({type: success, payload: result})
        })
    }, [url])

    return state.order
}