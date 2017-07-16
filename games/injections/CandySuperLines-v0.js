(function() {

  var LOAD_TIMEOUT = 60000;

  localStorage.clear();

  function disableActions(acts) {
    hijackActions(acts, () => false);
  }

  function hijackActions(acts, cb) {
    acts.forEach((x) => cr_getC2Runtime().Ve[x].cb = cb);
  }

  function globalVar(name) {
    return cr_getC2Runtime().Nx.find((x) => x.name === name);
  }

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().za &&
          cr_getC2Runtime().za.name === 'StartMenu';
      }).then(function() {
        faketime.pause();
        cr_getC2Runtime().Jo.StartMenu.CB();
        cr_getC2Runtime().Jo.Game.Ip();
        faketime.advance(100);

        // Used to detect game over.
        globalVar('HighScore').data = -1;

        // Fallback game over detection.
        var oldLog = console.log;
        console.log = function(msg) {
          if (msg === 'Stage: GAMEOVER') {
            gameOver = true;
          }
          oldLog.apply(this, arguments);
        };

        // Restart, clear, and back buttons.
        var actionButtons = [
          "8960497066753772",
          "3011353272917485",
          "6978793658908829",
          "3390470554991227",
          "263282990387988",
          "3279336312140568",
          "4169097509801835",
          "3088885638640952",
          "3127665510615554",
          "7463438390506662",
          "3825306065874327",
          "175820019859882",
          "2975774470545989",
          "7161881082112804",
          "3892100491949189",
          "2077467982518715",
          "1612884278005609",
          "4322692431982454",
          "4368163676323516"
        ];
        disableActions(actionButtons);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver || globalVar('HighScore').data >= 0);
    },
    score: function() {
      var hs = globalVar('HighScore').data;
      return Promise.resolve(hs >= 0 ? hs : globalVar('Score').data);
    }
  }

})();
