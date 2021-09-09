# lineation

Mind mapping application

![simple screenshot](https://codeation.github.io/pages/images/lineation-test.png)

## Proof of Concept Version

Notes:

- This project is still in the early stages of development and is not yet in a usable state.
- The project tested on Ubuntu 21.04 and MacOS Big Sur (11.5)


## Installation

The application uses [a separate GUI driver](https://github.com/codeation/it) for drawing
instead of binding low-level library to a Golang.

You can [download](https://github.com/codeation/it/releases)
the compiled binary file or make it again from [the source](https://github.com/codeation/it).

You can specify the full path and name for the GUI driver via the environment variable, for example:

```
IMPRESS_TERMINAL_PATH=/path/it go run ./examples/simple/simple.go
```

or just copy the downloaded GUI driver to the working directory and run:

```
go run ./cmd test.xml
```

## Mind-map format

To create a simple mind map copy to `test.xml`:

```
<map>
    <node text="Main theme"></node>
</map>
```
