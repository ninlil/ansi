# columns

a small package to generate ansi-styling and coloring strings

## Getting Started

### Installing

Just do a `go get` on this package and your good to go.
```
go get github.com/ninlil/ansi
```

## Features

* Static or dynamic effect support
* No dependencies
* Stringer-interface compatible

## Examples

Static example:
```go
fmt.Printf("Here is a %stext in green%s", ansi.Green, ansi.Default)
```

Dynamic styling:
```go
func checkTemp(temp float64) {
	fmt.Printf("temp is %s%v%s\n", ansi.Green.YellowFore(temp > 65).RedFore(temp > 100), temp, ansi.Default)
}
...
checkTemp(37.5)  // will print value in green
checkTemp(73.5)  // will print value in yellow
checkTemp(105.5) // will print value in red
```

Different coding styles:
```go
error1 := ansi.Red.Background() | ansi.Yellow | ansi.Bright
error2 := ansi.NewStyle(ansi.Red.Background(), ansi.Yellow, ansi.Bright)
error3 := (ansi.Yellow | ansi.Bright).RedBack(true)
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/ninlil/columns/tags). 

## Authors

* **ninlil** - *Initial work* - [ninlil](https://github.com/ninlil)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
