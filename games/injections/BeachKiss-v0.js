(function() {

  var LOAD_TIMEOUT = 60000;

  function rawScore() {
    return parseFloat(Game.instance.countText.text) || 0;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.RES && RES.isGroupLoaded('preload');
      }).then(function() {
        faketime.pause();
        faketime.advance(1000);
        Game.instance.onTouchPlayHandler();
        faketime.advance(7000);
        Game.instance.onclickPlay();
        faketime.advance(100);

        Game.instance.winFun = Game.instance.loseFun = () => {
          gameOver = true;
          endgameScore = rawScore();
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
