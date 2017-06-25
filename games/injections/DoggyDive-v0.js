(function() {

  var PAUSE_BUTTON = 65;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('MainMenu').then(function() {
        faketime.pause();
        cr_getC2Runtime().layouts.MainMenu.stopRunning();
        cr_getC2Runtime().layouts.InGame.startRunning();
        faketime.advance(100);

        // Hide and disable pause button.
        cr_getC2Runtime().objectsByUid[PAUSE_BUTTON].visible = false;
        c2.disableActions([
          "9774798119969012",
          "3228471968258825",
          "3446333115600463",
          "536174030423607",
          "9155803286792680",
          "361534206431897",
          "4601013188734914",
          "2273611736331103",
          "2097141192616796",
          "673818012721822",
          "3037788876766021",
          "1457415893568723",
          "1173004408349681"
        ]);

        faketime.advance(100);

        // Detect game over.
        c2.globalVar('HighScore').setValue(-1);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('HighScore').getValue() >= 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('Score').getValue());
    }
  };

})();
