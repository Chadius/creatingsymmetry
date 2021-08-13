# Formula file
By default, this program looks for `data/formula.yml`. This should contain the settings and the formula of the symmetry pattern you want to make.

## How does it work?
1. This program samples an *input* coordinate from the input space.
2. Then it *transforms* the input.
3. The transformed point is used to figure out what color it should *sample* from the input image.

## Common options
Every formula file contains these options.

### Input image
The name of the image file. JPG and PNG are supported, as well as any format Go lang’s `Image` library supports.

Examples:
`example/rainbow_stripe.png`
`input/iceCreamSundae.jpg`

### Output Resolution
How big do you want the resulting image?
- Bigger images give more detail.
- Smaller images render faster.

Write an object with `width` and `height` fields, both in pixels.

Examples:
```output_size:
  width: 500
  height: 300
```

### Input space
Sample mathematical values in this range.

Symmetry starts from `(minx, miny)`.
It applies the formula you specify and saves the result.
It uses the sample space (see below) to figure out what color it should pick from the input image.

Then it moves on from minx to maxx and applies the formula again. It will sample one row at a time until it reaches `(maxx, maxy)`.

What is the “right” input space varies by formula. If your input space is too large, you may pick uninteresting points that either converge to 0 or escape to infinity (so you’ll get one color or transparency.)

#### Help! My output is a solid blob of just one color!
Many values you can plug into the formula will converge towards 0. If your input space is too large, your result will be one solid color, usually at the center of your sample space.

Use a smaller input space so you remove the uninteresting 0 values and converge on the fun stuff.

#### Help! My output is one tiny blob in a transparent background!

Conversely, many values will instead grow so large they will escape to infinity. Your image will have a lot of transparency and maybe a blip of color at the center.

### Sample space
The program applied the formula to your input space (see above) and now has a lot of transformed points. Now it has to use those points somehow.

Sample space lets you say: “If the transformed point falls in this range, choose a color from the input image.” The transformed point is linearly mapped across the image, so values close to `(minx, miny)` will pick the upper left corners of the image, while `(maxx, maxy)` will pick the lower right.

Symmetry generator won’t pick a color (and keep it transparent) if the transformed point:
- escaped to infinity (or minus infinity) in X or Y coordinates.
- is outside of the sample space.

Like the input space, there is no “right” sample space. When you run this script it will show you the range of the transformed points.

```Example:
TODO insert example here
```

#### Help, my colors are really thin with lots of transparency in the middle.

If your sample space is too small, you may be rejecting too many transformed values. Expand your sample space so it captures more transformed points.

#### Help, the output has too many colors and looks messy.
If your sample space is too big, the transformed values will come from all over the photo. This might be too noisy.

Shrink your sample space or move the sample space so it covers a more homogenous area.

#### Help, every output focuses on the center of my input image.
These formulas tend to transform most values to near 0. If your sample space ranges from (-x, x) and (-y, y), this means (0,0) is the center of the sample space. Most of your output image will have those colors.

Move your sample space over to a different part of the image, so that the center is towards one of the corners. A range of (0, y) will pull the center output color to the top, for example.

## Formula
The transformation part of the file itself varies based on the type. If there are multiple keys it will only transform with one of them. Here’s the priority:

- `lattice` (TODO: Add link to lattice)
- `frieze` (TODO: Add link to frieze)
- `rosette` (TODO: Add link to rosette)

So if you have a `lattice` and `rosette` section, this will create a lattice pattern.