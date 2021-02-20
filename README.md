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

Scribble
Rosette
Frieze
17 wallpaper types
3D shapes
Hyperbolic

Color changing

### Research notes
#### Saving a bytestream to an image:
https://stackoverflow.com/questions/39927182/looking-for-better-way-to-save-an-in-memory-image-to-file

f, err := os.Create("img.jpg")
if err != nil {
    panic(err)
}
defer f.Close()
jpeg.Encode(f, target, nil)

### Opening an image
import	_ "image/jpeg"
)

reader, err := os.Open("testdata/video-001.q50.420.jpeg")
if err != nil {
     log.Fatal(err)
}

https://blog.golang.org/image

### Complex Math library
https://golang.org/pkg/math/cmplx/#Abs
