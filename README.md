# colt: a tool for getting terminal color codes

`colt` is a tool for getting terminal color codes. It uses a trivial syntax; aims to be similar to `tput` for terminals. Embed `colt` in shell scripts for maximum effect.

## Example

Pass a color to `colt` in a shell to print its terminal definition ("escape code"), like so:

```bash
#!/bin/bash

echo "$(colt red)This is red text! $(colt blue on white)This is blue with a white background!"
```

## Syntax

There are three positions in the grammar that are relevant:

Foreground is always in the first position.

If you want to specify a background, specify it or use "on" to specify it with
a bit more of an english twist.

If you want to specify text attributes such as bold, underline, etc, specify them with "with", such as:

```
colt red with bold underline
colt red on white with bold
```

## Author

Erik Hollensbe <github@hollensbe.org>

## License

MIT
