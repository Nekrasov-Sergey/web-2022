// import {Link} from "react-router-dom"
// import {logoutUser} from "../modules";
//
// export function Navbar() {
//     return (
//         <nav className="py-4 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
//             <p className="pl-40 font-bold text-2xl inline text-green-400">
//                 <Link to="/store">ГЛАВНАЯ</Link>
//             </p>
//
//             <p className="pl-8 font-bold text-2xl inline text-green-400">
//                 <Link to="/store/cart">КОРЗИНА</Link>
//             </p>
//
//             <p className="pl-8 font-bold text-2xl inline text-green-400">
//                 <Link to="/store/info">О НАС</Link>
//             </p>
//         </nav>
//     )
// }

import { Fragment } from 'react'
import { Disclosure, Menu, Transition } from '@headlessui/react'
import {NavLink} from "react-router-dom";
import {logoutUser} from "../modules";



const navigation = [
    { name: 'Главная', href: '/store'},
    { name: 'Корзина', href: '/cart'},
    { name: 'О нас', href: '/info'},
]

function classNames(...classes: string[]) {
    return classes.filter(Boolean).join(' ')
}

export  function Navbar() {
    return (
        <Disclosure as="nav" className="bg-gray-800">
            {() => (
                <>
                    <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
                        <div className="relative flex h-16 items-center justify-between">
                            <div className="flex flex-1 items-center justify-center sm:items-stretch sm:justify-start">
                                <div className="hidden sm:ml-6 sm:block">
                                    <div className="flex space-x-4">
                                        {navigation.map((item) => (
                                            <NavLink
                                                key={item.name}
                                                to={item.href}
                                                className={({ isActive }) =>
                                                    [
                                                        "px-3 py-2 rounded-md text-lg font-medium",
                                                        isActive ? "bg-gray-900 text-white" : "text-gray-300 hover:bg-gray-700 hover:text-white",
                                                    ]
                                                        .filter(Boolean)
                                                        .join(" ")
                                                }
                                            >
                                                {item.name}
                                            </NavLink>
                                        ))}
                                    </div>
                                </div>
                            </div>
                            <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">

                                {/* Profile dropdown */}
                                <Menu as="div" className="relative ml-3">
                                    <div>
                                        <Menu.Button className="flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
                                            <span className="sr-only">Open user menu</span>
                                            <img
                                                className="h-8 w-8 rounded-full"
                                                src="https://res.cloudinary.com/dl0tawm7w/image/upload/v1667748734/logo/chrome_YfuwE6ZZWj_cnrbky.png"
                                                alt=""
                                            />
                                        </Menu.Button>
                                    </div>
                                    <Transition
                                        as={Fragment}
                                        enter="transition ease-out duration-100"
                                        enterFrom="transform opacity-0 scale-95"
                                        enterTo="transform opacity-100 scale-100"
                                        leave="transition ease-in duration-75"
                                        leaveFrom="transform opacity-100 scale-100"
                                        leaveTo="transform opacity-0 scale-95"
                                    >
                                        <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                                            <Menu.Item>
                                                {({ active }) => (
                                                    <a
                                                        href="#"
                                                        className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
                                                    >
                                                        Ваш профиль
                                                    </a>
                                                )}
                                            </Menu.Item>
                                            <Menu.Item>
                                                {({ active }) => (
                                                    <a
                                                        onClick={() => logoutUser("logout")}
                                                        className={classNames(active ? 'bg-gray-100' : '', 'block px-4 py-2 text-sm text-gray-700')}
                                                    >
                                                        Выход
                                                    </a>
                                                )}
                                            </Menu.Item>
                                        </Menu.Items>
                                    </Transition>
                                </Menu>
                            </div>
                        </div>
                    </div>

                </>
            )}
        </Disclosure>
    )
}