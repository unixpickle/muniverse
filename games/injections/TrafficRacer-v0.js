(function() {

  // This game has a levels and achievements system.
  // This system does not seem to interfere with gameplay,
  // and only chimes in once the user dies.

  var PAUSE_UID = 17;

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('Score').getValue();
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game)
        faketime.advance(50);

        c2.globalVar('Death').setValue = () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        };

        var pauseButton = cr_getC2Runtime().objectsByUid[PAUSE_UID];
        pauseButton.y = -100;
        pauseButton.set_bbox_changed();
        faketime.advance(50);
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
