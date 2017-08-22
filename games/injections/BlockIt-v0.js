(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('Score').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();

        var rt = cr_getC2Runtime();

        // Avoid pause+instructions on first point.
        c2.globalVar('NewGame').setValue(1);

        rt.doChangeLayout(rt.layouts.Game);

        faketime.advance(50);

        rt.doChangeLayout = () => {
          if (!gameOver) {
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
