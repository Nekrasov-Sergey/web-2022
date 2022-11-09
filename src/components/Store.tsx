import {IStore} from '../models'
import {Link} from "react-router-dom"

export function Store(props: { store: IStore }) {
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={props.store.Image}
                 className="place-self-center object-contain h-20 w-20" alt={props.store.Name}
            />

            <p className="text-green-500 place-self-center text-2xl font-bold">
                Скидка {props.store.Discount} рублей

            </p>

            <Link to={`/store/${props.store.Name}`}
                  className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 place-self-center rounded-full text-2xl font-bold"
                  state={{Store: props.store.Name, UUID: props.store.UUID}}
            >
                Купить
            </Link>

            <p className="text-red-600 place-self-center text-3xl font-bold">
                {props.store.Name}
            </p>

            <p className="text-yellow-400 place-self-center text-2xl font-bold">
                {props.store.Price} ₽/шт
            </p>

            <p className="text-orange-500 place-self-center text-2xl font-bold">
                Остаток: {props.store.Quantity} шт
            </p>
        </div>
    )
}

