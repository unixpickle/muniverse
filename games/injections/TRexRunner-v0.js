(function() {

  var c2 = construct2api;

  function rawScore() {
    return c2.globalVar('game_points').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return init4j().then(function() {
        faketime.pause();
        ['mousedown', 'mouseup'].forEach((evt) => {
          var canvas = $('#gameCanvas');
          var e = new jQuery.Event(evt);
          e.pageX = canvas.width() / 2;
          e.pageY = canvas.height() / 2;
          canvas.trigger(e);
        });
        faketime.advance(500);

        // Catch game over when they reset our score.
        var scoreVar = c2.globalVar('game_points');
        var oldSetValue = scoreVar.setValue;
        scoreVar.setValue = function(val) {
          if (val === 0 && scoreVar.data > 0) {
            gameOver = true;
            endgameScore = scoreVar.data;
          }
          return oldSetValue.apply(this, arguments);
        };

        // Backup plan in case the score was 0 on game over.
        // Can this ever happen???
        var recordVar = c2.globalVar('game_record');
        recordVar.data = -1;
        recordVar.setValue = (val) => {
          if (!gameOver) {
            gameOver = true;
            endgameScore = val;
          }
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
