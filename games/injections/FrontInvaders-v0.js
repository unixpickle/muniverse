(function() {

  // Game speed depends on frame rate.
  faketime.animationFrameRate = 60

  var gameOver = false;
  window.muniverse = {
    init: function(options) {
      // Give the assets a moment to load.
      return (new Promise((resolve) => {
        setTimeout(resolve, 1000);
      })).then(function() {
        faketime.pause();
        GAME.$id('menu-start').click();

        // Add one level at a time to increase speed.
        for (var i = 1; i < options.level; ++i) {
          GAME.state.level++;
          GAME.Utils.NewLevel();
        }
        faketime.advance(100);

        GAME.Utils.Alert = () => gameOver = true;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(GAME.state.points);
    }
  }

})();
