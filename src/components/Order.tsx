import {useContext} from "react";
import {MyContext} from "./OrderPage";
import {GetUser} from "../requests/GetUser";


export function Order() {
    const ctx = useContext(MyContext)
    return (
        <div
            className="border-2 border-teal-200 mx-auto mt-4 w-5/6 h-16 py-5 px-5  rounded-lg grid grid-cols-5 bg-white"
        >
            <p className="pl-2 place-self-center text-lg">
                {ctx.Store}
            </p>

            <p className="pl-2 place-self-center text-lg">
                {ctx.Quantity}
            </p>

            <p className="place-self-center text-lg">
                {GetUser(ctx.UserUUID)}
            </p>

            <p className="place-self-center text-lg">
                {ctx.Date.replace("T", " ").split(".")[0]}
            </p>

            <p className="place-self-center text-lg">
                {ctx.Status}
            </p>

        </div>
    )
}