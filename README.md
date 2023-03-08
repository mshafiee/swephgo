# SwEphGo
Golang interface binding to the Swiss Ephemeris C library v2.10.01.

Swiss Ephemeris is a powerful and versatile astronomical calculation library that serves the needs of astrologers, archeoastronomers, and, depending on the purpose, also astronomers. It includes long-term ephemerides for the Sun, the Moon, the planets, over 300,000 asteroids, historically relevant fixed stars, and some hypothetical objects.

## Installation

To get started with using SwEphGo, you'll first need to install the Swiss Ephemeris Library. You can download it [here](https://www.astro.com/ftp/swisseph/) and then follow these installation steps:

1. To compile the library, navigate to the `src` folder of the downloaded library and execute `make libswe.so`. In some environments, you may need to modify the build command by adding `-lm -ldl` to `libswe.so` in the `Makefile`. Here's the updated Makefile code:
```makefile

libswe.so: $(SWEOBJ)
	$(CC) -shared -o libswe.so $(SWEOBJ) -lm -ldl
```

2. After compiling the library, copy the `libswe.so` file to `/usr/local/lib/`.
```sh
$ cp libswe.so /usr/local/lib/
```
3. Get the library with the following command:
```sh
$ go get github.com/mshafiee/swephgo
```

# Usage

Once you have the library installed, you can use it in your Go projects by importing it in the main package and calling the Version function to print out the version of the library being used. For example, you can use the following code:

````
package main

import (
	"fmt"
	"github.com/mshafiee/swephgo"
)

func main() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", sweVer)
}

````

To run the code above:

`````
$ go run main.go
`````

The output will be:
`````
Library used: Swiss Ephemeris v2.10.01
`````

# Disclaimer
This project includes C header files taken from the SWISSEPH source code. These rights belong to Astrodienst AG, Switzerland. For information on how to use the rights of these files, please visit the [Swiss Ephemeris](https://www.astro.com/swisseph/swephinfo_e.htm) website. The developer of this library is not affiliated with Astrodienst AG and this library is provided as is, without warranty of any kind, express or implied.


