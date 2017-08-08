(function() {

  var EXIT_IMAGE = 'images/buttonexit-sheet0.png';

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('PlayerPoints').data
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        cr_getC2Runtime().doChangeLayout(cr_getC2Runtime().layouts.Game);

        // Skip welcome menu.
        faketime.advance(4000);

        var gameState = construct2api.globalVar('GameState');
        var oldSetVal = gameState.setValue;
        gameState.setValue = function(val) {
          if (val === 'gameover') {
            if (!gameOver) {
              gameOver = true;
              endgameScore = rawScore();
            }
          }
          oldSetVal.apply(this, arguments);
        }

        // Disable exit button.
        c2.disableActions([
          "4958226314206813",
          "4975497375852706",
          "1156904674104345",
          "8054642562217905",
          "6922031131079974",
          "3214877790243649",
          "895958658397325",
          "4611794849231229",
          "5415737267734318",
          "3394980371386919",
          "420087145010897",
          "704704524328438",
          "206343347921578"
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
