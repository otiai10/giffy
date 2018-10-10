# Create Animated GIF

```sh
% go get -u github.com/otiai10/giffy
```

```sh
% giffy -i 'example/*.png'
Decoding example/01.png
Decoding example/02.png
Decoding example/03.png
Decoding example/04.png
Decoding example/05.png
Decoding example/06.png
Decoding example/07.png
Decoding example/08.png
Decoding example/09.png
Encoded successfully: animated.gif
%
```

```
% giffy -h
Usage of giffy:
  -delay int
    	Delay in milliseconds (default 1000)
  -i string
    	Input source images
  -loop int
    	Loop count (0 == infinite)
  -o string
    	Output file name (default "animated.gif")
```
