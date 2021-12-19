# Mathematical Abstract Art Generator
This program takes images and transforms them into abstract art pieces.

```shell script
cp data/formula.yml.example data/formula.yml
make run
```

Will turn

![Image with 7 horizontal stripes creating the rainbow with white on top and black on the bottom. Rainbow Stripe](example/rainbow_stripe.png)

into:

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](example/rosettes/rainbow_stripe_rosette_1.png)

# How does it work?
`Image + Transformation = New Image`

1. You supply a source image
2. The transformation applies a lot of mathematics to figure out which pixel of the source should be used
3. The output image uses the pixel as a color. 

The transformation offers the most customization. 

# Next topics
- How to install
- Command Line Options
- Transformations