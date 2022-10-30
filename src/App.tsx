import {Routes, Route, BrowserRouter} from 'react-router-dom'
import {Navbar} from "./components/Navbar";
import {HomePage} from "./components/HomePage";
import {Info} from "./components/Info";
import {Payment} from "./components/Payment";

export default function App() {
    return (
        <BrowserRouter basename="/">
            <Navbar/>
            <Routes>
                <Route path="/" element={<HomePage/>}/>
                <Route path="/info" element={<Info/>}/>
                <Route path="/payment" element={<Payment/>}/>
            </Routes>
        </BrowserRouter>
    )
}

