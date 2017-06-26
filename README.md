# μniverse [![GoDoc](https://godoc.org/github.com/unixpickle/muniverse?status.svg)](https://godoc.org/github.com/unixpickle/muniverse)

μniverse is my attempt to wrap a bunch of HTML5 games in a single API. This is similar to [openai/universe](https://github.com/openai/universe), except that it avoids Flash and VNC. For more on this, see [Advantages over Universe](#advantages-over-universe)

This is a work in progress. See [Roadmap](#roadmap) for more details.

# Getting Started

µniverse depends on a few software projects. It requires fairly new versions, so package managers like `apt-get` may not be suitable. Here are the packages and links to downloads:

 * [Docker >= 17.03](https://docs.docker.com/engine/installation/#time-based-release-schedule)
 * [Go >= 1.8](https://golang.org/dl/)

Installing Go is particularly involved, as you will want to create a `GOPATH`. See [here](https://golang.org/doc/code.html#GOPATH) for details on that.

Once you have Go and Docker, you can download µniverse in one easy step:

```
go get github.com/unixpickle/muniverse
```

The first time you use µniverse, it will pull a somewhat large Docker container. I recommend you do this manually before you use µniverse:

```
docker pull unixpickle/muniverse:0.54.0
```

For an example of how to use µniverse to train an RL agent, see [muniverse-agent](https://github.com/unixpickle/muniverse-agent). You may also want to checkout the [GoDoc](https://godoc.org/github.com/unixpickle/muniverse) for API details.

# Advantages over Universe

Compared to OpenAI Universe, μniverse will give the following advantages:

 * No need to play games in real-time.
   * No "falling behind" on slower computers
   * Play games faster than real-time on fast machines.
 * No need for a neural network to read scores from screenshots.
 * Fewer [glitches](https://github.com/openai/universe/issues/187) due to menu automation.
 * No unsafe actions (e.g. pressing the Main Menu button).
 * No Docker containers without accompanying source code
   * All code to generate containers is included
   * Open source scripts to download & package games

Most of the above advantages come from focusing on HTML5 games rather than Flash games.

# Contents

 * **this directory** - high-level Go API for controlling environments
 * [chrome/](chrome) - Go API for controlling a headless Chrome instance
 * [games/](games) - scripts for downloading & packaging games
 * [container/](container) - build files for the Docker container
 * [codegen/](codegen) - small program to auto-generate Go games registry
 * [util/](util) - small tools which come in handy while using µniverse.
 * [bindings/](bindings) - bindings for other programming languages.

# Roadmap

Here's what is done so far:

 * Go API for interfacing with headless Chrome.
 * Mechanism for downloading & packaging games.
 * JavaScript interface for controlling time (timers, Date, etc.)
 * Docker container for running headless Chrome.
 * Go API for controlling Docker containers.
 * Simple interface for gym-like environment control.

Here's a (non-exhaustive) to-do list:

 * Get more games.
 * Python bindings & Gym integration.
 * Get WebGL to work without occasional failures.
 * Better way to verify connection to keep-alive socket.
