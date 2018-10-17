(function() {

  var LOAD_TIMEOUT = 60000;
  var PLAY_BUTTON = [18, 1];
  var START_BUTTON = [20, 1];

  localStorage['save'] = JSON.stringify({
    "maxH": 11,
    "mon": 0,
    "lang": -1,
    "gr": [0, 0, 0, 0, 10000]
  });

  function clickButton(path) {
    var obj = game.stage.children[0];
    path.forEach((i) => obj = obj.children[i]);
    ['onInputDown', 'onInputUp'].forEach((evt) => {
      var ctx = obj.events[evt]._bindings[0].context;
      obj.events[evt].dispatch(ctx);
    });
  }

  function rawScore() {
    return game.state.states.game.tower.towerHeight;
  }

  var endgameScore = 0;
  var gameOver = false;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          window.game.state &&
          window.game.state.current === 'mainMenu' &&
          window.game.load.hasLoaded &&
          window.game.stage.children &&
          [1, 2].indexOf(window.game.stage.children.length) !== -1 &&
          window.game.stage.children[0].children &&
          window.game.stage.children[0].children.length === 20;
      }).then(function() {
        window.faketime.pause();
        clickButton(PLAY_BUTTON);
        window.faketime.advance(1000);
        gt15 = options.level;
        game.state.start('game');
        window.faketime.advance(1000);
        clickButton(START_BUTTON);
        window.faketime.advance(500);

        // Hide pause button.
        game.stage.children[0].children[14].visible = false

        window.faketime.advance(250);

        window.analyticsOnGameOver = () => {
          gameOver = true;
          endgameScore = rawScore();
        };
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(endgameScore || rawScore());
    }
  };

})();
