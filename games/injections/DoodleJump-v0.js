(function() {

  var LOAD_TIMEOUT = 60000;

  localStorage.clear();

  function gameObj() {
    return ((window.Phaser || {}).GAMES || [])[0] || null;
  }

  function rawScore() {
    return parseInt(gameObj().state.states.Game.coinsCountLabel._text) || 0;
  }

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return gameObj() &&
          gameObj().isBooted &&
          gameObj().isRunning &&
          gameObj().state &&
          gameObj().state.getCurrentState().key === 'Menu';
      }).then(function() {
        faketime.pause();
        var playButton = gameObj().stage.children[0].children[6];
        ['onInputDown', 'onInputUp'].forEach((event) => {
          playButton.events[event].dispatch();
          faketime.advance(100);
        });
        for (var i = 0; i < LOAD_TIMEOUT/100; ++i) {
          faketime.advance(100);
          if (Doodle.GameState.pauseButton) {
            break;
          }
        }
        if (!Doodle.GameState.pauseButton) {
          throw 'no pause button';
        }
        Doodle.GameState.gameOver = () => {
          gameOver = true;
          endgameScore = rawScore();
        };
        Doodle.GameState.pauseButton.visible = false;
        faketime.advance(100);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  };

})();
