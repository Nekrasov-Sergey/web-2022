import {AddToCart} from "../requests/AddToCart";
import {useContext} from "react";
import {MyContext} from "../requests/GetStores";

export function Store() {
    const ctx = useContext(MyContext)
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={ctx.Image}
                 className="place-self-center object-contain h-20 w-20" alt={ctx.Name}
            />

            <p className="text-green-500 place-self-center text-2xl font-bold">
                Скидка {ctx.Discount} ₽
            </p>

            <p className="text-blue-700 place-self-center text-2xl font-bold">
                В корзину:{AddToCart(0, ctx.UUID)}
            </p>

            <p className="text-red-600 place-self-center text-3xl font-bold">
                {ctx.Name}
            </p>

            <p className="text-yellow-400 place-self-center text-2xl font-bold">
                {ctx.Price} ₽/шт
            </p>

            <p className="text-orange-500 place-self-center text-2xl font-bold">
                Остаток: {ctx.Quantity} шт
            </p>
        </div>
    )
}

