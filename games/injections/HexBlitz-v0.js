(function() {

  var MENU_TIMEOUT = 60000;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return window.gameState === 'start';
      }).then(function() {
        window.gameComplete = () => gameOver = true;

        window.faketime.pause();
        window.faketime.advance(100);
        window.butEventHandler('startGame', null);

        window.userInput.removeHitArea("pause");
        var oldHandler = window.butEventHandler;
        window.butEventHandler = function(a, b) {
          if (a !== 'pause') {
            return oldHandler(a, b);
          }
        };

        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(score || 0);
    }
  };

})();
