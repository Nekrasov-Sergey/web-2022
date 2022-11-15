import {Routes, Route, BrowserRouter} from 'react-router-dom'
import {Navbar} from "./components/Navbar"
import {HomePage} from "./components/HomePage"
import {Info} from "./components/Info"
import {Payment} from "./components/Payment"
import {NotFound} from "./components/NotFound";
import {CartPage} from "./components/CartPage";

export const ENDPOINT = "http://localhost:8080"

export default function App() {
    return (
        <BrowserRouter>
            <Navbar/>
            <Routes>
                <Route path="/store" element={<HomePage/>}/>
                <Route path="/store/cart" element={<CartPage/>}/>
                <Route path="/store/info" element={<Info/>}/>
                <Route path="/store/:payment" element={<Payment/>}/>
                <Route path="/store*" element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}

