(function() {

  var LOAD_TIMEOUT = 60000;

  localStorage.clear();

  function rawScore() {
    return parseInt(cr_getC2Runtime().ni[28].text) || 0;
  }

  function disableActions(acts) {
    hijackActions(acts, () => false);
  }

  function hijackActions(acts, cb) {
    acts.forEach((x) => cr_getC2Runtime().bj[x].xd = cb);
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().Wm.length === 51 &&
          cr_getC2Runtime().Wm.every((x) => x.complete);
      }).then(function() {
        return pollAndWait(LOAD_TIMEOUT, function() {
          if (cr_getC2Runtime().ec.name === 'Game') {
            return true;
          }

          // Attempt to trigger game start.
          ['mousedown', 'mouseup'].forEach((evt) => {
            var e = new jQuery.Event(evt);
            e.pageX = window.innerWidth / 2;
            e.pageY = window.innerHeight / 2;
            jQuery(document).trigger(e);
          });

          return false;
        });
      }).then(function() {
        faketime.pause();
        faketime.advance(500);

        var menuActions = [
          "3437165976023613",
          "5409343724804213",
          "904653968323165",
          "8155779398581965",
          "6083515226784885",
          "9598613298151746",
          "1018933982423517",
          "5020659027107511"
        ];
        disableActions(menuActions);

        var restartActions = [
          "4069040382989227",
          "3118625563516727",
          "2853997813004963",
          "7618841920666676",
          "9946309224689924"
        ];
        disableActions(restartActions);

        // Hijack game over and "Stage clear".
        hijackActions(["5429103411728184", "2520081502387049"], () => {
          gameOver = true;
          endgameScore = rawScore();
        });
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
