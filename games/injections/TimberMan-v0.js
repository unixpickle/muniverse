(function() {

  var c2 = window.construct2api;

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return c2.waitStart('Game').then(function() {
        window.faketime.pause();
        c2.hijackAction('8916518275362286', () => gameOver = true);
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
