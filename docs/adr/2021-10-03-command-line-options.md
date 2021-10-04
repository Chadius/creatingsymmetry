# Why
The formula doesn't care where it gets the data from.
A file, or a network transfer, typed in by hand, a stream, it shouldn't matter.
It is also easy to swap out input files, output files and options without typing the file.

# What is it
The output options is split.

- CLI uses a command line for the source file.
- CLI uses a command line for the output file.
- CLI uses a command line for the output file's resolution.
- CLI uses a command line for the formula file.

Change mental models so the program answers 3 questions:
- What is the source image?
- How do we transform the image?
- Where do we output the results?

# What can we do now?
- Indicate if the formula file is JSON or YAML
- Allow data streams as source or output. Don't need to make it a stream. 