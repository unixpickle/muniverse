(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('Playerscore').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        return c2.waitStart('menu');
      }).then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.game);
        faketime.advance(100);

        cr_getC2Runtime().doChangeLayout = () => {
          gameOver = true;
          endgameScore = rawScore();
        };

        // Disable exit button.
        construct2api.disableActions([
          "2491911938410271",
          "7010802820274644",
          "8617698266386244"
        ]);

        // Disable pause button.
        construct2api.disableActions([
          "2559514485315485",
          "953542598173984",
          "3181902191951365",
          "1484471703910257",
          "8872210837809706",
          "8820377280819342",
          "9154001714607094"
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
  };

})();
