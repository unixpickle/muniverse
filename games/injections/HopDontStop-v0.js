(function() {

  var LOAD_TIMEOUT = 60000;

  var scoreMode = 'distance';
  function rawScore() {
    return (window.oGameData || {})[scoreMode] || 0;
  }

  var gameOver = false;
  var endgameScore = 0;

  window.muniverse = {
    init: function(options) {
      scoreMode = options.scoreMode;
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.gameState === 'start';
      }).then(function() {
        faketime.pause();

        // Avoid tutorial.
        firstRun = false;

        initGameEnd = () => {
          gameOver = true;
          endgameScore = rawScore();
        };
        toggleManualPause = () => false;

        initGame();
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
