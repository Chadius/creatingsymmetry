# Why
Unit tests for the process should be kept lightweight.
We should use mocks and stubs for the inner objects so we don't need to know details.
We also don't want these tests to fail when the Eyedropper fails, for example.

# What does this look like
Change Eyedropper to RectangularEyedropper
Change CoordinateThreshold to RectangularCoordinateThreshold
