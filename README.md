# Symmetrical Pattern Generator
This program lets you transform an image in to an image with symmetrical pattern. Just supply a base image and a formula to get started.

![Image with 7 horizontal stripes creating the rainbow with white on top and black on the bottom. Rainbow Stripe](example/rainbow_stripe.png)

![Transformed rainbow stripe image into rosette with 3 rotational symmetry](example/rainbow_stripe_rosette_1.png)
![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/rainbow_stripe_frieze_p11g.png)
![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/rainbow_stripe_lattice_hexagonal_p31m.png)

## Installation
This program is written in [Go](https://golang.org/), so download that first.

`go install` will download the other required libraries:
- yaml
- ginkgo
- gomega
- check

You will also need a source image to generate patterns with. I included one in `source/rainbow_stripe.png`.

## How to run
`make run` Looks for the file `data/formula.yml` to find the source image, the type of pattern to use and other settings.

Rename the `data/formula.yml.example` file to see it in action.

Look over [HERE](docs/common_options.md) to see common options used in every symmetry pattern file.

## Types of patterns
**Rosette** patterns surround the center of the image, expanding outward forever.
![Transformed rainbow stripe image into rosette with 3 rotational symmetry](example/rainbow_stripe_rosette_1.png)

p3 symmetry [(Link to source file)](example/rainbow_stripe_rosette_1.yml)

**Frieze** patterns expand horizontally forever but usually have a finite height.
![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/rainbow_stripe_frieze_p11g.png)
p11g symmetry [(Link to source file)](example/rainbow_stripe_frieze_p11g.yml)

**Lattice** patterns tend to repeat using a given 4 sided shape called a lattice. They expand horizontally and vertically.
![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/rainbow_stripe_lattice_hexagonal_p31m.png)

Hexagonal lattice with p31m symmetry [(Link to source file)](example/rainbow_stripe_lattice_hexagonal_p31m.yml)

## How to test
- `make test` Runs the unit tests.
- `make lint` Runs the linter

## Inspiration
[Creating Symmetry](https://www.amazon.com/Creating-Symmetry-Mathematics-Wallpaper-Patterns/dp/0691161739) by Frank Farris shows the 
math behind the patterns and inspired me to make this. Prepare for group theory and lots of complex numbers.
