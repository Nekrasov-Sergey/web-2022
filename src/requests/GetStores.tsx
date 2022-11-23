import {createContext, useEffect, useReducer} from "react";
import {getFromBackend} from "../modules";
import {StoreContext} from "../context";


export const ContextStore = createContext(StoreContext);
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

export function GetStores(sort: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `store/${sort}`

    useEffect(() => {
        getFromBackend(url).then((result) => {
            dispatch({type: success, stores: result})
        })
    }, [url])

    return state.stores
}