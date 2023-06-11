import React from 'react'

const Hero = () => {
  return (
    <div className='text-[#DE402F]'>
      <div className='text-5xl text-center font-bold p-4 m-4'>Why are we <span className='text-[#CEDE2F] animate-pulse'>different</span>?</div>
      <div className='grid grid-cols-4 gap-2'>
          <div className="p-1 m-1 text-justify tracking-tight">
            <h1 className='text-xl font-bold'>Quality over quantity.</h1>
            <p>
              At Digby, we believe in quality over quantity. We source only the finest materials and work with skilled artisans to create garments that are built to last. Every piece in our collection is carefully crafted with attention to detail, ensuring that they are not only stylish, but also comfortable and durable. By choosing our clothing, you're investing in a piece that you'll love wearing for years to come. Shop now and experience the difference of quality.
            </p>
          </div>

          <div>
            <h1 className='text-xl font-bold'>Ethically Selected Manufacturers</h1>
            <p>
              We believe in fashion that feels good and looks good. That's why we only work with ethically selected manufacturers who share our commitment to sustainability and fair labor practices. Every garment in our collection is made with the highest standards of social and environmental responsibility. You can feel good about your purchase knowing that it was made in a way that respects the rights and well-being of workers and the planet. Shop now and make a conscious choice for a better fashion future.
            </p>
          </div>

          <div>
            <h1 className='text-xl font-bold'>Your voice matters!</h1>
            <p>
              We value your feedback and take it to heart in everything we do. We know that you, our valued customers, are our best source of inspiration and ideas. You shape our collections and influence our decisions. We also believe that we have a lot to learn from you and we are always eager to listen and improve. Your satisfaction is our top priority, and we strive to make every shopping experience with us a positive one.
            </p>
          </div>

          <div>
            <h1 className='text-xl font-bold'>We take care about our Mother Earth!</h1>
            <p>
              All people working at Digby are committed to protecting the planet and preserving it for future generations. We use eco-friendly materials and implement sustainable practices throughout our production process. From sourcing fabrics to packaging, we are constantly looking for ways to reduce our impact on the environment. Our clothing is not just stylish, but also responsible. We care about the earth and we hope you will too by choosing our clothing.
            </p>
          </div>

      </div>
    </div>
  )
}

export default Hero