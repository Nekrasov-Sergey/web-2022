import {useEffect, useReducer} from "react";
import {getJsonStore} from "../modules";

const initialState = {store: ""}
const success = "Success"

function reducer(state: any, action: { type: any; store: any; }) {
    switch (action.type) {
        case success:
            return {
                store: action.store
            }
        default:
            return state
    }
}

export function GetStore(uuid: string) {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `store/${uuid}`

    useEffect(() => {
        getJsonStore(url).then((result) => {
            dispatch({type: success, store: result})
        })
    }, [url])

    return state.store

}