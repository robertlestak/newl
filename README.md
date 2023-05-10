# newl

replace new lines with a literal \n because I can't remember the sed or awk commands. that's it.

See discussion [here](https://stackoverflow.com/questions/1251999/how-can-i-replace-a-newline-n-using-sed)... and [here](https://unix.stackexchange.com/questions/140763/replace-n-by-a-newline-in-sed-portably)..... aaaand [here](https://stackoverflow.com/questions/38672680/replace-newlines-with-literal-n).

Unix purists will probably hate this, but I'm tired of looking up the sed/awk commands to do this.

## Usage

```bash
$ newl <file> <opt: output file>
```

If no output file is specified, the output will be printed to stdout.

If no input file is specified, the input will be read from stdin.

## Example

```bash
$ newl README.md
```