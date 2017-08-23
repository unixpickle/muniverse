(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('gem').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();

        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts['Layout 2']);
        faketime.advance(50);

        c2.globalVar('GamePro').setValue = () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        }
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
