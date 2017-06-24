(function() {

  var LOAD_TIMEOUT = 60000;

  var c2 = window.construct2api;
  var scoreMode = null;

  window.muniverse =  {
    init: function(options) {
      scoreMode = options.scoreMode;
      return c2.waitStart('CoverScreen').then(function() {
        faketime.pause();
        var layouts = cr_getC2Runtime().layouts;
        layouts.CoverScreen.stopRunning();
        layouts.levela.startRunning();
        faketime.advance(1000);

        // Disable pause button (and other buttons).
        construct2api.disableActions([
          "431423261264543",
          "4831162947952099",
          "2517372025361085",
          "8014570598893449"
        ]);

        // Used to detect game over.
        c2.globalVar('Highscore').setValue(-1);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('Highscore').getValue() >= 0);
    },
    score: function() {
      if (scoreMode === 'stars') {
        return Promise.resolve(c2.globalVar('gemsScore').getValue());
      } else if (scoreMode === 'distance') {
        return Promise.resolve(c2.globalVar('CurrentScore').getValue());
      } else {
        return Promise.reject('unknown score mode: ' + scoreMode);
      }
    }
  };

})();
