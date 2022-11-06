import {Routes, Route, BrowserRouter} from 'react-router-dom'
import {Navbar} from "./components/Navbar"
import {HomePage} from "./components/HomePage"
import {Info} from "./components/Info"
import {Payment} from "./components/Payment"
import {NotFound} from "./components/NotFound";

export const ENDPOINT = "http://localhost:8080"

export default function App() {
    return (
        <BrowserRouter>
            <Navbar/>
            <Routes>
                <Route path="/" element={<HomePage/>}/>
                <Route path="/info" element={<Info/>}/>
                <Route path="/:payment" element={<Payment/>}/>
                <Route path="/*" element={<NotFound/>}/>
            </Routes>
        </BrowserRouter>
    )
}

