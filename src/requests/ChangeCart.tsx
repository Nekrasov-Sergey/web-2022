import {useReducer} from "react";
import {deleteFromBackendToken, getFromBackendToken} from "../modules";
import DeleteIcon from '@mui/icons-material/Delete';

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

export function ChangeCart(Store: string) {
    const [dispatch] = useReducer(reducer, {count: 0});
    const url1 = `cart/increase/${Store}`
    const url2 = `cart/decrease/${Store}`
    const url3 = `cart/delete/${Store}`

    function Incr() {
        getFromBackendToken(url1).then(result => {
            dispatch({type: increase, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Decr() {
        getFromBackendToken(url2).then(result => {
            dispatch({type: decrease, payload: result})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    function Del() {
        deleteFromBackendToken(url3).then(() => {
            dispatch({type: del})
        }).catch(() => {
            dispatch({type: failure})
        })
    }

    return (
        <form className="inline">
            <button onClick={() => Decr()}>-</button>
            <button onClick={() => Incr()}>+</button>
            <button onClick={() => Del()}>
                <DeleteIcon/>
            </button>
        </form>
    )
}


