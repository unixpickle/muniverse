(function() {

  var LOAD_TIMEOUT = 60000;

  // The silly application uses fractional time
  // intervals, like 0.06.
  var oldSetInterval = setInterval;
  setInterval = function(a, b) {
    if (b < 1 && b > 0) {
      // Make it 30 FPS.
      arguments[1] = 1000/30;
    }
    oldSetInterval.apply(this, arguments);
  };

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          game.hoops.length > 0;
      }).then(function() {
        window.faketime.pause();
        game.click = true;
        window.faketime.advance(200);
        game.click = false;
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(game.state !== 'play');
    },
    score: function() {
      return Promise.resolve(game.score);
    }
  };

})();
