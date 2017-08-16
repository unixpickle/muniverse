(function() {

  var SCORE_UID = 8;

  var c2 = construct2api;

  // Game difficulty depends on framerate.
  faketime.animationFrameRate = 60;

  function rawScore() {
    return parseInt(cr_getC2Runtime().objectsByUid[SCORE_UID].text);
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts.Game);
        faketime.advance(50);

        construct2api.hijackAction('373931229090685', () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        });

        // Disable pause button.
        c2.disableActions(["6745344001251162", "722163754407836"]);

        // Disable back button.
        c2.disableActions(["5569368277405306", "666935170108070"]);
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
