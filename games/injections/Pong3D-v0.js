(function () {

  let gameOver = false;
  let score = 0;

  window.muniverse = {
    init: async function () {
      await pollAndWait(10000, () => 'undefined' !== typeof window.game);

      window.faketime.pause();

      const oldStop = game.stop;
      game.stop = () => {
        gameOver = true;
        if (game.playerWins()) {
          score = 1;
        } else {
          score = -1;
        }
        oldStop.apply(game);
      }
      game.start();
      window.faketime.advance(100);
    },
    step: function (millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function () {
      return Promise.resolve(score);
    }
  };

})();