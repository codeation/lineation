# lineation

Mind mapping application

![demo screenshot](https://codeation.github.io/pages/images/lineation-test.png)

## Alpha Version

Notes:

- This project is still in the early stages of development.
- The project tested on Debian 11.6 and macOS Big Sur (11.5)

## Driver installation

The application uses [a separate GUI driver](https://github.com/codeation/it) for drawing
instead of binding low-level library to a Golang.

You can [download](https://github.com/codeation/it/releases)
the compiled binary file or make it again from [the source](https://github.com/codeation/it).

## Mind-map format

To create a simple mind map copy to `test.xml`:

```
<map>
    <node text="Main theme"></node>
</map>
```

## Run

You can specify the full path and name for the GUI driver via the environment variable, for example:

```
IMPRESS_TERMINAL_PATH=/path/it go run ./cmd test.xml
```

or just copy the downloaded GUI driver to the working directory and run:

```
go run ./cmd test.xml
```
