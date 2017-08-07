(function() {

  // Addresses from https://datacrystal.romhacking.net/wiki/Pac-Man:RAM_map
  var LIVES_ADDR = 0x67;
  var SCORE_ADDR = 0x70;

  setupJSNES({
    ROM: 'roms/Pac-Man (U) [!].nes',
    init: function() {
      faketime.advance(5000);
      triggerKey(ENTER_KEY);
    },
    step: function(millis) {
      // Check if we have zero lives left.
      return globalJSNES.cpu.mem[LIVES_ADDR] == 0;
    },
    score: function() {
      var score = 0;
      for (var i = 0; i < 6; ++i) {
        var digit = globalJSNES.cpu.mem[SCORE_ADDR + i];
        score += digit * Math.pow(10, i+1);
      }
      return score;
    }
  });

})();
