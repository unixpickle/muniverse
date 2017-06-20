(function() {

  var PAUSE_BUTTON_UID = 453;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('CoverScreen').then(function() {
        window.faketime.pause();
        cr_getC2Runtime().layouts.CoverScreen.stopRunning();
        cr_getC2Runtime().layouts.level5.startRunning();
        cr_getC2Runtime().objectsByUid[PAUSE_BUTTON_UID].visible = false;

        // Highscore will get set to a positive number
        // when the game ends.
        c2.globalVar('Highscore').setValue(-1);

        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(c2.globalVar('Highscore').getValue() >= 0);
    },
    score: function() {
      var high = c2.globalVar('Highscore').getValue();
      var cur = c2.globalVar('CurrentScore').getValue();
      return Promise.resolve(high >= 0 ? high : cur);
    }
  };

})();
