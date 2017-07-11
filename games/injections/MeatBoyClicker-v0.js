(function() {

  var LOAD_TIMEOUT = 60000;

  function rawScore() {
    return game.state.states.main.currentScore;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      document.body.style.overflow = 'hidden';
      document.getElementsByClassName('sound')[0].style.display = 'none';
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game && game.state.current === 'main';
      }).then(function() {
        faketime.pause();

        var oldFinish = game.state.states.main.gameFinish;
        game.state.states.main.gameFinish = function() {
          gameOver = true;
          endgameScore = rawScore();
          oldFinish.call(this, arguments);
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
