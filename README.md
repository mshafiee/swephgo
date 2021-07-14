# SwEphGo
Golang interface binding to the Swiss Ephemeris C library v2.10.01.

Swiss Ephemeris is a function package of astronomical calculations that serves the needs of astrologers, archeoastronomers, and, depending on purpose, also the needs of astronomers. It includes long-term ephemerides for the Sun, the Moon, the planets, more than 300.000 asteroids, historically relevant fixed stars and some “hypothetical” objects.

# Installation
Download the Swiss Ephemeris Library [here](https://www.astro.com/ftp/swisseph/). After compiling the library, copy the libswe.so file to /usr/local/lib/

````
$ cp libswe.so /usr/local/lib/
````

Get the library with the following command:

````
$ go get github.com/mshafiee/swephgo
````

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
There are a number of C header files in this project that are taken from the SWISSEPH source code and their rights belong to Astrodienst AG, Switzerland.
For information on how to use the rights of these files, visit [Swiss Ephemeris](https://www.astro.com/swisseph/swephinfo_e.htm) website.
