(function() {

  var PAUSE_UID = 4;

  var c2 = construct2api;

  var scoreMode = 'distance';
  function rawScore() {
    if (scoreMode === 'distance') {
      return c2.globalVar('Score').getValue();
    } else if (scoreMode === 'coins') {
      return c2.globalVar('Coins').getValue();
    }
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function(options) {
      scoreMode = options.scoreMode;
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);
        faketime.advance(100);

        var pauseButton = cr_getC2Runtime().objectsByUid[PAUSE_UID];
        pauseButton.y = -500;
        pauseButton.set_bbox_changed();
        faketime.advance(100);

        cr_getC2Runtime().doChangeLayout = () => {
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
