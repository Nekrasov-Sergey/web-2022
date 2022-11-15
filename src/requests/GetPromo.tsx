import {useEffect, useReducer} from "react";
import {useLocation} from "react-router-dom";
import {getJsonPromo} from "../modules";

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
    const url = `store/promo/${useLocation().state.Quantity}/${useLocation().state.Store}`

    useEffect(() => {
        getJsonPromo(url).then(result => {
            dispatch({type: success, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }, [url])

    return (state.promo)
}
