# creating-symmetry
Create symmetrical wallpapers using math.

There are many types of symmetries you can create using a single image.
- Rotational: You can rotate an image around a single point.
- Translational: You can move and shift an image in a direction.
- Mirror: You can make a reflection of an image so everything perpendicular to the mirror line is the same distance away.

After reading [Creating Symmetry](https://www.amazon.com/Creating-Symmetry-Mathematics-Wallpaper-Patterns/dp/0691161739), I looked around and didn't find a program that would do it. So time to build my own program, in Go.

Give it a base image, tweak some parameters, wait for the image to render, keep tweaking it until it looks good.

## NOTES
Types to support:

Rosette
Frieze
17 wallpaper types
3D shapes
Hyperbolic

Color changing

### Research notes
#### Program min and max range
After calculating the destination you need to
- Calculate the min & max bounds
- Interpolate against those ranges to pick from the color wheel

#### Export preview pics
Export to a 200x200 stamp that is quick to calculate
Export with a rotating file name so you can look at old previews

#### Formula stuff
Make formula objects

### Objects
Transformation - Rosette, Frieze, Wallpaper, 3D projection
Merge options
Projection - Euclidean or Hyperbolic?
Palette Selection - chooses the color range