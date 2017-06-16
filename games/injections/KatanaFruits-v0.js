(function() {

  var MENU_TIMEOUT = 60000;

  // Hijack score.
  var oldCInterface = window.CInterface;
  window.CInterface = function() {
    oldCInterface.call(this);

    var oldRefScore = this.refreshScore;
    this.refreshScore = function(x) {
      oldRefScore.call(this, x);
      window.muniverse.score = function() {
        return Promise.resolve(x);
      };
    }
  }

  // Needed to skip the help pane that appears
  // when you first hit play.
  var oldCHelpPanel = window.CHelpPanel;
  var globalCHelpPanel = null;
  window.CHelpPanel = function() {
    oldCHelpPanel.call(this);
    globalCHelpPanel = this;
  };

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        // Wait for the main menu.
        return 'undefined' !== typeof s_oMenu;
      }).then(function() {
        s_oMenu._onButPlayRelease();
        return pollAndWait(MENU_TIMEOUT, function() {
          // Wait for the help panel to popup.
          return globalCHelpPanel !== null &&
            'undefined' !== typeof s_oGame;
        });
      }).then(function() {
        // Hijack "game over" logic so step() works.
        var gameOverBackup = s_oGame.gameOver;
        s_oGame.gameOver = function() {
          gameOver = true;
          gameOverBackup.call(this);
        };

        // Block the exit button.
        s_oGame.onExit = function() {};

        // Close help panel, exposing the game.
        globalCHelpPanel._onExit();

        // From now on, time only passes due to step().
        window.faketime.pause();

        // Hide popup UI.
        window.faketime.advance(100);
      });
    },

    score: function() {
      return Promise.resolve(0);
    },

    step: function(millis) {
      return new Promise(function(resolve) {
        window.faketime.advance(millis);
        resolve(gameOver);
      });
    }
  };

})();
