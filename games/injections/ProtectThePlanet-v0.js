(function() {

  var LOAD_TIMEOUT = 60000;

  // Skip tutorial.
  window.localStorage.pt_user_id = '5e5c5e7d874ffaaa';
  window.localStorage.ptpgame_mute = '0';
  window.localStorage.ptpgame_tutorial='3';

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.globalGameObj &&
          globalGameObj.state &&
          globalGameObj.state.current === 'MainMenu';
      }).then(function() {
        // Skip menu fade in.
        window.faketime.pause();
        window.faketime.advance(4000);

        globalGameObj.state.states.GameMain.gameOver = () => gameOver = true;

        // Get into the game and skip animations.
        globalGameObj.stage.children[0].children[5].events.onInputDown.dispatch();
        window.faketime.advance(3000);

        // Remove pause button.
        for (var i = 0; i < 10; i++) {
          if (globalGameObj.stage.children[0].children[9].onInputDown) {
            break;
          }
          window.faketime.advance(100);
        }
        if (!globalGameObj.stage.children[0].children[9].onInputDown) {
          throw 'Pause button was not setup.';
        }
        globalGameObj.stage.children[0].children[9].visible = false;
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      var text = globalGameObj.state.callbackContext.scoreText._text;
      return Promise.resolve(parseInt(text) || 0);
    }
  };

})();
