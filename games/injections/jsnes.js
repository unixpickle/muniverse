(function() {

  // Thanks to the following:
  // - Ben Firshman for sharing jsnes
  // - This memory map: https://datacrystal.romhacking.net/wiki/Pac-Man:RAM_map

  var LOAD_TIMEOUT = 60000;

  var ENTER_KEY = 13;

  function triggerKey(code) {
    var evt = {keyCode: code};
    ['keyDown', 'keyUp'].forEach((f) => {
      globalJSNES.keyboard[f](evt);
      faketime.advance(100);
    });
  }

  var games = {
    'PacMan-v0': {
      ROM: 'roms/Pac-Man (U) [!].nes',
      init: function() {
        faketime.advance(5000);
        triggerKey(ENTER_KEY);
      },
      step: function(millis) {
        // Check if we have zero lives left.
        return Promise.resolve(globalJSNES.cpu.mem[0x67] == 0);
      },
      score: function() {
        var score = 0;
        for (var i = 0; i < 6; ++i) {
          var digit = globalJSNES.cpu.mem[0x70 + i];
          score += digit * Math.pow(10, i+1);
        }
        return Promise.resolve(score);
      }
    }
  };

  var runningGame = null;
  window.muniverse = {
    init: function(options) {
      // Put the canvas in the top-left corner.
      var elem = $('.nes');
      document.body.textContent = '';
      $(document.body).append(elem);
      $('canvas').css({position: 'fixed'});
      $(document.body).css({margin: 0, overflow: 'hidden'});

      // Start loading the ROM (uses AJAX)
      runningGame = games[options.game];
      $('select').val(runningGame.ROM);
      globalJSNES.ui.loadROM();

      return pollAndWait(LOAD_TIMEOUT, () => {
        return globalJSNES.isRunning
      }).then(() => {
        faketime.pause();
        runningGame.init();
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return runningGame.step();
    },
    score: function() {
      return runningGame.score();
    }
  };

})();
