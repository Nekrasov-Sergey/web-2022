import {useEffect, useReducer} from "react";
import {useLocation} from "react-router-dom";
import {getFromBackendToken} from "../modules";

const initialState = {promo: []}
const success = "Success"
const failure = "Failure"

function reducer(state: any, action: { type: any; payload?: any; }) {
    switch (action.type) {
        case success:
            return {
                promo: action.payload
            }
        case failure:
            return {
                promo: "ВСЁ!"
            }
        default:
            return state
    }
}

export function GetPromo() {
    const [state, dispatch] = useReducer(reducer, initialState)
    const url = `store/${useLocation().state.Store}/${useLocation().state.Quantity}`

    useEffect(() => {
        getFromBackendToken(url).then(result => {
            dispatch({type: success, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }, [url])

    return (state.promo)
}
