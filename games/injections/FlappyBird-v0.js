(function() {

  var LOAD_TIMEOUT = 60000;

  // Prevent exceptions for google analytics.
  window.ga = () => false;

  // Falling physics is independent of animation frame
  // rate, so the game dynamics change based on this.
  faketime.animationFrameRate = 60;

  // Clear away high score.
  localStorage.clear();

  function rawScore() {
    return counter.text;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.loader &&
          window.loader.loaded &&
          stage.children.length === 8;
      }).then(function() {
        faketime.pause();

        var canvas = $('#testCanvas');
        document.body.textContent = '';
        $(document.body).append(canvas).css({overflow: 'hidden'});

        canvas.css({
          width: '320px',
          height: '512px',
          position: 'fixed',
          top: 0,
          left: 0
        });

        canvas.on('gameEnd', () => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
        });
      });
    },
    step: function(millis) {
      // Slow things down to a playable speed.
      // This is the actual speed I observe on the site.
      faketime.advance(millis/2);

      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(gameOver ? endgameScore : rawScore());
    }
  };

})();
