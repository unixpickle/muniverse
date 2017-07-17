(function() {

  function rawScore() {
    return construct2api.globalVar('Score').getValue();
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);
        faketime.advance(100);

        // There is a log message on game over.
        var oldLog = console.log;
        console.log = function(msg) {
          if (msg === 'GameOver') {
            gameOver = true;
            endgameScore = rawScore();
          }
          oldLog.apply(this, arguments);
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
