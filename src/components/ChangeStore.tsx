import React, {useState} from "react"
import {Navbar} from "./Navbar";
import {ChangingStore} from "../requests/ChangeStore";
import {Link, useLocation} from "react-router-dom";

export function ChangeStore() {
    const [name, setName] = useState(useLocation().state.Name);
    const handleChangeName = (event: { target: { value: any; }; }) => {
        setName(event.target.value);
    };

    const [discount, setDiscount] = useState(useLocation().state.Discount);
    const handleChangeDiscount = (event: { target: { value: any; }; }) => {
        setDiscount(Number(event.target.value));
    };

    const [price, setPrice] = useState(useLocation().state.Price);
    const handleChangePrice = (event: { target: { value: any; }; }) => {
        setPrice(Number(event.target.value));
    };

    const [quantity, setQuantity] = useState(useLocation().state.Quantity);
    const handleChangeQuantity = (event: { target: { value: any; }; }) => {
        setQuantity(Number(event.target.value));
    };

    const [promo, setPromo] = useState(useLocation().state.Promo);
    const handleChangePromo = (event: { target: { value: string; }; }) => {
        setPromo([event.target.value]);
    };

    const [image, setImage] = useState(useLocation().state.Image);
    const handleChangeImage = (event: { target: { value: any; }; }) => {
        setImage(event.target.value);
    };

    return (
        <>
            <Navbar/>

            <div className="bg-yellow-50 min-h-screen">
                <p className="ml-4 sm:text-2xl text-1xl font-normal text-black">
                    <Link to="/store" className="mr-2">
                        Freebie shop
                    </Link>
                    / changing
                </p>

                <p className="text-center sm:text-5xl text-3xl font-bold text-pink-500">
                    Изменение магазина
                </p>

                <div className="mt-10 mx-5 bg-white rounded-lg border-2 border-teal-200">
                    <div className="grid grid-cols-4 grid-rows-2 gap-10 p-8">
                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Название
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeName}
                                value={name}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Скидка
                            </label>
                            <input
                                type="number"
                                min="20"
                                max="2000"
                                onChange={handleChangeDiscount}
                                value={discount}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Цена
                            </label>
                            <input
                                type="number"
                                min="10"
                                max="1000"
                                onChange={handleChangePrice}
                                value={price}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Количество
                            </label>
                            <input
                                type="number"
                                min="1"
                                onChange={handleChangeQuantity}
                                value={quantity}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Промокоды
                            </label>

                            <input
                                type="text"
                                onChange={handleChangePromo}
                                value={promo}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>

                        <div className="col-span-2">
                            <label htmlFor="first-name"
                                   className="block text-base font-medium text-gray-700">
                                Изображение
                            </label>
                            <input
                                type="text"
                                onChange={handleChangeImage}
                                value={image}
                                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-base"
                            />
                        </div>
                    </div>

                    <div className="text-center mb-6">
                        {ChangingStore(useLocation().state.UUID, name, discount, price, quantity, promo, image)}
                    </div>
                </div>

                <p className="py-8 text-center">
                    <Link to="/store"
                          className="border-4 border-blue-700 text-blue-700 hover:bg-blue-700 hover:text-white py-1 px-3 rounded-full text-2xl font-bold"
                    >
                        Обратно на главную
                    </Link>
                </p>
            </div>
        </>
    )
}