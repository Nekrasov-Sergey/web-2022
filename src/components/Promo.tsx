import {IPromo} from '../models'
import {Link} from "react-router-dom"

export interface PromoProps {
    promo: IPromo
}

export function Promo(props: PromoProps) {
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={props.promo.Image}
                 className="place-self-center object-contain h-20 w-20" alt={props.promo.Store}
            />

            <p className="text-green-500 place-self-center text-2xl font-bold">
                Скидка {props.promo.Discount} рублей
            </p>

            <Link to={`${props.promo.Store}`}
                  className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 place-self-center rounded-full text-2xl font-bold"
                  state={{Promo: props.promo.Promo[0], Store: props.promo.Store}}
            >
                Купить
            </Link>

            <p className="text-red-600 place-self-center text-3xl font-bold">
                {props.promo.Store}
            </p>

            <p className="text-yellow-400 place-self-center text-2xl font-bold">
                {props.promo.Price} ₽/шт
            </p>

            <p className="text-orange-500 place-self-center text-2xl font-bold">
                Остаток: {props.promo.Quantity} шт
            </p>
        </div>
    )
}

