(function () {

  const LOAD_TIMEOUT = 60000;
  let gameOver = false;
  let score = 0;

  localStorage.clear();

  window.muniverse = {
    init: function (options) {
      return pollAndWait(LOAD_TIMEOUT, () => !!document.getElementById('start'))
        .then(function () {
          document.getElementById('start').click();
          return pollAndWait(LOAD_TIMEOUT, () => !!document.getElementById('play'));
        }).then(function () {
          document.getElementById('play').click();

          const oldEnded = Game.Scene.prototype.ended;
          Game.Scene.prototype.ended = function () {
            const res = oldEnded.apply(this);
            gameOver = gameOver || res;
            return res;
          };

          const oldCoin = Game.Hero.prototype.coin;
          Game.Hero.prototype.coin = function () {
            oldCoin.apply(this);
            ++score;
          };

          window.faketime.pause();
          window.faketime.advance(1000);
        });
    },
    step: async function (millis) {
      // For some reason, the game actually requires that
      // time passes. It could be because the audio tracks
      // timestamps separately, but I suspect it's WebGL
      // weirdness.
      for (let i = 0; i < millis; i += 30) {
        const increment = (i + 30 > millis ? millis - i : 30);
        window.faketime.advance(increment);
        const promise = new Promise((resolve) => {
          faketime._backups.setTimeout.call(window, resolve, 30)
        });
        await promise;
      }
      return gameOver;
    },
    score: function () {
      return Promise.resolve(score);
    }
  };

})();