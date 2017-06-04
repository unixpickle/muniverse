(function() {

  var LOAD_TIMEOUT = 60000;

  // Just to prevent an error.
  window.famobi.getMoreGamesButtonImage = () => 'res/logo.png';

  // Prevent wav files from downloading.
  window.famobi_supportedAudio = 0;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return 'undefined' !== typeof window.globalAppObject;
      }).then(function() {
        window.globalAppObject.startGame();

        window.globalAppObject.gameOver = () => gameOver = true;

        window.faketime.pause();
        window.faketime.advance(50);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(globalAppObject.logic.get_score());
    }
  };

})();
