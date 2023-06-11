import React from 'react'
import Navbar from './Navbar';
import Hero from './Hero';
import Footer from './Footer';
import Carousel from './Carousel';

const Home = () => {
  return (
    <>
        <Navbar />
        <Carousel />
        <Hero />
        <Footer />
    </>
  )
}

export default Home