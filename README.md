# Titan
A Booster Rocket for Gemini Apps.

Leverage [a-h's Go Gemini library](https://github.com/a-h/gemini) and [Go Templates](https://golang.org/pkg/text/template/) to make simple interactive capsules.

For a full example, check out Titandon, a gemini mastodon client.

# Features
- Pages
- Components
- Actions
- Inputs
- HTTP compatibility


## Usage

Define your Backend:
```
package main

import (
	"log"

	"github.com/a-h/gemini"
	"github.com/mibzman/titan"
)

func main() {
	booster := titan.Startup()

	booster.AddPage("/helloworld", "helloWorld.ttn", func(w gemini.ResponseWriter, r *gemini.Request) interface{} {
		return struct{ Message string }{"Hello World!"}
	})

	//generate TLS certs
	cert, err := CreateCert()
	if err != nil {
		log.Fatal("error:", err)
	}

	booster.Launch("localhost", cert)

}
```

Define your page in the `/pages` directory using golang templates
```
# A message from the Titan Server:
{{.Message}}
```

And you're flying!

## Components

Define components in the `/components` directory:
```
{{define "Example"}}
{{.Content}}
{{end}}
```

And use them in any page:
```
{{template "Example" .ComponentData}}
```

## Actions

When you want to preform an action without taking any input or displaying a page, define an action:

```
server.AddAction("/example", func (w gemini.ResponseWriter, r *gemini.Request) {
	log.printf("user hit example action")
})
```

## Input

When you need user input, define an input:

```
server.AddInput("/exampleinput", "Prompt Text", func (w gemini.ResponseWriter, r *gemini.Request) {
	UsersInput := titan.GetQuery(r)

   	log.printf("user entered: %v", UsersInput)

	w.SetHeader(gemini.CodeRedirect, "/timeline")
})
```

## HTTP compatibility
by default, Titan spins up both a gemini server and an http proxy server.  If desired, you can disable either by setting the relevant `LaunchConfig` setting to false.

# FAQ

## Why does this exist?
Gemini as a protocol seems perfect for creating extremely simple apps very quickly. Since the user interactions are so simple, Titan was created in ~4 hours of programming, and applications can be made in minutes.

## Why is it called Titan?
During NASA's Gemini program, Gemini capsules were launched on a booster called the [Titan II](https://en.wikipedia.org/wiki/Titan_II_GLV).  The Titan framework preserves Gemini's simplicity, but gives capsules a little boost to get them going.

## Why on earth does Titan also serve HTTP pages?
I didn't see the point in only supporting Gemini.  Pages, Actions, and Inputs all make sense in http, so why not?  The pages are ridiculously fast since there is no javascript.  If desired, you can always disable it.
