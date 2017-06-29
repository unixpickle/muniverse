(function() {

  var gameOver = false;
  var endgameScore = null;

  function rawScore() {
    return Math.round(Math.ceil(globalRunner.distanceRan) * 0.025);
  }

  window.muniverse = {
    init: function() {
      var runner = document.getElementsByClassName('runner-container')[0];
      document.body.appendChild(runner);
      document.getElementsByTagName('header')[0].remove();
      document.body.style.overflow = 'hidden';

      globalRunner.gameOver = () => {
        gameOver = true;
        endgameScore = rawScore();
      };

      faketime.pause();
      globalRunner.onKeyDown({keyCode: 38});
      faketime.advance(600);
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
