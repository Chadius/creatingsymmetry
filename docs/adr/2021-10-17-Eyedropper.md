# Why
What do we do after applying the formula and transforming the color?
The current `eyedropper_boundary` does 2 things.
- It acts as a filter. Values that are out of range are replaced with transparency.
- It maps to a color in the source image, using the entire thing as the target.

The first item is a CoordinateFilter.
The second item is an Eyedropper.

`eyedropper_boundary` was actually a filter that eyedropped against the entire source image. Let's separate those concepts.

# What is it
## New classes
### MappedCoordinate
Holds an x and y coordinate
Knows if it was filtered
Knows if it is infinity
Holds an eyedropped x and y coordinate

### MappedCoordinates
Holds an array of Coordinates
It can tell the minimum and maximum non-filtered, non-infinity Coordinates

### CoordinateFilter
Has settings for min x/y and max x/y
Can apply filter against MappedCoordinates
Will mark filtered Coordinates
If omitted, include all values

### Eyedropper
Only looks at Coordinates that were not filtered and have no infinity
Gets minimum/maximum x/y coordinates to define range
Maps against non-filtered, non-infinity

### ColorSampler
Uses an output image
Works with MappedCoordinates and gets the color from the image at that point
Returns an image

# What can we do now?
Can shape the eyedropper to be a circle, or to loop when it is outside the image.
Need to rewrite docs to explain both objects.
Add Null Objects:
- A Threshold Filter that accepts all non-Inf or non-NaN numbers
- An Identity Formula

# Caveats that will trigger future change
Performance hit is too great

# NOTES
