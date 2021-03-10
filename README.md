# htmlrepl

With this small library you can remove all HTML tags from a string. How this works I show you here.

## Install

```console
go get github.com/jojojojonas/htmlrepl
```

## How to use?

The function call is quite simple. You simply call the following function and set the string in the function.

```go
// Replace test html string
text, err := htmlrepl.Replace("html")
if err != nil {
	fmt.Println(err)
}

// Print out to terminal
fmt.Println(text)
```

## Help
If you have any questions or comments, please contact us by e-mail at [jonas.kwiedor@jj-ideenschmiede.de](mailto:jonas.kwiedor@jj-ideenschmiede.de)