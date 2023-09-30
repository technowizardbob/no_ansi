# No ANSI
## Scripts to remove ANSI escape codes
Purpose to stop log files from doing bad things
## Build
```
$ go build
$ sudo ln -s $(pwd)/no_ansi /usr/local/bin/NoANSI
```
## Example useage: 
```
$ cat file.ansi | NoANSI
```
