# Create Animated GIF

Install

```sh
% go get -u github.com/otiai10/giffy
```

Example

```sh
% giffy -i 'example/*.png'
9 files found. Decoding...
✔ ✔ ✔ ✔ ✔ ✔ ✔ ✔ ✔
Encoded successfully to animated.gif
%
```

Output

<img src="https://raw.githubusercontent.com/otiai10/giffy/master/example/animated.gif" width="40%" />

Options

```
$ giffy -h
Usage of giffy:
  -delay int
    	Delay in milliseconds (default 1000)
  -i string
    	Input source images
  -loop int
    	Loop count (0 == infinite)
  -o string
    	Output file name (default "animated.gif")
  -quiet
      Do not output verbose log
```
