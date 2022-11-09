import {Store} from "./Store"
import {AllStores} from "../modules";

export function HomePage() {
    return (
        <div className="bg-yellow-50 min-h-screen">
            <p className="ml-4 text-2xl font-normal text-black">
                Freebie shop
            </p>

            <p className="text-center text-6xl font-bold text-pink-500">
                Доступные промокоды на Ноябрь 2022
            </p>
            {AllStores().map((store) => {
                return <Store store={store}/>
            })}
        </div>
    )
}