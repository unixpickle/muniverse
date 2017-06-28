(function() {

  var LOAD_TIMEOUT = 60000;
  var MAIN_PATH = [0, 0];
  var MAIN_MENU_PATH = [0, 0, 1];

  fg_api.EMPTY_IMAGE = 'img/TranspPixel.png';

  function childAtPath(path) {
    var obj = (openfl_Lib.current || {}).stage;
    path.forEach((idx) => obj = ((obj || {}).__children || [])[idx]);
    return obj;
  }

  function rawScore() {
    return childAtPath(MAIN_PATH).desk.score;
  }

  var gameOver = false;
  var endgameScore = null;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.openfl_Lib &&
          childAtPath(MAIN_MENU_PATH);
      }).then(function() {
        faketime.pause();

        // Simple logic to start the game.
        childAtPath(MAIN_MENU_PATH).playClick();
        faketime.advance(1000);

        childAtPath(MAIN_PATH).desk.finishGame = () => {
          gameOver = true;
          endgameScore = rawScore();
        };
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  }

})();
