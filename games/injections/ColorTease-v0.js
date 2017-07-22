(function() {

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
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);
        var pauseButton = cr_getC2Runtime().objectsByUid[4];
        pauseButton.y = -1000;
        pauseButton.set_bbox_changed();
        faketime.advance(100);

        cr_getC2Runtime().layouts.GameOver.startRunning = () => {
          gameOver = true;
          endgameScore = rawScore();
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
