import Home from "./components/Home";
import ProductDisplay from "./components/ProductDisplay";
import Cart from "./components/Cart";


import {
  Routes,
  Route
} from 'react-router-dom'

function App() {

  return (
    <>
      <Routes>
        <Route path='/home' element={<Home />} />
        <Route path='/' element={<Login/>} />
        <Route path='/products' element={<ProductDisplay />} />
        <Route path='/cart' element={<Cart />} />
      </Routes>
    </>
  )
}

export default App
