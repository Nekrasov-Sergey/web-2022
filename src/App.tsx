import React from 'react';
import {Routes, Route, BrowserRouter} from 'react-router-dom'
import {Navbar} from "./components/Navbar";
import {HomePage} from "./components/HomePage";
import {Info} from "./components/Info";


function App() {
    return (
        <BrowserRouter basename="/">
            <Navbar/>
            <Routes>
                <Route path="/" element={<HomePage/>}/>
                <Route path="/info" element={<Info/>}/>
            </Routes>
        </BrowserRouter>
    )
}

export default App;