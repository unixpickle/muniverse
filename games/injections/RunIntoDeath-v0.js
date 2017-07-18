(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('Points').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();

        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts['Layout 1']);
        faketime.advance(100);

        c2.globalVar('GameOver').setValue = () => {
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
