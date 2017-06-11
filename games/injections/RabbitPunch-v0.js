(function() {

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        cr_getC2Runtime().layouts.menu.stopRunning();
        cr_getC2Runtime().layouts.game.startRunning();
        window.faketime.pause();
        window.faketime.advance(1000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(c2.globalVar('time130').getValue() === 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
