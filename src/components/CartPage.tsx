import {Link} from "react-router-dom";
import {ICart} from "../models";
import {Cart} from "./Cart";
import {GetCart} from "../requests/GetCart";

export function CartPage() {
    return (
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/store" className="mr-2">
                    Freebie shop
                </Link>
                / cart
            </p>

            <p className="text-center font-bold text-6xl text-pink-500">
                Корзина
            </p>

            {GetCart().map((cart: ICart) => {
                return <Cart cart={cart}/>
            })}
        </div>
    )
}

