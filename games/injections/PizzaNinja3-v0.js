(function() {

  var LOAD_TIMEOUT = 60000;

  var PLAY_BUTTON = [0, 15, 0];

  // Prevent error due to empty image.
  window.famobi.getMoreGamesButtonImage = () => 'assets/img/1up.png';

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game &&
          window.r825 &&
          game.state &&
          game.state.current === 'PreloadState' &&
          game.load.hasLoaded;
      }).then(function() {
        window.faketime.pause();

        r825.prototype.onr825Over = () => gameOver = true;

        // Go to the main menu.
        game.state.start('GameManager');
        window.faketime.advance(2000);

        // Hit the play button.
        r786 = r83;
        window.faketime.advance(1000);

        // Hit the Arcade button.
        r786 = r89;
        window.faketime.advance(1000);

        // Hit checkbox to begin game.
        r793.r607();
        window.faketime.advance(3000);

        // Disable pause button.
        game.stage.children[0].children[23].clbck = () => 0;
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(parseInt(game.world._hash[1].children[1]._text) || 0);
    }
  };

})();
