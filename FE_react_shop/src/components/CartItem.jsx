import React from "react"


const CartItem = ({ data }) => {
  return (
    <div className='border border-gray-500 max-w-[256px] p-1 ,m-2 bg-gray-200'>
        <div className="flex flex-col text-center">
          <h1 className='font-bold'>{data.product.name}</h1>
          <h2 className='text-sm text-opacity-40'>{data.product.description}</h2>
          <h2 className='text-2xl'>{data.product.price} ZŁ</h2>
          <h2 className='text-xl'>{data.amount} sztuk/i</h2>
      </div>
    </div>
  )
}

export default CartItem