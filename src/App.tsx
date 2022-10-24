import React from 'react';
import {Promo} from './components/Promo'
import {Promos} from './repository/Promos'


function App() {
    return (
        <div className="container mx-auto max-w-2xl pt-5">
            {Promos.map((promo, key) => {
                return <Promo promo={promo} key={key}/>
            })}
        </div>
    )
}

export default App;