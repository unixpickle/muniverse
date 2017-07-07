(function() {

  var LOAD_TIMEOUT = 60000;

  function pressKey(key) {
    codes = {'a': 65, 'd': 68, 'space': 32};
    ['keydown', 'keyup'].forEach((evt) => {
      var e = new jQuery.Event(evt);
      e.key = key;
      e.keyCode = codes[key];
      e.which = codes[key];
      jQuery(document).trigger(e);
    });
  }

  function tapScreen() {
    ['mousedown', 'mouseup'].forEach((evt) => {
      var e = new jQuery.Event(evt);
      e.pageX = document.body.offsetWidth / 2;
      e.pageY = document.body.offsetHeight / 2;
      jQuery(document).trigger(e);
    });
  }

  function readVariable(name) {
    return cr_getC2Runtime().ht.find((x) => x.name === name).data
  }

  function rawScore() {
    // A different global variable is used depending on
    // the number of colors.
    var vars = ['score', 'score_1', 'score_2'];
    return vars.map(readVariable).find((x) => x != 0) || 0;
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().Vg &&
          cr_getC2Runtime().Vg.length === 30 &&
          cr_getC2Runtime().Vg.every((x) => x.complete);
      }).then(function() {
        kongregateAPI.getAPI().stats.submit = (name, score) => {
          gameOver = true;
          endgameScore = score;
        };
        cr_getC2Runtime().ve['426443451296983'].Ra = () => {
          gameOver = true;
          endgameScore = rawScore();
        };

        faketime.pause();

        // Hide layout 2, show layout 1.
        cr_getC2Runtime().Lm['Layout 2'].Ow();
        cr_getC2Runtime().Lm['Layout 1'].In();

        faketime.advance(100);

        // 'd' increases the number of colors.
        for (var i = 2; i < options.colors; ++i) {
          pressKey('d');
          faketime.advance(100);
        }

        // Skip paste intro screen & tutorial.
        ['space', 'a', 'd'].forEach((k) => {
          pressKey(k);
          faketime.advance(100);
        })
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
