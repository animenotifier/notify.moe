# Task for new contributors

This task assumes that you have [installed](https://github.com/animenotifier/notify.moe#installation) notify.moe already, started the server with the `run` tool and have the code open in Visual Studio Code. It will teach you the basics by creating an entirely empty page within notify.moe.

# Step 1: Create a new page

Let's call it `foobar`. Create a new directory under `pages`, called `foobar`. Then create the following files inside it:

* foobar.go (controller)

```go
package foobar

import (
	"github.com/aerogo/aero"
)

// Get returns the contents of our amazing page.
func Get(ctx aero.Context) error {
	return ctx.HTML("Hey it's me, foobar!")
}
```

* foobar.pixy (template)

```pixy
component FooBar
	h1 Hi!
```


* foobar.scarlet (styles)

```scarlet
.foobar
	// Will be used later!
```

`foobar.pixy` and `foobar.scarlet` are currently not used but we'll deal with that later.

## Step 2: Route your page

Your page needs to become available on the `/foobar` route. Let's add it to `pages/index.go`, inside `Configure`:

```go
page.Get(app, "/foobar", foobar.Get)
```

Your IDE should automatically insert the needed package import upon saving the file.

## Step 3: Add sidebar button

Inside `layout/sidebar/sidebar.pixy`, add a new button inside the `Sidebar` component:

```pixy
SidebarButton("Foobar", "/foobar", "plus")
```

## Step 4: Confirm it's there!

Navigate to `beta.notify.moe` and you should see the button to access your newly made page! Yay!

## Step 5: Play around!

Feel free to play around with the code now. You can utilize pixy components by using the `components` package inside your controller:

```go
package foobar

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Get returns the contents of our amazing page.
func Get(ctx aero.Context) error {
	return ctx.HTML(components.FooBar())
}
```

This would now return the contents of your previously defined pixy component instead of a hard-coded string.