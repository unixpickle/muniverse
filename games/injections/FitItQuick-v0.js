(function() {

  var LOAD_TIMEOUT = 60000;
  var PLAY_BUTTON = [0, 8];
  var LEVEL_BUTTONS = [];
  for (var i = 0; i < 6; ++i) {
    LEVEL_BUTTONS.push([0, 6+i*3]);
  }

  // Unlock all levels.
  for (var i = 1; i < 30; ++i) {
    document.cookie = 'level_' + i + '=0';
  }

  function gameObj() {
    return ((window.Phaser || {}).GAMES || [])[0];
  }

  function buttonAtPath(path) {
    var game = gameObj();
    if (!game) {
      return null;
    }
    var obj = game.stage._stage;
    path.forEach((idx) => obj = ((obj || {}).children || [])[idx]);
    return obj;
  }

  function clickButton(button) {
    if (!button || !button.events) {
      throw 'cannot click button';
    }
    ['onInputDown', 'onInputUp'].forEach((evt) => {
      button.events[evt].dispatch(button.events[evt].context);
    });
  }

  function rawScore() {
    var time = gameObj().state.states.GameLocation.gameMenu.levelTime || 1;
    return parseInt(Math.round(gameObj().data.levelNum *
      parseInt(Constants.scoreValue) / time));
  }

  var gameOver = false;
  var endScore = null;
  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return gameObj() &&
          gameObj().state &&
          gameObj().state.current === 'MainMenu' &&
          gameObj().load &&
          gameObj().load.hasLoaded;
      }).then(function() {
        FitItQuick.GameMenu.prototype.pauseClick = () => false;
        FitItQuick.GameMenu.prototype.retryClick = () => false;
        FitItQuick.GameLocation.prototype.addVictoryDialog = () => {
          gameOver = true;
          endScore = rawScore();
        };

        window.faketime.pause();
        clickButton(buttonAtPath(PLAY_BUTTON));
        window.faketime.advance(1000);
        clickButton(buttonAtPath(LEVEL_BUTTONS[options.level-1]));
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      if (endScore !== null) {
        return Promise.resolve(endScore);
      }
      return Promise.resolve(rawScore());
    }
  };

})();
