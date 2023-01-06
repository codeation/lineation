# lineation

Mind Map Editor. Ready to catch your thoughts.

![demo screenshot](https://codeation.github.io/pages/images/lineation-test.png)

## Alpha Version

Notes:

- This project implements basic functions for working with Mind Map.
- The project tested on Debian 11.6 and macOS Big Sur

## Driver installation

The application uses [a separate GUI driver](https://github.com/codeation/it) for drawing
instead of binding low-level library to a Golang.

You can [download](https://github.com/codeation/it/releases)
the compiled binary `it` file or make it again from [the source](https://github.com/codeation/it).

## Run

To download Mind Map Editor run:

```
git clone https://github.com/codeation/lineation.git
cd lineation
```

You can specify the full path and name for the GUI driver via the environment variable, for example:

```
IMPRESS_TERMINAL_PATH=/path/it go run ./cmd demo/demo.xml
```

or just copy the downloaded GUI driver (file with name `it`) to the working directory and run:

```
go run ./cmd demo/demo.xml
```

## Mind-map format

To create a blank mind map copy to `blank.xml`:

```
<map>
    <node text="Main theme"></node>
</map>
```
