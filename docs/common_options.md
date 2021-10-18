# Where is the formula file located?
By default, this program looks for `data/formula.yml`.
It should contain:
- The [formula](#transformation-formula) that transforms the source into the output
- Numerical ranges used with the formula
  - [pattern viewport](#pattern-viewport)
  - [eyedropper boundary](#eyedropper-boundary)

## How does it work?
1. The input numerical range is used to create input pattern viewport.
2. Using the formula, it transforms the input.
3. A threshold is applied to filter transformed points that are out of range.
4. Using an eyedropper, the transformed points are mapped to the source image to figure out what color should be used.

### Example
- We have a `100x100` sample image. 
- We want a `200x200` output image.
- Our formula doubles the given `(x,y)` coordinates.
- The pattern viewport is from `(0,0)` to `(10,10)`.
- The threshold filter is from `(0,0)` to `(40,40)`
- The eyedropper boundary is from `(0,0)` to `(200,200)`

The pattern viewport will be sampled 200x200 times since that's the resolution of the output image.
So each sample will have a difference of 1/20 from the previous.  

Starting with `(0,0)`:
- The formula doubles the coordinates `(0,0)` to `(0,0)`.
- This satisfies the threshold filter.
- The eyedropper maps `(0,0)` to `(0,0)`.
- `(0,0)` in the source image lines up with the top left corner.
- This will draw the same color as the top left corner of the source image.

Next sample is `(1/20,0)`:
- The formula doubles the coordinates `(1/20,0)` to `(2/20,0)`.
- This satisfies the threshold filter.
- The eyedropper will map `(2/20,0)` to somewhere near the top left corner.
- This will draw the same color as near the top left corner of the source image.

Let's look at this sample `(5,2)`: 
- The formula doubles the coordinates `(5,2)` to `(10,4)`.
- This satisfies the threshold filter.
- The eyedropper maps `(10,4)` to somewhere in the top left quadrant.
- This will draw the same color of as the one spot in the top left quadrant of the source image.

The last sample `(10, 10)`:
- The formula doubles the coordinates `(10, 10)` to `(20, 20)`.
- This satisfies the threshold filter.
- Eyedropper maps `(20,20)` to the center of the image.

The samples chosen will draw from the top left corner of the source image to the center.

## Common options
Every formula file contains these options.

### pattern viewport
Which part of the output image will we look at?
Control the pattern viewport to move around and zoom in and out.

You enter the mathematical values you would like to transform. 

Based on the resolution of the output image, it will break the pattern viewport into a number of discrete values.
It starts from `(x_min, y_min)`, moves in the x direction to `(x_max, y_min)`.
Then it goes back to `x_min` but moves in the y direction slightly towards `y_max`.
It stops when it reaches `(x_max, y_max)`.

It transforms each discrete sample value using the formula.
Then maps the result against the [eyedropper boundary](#eyedropper-boundary) 
to figure out what color to use from the input image.

What is the “right” input space varies by formula.
- If your range is too small, you may not see the repeating pattern.
- If your range is too large, the formula might converge everything to 0. You'll see mostly one color in your output.
- If your range is too large, the formula might move everything to infinity. In this case nothing will be drawn, and you'll have a transparent output image.

#### Examples
Let's start with a rosette. They tend to have a hole at the center, a ring, and then some more detail around it.
The pattern viewport was selected to capture the "petals" that surround the ring.

```yaml
pattern_viewport:
  x_min: -8e-1
  y_min: -8e-1
  x_max: 8e-1
  y_max: 8e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2.png)

[(Formula)](../example/rosettes/rainbow_stripe_rosette_2.yml)

I want to stretch the horizontal pattern viewport and see more detail.
I will reduce the distance between x_min and x_max by half. In exchange for more detail, I don't get to see as much of the overall image.

```yaml
pattern_viewport:
  x_min: -4e-1
  y_min: -4e-1
  x_max: 8e-1
  y_max: 8e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2_sample_space_1.png)

[(Link to formula)](../example/rosettes/rainbow_stripe_rosette_2_sample_space_1.yml)

Let's zoom out and see the extremes. Make a large distance between x_min and x_max. Same thing for y_min and y_max.
It's hard to see anything about the central ring, but now I get to see more of the pattern.

```yaml
pattern_viewport:
  x_min: -64e-1
  y_min: -64e-1
  x_max: 64e-1
  y_max: 64e-1
```

![Transformed rainbow stripe image into rosette with 4 rotational symmetry, creating purple and green petals on a blue background](../example/rosettes/rainbow_stripe_rosette_2_sample_space_2.png)

[(Link to formula)](../example/rosettes/rainbow_stripe_rosette_2_sample_space_2.yml)

### Coordinate Threshold
The transformed [pattern viewport](#sample-space) has many results, covering a wide numerical range. Sometimes you want to focus on a single mathematical range and ignore the rest. Coordinate Threshold to the rescue.

```yaml
coordinate_threshold:
  x_min: -4e1
  x_max: 4e1
  y_min: -6e1
  y_max: 2e1
```

`x_min`, `x_max`, `y_min`, `y_max` define the limits of the threshold.

These coordinates will not satisfy the filter and will turn transparent:
- X is less than `x_min`.
- X is more than `x_max`.
- Y is less than `y_min`.
- Y is more than `y_max`.
- X or Y is Infinity.
- X or Y is Not a Number (or undefined).

You'll need to pay attention to the terminal output to see the absolute ranges of the transformed points.

Like the pattern viewport, there is no “right” Coordinate Threshold.
- If your Coordinate Threshold is too small, the transformed values will fall outside, and you'll have a transparent image.
- If your Coordinate Threshold is too big, there is too much uniform noise. Your image will be a single color.

##### Examples
Eyedropper Boundary is easier to explain in one dimension, so these examples focus on `y_min` and `y_max`.

Let's take a look at a frieze file. Because `y_min` and `y_max` have the same distance from the center, the pattern's main color should be the same as the center of the source image.

The green stripe is at the center, so the frieze pattern should be mostly green.
```yaml
coordinate_threshold:
  x_min: -1.1e1
  x_max: 1.1e1
  y_min: -1.8e1
  y_max: 1.8e1
```

![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a green background](../example/friezes/rainbow_stripe_frieze_p2mg_alpha.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg_alpha.yml)

Let's push the center towards the orange/red part of the source image. That lies near the bottom, so `y_min` and `y_max`'s midpoint should be negative.

```yaml
coordinate_threshold:
  x_min: -1.1e1
  x_max: 1.1e1
  y_min: -2.8e1
  y_max: 0.8e1
```
![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from a checker board orange and red background](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_orange.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_orange.yml)

Values that fall out of this range become transparent. So if we increase the size of `y_min` and `y_max` then more pixels will be drawn.
Note how the valleys and peaks are more extreme.
Also note how the orange stripe is dominant. If we expanded the range more, we would get more orange.

```yaml
coordinate_threshold:
  x_min: -1.1e1
  x_max: 1.1e1
  y_min: -5.8e1
  y_max: 1.8e1
  ```
![Transformed rainbow stripe image into frieze with p2mg symmetry, with multicolored spikes emerging from an orange background with red valleys](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_extra_thick.png)

[(Link to formula)](../example/friezes/rainbow_stripe_frieze_p2mg_sample_space_extra_thick.yml)

### Eyedropper Boundary
Which part of the source image do you want to use? Pick the part of the source image to sample by adjusting the Eyedropper.

```yaml
eyedropper:
  left: 0
  right: 200
  top: 0
  bottom: 250
```

In this example, `(100, 125)` is the center of the image. Most of your transformed coordinates will be mapped near the center, so most of the pattern will look like the center.

## Transformation Formula
Only one formula will be rendered at a time. Use exactly one of these keys, based on the transformation formula you want:

* [Lattice](pattern_lattice.md)
* [Frieze](pattern_frieze.md)
* [Rosette](pattern_rosette.md)

The formulas are listed in priority order. So if you include multiple, it will look for a lattice formula first, then frieze and finally rosette.