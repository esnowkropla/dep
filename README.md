dep
===
A quick and dirty little graphics engine; draws billboard sprites in 3D space.  I wrote it to learn both [go](http://golang.org/) and OpenGL at the same time.  As a result, it's both horribly un-idomatic Go and old-style OpenGL.  Pretty sure the frame timing actually throws frames off rather than timing them consistently too.  Ah well.  

At one point I planned to play around with some stuff in it like running the simulation and the renderer in different goroutines so that the world would generate updates and the renderer could slurp them up (hooray concurrency!), but then [thesis](https://github.com/esnowkropla/eskthesis.git) happened.
