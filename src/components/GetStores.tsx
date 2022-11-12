import {useEffect, useReducer} from "react";
import {getJsonStores} from "../modules";

const initialState = {
    stores: [],
}

const reducer = (state: any, action: { type: any; stores: any; }) => {
    switch (action.type) {
        case "FETCH_SUCCESS":
            return {
                stores: action.stores
            }
        default:
            return state
    }
}

export function GetStores() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `store`

    useEffect(() => {
        getJsonStores(url).then((result) => {
            dispatch({type: "FETCH_SUCCESS", stores: result})
        })
    }, [url])

    return state.stores
}