import {useReducer} from "react";
import {getFromBackend} from "../modules";

const increase = "Increase"
const decrease = "Decrease"
const del = "Delete"
const failure = "Failure"

function reducer(state: any, action: { type: any; payload?: any; }) {
    switch (action.type) {
        case increase:
            return {
                count: action.payload
            }
        case decrease:
            return {
                count: action.payload
            }
        case del:
            return {
                count: 0
            }
        case failure:
            return {
                count: 0
            }
        default:
            return state
    }
}


export function AddToCart(Quantity: number, Store: string) {
    const [state, dispatch] = useReducer(reducer, {count: Quantity});
    const url1 = `cart/increase/${Store}`
    const url2 = `cart/decrease/${Store}`
    const url3 = `cart/delete/${Store}`

    function Incr() {
        getFromBackend(url1).then(result => {
            dispatch({type: increase, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Decr() {
        getFromBackend(url2).then(result => {
            dispatch({type: decrease, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Del() {
        getFromBackend(url3).then(() => {
            dispatch({type: del})
        }).catch(() => {
            dispatch({type: failure})
        })
    }
    return (
        <>
            {" "}{state.count}{" "}
            <button onClick={() => Decr()}>-</button>
            <button onClick={() => Incr()}>+</button>
            <button onClick={() => Del()}>DEL</button>
        </>
    );
}


