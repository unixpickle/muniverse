(function() {

  var LOAD_TIMEOUT = 60000;

  window.famobi.onOrientationChange = () => false;

  var gameOver = false;
  window.famobi.gameOver = () => gameOver = true;

  var score = 0;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.s_oStage &&
          window.s_oStage.children.length === 7;
      }).then(function() {
        faketime.pause();

        // Fade in.
        faketime.advance(1000);

        oMain.gotoGame();
        faketime.advance(100);

        // The score itself is stored in an annoyingly
        // hidden variable, so we hijack a callback.
        var oldVS = s_oInterface.viewScore;
        s_oInterface.viewScore = function(s) {
          if (!gameOver) {
            score = s;
          }
          oldVS.apply(this, arguments);
        };

        // Disable the exit button.
        // Prevents the game from starting over again,
        // but that is of no consequence.
        s_oGame.onExit = () => false;

        // Allow some time for font to load.
        // Not super important, but will hopefully make
        // things more consistent.
        return new Promise(function(resolve) {
          faketime._backups.setTimeout.call(window, resolve, 1000);
        });
      }).then(function() {
        faketime.advance(100);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(score);
    }
  };

})();
