(function() {

  var LOAD_TIMEOUT = 60000;

  var gameOver = false;
  famobi.gameOver = () => gameOver = true;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.gameEngine &&
          window.gameEngine.scenes[0].state === 2;
      }).then(function() {
        faketime.pause();
        gameEngine.changeSceneTo(GAME_MODE.MENU);
        faketime.advance(1500);
        gameEngine.changeSceneTo(GAME_MODE.INGAME);
        faketime.advance(3000);
        gameEngine.scenes[3].gameHUD.pauseButton.visible = false;

        // Wait for audio to load (or fail to do so).
        faketime.play(2);
        return pollAndWait(LOAD_TIMEOUT*2, function() {
          return gameEngine.scenes[3].state >= 6;
        });
      }).then(function() {
        faketime.pause();
        faketime.advance(2000);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameEngine.scenes[3].gameHUD.finalScore);
    }
  };

})();
