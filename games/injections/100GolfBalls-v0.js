(function() {

  var LOAD_TIMEOUT = 60000;

  var isGameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.TTLoader && window.TTLoader.loadedData;
      }).then(function() {
        faketime.pause();

        // Hit the first play button.
        TTLoader.close();
        faketime.advance(1000);

        // Hit the second play button.
        onPlayClick();
        faketime.advance(1000);

        // Hit the third play button.
        stage.objects.forEach((view) => {
          if (view.onclick === goGame) {
            view.onclick.call(window, {target: view});
          }
        });
        faketime.advance(900);

        // Hide and disable the pause button.
        pauseBtn.visible = false;
        faketime.advance(100);

        window.gameOver = () => isGameOver = true;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(isGameOver);
    },
    score: function() {
      return Promise.resolve(gameObjects.score);
    }
  }

})();
