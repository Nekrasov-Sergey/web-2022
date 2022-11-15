import {ICart} from '../models'
import {AddToCart} from "../requests/AddToCart";
import {GetStore} from "../requests/GetStore";


export function Cart(props: { cart: ICart }) {
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={GetStore(props.cart.Store).Image}
                 className="place-self-center object-contain h-20 w-20" alt={GetStore(props.cart.Store).Name}
            />

            <p className="text-green-500 place-self-center text-2xl font-bold">
                Скидка {GetStore(props.cart.Store).Discount} рублей
            </p>

            <p className="text-blue-700 py-1 px-3 place-self-center text-2xl font-bold">
                Кол-во:{AddToCart(props.cart.Quantity, props.cart.Store)}
            </p>

            <p className="text-red-600 place-self-center text-3xl font-bold">
                {GetStore(props.cart.Store).Name}
            </p>

            <p className="text-yellow-400 place-self-center text-2xl font-bold">
                {GetStore(props.cart.Store).Price} ₽/шт
            </p>

            <p className="text-orange-500 place-self-center text-2xl font-bold">
                Сумма: {GetStore(props.cart.Store).Price*props.cart.Quantity} шт
            </p>
        </div>
    )
}

