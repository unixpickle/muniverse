(function () {

  let gameOver = false;

  window.muniverse = {
    init: function () {
      window.faketime.pause();

      document.getElementById('score').style.display = 'none';
      const gameElem = document.getElementById('game');
      gameElem.style.left = '0';
      gameElem.style.top = '0';

      document.body.style.margin = '0';
      document.body.style.padding = '0';

      game.gameOver = () => gameOver = true;
      game.newGame();
    },
    step: function (millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function () {
      return Promise.resolve(game.score);
    }
  };

})();