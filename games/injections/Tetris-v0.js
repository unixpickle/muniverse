(function() {

  // Addresses from https://datacrystal.romhacking.net/wiki/Tetris_(NES):RAM_map
  var GAME_OVER_ADDR = 0x58;
  var SCORE_ADDR = 0x53

  setupJSNES({
    ROM: 'roms/Tetris (U) [!].nes',
    init: function() {
      faketime.advance(5000);

      // Skip copyright and homescreen, then select A-type
      // and level 0.
      for (var i = 0; i < 4; ++i) {
        triggerKey(ENTER_KEY);
      }
    },
    step: function(millis) {
      return globalJSNES.cpu.mem[GAME_OVER_ADDR] !== 0;
    },
    score: function() {
      return parseLittleBCD(SCORE_ADDR, 3);
    }
  });

})();
