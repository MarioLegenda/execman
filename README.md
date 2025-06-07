# Introduction

**execman** is a tool to execute code in different programming languages (for now, 14
different programming languages). It came from my abandoned project; a blogging platform
specifically designed for software developers where developers could write blogs with
code snippets that could be executed on the backend and for the result to be shown on
the frontend. The project didn't go anywhere but the microservice that I created to execute
code on the backend is this project. 

# Why?

From what I know, most of the websites that implement some kind of code execution, boot up
a container and execute it inside it. In my case (the blogging platform from above), that proved
as wasteful since booting up a container took time and memory. Roughly, every container (every code
execution) took around 50MB of memory to boot the container into a running state and that is before starting to run
the code. Also, some language like _C_ or _Go_ require you to compile them and then run the code which
takes time and memory. 

# How execman works

**execman** creates containers before you run your code and then runs them concurrently with
**docker exec** command. In front of the containers (you specify the number of containers to run), 
there is a load balancer with a certain number of workers (which you also specify). Running code
is load balanced on the created containers and executed with **docker exec**. That means that if users
send 10 requests in parallel, those requests will be executed in parallel on any number of containers
that you specify to boot up. 

For example, if you have 10 workers and 1 container for the Ruby language, 10 requests will be executed
in parallel on this one container in roughly the time that it takes to actually execute them on your own local
machine without booting up your the container since it is already running. 

# How to use it


