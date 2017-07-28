(function() {

  var c2 = construct2api;

  function rawScore() {
    var suffixes = ['', '2', '3', '4', '5', '6'];
    var sum = 0;
    suffixes.forEach((suffix) => {
      sum += c2.globalVar('points_of_food' + suffix).getValue();
    });
    return sum;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function(options) {
      return init4j().then(function() {
        faketime.pause();
        var layout = cr_getC2Runtime().layouts['tela' + options.level];
        cr_getC2Runtime().doChangeLayout(layout);
        faketime.advance(50);

        var oldLog = console.log;
        console.log = function(msg) {
          if (msg === '游戏结束') {
            if (!gameOver) {
              gameOver = true;
              endgameScore = rawScore();
            }
          }
          oldLog.apply(console, arguments);
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
