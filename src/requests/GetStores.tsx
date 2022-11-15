import {useEffect, useReducer} from "react";
import {getJsonStores} from "../modules";

const initialState = {stores: []}
const success = "Success"

function reducer(state: any, action: { type: any; stores: any; }) {
    switch (action.type) {
        case success:
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
            dispatch({type: success, stores: result})
        })
    }, [url])

    return state.stores
}