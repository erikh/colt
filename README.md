# colt: a tool for getting terminal color codes

`colt` is a tool for getting terminal color codes. It uses a trivial syntax; aims to be similar to `tput` for terminals. Embed `colt` in shells or use `colt env` to get all environment variables for your terminal definition in one easy way.

## Example

Pass a color to `colt` in a shell to print its terminal definition ("escape code"), like so:

```bash
#!/bin/bash

echo "$(colt red)This is red text! $(colt blue on white)This is blue with a white background!"
```

## Syntax

TBD

## Author

Erik Hollensbe <github@hollensbe.org>

## License

MIT
