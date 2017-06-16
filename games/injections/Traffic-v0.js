(function() {

  var c2 = window.construct2api;

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        window.faketime.pause();

        // Start the game.
        construct2api.runActions(['6171016947255058']);

        var pauseActions = [
          "6858239657436337",
          "6328526297893416",
          "4669410569959028",
          "1041573101300278",
          "1506440533690307",
          "1441149893660512",
          "7215227716986535",
          "9315631006287768",
          "2374544501664086",
          "5996620403402992"
        ];
        c2.disableActions(pauseActions);

        c2.hijackAction('8241123120201745', () => gameOver = true);
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver || !!c2.globalVar('gameover').getValue());
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
