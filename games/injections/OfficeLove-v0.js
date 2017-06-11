(function() {

  var LOAD_TIMEOUT = 60000;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          game.state &&
          game.state.current === 'cmainmenu';
      }).then(function() {
        var oldPauseFunc = cgame.prototype.downPause;
        cgame.prototype.downPause = function() {
          if (!this.isPaused) {
            return;
          }
          return oldPauseFunc.apply(this, arguments);
        };
        cgame.prototype.showGameOver = () => gameOver = true;

        game.state.start("cgame");

        // Wait for other assets to load.
        return pollAndWait(LOAD_TIMEOUT, () => game.state.states.cgame.bg);
      }).then(function() {
        window.faketime.pause();
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(game.state.states.cgame.score);
    }
  };

})();
