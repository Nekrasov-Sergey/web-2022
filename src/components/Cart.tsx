import {Link} from "react-router-dom";

export function Cart() {
    return (
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/store" className="mr-2">
                    Freebie shop
                </Link>
                / cart
            </p>
            <p className="text-center font-bold text-6xl text-pink-500">
                Это корзина покупок
            </p>
        </div>
    )
}