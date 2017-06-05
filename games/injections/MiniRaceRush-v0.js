(function() {

  var MENU_TIMEOUT = 60000;

  window.gameState = '';
  window.levelStars = 0;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return window.gameState === 'start';
      }).then(function() {
        var oldEnd = window.initGameEnd;
        window.initGameEnd = function() {
          gameOver = true;
          oldEnd();
        };
        window.toggleManualPause = function() {
        };
        initGame();
        window.faketime.pause();
      });
    },
    score: function() {
      return Promise.resolve(window.levelStars);
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    }
  };

})();
