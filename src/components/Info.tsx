import {Link} from "react-router-dom";

export function Info() {
    return (
        <div className="bg-yellow-50">
            <p className="ml-4 text-2xl font-normal text-black">
                <Link to="/" className="mr-2">
                    Freebie shop
                </Link>
                / info
            </p>

            <p className="text-center font-bold text-6xl text-pink-500">
                Freebie shop
            </p>

            <p className="text-center mt-4 mx-8 font-medium text-3xl text-indigo-700">
                Это магазин промокодов, где вы можете купить промокоды для многих популярных магазинов.
                Экономьте деньги вместе с нами!
            </p>

            <img src="https://res.cloudinary.com/dh4qv3hob/image/upload/v1667665906/Promos/Discount_hgs7rl.png" width="29%" className="mx-auto" alt="Discount"/>
        </div>
    )
}

