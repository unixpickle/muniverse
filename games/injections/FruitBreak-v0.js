(function() {

  var c2 = window.construct2api;

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return c2.waitStart('Main').then(function() {
        window.faketime.pause();

        cr_getC2Runtime().layouts.Main.stopRunning();
        cr_getC2Runtime().layouts.Game.startRunning();

        c2.hijackAction('933214569839576', () => gameOver = true);
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('Score').getValue());
    }
  };

})();
