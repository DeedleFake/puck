puck
====

puck is an experimental statically-linked package management system.

Why another package manager?
----------------------------

A very good question. There are a ton of package managers out there already, so why write another one? In fact, that exact question had been bothering me when I think I came up with an answer.

I had been wanting to try writing a simple package manager in Go just for the fun of it, but I hadn't been able to come up with a use case. I was experimenting with Alpine Linux, and I suddenly thought of a potential niche for a new package manager. It's generally a good thing when Docker images can be made small. Alpine helps with this a lot, averaging sizes of around 8 to 10 MB, vs. Ubuntu's multiple hundreds, but there's still essentially a full GNU/Linux system under the hood, minus the kernel. Whenever I make a Docker image for a Go application, I generally try to make it `FROM scratch`. However, this is generally a bit of a pain, as I usually need to compile the application in question under special conditions, and it might not really be possible on a given system.

So how about a package manager designed specifically for that? The package manager itself is written in Go, and thus can be easily fully statically linked. The goal is to provide a small set of basic packages with as many of them statically-linked as possible. Even better, this should cut down the number of dependencies each package has by extremely large amounts. While it won't eliminate them completely, as some packages may require other ones in order to run programs that they contain, it should make for a very clean, if slightly larger, ecosystem.

Design
------

* Provide `puck` Docker image for basing other images on.
* Should be completely capable of uninstalling itself. This is so that a `FROM puck` image can be as clean as a `FROM scratch` image once it's set up.
* Simple bootstrap installation script.
* JSON-based and very minimal configuration format.
* Some form of compressed tar files for actual packages.
* JSON-based package creation script system.
	* Single JSON file which describes the package.
	* Designates files to execute which put package files into directory for tarring.
* No local package cache by default.
* Very simple remote repository scheme.
* Allows local repositories.
* Some error avoidance.
	* Check for conflicting files when installing.
	* Check dependency versions.
