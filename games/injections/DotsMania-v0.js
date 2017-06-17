(function() {

  var LOAD_TIMEOUT = 60000;

  function gameObj() {
    return ((window.Phaser || {}).GAMES || [])[0];
  }

  function rawScore() {
    var manager = gameObj().state.states.game.manager;
    return manager.scrPlaying.score;
  }

  // Avoid pre-game tutorial popups.
  ['mvs', 'tme', 'endls'].forEach((x) => localStorage['dts_tut'+x] = '1');

  var gameOver = false;
  var endgameScore = 0;

  window.muniverse = {
    init: function(options) {
      // The loading menu goes faster if we speed
      // up time (because it's a fake loader).
      window.faketime.pause();
      window.faketime.play(2);
      return pollAndWait(2*LOAD_TIMEOUT, function() {
        return gameObj() &&
          gameObj().stage.children[0].children.length === 5 &&
          gameObj().stage.children[0].children[4].text === 'TAP TO CONTINUE';
      }).then(function() {
        window.faketime.pause();

        // "TAP TO CONTINUE"
        gameObj().stage.children[0].children[4].events.onInputUp.dispatch();
        window.faketime.advance(1000);

        var manager = gameObj().state.states.game.manager;

        // "Play"
        manager.changeScreen(GameScreen.SelectGame);
        window.faketime.advance(1000);

        // Start the corresponding game mode.
        var btnNames = {
          'moves': 'btnMoves',
          'time': 'btnTime',
          'endless': 'btnEndless'
        };
        var btnName = btnNames[options.mode];
        var selectGame = manager.scrSelectGame;
        selectGame[btnName].onClick(selectGame[btnName].context);

        // Disable pause button.
        manager.scrPlaying.buttonPause.events.onInputUp.removeAll();

        // The game can end without time passing due to mouse
        // events, so we *must* catch the event and set a
        // global variable.
        manager.scrPlaying.board.gameOver = function() {
          gameOver = true;
          endgameScore = rawScore();
        };

        window.faketime.advance(400);

        if (!options.bonus) {
          ['Color', 'Dot', 'Moves', 'Time'].forEach((bonus) => {
            var btn = manager.scrPlaying['bonus' + bonus];
            if (btn && btn.icon && btn.icon.events) {
              btn.icon.events.onInputUp.removeAll();
            }
            manager.scrPlaying['onBonus'+bonus+'Click'] = () => false;
          });
        }
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
