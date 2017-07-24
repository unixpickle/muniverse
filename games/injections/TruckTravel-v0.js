(function() {

  var FONT_LOAD_TIME = 1000;
  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('score').getValue();
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function(options) {
      return init4j().then(function() {
        return new Promise((resolve) => setTimeout(resolve, FONT_LOAD_TIME));
      }).then(function() {
        faketime.pause();
        var rt = cr_getC2Runtime();
        rt.doChangeLayout(rt.layouts.character);
        faketime.advance(50);

        var gameCanvas = $('#gameCanvas');
        var y = gameCanvas.height() / 2;
        var x = (gameCanvas.width() / 3) * (options.character + 1);
        ['mousedown', 'mouseup'].forEach((evt) => {
          var e = new jQuery.Event(evt);
          e.pageX = x;
          e.pageY = y;
          gameCanvas.trigger(e);
        });
        faketime.advance(50);
        var oldSetValue = c2.globalVar('gameover').setValue
        c2.globalVar('gameover').setValue = function(val) {
          if (val && !gameOver) {
            gameOver = true;
            endgameScore = rawScore();
          }
          return oldSetValue.apply(this, arguments);
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
