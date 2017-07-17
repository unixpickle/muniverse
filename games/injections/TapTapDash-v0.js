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
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Level);
        faketime.advance(100);

        c2.globalVar('GameOver').setValue = () => {
          gameOver = true;
          endgameScore = rawScore();
        };

        // Disable menu button.
        construct2api.disableActions([
          "311337503160521",
          "103607899942175",
          "178869730448082",
          "882864135313993",
          "645512358232627"
        ]);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  }

})();
