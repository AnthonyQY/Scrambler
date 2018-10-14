# Scrambler
Free, Open-Source productivity negation program coded in Golang

This annoying little program is guaranteed to diminish both the sanity and 
productivity of your victim(s) at the same time!

## Disclaimer 
This code is provided for educational purposes only.
I am not responsible for any damage(s) caused by any use of this code.
Please use with caution.


# Features

### This program periodically...
* Clicks/Double clicks randomly!
* Types random strings of up to 100 characters!
* Scrolls up/down randomly!
* Moves the cursor to a random location on the screen!

### It is also...

* Fully customizable, fork and edit the code to your liking!

# Requirements

### Golang Version 1.10
https://golang.org

### RobotGo Package
https://github.com/go-vgo/robotgo

### GCC
https://gcc.gnu.org
Note: If you are using a 64-bit version of Golang, you will need to install a 64-bit version of GCC such as MinGW-w64 as well.
This same logic also applies to 32-bit versions of Golang.

# Build
### Windows
Note: The flags '-ldflags' and '-H=windowsgui' are used to hide the console of the program
```
go build -ldflags -H=windowsgui -o scrambler.exe scrambler.go
```

### Other Platforms
```
go build -o scrambler.exe scrambler.go
```

# Credits
### RobotGo

https://github.com/go-vgo/robotgo


