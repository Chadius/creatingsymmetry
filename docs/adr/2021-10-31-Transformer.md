# Why
Extract concepts and ideas into a new class
Dependency Injection to make the process more flexible

# What is this
A new interface, Transformer, turns one image stream into another
A new class, FormulaTransformer, turns one image stream into another using a formula

Dependency Injects:
- Input Image
- Filter
- Eyedropper
- Output Image size

- Formula (Eventually)

# Caveats
Formula doesn't fit into a single interface yet. Best to pass around a wallpaper command for now.