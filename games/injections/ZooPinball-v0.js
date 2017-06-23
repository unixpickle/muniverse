(function() {

  var PAUSE_BUTTON = '14';

  // Needed to play in real-time.
  faketime.animationFrameRate = 60;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('CoverScreen').then(function() {
        var layouts = cr_getC2Runtime().layouts;
        layouts.CoverScreen.stopRunning();
        layouts.level1.startRunning();

        // Used to detect game-over.
        c2.globalVar('Highscore').setValue(-1);

        faketime.pause();
        faketime.advance(500);

        cr_getC2Runtime().objectsByUid[PAUSE_BUTTON].visible = false;

        faketime.advance(5000);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('Highscore').getValue() >= 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('CurrentScore').getValue());
    }
  }

})();
