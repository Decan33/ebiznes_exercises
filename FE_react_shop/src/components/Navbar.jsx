import React from 'react'
import { Link, Navigate } from 'react-router-dom';

const Navbar = () => {
  return (
    <div className="bg-[#DE972F] flex justify-between items-center h-16 w-full mx-auto px-4">
      <h1 className='w-full text-4xl font-bold text-[#DE402F] hover:transition duration-300 delay-75 hover:text-[#f8503e] animate-bounce hover:ease-linear cursor-pointer'>
        <Link to='/' replace className='drop-shadow-lg'>Digby</Link>
      </h1>
      <ul className='md:flex text-[#DE402F]'>
        <li className='p-4 font-bold'>
          <Link to='/products' replace>Products</Link>
        </li>
        <li className='p-4 font-bold'>
          <Link to='/cart' replace>Cart</Link>
        </li>
      </ul>
    </div>
  )
}

export default Navbar