(function() {

  var LOAD_TIMEOUT = 60000;

  // Game speed depends on the framerate.
  faketime.animationFrameRate = 30;

  function click(x, y) {
    var canvas = document.getElementById('canvas');
    var evt = {
      button: 0,
      pageX: (window.innerWidth-640)/2 + x,
      pageY: (window.innerHeight-512)/2 + y
    };
    canvas.onmousemove(evt);
    faketime.advance(100);
    canvas.onmousedown(evt);
    faketime.advance(100);
    canvas.onmouseup(evt);
    faketime.advance(100);
  }

  function rawScore() {
    return (window._h6 || {})._i6 || 0;
  }

  var gameOver = false;
  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        // These fields seem to get set after the load has
        // completed and the menu has been shown.
        var loadFields = ["_xM1", "_A7", "_Ys1", "_fs1", "_Zt1", "_uu1"];
        return loadFields.every((k) => window.hasOwnProperty(k));
      }).then(function() {
        faketime.pause();

        // Make sure the menu is visible.
        faketime.advance(500);

        // Hit "Start Game".
        click(330, 232);

        // Hit "Next Page".
        click(330, 440);

        // Hit "Why Not?"/"Heck Yeah"/etc.
        click(512, 450);

        // Click the button for the difficulty.
        var difficulty = ['easy', 'medium', 'hard'].indexOf(options.difficulty);
        click(130, 160 + 100*difficulty);

        window.kongGetAPI = () => {
          gameOver = true;
          endgameScore = rawScore();
          return -1;
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
  }

})();
