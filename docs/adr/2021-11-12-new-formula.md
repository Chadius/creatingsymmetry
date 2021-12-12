# Why
There are many formula objects that are very similar but slightly different.
We have a lot of if statements floating around to denote when to use what object.

# What is this
We should instead create one Formula interface, and then subclass it for the different types.
We also need a new formula, Identity, so we can stop checking for nil.
We should create a formula builder that can work off of commands and off of JSON/YAML to make formulas.
- A builder can also serve as a factory, especially when the object we need to make is ambiguous.

# What can we do now?
- Remove objects that were similar but distinct
- Support crossfade functions (it's just a formula inside another formula)
- Consumers have to change their format:

1. replace `*_formula` key with `formula`
2. Add a `type: X` key where X is the type of formula. For example `type: rosette`.