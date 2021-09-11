# Symmetrical Pattern Generator
This program lets you transform an image in to an image with symmetrical pattern. Just supply a base image and a formula to get started.

![Image with 7 horizontal stripes creating the rainbow with white on top and black on the bottom. Rainbow Stripe](example/rainbow_stripe.png)

![Transformed rainbow stripe image into rosette with 3 rotational symmetry](example/rosettes/rainbow_stripe_rosette_1.png)
![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/friezes/rainbow_stripe_frieze_p11m_and_p11g.png)
![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.png)

I assume you're comfortable with [Go](https://golang.org/), YAML and a command line. You'll install this in go, write the formulas in YAML, and run `make run` on a command line to generate pretty patterns.

## Important pages
* [Common Options](docs/common_options.md)
* [Rosette Patterns](docs/pattern_rosette.md)
* [Frieze Patterns](docs/pattern_frieze.md)
* [Lattice Patterns](docs/pattern_lattice.md)

## Installation
This program is written in [Go](https://golang.org/), so download that first.

Once that is done, you can check out this repository, open a terminal and navigate to this codebase.

Run `go install` and it will download the other required libraries:
- yaml
- ginkgo
- gomega
- check

You will need a source image. [A rainbow stripe pattern is included](example/rainbow_stripe.png) in this repository and used throughout these docs.

![Rainbow Stripe with white stripe on top and black stripe on the bottom](example/rainbow_stripe.png)

## How to run
Open your command line and type:
`make run`

This will look for the file `data/formula.yml` to find the required settings:
- The name of the source image
- The name and size of the output image
- The formula that transforms the source into the output
- Numerical ranges used with the formula

All options (except the formula) are described [here.](docs/common_options.md)

### Example
If you learn better by example, try renaming [data/formula.yml.example](./data/formula.yml.example) to `data/formula.yml`.
When you run `make run`, it will generate the [orange and red pattern](#rosette) you see below.

## Types of patterns
### Rosette
**Rosette** patterns surround the center of the image, expanding outward. They can be rotated around the center without visually changing the look. 

![Transformed rainbow stripe image into rosette with 3 rotational symmetry, creating three yellow to purple petals on a orange and red striped background](example/rosettes/rainbow_stripe_rosette_1.png)

[3 way rotational symmetry](example/rosettes/rainbow_stripe_rosette_1.yml)

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](example/rosettes/rainbow_stripe_rosette_2.png)

[4 way rotational symmetry](example/rosettes/rainbow_stripe_rosette_2.yml)

![Transformed rainbow stripe image into rosette with 5 rotational symmetry, creating a 10 point mostly green hubcap](example/rosettes/rainbow_stripe_rosette_3.png)

[5 way rotational symmetry](example/rosettes/rainbow_stripe_rosette_3.yml)

[Click here](docs/pattern_rosette.md) to learn more about rosette-based patterns and design your own.

### Frieze
**Frieze** patterns expand horizontally forever but usually have a finite height.

![Transformed rainbow stripe image into frieze with p11g symmetry, with blue and yellow hourglasses in a green background](example/friezes/rainbow_stripe_frieze_p11m_and_p11g.png)

[Reflective symmetry across the x-axis](example/friezes/rainbow_stripe_frieze_p11m_and_p11g.yml)

![Transformed rainbow stripe image into frieze with p2111 symmetry, like an orange branch with black notches separated by blue and white droplets](example/friezes/rainbow_stripe_frieze_p211.png)

[2 way Rotational symmetry](example/friezes/rainbow_stripe_frieze_p211.yml)

![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a green background](example/friezes/rainbow_stripe_frieze_p2mg.png)

[Glide symmetry: move the pattern horizontally a bit and then flip it over the x-axis](example/friezes/rainbow_stripe_frieze_p2mg.yml)

[Click here](docs/pattern_frieze.md) to learn more about frieze-based patterns and how to build your own.

### Lattice
**Lattice** patterns transform using a 4 sided shape called a lattice. They expand horizontally and vertically forever. Sometimes these are called wallpaper patterns.

![Transformed rainbow stripe image into rectangular lattice with pmg symmetry, with green and orange round rectangules on a yellow background](example/lattices/rainbow_stripe_lattice_rectangular_pmg.png)

[Rectangular lattice](example/lattices/rainbow_stripe_lattice_rectangular_pmg.yml).
The lattice connects the 4 centers of the green and orange ovals. Then the lattice is repeated to create the image.

![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.png)

[Hexagonal lattice](example/lattices/rainbow_stripe_lattice_hexagonal_p31m.yml)
Creates 3 way rotational symmetry. This lattice can be stacked 7 times to create a hexagonal pattern. Count the solid blue dots to see the corners of the lattices.

![Transformed rainbow stripe image into rhombic lattice with cmm symmetry. Red and orange blobs sit interlocked against a transparent background](example/lattices/rainbow_stripe_lattice_rhombic_cmm.png)

[Rhombic lattice](example/lattices/rainbow_stripe_lattice_rhombic_cmm.yml)
The lattice is based on a rhombus, where all sides are the same length but not at a square.
All red blobs should look the same, but due to the low resolution rounding errors appear.

[Click here](docs/pattern_lattice.md) to learn more about lattice-based patterns. (Still a Work In Progress!)

## How to test
If you plan to mess around with the code itself, here are 2 more make commands that will come in handy:
- `make test` Runs the unit tests.
- `make lint` Runs the linter.

## Inspiration
[Creating Symmetry by Frank Farris](https://www.amazon.com/Creating-Symmetry-Mathematics-Wallpaper-Patterns/dp/0691161739) 
merges math and art to create beautiful patterns and is worth the read. Prepare for Group Theory and Complex Numbers!
