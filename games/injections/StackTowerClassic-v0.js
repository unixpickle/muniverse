(function() {

  var LOAD_TIMEOUT = 60000;

  // Prevent error which prevents game from loading.
  window.famobi.getMoreGamesButtonImage = () => 'img/bg.jpg';

  function rawScore() {
    return window.globalMain.game.curStackSize;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.famobi.gameOver = function() {
    gameOver = true;
    endgameScore = rawScore();
  };

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.globalMain;
      }).then(function() {
        window.faketime.pause();
        globalMain.__children[4].playClick();
        window.faketime.advance(500);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(endgameScore || globalMain.game.curStackSize);
    }
  };

})();
