package com.ebiznes.exercises.demo;

public enum WebsiteTexts {
    QUALITY_OVER_QUANTITY(
            "Quality over quantity.\n" +
            "At Digby, we believe in quality over quantity. We source only the finest materials and work with skilled artisans to create garments that are built to last. Every piece in our collection is carefully crafted with attention to detail, ensuring that they are not only stylish, but also comfortable and durable. By choosing our clothing, you're investing in a piece that you'll love wearing for years to come. Shop now and experience the difference of quality."
    ),
    ETHICALLY_SELECTED_MANUFACTURERS(
            "Ethically Selected Manufacturers\n" +
                    "We believe in fashion that feels good and looks good. That's why we only work with ethically selected manufacturers who share our commitment to sustainability and fair labor practices. Every garment in our collection is made with the highest standards of social and environmental responsibility. You can feel good about your purchase knowing that it was made in a way that respects the rights and well-being of workers and the planet. Shop now and make a conscious choice for a better fashion future."
    ),
    YOUR_VOICE_MATTERS(
            "Your voice matters!\n" +
                    "We value your feedback and take it to heart in everything we do. We know that you, our valued customers, are our best source of inspiration and ideas. You shape our collections and influence our decisions. We also believe that we have a lot to learn from you and we are always eager to listen and improve. Your satisfaction is our top priority, and we strive to make every shopping experience with us a positive one."
    ),
    WE_TAKE_CARE(
            "We take care about our Mother Earth!\n" +
                    "All people working at Digby are committed to protecting the planet and preserving it for future generations. We use eco-friendly materials and implement sustainable practices throughout our production process. From sourcing fabrics to packaging, we are constantly looking for ways to reduce our impact on the environment. Our clothing is not just stylish, but also responsible. We care about the earth and we hope you will too by choosing our clothing."
    ),
    HOME_FOOTER(
            "About this project\n" +
                    "Created by Kacper Cygan with the help of TailwindCSS, ReactJS and many others useful libraries and packages"
    );


    final String text;

    WebsiteTexts(String text) {
        this.text = text;
    }

    public String getText() {
        return text;
    }
}
