import React, {useState} from 'react'
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  XMarkIcon,} from '@heroicons/react/24/solid'

function Carousel() {
  const slides = [
    {
      url: 'https://images.unsplash.com/photo-1594882645126-14020914d58d?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1485&q=80',
    },
    {
      url: 'https://images.unsplash.com/photo-1523381210434-271e8be1f52b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=870&q=80',
    },
    {
      url: 'https://images.unsplash.com/photo-1560243563-062bfc001d68?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80',
    },
  ]

  const [currentIndex, setCurrentIndex] = useState(0)

  const prevSlide = () => {
    const isFirstSlide = currentIndex === 0;
    const newIndex = isFirstSlide ? slides.length - 1 : currentIndex -1;
    setCurrentIndex(newIndex);
  };

  const nextSlide = () => {
    const isLastSlide = currentIndex === slides.length - 1;
    const newIndex = isLastSlide ? 0 : currentIndex + 1;
    setCurrentIndex(newIndex);
  };

  const goToSlide = (slideIndex) => {
    setCurrentIndex(slideIndex)
  }

  return (
    <div className='max-w-[1500px] h-[565px] w-full m-auto py-1 px-1 relative'>
      <div 
        style={{backgroundImage: `url(${slides[currentIndex].url})`}}
        className='w-full h-full rounded-md bg-center bg-cover duration-500'
        ></div>
      {/* Left Arrow */}
      <div className='absolute top-[50%] -translate-x-0 translate-y-[-50%] left-5 text-2x1 rounded-full p-2 bg-black/40 hover:bg-black/80 text-white cursor-pointer'>
        <ChevronLeftIcon onClick={prevSlide} className='h-6 w-6'/>
      </div>
      {/* Right Arrow */}
      <div className='absolute top-[50%] -translate-x-0 translate-y-[-50%] right-5 text-2x1 rounded-full p-2 bg-black/40 hover:bg-black/80 text-white cursor-pointer'>
        <ChevronRightIcon onClick={nextSlide} className='h-6 w-6' />
      </div>
    </div>
  )
}

export default Carousel