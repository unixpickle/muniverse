(function() {

  var LOAD_TIMEOUT = 60000;

  function pressZ() {
    ['keydown', 'keyup'].forEach((evt) => {
      var e = new jQuery.Event(evt);
      e.key = 'z';
      e.keyCode = 90;
      e.which = 90;
      jQuery(document).trigger(e);
    });
  }

  function rawScore() {
    return cr_getC2Runtime().Qq.find((x) => x.name === 'score').data;
  }

  function loadingPercent() {
    return cr_getC2Runtime().Qq.find((x) => x.name === 'loading').data;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().lg &&
          cr_getC2Runtime().lg.length === 31 &&
          cr_getC2Runtime().lg.every((x) => x.complete) &&
          loadingPercent() === 100;
      }).then(function() {
        kongregateAPI.getAPI().stats.submit = (name, score) => {
          gameOver = true;
          endgameScore = score;
        };

        faketime.pause();

        pressZ();
        faketime.advance(100);

        if (cr_getC2Runtime().ea.name !== 'pl_r') {
          throw 'tap did not start game';
        }

        pressZ();
        faketime.advance(100);
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
