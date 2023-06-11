import { useState } from 'react'
import Home from "./components/Home";
import ProductDisplay from "./components/ProductDisplay";
import Cart from "./components/Cart";


import { BrowserRouter as Router,
  Routes,
  Route,
  Link} from 'react-router-dom'

function App() {

  return (
    <>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/products' element={<ProductDisplay />} />
        <Route path='/cart' element={<Cart />} />
      </Routes>
    </>
  )
}

export default App
