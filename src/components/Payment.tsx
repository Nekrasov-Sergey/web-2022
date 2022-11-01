import {Link} from "react-router-dom"
import {useLocation} from "react-router-dom"

export function Payment() {
    return (
        <div className="bg-yellow-50">
            <p className="ml-4 text-2xl font-normal text-black">
                Freebie shop / {useLocation().state.store}
            </p>

            <p className=" font-bold text-6xl text-pink-500 text-center">
                Страница оплаты
            </p>

            <p className="mt-8 font-medium text-4xl text-green-500 text-center">
                Ваш промокод:
                <p className="font-bold italic text-4xl text-red-700">
                    {useLocation().state.promo}
                </p>
            </p>

            <p className="my-8 text-center">
                <Link to="/"
                      className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                >
                    Обратно на главную
                </Link>
            </p>

            <img src="/resourses/Payment.jpg" width="23%" className="mx-auto" alt="Payment"/>
        </div>
    )
}