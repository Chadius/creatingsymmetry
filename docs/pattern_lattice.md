# What is a lattice pattern?
Imagine a 4 sided figure, like a square or rectangle. Tile them on top of each other and to the side.
Now, stretch the source image across those four points. That's a lattice. 

# Symmetry Types
## Translational
All lattice based patterns have translational symmetry.
The lattice is always aligned across 2 axes, so if you scroll the pattern along those directions you will see the pattern repeat.

![Transformed rainbow stripe image into rectangular lattice with pmg symmetry, with green and orange round rectangules on a yellow background](../example/lattices/rainbow_stripe_lattice_rectangular_pmg.png)

In this example, the axes are horizontal and vertical. You can move up, down, left or right.

![Preceding image with arrows drawn, showing the center of the green rectangles](../docs/lattice_symmetry/rainbow_stripe_lattice_rectangular_pmg_symmetry_translational.png)

## Reflective
These types look the same if you flip them over a line.
- Some reflective lines are along the x-axis. You can fold the top of the image to the bottom of the image.
- Some reflective lines are along the y-axis. You can fold the left of the image to the right of the image.

![Transformed rainbow stripe image into a square lattice with p4m symmetry, forming a tilted purple and blue checkerboard pattern](../example/lattices/rainbow_stripe_lattice_square_p4m.png)

In this example, you can draw a line across the middle, from top to bottom. You'll see the left side matches the right side.

![Preceding image with a vertical line, showing the left is a reflection of the right](../docs/lattice_symmetry/rainbow_stripe_lattice_square_p4m_reflective_yaxis.png)

You can also draw a line from left to right. The top matches the bottom.

![Preceding image with a horizontal line, showing the top is a reflection of the bottom](../docs/lattice_symmetry/rainbow_stripe_lattice_square_p4m_reflective_xaxis.png)

This example also has diagonal reflections.

![Preceding image with a diagonal line, showing the corners reflect on the corresponding corner](../docs/lattice_symmetry/rainbow_stripe_lattice_square_p4m_reflective_diagonal.png)

## Rotational
All patterns can be rotated 360 degrees without visual changes.
But some can be rotated at 60, 90, 120 or 180 degrees and look the same.

![Transformed rainbow stripe image into hexagonal lattice with P6 symmetry, 6 circles with 3 holes surrounding a single holed circle](../example/lattices/rainbow_stripe_lattice_hexagonal_p6.png)

This pattern has 6 way rotational symmetry. Rotate it 60 degrees and the image will look the same.

![Transformed rainbow stripe image into hexagonal lattice with P6 symmetry, 6 circles surround a central circle. 3 above and 3 below. The surrounding circles have 3 holes like a power outlet and the central circle has one hole.](../docs/lattice_symmetry/rainbow_stripe_lattice_hexagonal_p6_symmetry_6_rotation.png)

# Lattice types
There are 5 lattice-based patterns that lead to 17 types of symmetry.

## Hexagonal

![Transformed rainbow stripe image into hexagonal lattice with p31m symmetry, with purple, indigo and blue nodes against a transparent background](../example/lattices/rainbow_stripe_lattice_hexagonal_p31m.png)

Hexagonal lattice with p31m symmetry [(link to formula)](../example/lattices/rainbow_stripe_lattice_hexagonal_p31m.yml)
The four sided lattice is tilted, so look for the solid blue points. Stacked enough times it connects 7 of them.

Learn more about hexagonal lattices [here.](lattice_hexagonal.md)

## Rectangular
![Transformed rainbow stripe image into rectangular lattice with pmg symmetry, with green and orange round rectangules on a yellow background](../example/lattices/rainbow_stripe_lattice_rectangular_pmg.png)

Rectangular lattice with pmg symmetry [(link to formula)](../example/lattices/rainbow_stripe_lattice_rectangular_pmg.yml).
There is a lattice that connects the centers of 4 orange and green rectangles. The lattice is stacked throughout the pattern to make it repeat.

## Square
![Transformed rainbow stripe image into a square lattice with p4m symmetry, forming a tilted purple and blue checkerboard pattern](../example/lattices/rainbow_stripe_lattice_square_p4m_and_p4g.png)

Square lattice with p4m symmetry [(link to formula)](../example/lattices/rainbow_stripe_lattice_square_p4m_and_p4g.yml).
Similar to the rectangular lattice, this lattice meets at right angles and has the same lengths of each side.
Imagine the square's corners inside each of the holes of the purple diamonds, and you'll see the lattice.

You can find out more about Square patterns [here.](lattice_square.md)

## Rhombic
![Transformed rainbow stripe image into rhombic lattice with cmm symmetry. Red and orange blobs sit interlocked against a transparent background](../example/lattices/rainbow_stripe_lattice_rhombic_cmm.png)

Rhombic lattice with cmm symmetry [(link to formula)](../example/lattices/rainbow_stripe_lattice_rhombic_cmm.yml)
The lattice is based on a rhombus, where all sides are the same length but not at a square. There are rounding errors since the resolution is so small, but all of the red shapes should be exactly the same.

Build your own Rhombic patterns [here.](lattice_rhombic.md)

## Generic
![Transformed rainbow stripe image into generic lattice with p2 symmetry. Orange pattern with tilted holes are like a grilled cheese sandwich](../example/lattices/rainbow_stripe_lattice_generic_p2.png)

Generic lattice with p2 symmetry [(link to formula)](../example/lattices/rainbow_stripe_lattice_generic_p2.yml)
Generic lattices are freeform. They are guaranteed to be stackable and may have 180-degree symmetrical rotation.

Want to make your own generic lattices? Click [here.](lattice_generic.md)
