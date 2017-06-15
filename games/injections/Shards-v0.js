(function() {

  var LOAD_TIMEOUT = 60000;

  // Make the app run at the correct speed.
  window.faketime.animationFrameRate = 60;

  // Unlock every level.
  var saveData = {
    levelStats: [],
    musicOn: true,
    soundOn: true,
    currentLanguage: 'en',
    shardiensOn: true,
    lastLoadedLevel: 0,
    difficulty: 1,
    hiScore: 0,
  };
  for (var i = 0; i < 80; ++i) {
    saveData.levelStats.push({unlocked: true, time: 0, stars: 0});
  }
  localStorage.shards_save = JSON.stringify(saveData);

  var gameOver = false;

  window.famobi_analytics.trackEvent = function(name) {
    if (name === 'EVENT_LEVELFAIL' || name === 'EVENT_LEVELSUCCESS') {
      gameOver = true;
    }
  };

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.globalGameObj &&
          globalGameObj.state.current === 'Menu' &&
          globalGameObj.load.hasLoaded;
      }).then(function() {
        // Disable pause button.
        Shards.Board.prototype.onPause = () => false;

        // Skip startup animation.
        window.faketime.advance(2000);
        globalGameObj.state.start('Levels', true, false);
        window.faketime.advance(100);

        // Begin game.
        var levelObj = {parent: {level: options.level}};
        globalGameObj.state.states.Levels.onStartLevel(levelObj, null, true);

        return pollAndWait(LOAD_TIMEOUT, function() {
          return globalGameObj.state.current === 'Play';
        });
      }).then(function() {
        window.faketime.pause();

        // Skip initial animation.
        window.faketime.advance(3000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      var board = globalGameObj.state.states.Play._board;
      return Promise.resolve(board._score._score);
    }
  };

})();
