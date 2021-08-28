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
```yaml
output_size:
  width: 500
  height: 300
```

### Sample space
Sample mathematical values in this range. You can think of it as zooming in/out your picture.

Symmetry starts from `(minx, miny)`.
It applies the formula you specify and saves the result.
It uses the sample space (see below) to figure out what color it should pick from the input image.

Then it moves on from minx to maxx and applies the formula again. It will sample one row at a time until it reaches `(maxx, maxy)`.

What is the “right” input space varies by formula. If your input space is too large, you may pick uninteresting points that either converge to 0 or escape to infinity (so you’ll get one color or transparency.)

#### Examples
Let's start with a rosette. They tend to have a hole at the center, a ring, and then some more detail around it.
The sample space was selected to capture the "petals" that surround the ring.

```yaml
sample_space:
  minx: -8e-1
  miny: -8e-1
  maxx: 8e-1
  maxy: 8e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2.png)

[(Link to formula)](../example/rosettes/rainbow_stripe_rosette_2.yml)

Let's say I want to stretch the horizontal sample space. I want more detail along the horizontal space.
I will reduce the distance between minx and maxx by half. I won't be able to see as much, though.

```yaml
sample_space:
  minx: -4e-1
  miny: -4e-1
  maxx: 8e-1
  maxy: 8e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2_sample_space_1.png)

[(Link to formula)](../example/rosettes/rainbow_stripe_rosette_2_sample_space_1.yml)

Let's zoom out and see what the extremes. Give a large distance between the x and y axes.
We will lose detail but we get to see more of the pattern.
```yaml
sample_space:
  minx: -64e-1
  miny: -64e-1
  maxx: 64e-1
  maxy: 64e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2_sample_space_2.png)

[(Link to formula)](../example/rosettes/rainbow_stripe_rosette_2_sample_space_2.yml)

### Color value space
The program applied the formula to your input space (see above) and now has a lot of transformed points. Now it has to use those points somehow.

Color value space lets you say: “If the transformed point falls in this range, choose a color from the input image.” The transformed point is linearly mapped across the image, so values close to `(minx, miny)` will pick the upper left corners of the image, while `(maxx, maxy)` will pick the lower right.
Most colors will be near (0,0), so choose your center where you want to see the majority of your colors.

Symmetry generator won’t pick a color (and keep it transparent) if the transformed point:
- escaped to infinity (or minus infinity) in X or Y coordinates.
- is outside of the sample space.

Like the sample space, there is no “right” color value space. When you run this script it will show you the range of the transformed points.

##### Examples
Because the rainbow stripe file is the same horizontally, it doesn't matter what we pick for minx and maxx, as long as our range isn't too narrow.
So for these examples, we'll focus on miny and maxy.

Let's take a look at a frieze file.

It has a balanced color value space, where (0,0) is near the center. The source image is a green stripe at the center, so the majority of the pattern is green.
```yaml
color_value_space:
  minx: -1.1e1
  maxx: 1.1e1
  miny: -1.8e1
  maxy: 1.8e1
```

![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a green background](../example/friezes/rainbow_stripe_frieze_p2mg.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg.yml)

The bottom of the source image is orange/red, so let's shift miny and maxy in a negative direction. This will put most of the 0 values in the orange/red section.

```yaml
color_value_space:
  minx: -1.1e1
  maxx: 1.1e1
  miny: -2.8e1
  maxy: 0.8e1
```
![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a checker board orange and red background](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_orange.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_orange.yml)

Now, let's say you want smaller valleys and you want to stretch them out more.
You can increase the distance between extremes.
More values will land in that range and will be drawn.

```yaml
color_value_space:
  minx: -1.1e1
  maxx: 1.1e1
  miny: -5.8e1
  maxy: 1.8e1
  ```
![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from an orange background with red valleys](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_extra_thick.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_extra_thick.yml)

This pulls the red away from the orange, maybe you don't like a checkerboard pattern.

## Formula
The transformation part of the file itself varies based on the type. If there are multiple keys it will only transform with one of them. Here’s the priority:

- `lattice`
- `frieze`
- `rosette`

So if you have a `lattice` and `rosette` section, this will create a lattice pattern.