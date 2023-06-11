import React from 'react';
import { useEffect, useState } from 'react';
import Navbar from "./Navbar";
import Card from "./Card";


const Shirts = () => {

  const [ backendData, setBackendData ] = useState([{}])

  useEffect(() => {
    fetch("http://localhost:5000/tshirts")
    .then( backendResponse => {
      console.log(backendResponse)
      backendResponse.json().then(data => setBackendData(data))
    })
  }, []);

  return (
    <>
    <Navbar/>
    <div className="flex items-center justify-center m-2 gap-4">
    {backendData.result.map((data, index) =>
            <Card key={index} data={data}></Card>
     )}
    </div>
    </>
  );
}

export default Shirts