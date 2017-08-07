// Various APIs for creating an NES environment using
// the JSNES emulator.
//
// The main routine is setupJSNES, which takes a set of
// routines for hooking into a specific game.
// See PacMan-v0.js for an example of usage.
//
// Thanks to Ben Firshman for sharing JSNES.
(function() {

  var LOAD_TIMEOUT = 60000;

  // APIs for triggering keys in the emulator.
  // Useful for navigating menus.
  window.ENTER_KEY = 13;
  window.triggerKey = function(code) {
    var evt = {keyCode: code};
    ['keyDown', 'keyUp'].forEach((f) => {
      globalJSNES.keyboard[f](evt);
      faketime.advance(100);
    });
  };

  // Utility for parsing BCD numbers.
  window.parseLittleBCD = function(addr, bytes) {
    var res = 0;
    for (var i = addr; i < addr+bytes; ++i) {
      var byteVal = globalJSNES.cpu.mem[i]
      var val1 = byteVal & 0xf;
      var val2 = byteVal >> 4;
      var offset = 2 * (i - addr);
      res += val1*Math.pow(10, offset) + val2*Math.pow(10, offset+1);
    }
    return res;
  }

  window.setupJSNES = function(gameImpl) {
    window.muniverse = {
      init: function(options) {
        // Put the canvas in the top-left corner.
        var elem = $('.nes');
        document.body.textContent = '';
        $(document.body).append(elem);
        $('canvas').css({position: 'fixed'});
        $(document.body).css({margin: 0, overflow: 'hidden'});

        // Start loading the ROM (uses AJAX)
        $('select').val(gameImpl.ROM);
        globalJSNES.ui.loadROM();

        return pollAndWait(LOAD_TIMEOUT, () => {
          return globalJSNES.isRunning
        }).then(() => {
          faketime.pause();
          gameImpl.init(options);
        });
      },
      step: function(millis) {
        faketime.advance(millis);
        return Promise.resolve(gameImpl.step(millis));
      },
      score: function() {
        return Promise.resolve(gameImpl.score());
      }
    };
  };

})();
