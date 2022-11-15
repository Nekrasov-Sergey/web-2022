import {useReducer} from "react";
import {decreaseQuantity, increaseQuantity} from "../modules";

const increase = "Increase"
const decrease = "Decrease"
const failure="Failure"

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
        case failure:
            return {
                count: 0
            }
        default:
            return state
    }
}

export let asd: number

export function AddToCart(Quantity: number, Store: string) {
    const [state, dispatch] = useReducer(reducer, {count: Quantity});
    const url1 = `cart/increase/${Store}`
    const url2 = `cart/decrease/${Store}`

    function Incr() {
        increaseQuantity(url1).then(result => {
            dispatch({type: increase, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Decr() {
        decreaseQuantity(url2).then(result => {
            dispatch({type: decrease, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    asd = state.count
    return (
        <>
            {" "}{state.count}{" "}
            <button onClick={() => Decr()}>-</button>
            <button onClick={() => Incr()}>+</button>
        </>
    );
}


