# Packaging games for μniverse

The [download](download) script downloads and prepares every game. In particular, it creates a directory called `downloaded_games` and makes a sub-directory for each game it downloads.

Each game sub-directory contains an `index.html` file and various game-specific resource files. Each subdirectory also contains an `info.json` file with specific information about that game.

# The `spec.json` file

The [spec.json](spec.json) file contains a JSON array of game objects. Each game object contains various information about a game. Here's what a game object might look like:

```json
{
  "name": "SomeGame-v0",
  "base": "http://someurl",
  "top_injections": [
    "faketime.js"
  ],
  "bottom_injections": [
    "SomeGame-v0.js"
  ],
  "resources": [
    "js/game.js",
    "js/jquery-2.0.3.min.js",
    "sounds/explosion.ogg",
    "css/reset.css",
  ],
  "processors": [
    "some_script"
  ],
  "metadata": {}
}
```

The fields in a game object have the following meaning:

 * **name:** a CamelCase name with a version number after it.
 * **base:** the base URL to fetch as `index.html`. Also the URL to which resource paths are relative.
 * **top_injections:** scripts from the [injections](injections) directory to inject into the top of `<head>`.
 * **bottom_injections:** like *top_injections*, but for the bottom of `<head>`.
 * **resources:** a list of paths to download. All paths are relative to *base*. This list should be exhaustive; it should cover every resource consumed by the game.
 * **processors:** scripts from [process](process) to run on the downloaded game. More on this in [The download process](#the-download-process).
 * **metadata:** game-specific meta-data. This may be used to provide information to processor scripts.

# The download process

When a game is downloaded, three basic steps are taken:

 * A game directory is made and the resources are downloaded, including `index.html`.
 * Each script in the game's `processors` array is run with the game directory as an argument.
 * The `index.html` file is modified to include the injected scripts from `top_injections` and `bottom_injections`.

A processor script can find out information about a game by looking at the game's `info.json` file. The `info.json` file is located in the root of the game directory and contains a copy of the game object from `spec.json`.

# Adding a game

The first thing to do is figure out what resources the game uses. Try to build a complete list of the files loaded by the game. This way, you can build up the `resources` array.

Every game must expose a `window.muniverse` object with three functions, all of which return Promises. These functions are:

 * **init:** get the game to a logical starting point, where an agent can immediately start making actions. The game should be paused in some way so that *step* needn't be called immediately after *init*.
 * **score:** get the current (numerical) score.
 * **step(millis):** advance the game by the given number of milliseconds. Return a boolean which is `true` if the game has ended.

In order to provide the above API, you will likely want to do a few things:

 * Inject the [injections/faketime.js](injections/faketime.js) API in order to control time.
 * Inject your own script to hijack the game and implement the `window.muniverse` API.
 * Write a processor script to make the code easier to handle.
   * Remove unwanted content like ads or analytics.
   * Patch JS code to manipulate it more easily.
