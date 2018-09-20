wanikaniwallpaper-go
====================

Installation with Go Dep
------------------------

Clone this repository, install dependencies and then compile for your platform:

```
dep ensure
go build
```

Command-line Options
--------------------

Run `./wanikaniwallpaper-go -h` to print the usage statement:

```
Usage of ./wanikaniwallpaper-go:
  -dpi float
    	screen resolution in Dots Per Inch (default 72)
  -fontfile string
    	filename of the ttf font (default "font/ipag.ttf")
  -height int
    	height of wallpaper (default 1200)
  -key string
    	API key
  -orderfile string
    	path to order file (default "data/order")
  -output string
    	path to ouput file (default "out.png")
  -size float
    	font size in points (default 12)
  -width int
    	width of wallpaper (default 1920)
```

Run `./wanikaniwallpaper-go -key YOUR_KEY_HERE` to execute the program with the defaults above to create `out.png`.
