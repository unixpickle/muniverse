(function() {

  var c2 = construct2api;

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();

        // Click the play button.
        c2.runActions([
          "5816348869269873",
          "8149662076093909",
          "6921579728720853",
          "5845007413182387",
          "4015238935747578",
          "3344508725328036",
          "7602279156690056",
          "9382913974970818",
          "127632827268954",
          "512069354578828"
        ]);

        // This variable is always set on game over.
        c2.globalVar('highScore').setValue = (score) => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = score;
          }
        };

        faketime.advance(1000);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      if (gameOver) {
        return Promise.resolve(endgameScore);
      } else {
        return Promise.resolve(c2.globalVar('score').getValue());
      }
    }
  };

})();
