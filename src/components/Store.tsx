import {useContext} from "react";
import {ContextStore} from "../requests/GetStores";
import {GetCart1} from "../requests/GetCart1";
import {ChangeCart} from "../requests/ChangeCart";

export function Store() {
    const ctx = useContext(ContextStore)
    let Cart = GetCart1(ctx.UUID)

    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 mob:mt-1 w-1/2 mob:w-11/12 h-40 mob:h-28 py-5 px-5 mob:py-2 mob:px-0 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={ctx.Image}
                 className="place-self-center object-contain h-20 w-20 mob:h-12 mob:w-12" alt={ctx.Name}
            />

            <p className="text-green-500 place-self-center sm:text-2xl text-1xl font-bold">
                Скидка {ctx.Discount} ₽
            </p>

            <p className="text-blue-700 place-self-center sm:text-2xl text-1xl font-bold mob:font-normal">
                В корзину:{" "}{Cart.Quantity}{" "}{ChangeCart(ctx.UUID)}
            </p>

            <p className="text-red-600 place-self-center sm:text-3xl text-1xl font-bold">
                {ctx.Name}
            </p>

            <p className="text-yellow-400 place-self-center sm:text-2xl text-1xl font-bold">
                {ctx.Price} ₽/шт
            </p>

            <p className="text-orange-500 place-self-center sm:text-2xl text-1xl font-bold">
                Остаток: {ctx.Quantity} шт
            </p>
        </div>
    )
}

