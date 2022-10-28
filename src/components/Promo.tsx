import {IPromo} from '../models'
import {useState} from "react";

interface PromoProps {
    promo: IPromo
}

export function Promo(props: PromoProps) {
    const [ShowDescription, setShowDescription] = useState<boolean>(false)
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-1/2 h-40 py-5 px-5 rounded-lg grid grid-rows-2 grid-cols-3 bg-white"
        >
            <img src={process.env.PUBLIC_URL + props.promo.image}
                 className="row-span-2 place-self-center object-contain h-36 w-36" alt={props.promo.store}/>
            <p className="text-red-600 place-self-center text-3xl font-bold">{props.promo.store}</p>
            <button
                className="border-2 border-blue-500 text-blue-700 hover:bg-blue-500 hover:text-white rounded-full font-bold"
                onClick={() => setShowDescription((prevState) => !prevState)}
            >
                {!ShowDescription ? <div>Показать описание</div> : <div>Скрыть описание</div>}
            </button>
            <p className="text-green-500 place-self-center text-2xl font-bold">Скидка {props.promo.discount} рублей</p>
            {ShowDescription && <div className="place-self-center">
                <p className="text-yellow-400 text-1xl font-bold">{props.promo.price} ₽/шт</p>
                <p className="text-yellow-400 text-1xl font-bold">Остаток: {props.promo.quantity} шт</p>
            </div>}
        </div>
    )

}