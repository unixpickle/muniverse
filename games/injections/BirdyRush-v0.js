(function() {

  var LOAD_TIMEOUT = 60000;
  var PLAY_BUTTON = [0, 173, 0];
  var PAUSE_BUTTON = [0, 172, 1];

  // Code taken from BurninRubber-v0.js.
  function childAtPath(path) {
    var child = ((window.game || {}).stage) || {};
    path.forEach((x) => child = ((child || {}).children || [])[x]);
    return child || null;
  }

  function clickButton(button) {
    if (!button || !button.events || !button.events.onInputDown) {
      throw 'cannot click null button';
    }
    button.events.onInputDown.dispatch();
  }

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          window.game.isBooted &&
          window.game.isRunning &&
          window.game.state &&
          window.game.state.getCurrentState().key === 'GameState' &&
          childAtPath(PLAY_BUTTON);
      }).then(function() {
        window.faketime.pause();

        GameState.prototype.onGameOver = () => gameOver = true;

        clickButton(childAtPath(PLAY_BUTTON));
        childAtPath(PAUSE_BUTTON).events.onInputDown.removeAll();
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      // The score is -5 before the game starts.
      return Promise.resolve(Math.max(score, 0));
    }
  };

})();
