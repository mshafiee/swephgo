# SwEphGo
Golang interface binding to the Swiss Ephemeris C library v2.10.01.

Swiss Ephemeris is a powerful and versatile astronomical calculation library that serves the needs of astrologers, archeoastronomers, and, depending on the purpose, also astronomers. It includes long-term ephemerides for the Sun, the Moon, the planets, over 300,000 asteroids, historically relevant fixed stars, and some hypothetical objects.

## Installation

1. Download the Swiss Ephemeris Library [here](https://www.astro.com/ftp/swisseph/).
2. After compiling the library, copy the `libswe.so` file to `/usr/local/lib/`.
```sh
$ cp libswe.so /usr/local/lib/
```
3. Get the library with the following command:
```sh
$ go get github.com/mshafiee/swephgo
```

# Usage

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

RUN:

`````
$ go run main.go
Library used: Swiss Ephemeris v2.10.01
`````

# Disclaimer
This project includes C header files taken from the SWISSEPH source code. These rights belong to Astrodienst AG, Switzerland. For information on how to use the rights of these files, please visit the [Swiss Ephemeris](https://www.astro.com/swisseph/swephinfo_e.htm) website.


