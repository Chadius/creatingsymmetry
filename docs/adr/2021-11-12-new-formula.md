# Why
There are many formula objects that are very similar but slightly different.
We have a lot of if statements floating around to denote when to use what object.

# What is this
We should instead create one Formula interface, and then subclass it for the different types.
We also need a new formula, Identity, so we can stop checking for nil.
We should create a formula builder that can work off of commands and off of JSON/YAML to make formulas.