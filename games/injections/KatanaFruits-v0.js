(function() {

  var MENU_TIMEOUT = 10000;

  var oldCInterface = window.CInterface;
  window.CInterface = function() {
    oldCInterface.call(this);

    var oldRefScore = this.refreshScore;
    this.refreshScore = function(x) {
      oldRefScore.call(this, x);
      window.muniverse.score = function() {
        return x;
      };
    }
  }

  var oldCHelpPanel = window.CHelpPanel;
  var globalCHelpPanel = null;
  window.CHelpPanel = function() {
    oldCHelpPanel.call(this);
    globalCHelpPanel = this;
  };

  function pollAndWait(timeout, check) {
    var INTERVAL = 100;
    return new Promise(function(resolve, reject) {
      var ival;
      ival = setInterval(function() {
        if (check()) {
          clearInterval(ival);
          resolve();
        }
        timeout -= INTERVAL;
        if (timeout < 0) {
          clearInterval(ival);
          reject('timeout surpassed');
        }
      }, INTERVAL);
    });
  }

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return 'undefined' !== typeof s_oMenu;
      }).then(function() {
        s_oMenu._onButPlayRelease();
        return pollAndWait(MENU_TIMEOUT, function() {
          return globalCHelpPanel !== null &&
            'undefined' !== typeof s_oGame;
        });
      }).then(function() {
        var gameOverBackup = s_oGame.gameOver;
        s_oGame.gameOver = function() {
          gameOver = true;
          gameOverBackup.call(this);
        };

        globalCHelpPanel._onExit();
        window.faketime.pause();
      });
    },

    score: function() {
      return 0;
    },

    step: function(millis) {
      return new Promise(function(resolve) {
        window.faketime.advance(millis);
        resolve(gameOver);
      });
    }
  };

})();
