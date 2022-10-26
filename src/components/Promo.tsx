import {IPromo} from '../models'
import {useState} from "react";

interface PromoProps {
    promo: IPromo
}

export function Promo(props: PromoProps) {
    const [showDetails, setShowDetails] = useState<boolean>(false)
    return (
        <div className="border w-1/2 py-0 px-0 rounded flex flex-col justify-between items-center mb-2 place-content-start">
            <img src={process.env.PUBLIC_URL + props.promo.image} className="w-20" alt={props.promo.store}/>
            <p>{ props.promo.store }</p>
            <p className="font-bold"> Скидка {props.promo.discount} рублей</p>
            <button
                className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
                onClick={() => setShowDetails((prevState) => !prevState)}
            >
                {!showDetails ? <div>Show Details</div> : <div>Hide Details</div>}
            </button>
            {showDetails && <div>
                <p>{props.promo.price} ₽/шт</p>
                <p>Остаток: {props.promo.quantity} шт.</p>
            </div>}
        </div>
    )
}