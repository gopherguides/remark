# Converting a Static Site to Buffalo

This is a `code` sample.

---

## Overview

<details>
overview: true
</details>

I recently had a friend come to me and mention that their website I created for them years ago needed some updating.  After looking at it, it was clear it was time for an update.  This article will walk through the design decisions and technical challenges that I needed to address in bringing their site up to some modern standards.

---

youtube: lmzpYYRHHbM

## Video

---

hidetitle: true

##  Article

I recently had a friend come to me and mention that their website I created for them years ago needed some updating.  After looking at it, it was clear it was time for an update.  This article will walk through the design decisions and technical challenges that I needed to address in bringing their site up to some modern standards.

---

### Related Content

Before getting too far, if you already watched the video and simply want to browse the code, here is the repository:

[github.com/gopherguides/convert-static-site-to-buffalo](https://github.com/gopherguides/convert-static-site-to-buffalo)

---

### The Problem

The old site had many problems. First, the layout was no longer mobile friendly.

<div class="text-center">
<img src="legacy-site.png" title="Legacy Site" class="img img-fluid mb-3 mx-auto"></img>
</div>

Because the site was no longer mobile friendly, using a device such as a tablet or phone, users would no longer find their business when searching for them with popular search engines.

The next problem was the image gallery.  I could simply create static pages for all those images, but I wanted a better solution than that.  Additionally, I wanted some modern navigation, etc.

To summarize, I had to accomplish the following:

- Update the layout to be mobile friendly
- Create a dynamic image gallery
- Create dynamic navigation
- Allow for future updates that may need to be `dynamic` in nature

---

### Picking A Technology

There are endless options to choose from for modern web development. However, I wanted something that could satisfy the following:

- Easy deployment
- Fun to work with
- Modern asset pipline (scss, asset signing, etc.)
- Live Reload
- Full Web Stack
- Quick to install and get running

As you can imagine, that quickly narrows down both the technology and the frameworks available.  We all know that Ruby on Rails can do all that (with varying degrees on deployment depending on your hosting solution).  However, I love [Go](https://golang.org), so naturally, I was biased towards a solution writtin in go.

---

### Enter Buffalo

[Buffalo](https://gobuffalo.io/en/) satisfied all of my criteria (and much more).  I didn't need a database for this project, but knowing that I can easily add support for that in the future is also a big win for this project.

Now that I've decided on a technology, it was time to get started.  The first thing I had to do was [download](https://github.com/gobuffalo/buffalo/releases/) the latest _Buffalo_ binary to generate my project.

Once you have downloaded the latest version, you can verify you have the correct version by running the following command:

```sh
$ buffalo version
INFO[0001] Buffalo version is: v0.15.3

```

---

### Generating The Project

The first thing we need to do is now create the initial project.  This will do things like:

- Create an asset pipeline
- Create a starting route and template to work from
- Create some basic configuration

To create the project, I ran the following command:

```sh
buffalo new craigscustomsound --skip-pop
```

The `--skip-pop` flag tells the project not to generate any database dependencies.

You will see something like the following output as the project creates all of the initial files needed (the output is abbreviated for purposes of this blog):

```sh
DEBU[2020-01-24T12:31:31-06:00] Step: e9ef182f
DEBU[2020-01-24T12:31:31-06:00] Chdir: /Users/corylanou/tmp/craigscustomsound
DEBU[2020-01-24T12:31:31-06:00] File: /Users/corylanou/tmp/craigscustomsound/.codeclimate.yml
DEBU[2020-01-24T12:31:36-06:00] File: /Users/corylanou/tmp/craigscustomsound/webpack.config.js
.
.
.
DEBU[2020-01-24T12:31:36-06:00] LookPath: yarnpkg
DEBU[2020-01-24T12:31:36-06:00] Exec: yarnpkg install --no-progress --save
DEBU[2020-01-24T12:31:37-06:00] yarn install v1.15.2

DEBU[2020-01-24T12:31:37-06:00] info No lockfile found.

DEBU[2020-01-24T12:31:37-06:00] [1/4] Resolving packages...

DEBU[2020-01-24T12:31:37-06:00] warning popper.js@1.16.1: Popper changed home, find its new releases at @popperjs/core

DEBU[2020-01-24T12:31:51-06:00] [2/4] Fetching packages...

DEBU[2020-01-24T12:31:54-06:00] [3/4] Linking dependencies...

DEBU[2020-01-24T12:31:57-06:00] [4/4] Building fresh packages...

DEBU[2020-01-24T12:32:04-06:00] success Saved lockfile.

DEBU[2020-01-24T12:32:04-06:00] Done in 27.45s.

DEBU[2020-01-24T12:32:04-06:00] Step: e19e0b2c
DEBU[2020-01-24T12:32:04-06:00] Chdir: /Users/corylanou/tmp/craigscustomsound
DEBU[2020-01-24T12:32:05-06:00] Step: 5d306080
DEBU[2020-01-24T12:32:05-06:00] Chdir: /Users/corylanou/tmp/craigscustomsound
DEBU[2020-01-24T12:32:05-06:00] File: /Users/corylanou/tmp/craigscustomsound/.gitignore
DEBU[2020-01-24T12:32:05-06:00] Exec: git init
Initialized empty Git repository in /Users/corylanou/tmp/craigscustomsound/.git/
DEBU[2020-01-24T12:32:05-06:00] Exec: git add .
DEBU[2020-01-24T12:32:05-06:00] Exec: git commit -q -m Initial Commit
INFO[2020-01-24T12:32:05-06:00] Congratulations! Your application, craigscustomsound, has been successfully built!
INFO[2020-01-24T12:32:05-06:00] You can find your new application at: /Users/corylanou/tmp/craigscustomsound
INFO[2020-01-24T12:32:05-06:00] Please read the README.md file in your new application for next steps on running your application.
```

The longest part of this task is actually creating the asset pipeline and downloading all the `node` dependencies.

Once that is done, you will have a directory called `craigscustomsound` (or whatever you used for your project name).  That directory will have the following structure:

```sh
.
├── Dockerfile
├── README.md
├── actions
├── assets
├── config
├── fixtures
├── go.mod
├── go.sum
├── grifts
├── inflections.json
├── locales
├── main.go
├── node_modules
├── package.json
├── public
├── templates
├── webpack.config.js
└── yarn.lock

```

---

### Running The Server

Now that the initial project is generated, I can start up the web server and actually see the new site that was generated.

First, change directories to the project you created:

```sh
cd craigscustomsound
```

Now, start the server with the following command:

```
buffalo dev
```

This will start up all services needed to browse the new site.  It starts up on [localhost:3000](http://localhost:3000).

You will now see the initial page that was created for the project:

<img src="initial-project.png" class="img img-fluid" title="Initial Project"></img>

---

### `How it Works`

There are several files that you will need to know about to start creating and working with your new project. First, we'll start with the `go` files that contain the code for the `handlers`.


<br>

There are two primary files to be aware of.  They are `app.go` and `home.go`, and reside in the `actions` folder:

```sh
actions/
├── app.go
└── home.go
```

In the `app.go` file, the initial `routes`, `middleware`, and more are wired into the system:

<code src="./src/initial/actions/app.go" section="app"></code>

Most notably, what we care about, is the `home` handler that is being wired up to serve the `/` default route:

<code src="./src/initial/actions/app.go" section="home"></code>

The next file, `home.go`, is the code that will actually serve the `/` route, as well as decide which template to use.

<code src="./src/initial/actions/home.go"></code>

The above code is the boilerplate for simply stating to load the `index.html` template, and serve it with a status code of `200` (ok).

Next, let's take a look at the templates for the site.

---

### Templates

The templates for the site reside in the `templates` directory.  There are three initial files:

```sh
templates/
├── _flash.plush.html
├── application.plush.html
└── index.plush.html
```

The file `application.plush.html` is your `site` template, and where you would add any `global` content, such as Google Analitics tracking, and things of that nature.

<code src="./src/initial/templates/application.plush.html"></code>

Notice it is using a `partial` file for showing `flash` messages.  This is very much like how Ruby on Rails does partials as well.  Also notice that the `partial` file of `_flash.plush.html` starts with an underscore (`_`).  This convention is used for all partials when creating template files.

Next is the `index.plush.html` file.  This is a fairly large file to start, but remember that it rendered the default page that showed all of your routes and initial site info.  You can feel free to replace all of the content in this page with your content for the site you are building.

---

### Creating the New Site

Now that we have a quick tour of the files we need to work with, we will want to add all of our pages.  The pages we are creating are as follows:

* Home Page
* About Us Page
* Contact Us Page
* Products & Services Page
* Image Gallery Page

To do that, I'll create the following new files in the `templates` directory:

```sh
about.plush.html
application.plush.html
contact.plush.html
gallery.plush.html
products.plush.html
```

And remember we already have this file to re-purpose:

```sh
index.plush.html
```

Next, I need to create routes for each of these files, handlers, and serve each of the new templates.

For the handlers, I'll create the following files in the `actions` directory:

```sh
about.go
contact.go
gallery.go
products.go
```

And remember we already have the `home.go` file that is already serving our index page.

Each of these files will have the same basic handler and serve the correct template for each one.  For example, this is the content of `about.go`:

<code src="./src/craigscustomsound/actions/about.go"></code>

Finally, we need to wire the routes up.  We do this in `app.go`

<code src="./src/craigscustomsound/actions/app.go" section="routes"></code>

This is now the initial set of files needed to start building out the entire site.

---

### Routing

To ensure that we have the correct routes, we can use the `buffalo` cli (command line interface) tool to list the actual routes.  This can be done with the following command:

```sh
 $ buffalo routes
 METHOD | PATH       | ALIASES | NAME         | HANDLER
 ------ | ----       | ------- | ----         | -------
 GET    | /          |         | rootPath     | craigscustomsound/actions.HomeHandler
 GET    | /about/    |         | aboutPath    | craigscustomsound/actions.AboutHandler
 GET    | /contact/  |         | contactPath  | craigscustomsound/actions.ContactHandler
 GET    | /gallery/  |         | galleryPath  | craigscustomsound/actions.GalleryHandler
 GET    | /products/ |         | productsPath | craigscustomsound/actions.ProductsHandler
```

You can see that it lists out the `Method`, `Path`, `Alias`, `Name`, and associated `Handler` for each route.

---

### Helpers

Because I want navigation to `track` and `highlight` the page it is on, I will need to create a helper. The site is using [bootstrap](https://getbootstrap.com/), and for navigation, the have an [active](https://getbootstrap.com/docs/4.4/components/navbar/#nav) class that will automatically highlight the page.

The first thing needed is to create a `partial` for the navigation html template.  I'll create the following file:

```sh
templates\_navbar.plush.html
```

I can then include that in my `application.plush.html` file with the following tag:

<code src="./src/craigscustomsound/templates/application.plush.html" section="nav"></code>

Now I only need to edit one file for all of the sites navigation.

For the helper, I effectively want to be able to specify the `route` that the navigation should match.  It should write out the `active` class for the navigation element if it matches that route.  So first, I'll create a helper.

We typicaly add our global helpers in `actions/render.go`.

This is the helper I created to look at the current route, and if it matches what was passed in via the template, it will return out the string `active`:

<code src="./src/craigscustomsound/actions/render.go" section="helper"></code>

Next, we need to update `templates/_navbar.plush.html` to use the helper:

<code src="./src/craigscustomsound/templates/_navbar.plush.html"></code>

We needed to put the helper in every navigation element and specify the corresponding route.

---

### The Static Content

Now that I have dynamic navigation working, I was able to finish out all the static pages by simply creating the content from the previous site.  This brings me to the following layout:

<img src="home.png" title="New Site" class="img img-fluid mb-3 mx-auto"></img>

---

### Image Gallery

The image gallery presents a little bit of code.  I wasn't concerned with needing any type of interface for this specific client to post new photos. Since they don't update their images very often, I just needed something that could read the current directory of images, and then list them out as thumbnails, with the ability to view the larger image.

The first thing we need to understand is that _Buffalo_ `packages` all static assets.  This means that it looks at files such as `css`, `images`, etc. and compiles them into a binary representation.  It then takes that compiled representation and builds it into the final binary that you ship to production.  To make this work, it creates as `assetBox` in the `actions` package:

By default, it packages up the `public` directory:

```go
var assetsBox = packr.New("app:assets", "../public")
```

And it serves that directory from `app.go` with the following command:

```go
app.ServeFiles("/", assetsBox) // serve files from the public directory
```

Now that we know how assets are stored, I can write my handler for the gallery page.

I have three directories I want to server images from:

* commercial
* theater
* prewire

I'll do the same pattern for all of these assets, but we'll show the example of how I got a list of images from the `commercial` directory:

```go
commercial := []string{}
assetsBox.Walk(func(nm string, f packd.File) error {
	if strings.HasPrefix(nm, "assets/images/gallery/commercial/medium/") {
		commercial = append(commercial, strings.TrimPrefix(nm, "assets/"))
	}
	return nil
})
```

Effectively, I'm walking the `assetsBox` that we created when we fired the site up. I match any file that has the desired path, and store it in a slice.  Then I can add that slice to my pages context with the following command:

```go
c.Set("commercial", commercial)
```

This now allows me to access this list of information from my site template.

Here is the entire handler for the gallery with all three directories being processed:

<code src="./src/craigscustomsound/actions/gallery.go"></code>

Next, we have to create the html layouts.  I'm using bootstraps pre-defined classes for both the images, and modal popups. To see how the final layout was accomplished, you can refer to the video at the top of this page as well as the repo.

By using _Bootstrap_ and _Buffalo_, I was able to create a dynamic image gallery in about an hour that satisfied all of the criteria for this project.

---

### Building the Binary

Now that the entire site is complete, we want to build and deploy the binary.  I'm hosting on a simple droplet at [Digital Ocean](https://www.digitalocean.com/products/droplets).  There are a number of ways to deploy a _Buffalo_ app and they are well [documented](https://gobuffalo.io/en/docs/deploy/building/), so I will only be showing how to build the binary itself.

You can use the following command to `build` a single binary for your target platform to deploy your project:

```sh
buffalo build
```

This will place a built binary called `craigscustomsound` and place it in the `bin` directory.

It is important to keep in mind that by default this binary is built for the operating system and architecture of the computer I am on. Since I am deploying this on a linux based machine, I need to specify the target operating system:

```sh
GOOS=linux buffalo build
```

You can specify any target architecture or operating system for _Buffalo_ the same way the Go build tools work.  For more information on what options are available, you can read the [Building Go Applications for Different Operating Systems and Architectures](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures) article.

Now I have a single static binary that I can place on my droplet and run without any additional files.

---

### Production Configuration

The last thing I need to do is configure the service to run on a different port.  By default, the `buffalo dev` command starts on port `3000`.  However, for my environment, that won't work.

I'm using a reverse proxy on my server, so I need it to run on port `8001`.  I could accomplish this by hard coding the port in my program, setting an environment variable on my host, or by using a configuration file. For this project, I opted to use the configuration file.

By default, all _Buffalo_ projects will look for a `.env` file in the root of the project. You can find information for more configuration options [here](https://gobuffalo.io/en/docs/getting-started/config-vars#support-for-env-files).

This is the config for this project in production:

<code src="./src/craigscustomsound/env"></code>

---

### Summary

<details>
summary: true
</details>

In this project we went from an outdated, static website, to a modern dynamic website.  By using the _Buffalo_ project, I was able to quickly scaffold my pages and use bootstrap to lay them out.  This project barely touches any of the features that the _Buffalo_ project provides, but also shows that you can customize _Buffalo_ for just the needs at hand.

Overall, the project took about 2 hours from start to finish.  As a reference, this blog article and video production took three days.  I hope you enjoyed this article.
Feel free to contact me with suggestions and feedback on this article on [twitter](https://twitter.com/corylanou).
