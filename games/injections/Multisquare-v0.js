(function() {

  var LOAD_TIMEOUT = 60000;

  var PLAY_BUTTON = [0, 0, 0, 4];

  function componentAtPath(path) {
    var obj = (window.globalMain || {}).stage;
    path.forEach((idx) => obj = ((obj || {}).__children || [])[idx]);
    return obj;
  }

  var gameOver = false;
  var endgameScore = 0;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return componentAtPath(PLAY_BUTTON);
      }).then(function() {
        faketime.pause();
        componentAtPath(PLAY_BUTTON).__eventMap.h.mouseDown[0].callback();

        // Skip moving part of the tutorial.
        faketime.advance(5000);

        // Start a fresh game without a tutorial.
        globalMain.fMousePause();
        faketime.advance(1000);
        globalMain.fMouseRestart()
        faketime.advance(1000);

        globalMain.fGameOver = (arg) => {
          if (arg === 1) {
            // This gets called if the user tries to pause.
            // By returning, we prevent the pause.
            return;
          }
          endgameScore = globalMain.score;
          gameOver = true;
        };
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : globalMain.score);
    }
  };

})();
