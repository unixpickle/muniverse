(function() {

  var LOAD_TIMEOUT = 60000;

  localStorage.clear();
  localStorage.SaveLevels = '20';

  function rawScore() {
    return parseInt(cr_getC2Runtime().Me[756].text);
  }

  function hijackActions(sids, cb) {
    sids.forEach((sid) => cr_getC2Runtime().gf[sid].ub = cb);
  }

  var gameOver = false;
  var endgameScore = 0;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().Ca &&
          cr_getC2Runtime().Ca.name === 'menu';
      }).then(function() {
        faketime.pause();

        cr_getC2Runtime().Ch['menu'].HA();
        cr_getC2Runtime().Ch[options.level].Op();

        faketime.advance(100);

        var pauseActions = [
          "9559227022866180",
          "2048550937050429",
          "26754623099581",
          "716029373865276",
          "3050427295258096"
        ];
        hijackActions(pauseActions, () => false);

        var endgameActions = [
          "9110674448305312",
          "2150286038020517"
        ];
        hijackActions(endgameActions, () => {
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
