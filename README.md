# Symmetrical Pattern Generator
This program lets you transform an image in to an image with symmetrical pattern. Just supply a base image and a formula to get started.

![Image with 7 horizontal stripes creating the rainbow with white on top and black on the bottom. Rainbow Stripe](example/rainbow_stripe.png)

![Transformed rainbow stripe image into rosette with 3 rotational symmetry](example/rosettes/rainbow_stripe_rosette_1.png)
![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/friezes/rainbow_stripe_frieze_p11g.png)
![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.png)

I assume you're comfortable with [Go](https://golang.org/), YAML and a command line. You'll install this in go, write the formulas in YAML, and run this program on the command line to generate patterns.

## Table of Contents
* [Common Options](docs/common_options.md)
* [Rosette Patterns](docs/pattern_rosette.md)
* [Frieze Patterns](docs/pattern_frieze.md)
* [Lattice Patterns](docs/pattern_lattice.md)

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
### Rosette
**Rosette** patterns surround the center of the image, expanding outward.
![Transformed rainbow stripe image into rosette with 3 rotational symmetry, creating three yellow to purple petals on a orange and red striped background](example/rosettes/rainbow_stripe_rosette_1.png)

p3 symmetry [(link to formula)](example/rosettes/rainbow_stripe_rosette_1.yml)

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](example/rosettes/rainbow_stripe_rosette_2.png)

p4 symmetry [(link to formula)](example/rosettes/rainbow_stripe_rosette_2.yml)

![Transformed rainbow stripe image into rosette with 5 rotational symmetry, creating a 10 point mostly green hubcap](example/rosettes/rainbow_stripe_rosette_3.png)

p5 symmetry [(link to formula)](example/rosettes/rainbow_stripe_rosette_3.yml)

[Click here](docs/pattern_rosette.md) to learn more about rosette-based patterns.

### Frieze
**Frieze** patterns expand horizontally forever but usually have a finite height.

![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/friezes/rainbow_stripe_frieze_p11g.png)

p11g symmetry [(link to formula)](example/friezes/rainbow_stripe_frieze_p11g.yml)

![Transformed rainbow stripe image into frieze with p2111 symmetry, like an orange branch with black notches separated by blue and white droplets](example/friezes/rainbow_stripe_frieze_p211.png)

p211 symmetry [(link to formula)](example/friezes/rainbow_stripe_frieze_p211.yml)

![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a green background](example/friezes/rainbow_stripe_frieze_p2mg.png)

p2mg symmetry [(link to formula)](example/friezes/rainbow_stripe_frieze_p2mg.yml)

[Click here](docs/pattern_frieze.md) to learn more about frieze-based patterns.

### Lattice
**Lattice** patterns tend to repeat using a given 4 sided shape called a lattice. They expand horizontally and vertically.

![Transformed rainbow stripe image into rectangular lattice with pmg symmetry, with green and orange round rectangules on a yellow background](example/lattices/rainbow_stripe_lattice_rectangular_pmg.png)

Rectangular lattice with pmg symmetry [(link to formula)](example/lattices/rainbow_stripe_lattice_rectangular_pmg.yml).
There is a lattice that connects the centers of 4 orange and green rectangles. The lattice is stacked throughout the pattern to make it repeat.

![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.png)

Hexagonal lattice with p31m symmetry [(link to formula)](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.yml)
The four sided lattice is tilted, so look for the solid blue points and you may see it. Stacked enough times it connects 7 of them.

![Transformed rainbow stripe image into rhombic lattice with cmm symmetry. Red and orange blobs sit interlocked against a transparent background](example/lattices/rainbow_stripe_lattice_rhombic_cmm.png)

Rhombic lattice with cmm symmetry [(link to formula)](example/lattices/rainbow_stripe_lattice_rhombic_cmm.yml)
The lattice is based on a rhombus, where all sides are the same length but not at a square. There are rounding errors since the resolution is so small, but all of the red shapes should be exactly the same.

[Click here](docs/pattern_lattice.md) to learn more about lattice-based patterns.

## How to test
- `make test` Runs the unit tests.
- `make lint` Runs the linter

## Inspiration
[Creating Symmetry](https://www.amazon.com/Creating-Symmetry-Mathematics-Wallpaper-Patterns/dp/0691161739) by Frank Farris shows the 
math behind the patterns and inspired me to make this. Prepare for group theory and lots of complex numbers.
