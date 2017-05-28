(function() {

  var MENU_TIMEOUT = 60000;

  // Deal with exception from drawing an image.
  window.famobi.getMoreGamesButtonImage = () => '';

  var mGame = null;
  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return 'undefined' !== typeof window.injectedPlay;
      }).then(function() {
        window.injectedPlay();
        return pollAndWait(MENU_TIMEOUT, function() {
          return 'undefined' !== typeof window.injectedGame;
        });
      }).then(function() {
        window.faketime.pause();
      });
    },
    step: function(millis) {
      mGame = mGame || window.injectedGame.mGame;
      window.faketime.advance(millis);
      over = (window.injectedGame.mGame === null);
      return Promise.resolve(over);
    },
    score: function() {
      if (!mGame) {
        return Promise.resolve(0);
      } else {
        return Promise.resolve(mGame.model.score);
      }
    }
  };

})();
