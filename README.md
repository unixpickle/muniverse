# μniverse

μniverse is my attempt to wrap a bunch of HTML5 games in a single API. This is similar to [openai/universe](https://github.com/openai/universe), except that it will avoid Flash and VNC.

Compared to OpenAI Universe, μniverse will give the following advantages:

 * No need to play games in real-time.
 * No need for a neural network to read scores from screenshots.
 * Fewer [bugs and glitches](https://github.com/openai/universe/issues/187) due to automation.
 * No shady Docker containers without accompanying source code

Most of the above advantages come from focusing on HTML5 games rather than Flash games.

# Contents

 * [chrome/](chrome) - Go API for controlling a headless Chrome instance.
