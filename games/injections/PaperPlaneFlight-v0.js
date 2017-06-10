(function() {

  var LOAD_TIMEOUT = 60000;

  // Don't show the pre-game tutorial.
  window.localStorage['flambe:tutShown'] = 't';

  var gameOver = false;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.globalGameObj &&
          globalGameObj.Main &&
          globalGameObj.Main.soundBtn &&
          globalGameObj.Main.soundBtn.get_visible();
      }).then(function() {
        var planeType = options.planeType || 'default';
        globalGameObj.MainHolder.state.chosenPlaneType = planeType;

        window.faketime.pause();
        globalGameObj.MainHolder.enterPlayingScene(false);
        globalGameObj.Main.pause.set_pointerEnabled(false);
        globalGameObj.MainHolder.state.endGame = () => gameOver = true;
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(globalGameObj.MainHolder.state.scoreCur);
    }
  };

})();
