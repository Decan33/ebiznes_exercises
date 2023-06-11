import React, { useEffect, useState } from "react"

import Navbar from "./Navbar"
import Card from "./Card"

const ProductDisplay = () => {
  const [backendData, setBackendData] = useState([])

  useEffect(() => {
    fetch("http://localhost:5000/products")
      .then((response) => response.json())
      .then(data => {
        console.log(data)
        setBackendData(data)
      })
  }, [])

  return (
    <>
      <Navbar />
      <div className="flex items-center justify-center m-2 gap-4">
        {backendData ? (
          backendData.map((data, index) => (
            <Card key={index} data={data} />
          ))
        ) : (
          <p>Loading...</p>
        )}
      </div>
    </>
  )
}

export default ProductDisplay