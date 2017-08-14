(function() {

  var c2 = window.construct2api;

  var scoreMode = null;
  function rawScore() {
    if (scoreMode === 'distance') {
      return c2.globalVar('Dist').getValue();
    } else {
      return c2.globalVar('Score').getValue();
    }
  }

  var gameOver = false;
  var endgameScore = null;
  window.muniverse = {
    init: function(options) {
      return init4j().then(function() {
        scoreMode = options.scoreMode;

        faketime.pause();
        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts.Game);
        faketime.advance(100);

        // Disable pause button.
        c2.disableActions(["5053699760084754"]);

        c2.globalVar('Isover').setValue = (over) => {
          if (over && !gameOver) {
            gameOver = true;
            endgameScore = rawScore();
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
