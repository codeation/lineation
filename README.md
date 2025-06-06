# lineation

Mind Map Editor. Ready to catch your thoughts.

<img src="https://codeation.github.io/images/lineation_demo.png" width="782" height="636" />

## Alpha Version

Notes:

- This project implements basic functions for working with Mind Map.
- The project was tested on Debian 12.10 and macOS 15.4.1.

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
