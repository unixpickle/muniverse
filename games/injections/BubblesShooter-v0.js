(function() {

  var MENU_TIMEOUT = 60000;

  var gameIsOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return 'undefined' !== typeof window.gameEngine &&
          gameEngine.scenes.length === 4 &&
          gameEngine.scenes[0].state === 2 &&
          'undefined' !== typeof gameEngine.scenes[3].levelManager.jsonString;
      }).then(function() {
        var cb = function() {
          gameIsOver = true;
        };
        GAME.GameScene.prototype.win = cb;
        GAME.GameScene.prototype.lose = cb;
        GAME.GameScene.prototype.pause = function() {
        };
        gameEngine.changeSceneTo(GAME_MODE.INGAME);
        window.faketime.pause();
        window.faketime.advance(6000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameIsOver);
    },
    score: function() {
      return Promise.resolve(gameEngine.scenes[3].gameHUD.score || 0);
    }
  };

})();
