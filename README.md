# Attic

Attic is a static website generator. It's a very simple tool that uses
Go's [template syntax](http://golang.org/pkg/html/template/) to build static websites.

## Install

Assuming that you already have Go installed and set up properly, use the standard
Go tools for installing attic: `go install github.com/ericlevine/attic`

## Using Attic

To get started with Attic, first you need to create a configuration file. It's
a JSON file with the following fields:

* **InputDir** - directory containing your layout and pages
* **OutputDir** - directory where your output will generate
* **Layout** - layout for the generated templates
* **Pages** - list of pages that will be generated

## Example

In this example, we'll be assuming we have a directory structure with this layout:

```
example/
  templates/
    contact.html
    index.html
    layout.html
    resume.html
  static/
  config.json
```

In this example, we'd set up `config.json` to look like this:

```javascript
{
  "InputDir": "templates",
  "OutputDir": "static",
  "Layout": "layout.html",
  "Pages": ["index.html", "contact.html", "resume.html"]
}
```

Here's an example layout:

```html
<!DOCTYPE html>
<html>
  <body>
    <div id="nav">
      <ul>
        <li{{if .IsActive "index.html"}} class="active"{{end}}>
          <a href="index.html">Home</a>
        </li>
        <li{{if .IsActive "contact.html"}} class="active"{{end}}>
          <a href="contact.html">Contact</a>
        </li>
        <li{{if .IsActive "resume.html"}} class="active"{{end}}>
          <a href="resume.html">Resume</a>
        </li>
      </ul>
    </div>
    <div id="content">{{template "content" .}}</div>
  </body>
</html>
```

Here's an example `contact.html`:
```html
{{define "content"}}
Contact me by phone at +1-555-555-5555.
{{end}}
```

## Disclaimer

While this project as defined above should be totally usable for simple
sites, this is not nearly as feature complete as other tools that do
similar things. If you're looking for something full featured, consider
using [Jekyll](http://jekyllrb.com/) or [Middleman](https://middlemanapp.com/) instead.
