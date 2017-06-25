(function() {

  var LOAD_TIMEOUT = 60000;

  function clickButton(btn) {
    btn.clbck.call(btn.clbckCtx);
  }

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          game.state.current === 'PreloadState';
          game.state.states.PreloadState.load &&
          game.state.states.PreloadState.load.hasLoaded;
      }).then(function() {
        faketime.pause();
        game.state.start("r344");
        faketime.advance(2000);

        // Start the game tutorial.
        r314 = r349;
        faketime.advance(2500);

        // Skip through tutorial.
        clickButton(game.stage.children[0].children[15].children[1]);
        faketime.advance(2000);
        clickButton(game.stage.children[0].children[15].children[1]);
        faketime.advance(2000);

        // Disable pause button.
        game.stage.children[0].children[8].clbck = r375.r74 = () => false;
        game.stage.children[0].children[8].x = -100;
        faketime.advance(100);

        window.r102 = () => gameOver = true;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      var field = game.stage.children[0].children[6].children[2].children[0];
      return Promise.resolve(parseInt(field._text) || 0);
    }
  }

})();
