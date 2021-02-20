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
### Complex Math library
https://golang.org/pkg/math/cmplx/#Abs

radius to the nth power
e to the negative i * angle