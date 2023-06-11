import React, { useEffect, useState } from "react"


const Card = ({ data }) => {
  const [isAddingToCart, setIsAddingToCart] = useState(false);
  
  const addToCart = () => {
    setIsAddingToCart(true);
  
    fetch("http://localhost:5000/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ name: data.name })
    })
    .then(response => response.json())
    .then(data => {
      console.log("Is added to cart? ", data)
    })
    .finally(() => {
      setIsAddingToCart(false);
    })
  }


  return (

    <div className='border border-gray-500 max-w-[256px] p-1 ,m-2 bg-gray-200'>
        <div className="flex flex-col text-center">
          <h1 className='font-bold'>{data.name}</h1>
          <h2 className='text-sm text-opacity-40'>{data.description}</h2>
          <h2 className='text-2xl'>{data.price} ZŁ</h2>
          <button
          className='cursor-pointer bg-[#DE402F] text-white text-xl border m-2 p-2 border-red-900 rounded-sm w-[234px] h-12'
          onClick={addToCart}
          disabled={isAddingToCart}>{isAddingToCart ? "Dodawanie..." : "Dodaj do koszyka"}
          </button>
      </div>
    </div>
  )
}

export default Card