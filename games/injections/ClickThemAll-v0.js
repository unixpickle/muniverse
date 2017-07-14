(function() {

  var LOAD_TIMEOUT = 60000;

  function globalVar(name) {
    return cr_getC2Runtime().Qi.find((v) => v.name === name);
  }

  var gameOver = false;
  var endgameScore = 0;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().uf &&
          cr_getC2Runtime().uf.length === 5 &&
          cr_getC2Runtime().uf.every((img) => img.complete) &&
          cr_getC2Runtime().F &&
          cr_getC2Runtime().F.name === 'Layout 1';
      }).then(function() {
        faketime.pause();
        cr_getC2Runtime().$j['Layout 1'].Dr();
        cr_getC2Runtime().$j['Layout 2'].Mk();
        faketime.advance(100);

        kongregateAPI.getAPI().stats.submit = (name, val) => {
          gameOver = true;
          if (name === 'score') {
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
      return Promise.resolve(gameOver ? endgameScore : globalVar('Point').data);
    }
  };

})();
