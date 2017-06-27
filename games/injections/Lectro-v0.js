(function() {

  window.indexedDB.deleteDatabase('localforage');

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('start_menu').then(function() {
        faketime.pause();
        cr_getC2Runtime().layouts.start_menu.stopRunning();
        cr_getC2Runtime().layouts.game.startRunning()
        faketime.advance(100);

        // Used to detect game over.
        c2.globalVar('high_score').setValue(-1);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('high_score').getValue() >= 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
