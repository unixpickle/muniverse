(function() {

  var LOAD_TIMEOUT = 60000;

  function globalVar(name) {
    return cr_getC2Runtime().Qi.find((v) => v.name === name);
  }

  function rawScore() {
    return globalVar('Point').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().uf &&
          cr_getC2Runtime().uf.length === 5 &&
          cr_getC2Runtime().uf.every((img) => img.complete) &&
          cr_getC2Runtime().F &&
          cr_getC2Runtime().F.name === 'Layout 1';
      }).then(function() {
        faketime.pause();
        cr_getC2Runtime().$j['Layout 1'].Dr();
        cr_getC2Runtime().$j['Layout 2'].Mk();
        faketime.advance(100);

        kongregateAPI.getAPI().stats.submit = (name, val) => {
          if (!gameOver) {
            gameOver = true;
            if (name === 'score') {
              endgameScore = rawScore();
            }
          }
        };

        // Hijack setValue to detect negative score.
        // The kongregate submit API is not called when the
        // game ends due to too many skull clicks.
        var pointsVar = globalVar('Point');
        var oldSet = pointsVar.md;
        pointsVar.md = function(val) {
          oldSet.apply(this, arguments);
          if (val < 0 && !gameOver) {
            gameOver = true;
            endgameScore = val;
          }
        };
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  };

})();
