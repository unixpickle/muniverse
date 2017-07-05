(function() {

  var LOAD_TIMEOUT = 60000;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        // Wait for play/volume control buttons to be
        // visible (i.e. for the load to complete).
        return window.globalWa &&
          window.globalWa.ia &&
          window.globalWa.ia.length === 3 &&
          window.globalWa.ia[0].D &&
          window.globalWa.ia[0].D.G;
      }).then(function() {
        // Trigger mouse handler to prevent the app from
        // thinking we are on a mobile device.
        globalWa.dk();

        // Go to the player select page.
        globalWa.play();

        return pollAndWait(LOAD_TIMEOUT, () => window.globallc);
      }).then(function() {
        faketime.pause();
        globallc.select([options.player]);
        faketime.advance(100);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(window.globalGameOver || false);
    },
    score: function() {
      return Promise.resolve(globalU.bh);
    }
  };

})();
