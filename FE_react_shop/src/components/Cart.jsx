import React, { useEffect, useState } from "react";

import Navbar from "./Navbar";
import CartItem from "./CartItem";

const Cart = () => {
  const [backendData, setBackendData] = useState([]);
  const [userPrice, setUserPrice] = useState(0);

  useEffect(() => {
    fetch("http://localhost:5000/cart")
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setBackendData(data);
      });
  }, []);

  // Calculate the sum of prices
  const calculateTotalPrice = () => {
    let total = 0;
    backendData.forEach((item) => {
      total += item.product.price * item.amount;
    });
    return total;
  };

  // Handle user price input
  const handleUserPriceChange = (event) => {
    setUserPrice(Number(event.target.value));
  };

  // Handle payment
  const handlePayment = () => {
    fetch("http://localhost:5000/pay", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        amount: userPrice,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        // Handle the response as needed
      })
      .catch((error) => {
        console.error("Error:", error);
        // Handle any errors that occurred during the request
      });
  };

  return (
    <>
      <Navbar />
      <div className="flex items-center justify-center m-2 gap-4">
        {backendData && backendData.length > 0 ? (
          backendData.map((data, index) => (
            <CartItem key={index} data={data} />
          ))
        ) : (
          <p>No products to pay for!</p>
        )}
      </div>

      {backendData && backendData.length > 0 && (
        <div className="flex flex-col items-center mt-4">
          <div className="bg-gray-200 px-4 py-2 rounded-md">
            Total Price: ${calculateTotalPrice()}
          </div>
          <form className="mt-4">
            <label htmlFor="userPrice" className="mr-2">
              Enter a price:
            </label>
            <input
              type="number"
              id="userPrice"
              value={userPrice}
              onChange={handleUserPriceChange}
              className="px-2 py-1 rounded-md"
            />
            <button
              type="button"
              className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-md"
              onClick={handlePayment}
            >
              Pay the Price
            </button>
          </form>
        </div>
      )}
    </>
  );
};

export default Cart;